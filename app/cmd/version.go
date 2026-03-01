package cmd

import (
	"bit-labs.cn/owl/utils"
	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	Use:   "admin:version",
	Short: "查看 admin 程序版本",
	Long:  "查看 admin 应用程序版本号，可以更详细的描述",
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintGreen("v1.0.0")
	},
}
