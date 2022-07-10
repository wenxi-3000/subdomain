/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"subdomain/brute"

	"github.com/spf13/cobra"
)

// bruteCmd represents the brute command
var bruteCmd = &cobra.Command{
	Use:   "brute",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runBrute,
}

func init() {
	rootCmd.AddCommand(bruteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bruteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bruteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runBrute(_ *cobra.Command, _ []string) error {
	brute.Brute(&options)
	return nil
}
