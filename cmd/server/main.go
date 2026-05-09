package main

import (
	"app/internal/limit"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/internal/config"
	"app/internal/handler"
	"app/internal/repo"
	"app/internal/router"
	"app/internal/service"
	"app/internal/setup"
	"app/pkg/snowid"
)

func main() {
	snowid.Init(1)

	cfg := config.Load_Config()

	//DI依赖注入，分层解耦思想,在main函数中组装各个层的组件，注入依赖，最后启动http服务和后台worker协程
	tb := limit.NewTokenBucket(cfg.RateLimitCapacity, cfg.RateLimitRefillRate)
	ratelimitmiddleware := router.NewRateLimitMiddleware(tb)

	db := setup.InitDB(cfg.DBPath)
	defer db.Close()
	redis := setup.InitResdis(cfg.RedisAddr)
	defer redis.Close()
	redisrepo := repo.NewRedisRepo(redis)
	taskrepo := repo.NewTaskRepo(db)
	taskservice := service.NewTaskService(taskrepo, redisrepo, cfg.WorkerPool, cfg.JobQueue, cfg.ProcessConcurrency)
	h := handler.NewTaskHandler(taskservice)

	//路由注册
	r := router.NewRouter()
	r.Use(router.LogMiddleware, router.RecoverMiddleware)

	r.HandleFunc("/EchoRequestHandler", h.EchoRequestHandler)
	r.HandleFunc("/Getstatus", h.Getstatus)
	r.HandleFunc("/HealthHandler", h.HealthHandler)
	r.HandleFunc("/SlowHandler", h.SlowHandler)
	r.HandleFunc("/Submit", ratelimitmiddleware(h.Submit))

	//组装http.Server
	Service := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go taskservice.StartWorkers(ctx)
	//单开协程来启动http服务，避免阻塞主线程，监听端口并返回错误	
	go func() {

		log.Printf("[Start] the programe running on the port \n %s", cfg.Port)

		if err := Service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[Fatal] listen failed %s", err)
		}
	}()

	//等待系统信号来停止阻塞，如果收到信号，则退出循环，停止阻塞
	<-ctx.Done()

	log.Printf("[shutdown] the programe is shutting down \n %s", cfg.Port)

	//创建一个5秒的超时context来优雅关闭http服务
	shutdownctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Service.Shutdown(shutdownctx); err != nil {
		log.Fatalf("[fatal] shutdown gracefully failed %s", err)
	}

	log.Printf("[Shutdown]  the programe shutdown successfully\n %s", cfg.Port)

	done := make(chan struct{})

	//单开协程等待wait完成之后再关闭done这个通道，然后完成结束，如果时间超过十秒钟的话就直接跳过了完成结束
	go func() {
		taskservice.Wait()
		close(done)
	}()

	//等待done这个通道被关闭，如果时间超过十秒钟的话就直接跳过了完成结束

	select {
	case <-done:
		log.Printf("[Shutdown] worker pool drained")
	case <-time.After(10 * time.Second):
		log.Printf("[Shutdown] worker pool drain timeout")
	}
	log.Printf("[Shutdown] server exited")

}
