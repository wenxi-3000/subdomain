/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"log"
	"os"
	"path"
	"subdomain/brute"
	"subdomain/passive"
	"subdomain/utils"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runRun(_ *cobra.Command, _ []string) error {
	passive.Passive(&options)
	brute.Brute(&options)

	for input := range options.Inputs {
		var results []string
		resultFile := path.Join(options.Paths.Result, input+"-total"+".txt")
		passivePath := path.Join(options.Paths.Result, input+"-passive"+".txt")
		brutePath := path.Join(options.Paths.Result, input+"-brute"+".txt")
		//结果处理
		//判断结果文件是否存在，存在就删除
		if utils.FileExists(resultFile) {
			os.Remove(utils.NormalizePath(resultFile))
		}
		//结果处理
		file, err := os.OpenFile(resultFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		w := bufio.NewWriter(file)
		// for domain := range options.Presults {
		// 	fmt.Println(domain)
		// 	w.Write([]byte(domain + "\n"))
		// }
		passiveFile, err := utils.FileSlice(passivePath)
		if err != nil {
			log.Println(err)
		}
		results = append(results, passiveFile...)
		bruteFile, err := utils.FileSlice(brutePath)
		if err != nil {
			log.Println(err)
		}
		results = append(results, bruteFile...)

		for _, domain := range results {
			w.Write([]byte(domain + "\n"))
		}
		w.Flush()

	}
	return nil

}
