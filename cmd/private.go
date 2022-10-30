/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// privateCmd represents the private command
var privateCmd = &cobra.Command{
	Use:   "private",
	Short: "For working with internal IPv4",
	Long: `Command for working with internal IPv4.
This command can do several different things,
so to learn more type in --help`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("private called")
	},
}

func init() {
	rootCmd.AddCommand(privateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// privateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// privateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
