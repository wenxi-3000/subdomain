/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"subdomain/passive"

	"github.com/spf13/cobra"
)

// passiveCmd represents the passive command
var passiveCmd = &cobra.Command{
	Use:   "passive",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runPassive,
}

func init() {
	rootCmd.AddCommand(passiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runPassive(_ *cobra.Command, _ []string) error {
	passive.Passive(&options)
	return nil
}
