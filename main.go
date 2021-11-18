package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	listenSignal(context.Background(), server)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sigs
	log.Println("notify sigs")
	httpSrv.Shutdown(ctx)
	log.Println("http shudown")

	// select {
	// case <-sigs:
	// 	log.Println("notify sigs")
	// 	httpSrv.Shutdown(ctx)
	// 	log.Println("http shudown")
	// }
}
