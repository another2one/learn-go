package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Verbose bool
var Source string

var testCmd = &cobra.Command{
	Use:     "test1",
	Short:   "test1",
	Long:    "test for cobra",
	Args:    cobra.MinimumNArgs(1), //至少一个参数
	Example: `.\main.exe version test1 1 -s=22	.\main.exe verion test1 1 -s=22 --flag1=2 -v=true`,
	PreRun: func(cmd *cobra.Command, args []string) { // run before  post: run after
		fmt.Println("pre run......")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("run test args %v %v %v ......", args, Verbose, Source)
	},
}

func init() {
	testCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output") // PersistentFlags 全部继承
	testCmd.Flags().String("flag1", "s", "flag1 usage")                                   // Flags 当前命令使用
	testCmd.Flags().StringVarP(&Source, "source", "s", "ss", "Source directory to read from")
	testCmd.MarkFlagRequired("source") // Flags默认是可选的。如果您希望命令在未设置Flags时报告错误，请将其标记为必需
	versionCmd.AddCommand(testCmd)     // version 子命令
}
