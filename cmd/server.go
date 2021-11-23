package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"webhook.com/bootstrap"
)

var Config struct {
	host string
	port string
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
	Config.host = conf.GetString("server.host") //监听主机ip
	Config.port = conf.GetString("server.port") //监听端口
}

// 服务根命令
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Service management",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// 启动服务
var startCmd = &cobra.Command{
	Use:   "start [flags]",
	Short: "Server start [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.ServerStart(Config.host, Config.port)
	},
}

// 停止服务
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Server stop ...",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// 初始化
func init() {
	serverCmd.AddCommand(startCmd) //启动服务命令
	serverCmd.AddCommand(stopCmd)  //停止服务
	rootCmd.AddCommand(serverCmd)  //添加服务
}
