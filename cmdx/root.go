package cmd

import (
	"log"
	"os"
	"subdomain/libs"

	"github.com/spf13/cobra"
)

var options = libs.Options{}

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	// RootCmd.PersistentFlags().StringSliceVarP(&options.Module, "module", "m", []string{"passive", "brute"}, "爆破还是被动扫描")
	RootCmd.PersistentFlags().StringVarP(&options.Domain, "domain", "d", "", "输入的域名")
	RootCmd.PersistentFlags().StringVarP(&options.InputFile, "file", "f", "targets.txt", "输入的域名文件")
	RootCmd.PersistentFlags().StringVarP(&options.ConfigFile, "config", "c", "config.yaml", "配置文件地址")

	// RootCmd.PersistentFlags().StringVarP(&options.OutputPath, "output", "o", "results", "配置文件地址")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	libs.InitOptions(&options)
}
