package main

import (
	"context"
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/gredis"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// gin 自己的日志记录, 这个日志只记录请求情况
	// 禁用控制台颜色， 写入日志的时候不需要颜色？（为了性能？）
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f)

	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

	router := routers.InitRouter()

	log.Printf("run server port : %v", setting.ServerSetting.HttpPort)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exitting")

}
