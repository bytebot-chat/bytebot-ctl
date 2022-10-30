/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dockerComposeCmd represents the dockerCompose command
var dockerComposeCmd = &cobra.Command{
	Use:   "dockerCompose",
	Short: "Deploy a new docker-compose based stack for Bytebot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker-compose called")
	},
}

func init() {
	newCmd.AddCommand(dockerComposeCmd)
}
