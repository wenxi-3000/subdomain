package cmd

import (
	"subdomain/brute"

	"github.com/spf13/cobra"
)

var bruteCmd = &cobra.Command{
	Use:   "passive",
	Short: "被动收集域名",
	Long:  `被动收集子域名`,
	RunE:  runBrute,
}

func init() {

	bruteCmd.Flags().IntVarP(&options.Thread, "thread", "t", 1000, "爆破并发量")
	bruteCmd.Flags().IntVar(&options.Retry, "retry", 3, "爆破重复次数")
	bruteCmd.Flags().StringVarP(&options.DictList, "dict", "w", "./brute/dict/53683.txt", "指定子域名字典")
	bruteCmd.Flags().StringVar(&options.DnsServer, "DnsServer", "", "指定dns解析的server")
	RootCmd.AddCommand(bruteCmd)
}

func runBrute(_ *cobra.Command, _ []string) error {
	brute.Brute(&options)
	return nil
}
