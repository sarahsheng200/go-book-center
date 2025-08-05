package main

import (
	"context"
	"fmt"
	"go-book-center/app/common"
	con "go-book-center/app/config"
	"go-book-center/app/database"
	"go-book-center/app/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger = common.Logger

func main() {
	database.MysqlConnection()

	routers := router.InitRouter()
	config := con.Conf

	//if config.Server.UseRedis {
	//	// 初始化redis服务
	//	database.InitRedis()
	//}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: routers,
	}
	log.Println(fmt.Sprintf("Listening and serving HTTP on Port: %d, Pid: %d", config.Server.Port, os.Getpid()))

	go func() {

		if err := server.ListenAndServe(); err != nil {
			logger.Fatalf("listen error: %s\n", err)
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	logger.Info("shutdown server...")

	// 创建5s的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {

		logger.Error("server shutdown:", err)
		log.Fatalf("server shutdown:", err)
	}
	logger.Info("server exiting...")
}
