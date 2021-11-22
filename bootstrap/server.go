package bootstrap

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"webhook.com/route"
)

// 启动服务
// port 启动端口号
func ServerStart(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetString("port")

	//路由初始化
	router := route.InitRouter()
	// 服务初始化
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		// 连接服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 平滑重启
	listenSignal(server)
}

// 监听退出信号
func listenSignal(httpSrv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	log.Println("Shutdown Server ...")

	//5秒后安全的退出程序
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shundown:", err)
	}
	log.Println("Server exiting")
}
