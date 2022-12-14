/*
Copyright © 2022 theredwiking <theredviking1@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theredwiking/scan-for-ipv4/cmd/calculator"
	"moul.io/banner"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "calculates all possible IPv4 addresses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(banner.Inline("calculating ip range"))
		file, _ := cmd.Flags().GetBool("file")
		path, _ := cmd.Flags().GetString("path")

		if file {
			if path != "" {
				data, err := calculator.CalcPath(path)

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(banner.Inline(data))
			} else {
				data, err := calculator.CalcFile()

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(banner.Inline(data))
			}
		} else {
			data, err := calculator.Calc()

			if err != nil {
				fmt.Println(err)
			}

			for _, ip := range data {
				fmt.Println(ip)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)
	flag := calculateCmd.Flags()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculateCmd.PersistentFlags().String("foo", "", "A help for foo")
	flag.BoolP("file", "f", false, "Outputs into a file")
	flag.String("path", "", "Where to store file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calculateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
