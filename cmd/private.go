/*
Copyright © 2022 theredwiking <theredviking1@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theredwiking/scan-for-ipv4/cmd/calculator"
	"github.com/theredwiking/scan-for-ipv4/cmd/private"
)

// privateCmd represents the private command
var privateCmd = &cobra.Command{
	Use:   "private",
	Short: "For working with internal IPv4",
	Long: `Command for working with internal IPv4.
This command can do several different things,
so to learn more type in --help`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")

		data, err := calculator.Calc()

		if err != nil {
			fmt.Println(err)
		}

		if path != "" {
			result := private.Scanner(data, path)
			fmt.Println(result)
		} else {
			result := private.Scanner(data, "./result.json")
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(privateCmd)
	flag := privateCmd.Flags()

	flag.String("path", "", "Where to store file")
}
