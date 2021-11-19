package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"webhook.com/route"
)

func main() {
	var port string //web端口
	flag.StringVar(&port, "p", "8080", "请输入端口号")
	flag.Parse() //解析参数
	router := route.InitRouter()

	// 服务初始化
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// 启动服务
	go server.ListenAndServe()

	// 平滑重启
	listenSignal(server)
}

// 监听退出信号
func listenSignal(httpSrv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shundown:", err)
	}
	log.Println("Server exiting")
}
