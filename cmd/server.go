package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"webhook.com/bootstrap"
)

// 是否开启daemon进程
var daemon bool

// 服务相关的配置
var Config struct {
	logfile string
	host    string
	port    string
}

// 初始化
func init() {
	conf := viper.New()
	conf.AddConfigPath("./")
	conf.SetConfigName("conf")
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	Config.logfile = conf.GetString("server.logfile") //监听的日志文件
	Config.host = conf.GetString("server.host")       //监听主机ip
	Config.port = conf.GetString("server.port")       //监听端口

}

// 服务根命令
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Service management",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// 开启守护进程
func daemonProcess(logfile string) {
	pid := syscall.Getppid()
	if pid == 1 {
		if err := os.Chdir("./"); err != nil {
			panic(err)
		}
		syscall.Umask(0)
		return
	}
	fmt.Println("WebHooks daemon start!")
	fp, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = fp.Close()
	}()
	c := exec.Command(os.Args[0], os.Args[1:]...)
	c.Stdout = fp
	c.Stderr = fp
	c.Stdin = nil
	c.SysProcAttr = &syscall.SysProcAttr{Setsid: true} //TODO TEST

	if err := c.Start(); err != nil {
		panic(err)
	}
	_, _ = fp.WriteString(fmt.Sprintf("[PID] %d Start At %s\n", c.Process.Pid, time.Now().Format("2006-01-02 15:04:05")))
	os.Exit(0)
}

// 启动服务
var startCmd = &cobra.Command{
	Use:   "start [flags]",
	Short: "Server start [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		if daemon { //是否启动daemon进程
			daemonProcess(Config.logfile)
		}
		bootstrap.ServerStart(Config.host, Config.port)
	},
}

// 停止服务
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Server stop ...",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.ServerStop(Config.host, Config.port)
	},
}

// 初始化
func init() {
	startCmd.Flags().BoolVarP(&daemon, "daemon", "d", false, "test info") //只能用于当前命令
	serverCmd.AddCommand(startCmd)                                        //启动服务命令
	serverCmd.AddCommand(stopCmd)                                         //停止服务
	rootCmd.AddCommand(serverCmd)                                         //添加服务
}
