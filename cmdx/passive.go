package cmd

import (
	"log"
	"subdomain/libs"
	"subdomain/passive"

	"github.com/spf13/cobra"
)

var passiveCmd = &cobra.Command{
	Use:   "passive",
	Short: "被动收集域名",
	Long:  `被动收集子域名`,
	RunE:  runPassive,
}

func init() {

	passiveCmd.Flags().StringSliceVarP(&options.Source, "source", "s", libs.Resources, "指定被动收集信息来源")
	RootCmd.AddCommand(passiveCmd)
}

func runPassive(_ *cobra.Command, _ []string) error {
	log.Println("passive......")
	passive.Passive(&options)
	return nil
}
