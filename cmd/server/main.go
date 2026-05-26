package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"myapp/internal/app/handler"
	"myapp/internal/app/repo"
	"myapp/internal/app/service"
	"myapp/internal/app/setup"
	"myapp/internal/config"
	"myapp/internal/observe"
	"myapp/internal/pkg"
	"myapp/internal/router"
	"myapp/pkg/snowid"
	"strings"
)

func main() {
	// 加载配置
	cfg := config.Load_Config()
	snowid.Init(uint16(cfg.MachineID))
	observe.ApplyRuntimeProfiling(int(cfg.BlockProfileRate), int(cfg.MutexProfileFraction))

	//DI依赖注入，分层解耦思想,在main函数中组装各个层的组件，注入依赖，最后启动http服务和后台worker协程

	db := setup.InitMysqlDB(cfg.MysqlPath, int(cfg.MaxConns), int(cfg.MaxIdleConns), time.Duration(cfg.ConnMaxLifetime)*time.Minute)
	defer db.Close()
	redis := setup.InitResdis(cfg.RedisAddr)
	defer redis.Close()
	/*
		// ========== 一键清空 Redis ==========
		result, err := redis.FlushAll(context.Background()).Result()
		if err != nil {
			fmt.Printf("清空 Redis 失败: %v\n", err)
		} else {
			fmt.Printf("Redis 已全部清空: %s\n", result)
		}
		// ===================================
	*/
	redisrepo := repo.NewRedisRepo(redis)
	taskrepo := repo.NewTaskRepo(db)
	taskservice := service.NewTaskService(taskrepo,
		redisrepo,
		cfg.WorkerPool,
		cfg.JobQueue,
		cfg.ProcessConcurrency)
	h := handler.NewTaskHandler(taskservice)

	tb := pkg.NewTokenBucket(cfg.RateLimitCapacity, cfg.RateLimitRefillRate)
	locallmmd := router.NewLocalLimitMiddleware(tb)
	distlmmd := router.NewDistLimitMiddleware(
		redisrepo,
		cfg.DistLimitMax,
		time.Duration(cfg.DistLimitWindow)*time.Millisecond,
		cfg.DistLimitFailOpen)

	authmiddleware := router.NewAuthMiddleware(cfg.AuthToken)
	signmiddleware := router.NewSignMiddleware(cfg.SignSecret, time.Duration(cfg.SignWindow)*time.Second, redisrepo)
	lmmd := func() router.Middleware {
		switch strings.ToLower(cfg.LimitModel) {
		case "local":
			return locallmmd
		default:
			return distlmmd
		}
	}()

	//路由注册
	r := router.NewRouter()
	r.Use("LogMiddleware", router.LogMiddleware, router.LogPriority)
	r.Use("RecoverMiddleware", router.RecoverMiddleware, router.RecoverPriority)
	r.Use("TraceMiddleware", router.TraceMiddleware, router.TracePriority)

	r.HandleFunc("/EchoRequestHandler", h.EchoRequestHandler)
	r.HandleFunc("/HealthHandler", h.HealthHandler)
	r.HandleFunc("/SlowHandler", h.SlowHandler)
	r.HandleFunc("/Getstatus", router.ChainFunc(h.Getstatus, lmmd, authmiddleware))
	r.HandleFunc("/Submit", router.ChainFunc(h.Submit, lmmd, signmiddleware, authmiddleware))

	//组装http.Server
	Service := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}
	//组装pprof服务
	Pprof := observe.NewPprofServer(cfg.PprofAddr)
	//启动pprof服务
	go func() {
		log.Printf("[Start] the programe running on the port \n %s", cfg.PprofAddr)
		if err := Pprof.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[Fatal] pprof listen failed %s", err)
		}
	}()
	//创建一个context来监听系统信号，如果收到信号，则退出循环，停止阻塞
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	//启动worker协程，避免阻塞主线程，监听端口并返回错误
	//启动主协程，监听端口并返回错误
	go taskservice.StartWorkers(ctx)
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
