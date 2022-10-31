/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>

*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Add a service to your Bytebot stack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service called")
	},
}

func init() {
	addCmd.AddCommand(serviceCmd)
}
