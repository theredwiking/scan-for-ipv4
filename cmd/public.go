/*
Copyright © 2022 theredwiking <theredviking1@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theredwiking/scan-for-ipv4/cmd/public"
)

// publicCmd represents the public command
var publicCmd = &cobra.Command{
	Use:   "public",
	Short: "Scans for public IPv4 addresses",

	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		public.Scanner(filename)
	},
}

func init() {
	rootCmd.AddCommand(publicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	publicCmd.Flags().String("file", "", "Name of output file")
}
