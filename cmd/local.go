/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Deploy a new Bytebot stack locally, without docker or docker-compose",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("local called")
	},
}

func init() {
	newCmd.AddCommand(localCmd)
}
