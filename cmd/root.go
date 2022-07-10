/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"subdomain/libs"

	"github.com/spf13/cobra"
)

var options = libs.Options{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "subdomain",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.subdomain.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&options.Domain, "domain", "d", "", "输入的域名")
	rootCmd.PersistentFlags().StringVarP(&options.InputFile, "file", "f", "targets.txt", "输入的域名文件")
	rootCmd.PersistentFlags().StringVarP(&options.ConfigFile, "config", "c", "config.yaml", "配置文件地址")

	//被动
	rootCmd.PersistentFlags().StringSliceVarP(&options.Source, "source", "s", libs.Resources, "指定被动收集信息来源")

	//主动
	rootCmd.PersistentFlags().IntVarP(&options.Thread, "thread", "t", 1000, "爆破并发量")
	rootCmd.PersistentFlags().IntVar(&options.Retry, "retry", 3, "爆破重复次数")
	rootCmd.PersistentFlags().StringVarP(&options.DictList, "dict", "w", "./brute/dict/53683.txt", "指定子域名字典")
	rootCmd.PersistentFlags().StringVar(&options.DnsServer, "DnsServer", "", "指定dns解析的server")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	libs.InitOptions(&options)
}
