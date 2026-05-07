package main

import (
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
	db := setup.InitDB(cfg.DBPath)
	redis := setup.InitResdis(cfg.RedisAddr)
	//DI依赖注入，分层解耦思想，handelr依赖Service，service依赖repo，repo依赖数据库db，最后在main里面进行组装

	redisrepo := repo.NewRedisRepo(redis)
	taskrepo := repo.NewTaskRepo(db)
	taskservice := service.NewTaskService(taskrepo, redisrepo)
	h := handler.NewTaskHandler(taskservice)

	//路由注册
	r := router.NewRouter()
	r.Use(router.LogMiddleware, router.RecoverMiddleware)

	r.HandleFunc("/EchoRequestHandler", h.EchoRequestHandler)
	r.HandleFunc("/Getstatus", h.Getstatus)
	r.HandleFunc("/HealthHandler", h.HealthHandler)
	r.HandleFunc("/SlowHandler", h.SlowHandler)
	r.HandleFunc("/Submit", h.Submit)

	//组装http.Server
	Service := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go taskservice.Worker(ctx)
	//单开协程来启动http服务，避免阻塞主线程
	go func() {

		log.Printf("[Start] the programe running on the port \n %s", cfg.Port)

		if err := Service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[Fatal] listen failed %s", err)
		}

	}()

	//等待系统信号来停止阻塞
	<-ctx.Done()

	log.Printf("[shutdown] the programe is shutting down \n %s", cfg.Port)

	//创建一个5秒的超时context来优雅关闭http服务
	shutdownctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Service.Shutdown(shutdownctx); err != nil {
		log.Fatalf("[fatal] shutdown gracefully failed %s", err)
	}

	log.Printf("[Shutdown]  the programe shutdown successfully\n %s", cfg.Port)

}
