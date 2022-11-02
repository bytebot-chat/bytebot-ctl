/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/bytebot-chat/bytebot-ctl/ctl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// stackCmd represents the stack command
var runStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Start or run a configured stack",
	Run: func(cmd *cobra.Command, args []string) {
		name := viper.GetString("name")
		createRedis := true

		for _, stack := range C.Stacks {
			switch stack.Kind {
			case "docker-compose":
				composeFile, err := ctl.NewDockerComposeFile(ctl.NewDockerComposeConfig(name, createRedis))
				if err != nil {
					fmt.Println(err)
					return
				}
				err = os.WriteFile("docker-compose.yaml", composeFile, 0644)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	},
}

func init() {
	runCmd.AddCommand(runStackCmd)

}
