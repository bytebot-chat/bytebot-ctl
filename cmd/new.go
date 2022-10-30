/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>

*/
package cmd

import (
	"fmt"

	"github.com/bytebot-chat/bytebot-ctl/ctl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new bytebot stack",
	Run: func(cmd *cobra.Command, args []string) {

		stack, err := ctl.NewStackWithPrompt()
		if err != nil {
			fmt.Println(err)
		}

		// this allows for stacks with duplicate names and i'm sure that won't ever be a problem
		C.Stacks = append(C.Stacks, *stack)
		viper.Set("Stacks", C.Stacks)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
