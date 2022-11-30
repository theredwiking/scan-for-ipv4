/*
Copyright © 2022 theredwiking <theredviking1@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theredwiking/scan-for-ipv4/cmd/calculator"
)

// privateCmd represents the private command
var privateCmd = &cobra.Command{
	Use:   "private",
	Short: "For working with internal IPv4",
	Long: `Command for working with internal IPv4.
This command can do several different things,
so to learn more type in --help`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetBool("file")
		path, _ := cmd.Flags().GetString("path")

		if file {
			if path != "" {
				data, err := calculator.Calc()

				if err != nil {
					fmt.Println(err)
				}

				for _, ip := range data {
					fmt.Println(ip)
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
	rootCmd.AddCommand(privateCmd)
	flag := privateCmd.Flags()

	flag.BoolP("file", "f", false, "Outputs into a file")
	flag.String("path", "", "Where to store file")
}
