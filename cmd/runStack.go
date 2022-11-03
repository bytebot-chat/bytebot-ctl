/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/bytebot-chat/bytebot-ctl/ctl"
	"github.com/spf13/cobra"
)

// stackCmd represents the stack command
var runStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Start or run a configured stack",
	Run: func(cmd *cobra.Command, args []string) {
		for _, stack := range C.Stacks {
			switch stack.Kind {
			case "docker-compose":
				composeFile, err := ctl.NewDockerComposeConfig(stack)
				if err != nil {
					fmt.Println(err)
					return
				}

				configYaml, err := composeFile.ToYaml()
				if err != nil {
					fmt.Println(err)
				}
				err = os.WriteFile("docker-compose.yaml", configYaml, 0644)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(configYaml))
			}
		}
	},
}

func init() {
	runCmd.AddCommand(runStackCmd)

}
