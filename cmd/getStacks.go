/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getStacksCmd represents the `bytebot-cli get stacks` command
var getStacksCmd = &cobra.Command{
	Use:   "stacks",
	Short: "Get a list of all Bytebot stacks you are managing",
	Run: func(cmd *cobra.Command, args []string) {
		for _, stack := range C.Stacks {
			fmt.Println(stack.Name)
		}
	},
}

func init() {
	getCmd.AddCommand(getStacksCmd)
}
