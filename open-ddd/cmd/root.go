package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "open-ddd",
		Short: "open-ddd脚手架，代码生成工具",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func init() {}
