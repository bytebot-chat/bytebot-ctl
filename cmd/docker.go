/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Spawn a new Docker stack for bytebot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker called")
	},
}

func init() {
	newCmd.AddCommand(dockerCmd)
}
