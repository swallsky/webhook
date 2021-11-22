package cmd

import (
	"webhook.com/bootstrap"

	"github.com/spf13/cobra"
)

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
	Run:   bootstrap.ServerStart,
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
	startCmd.Flags().String("port", "8080", "server port") //默认启动的端口
	serverCmd.AddCommand(startCmd)                         //启动服务命令
	serverCmd.AddCommand(stopCmd)                          //停止服务
	rootCmd.AddCommand(serverCmd)                          //添加服务
}
