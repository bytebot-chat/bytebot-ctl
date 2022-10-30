/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var stackName string

// stackCmd represents the stack command
var getStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Get information about a bytebot stack",
	Run: func(cmd *cobra.Command, args []string) {
		name := viper.GetString("name")
		for _, stack := range C.Stacks {
			if stack.Name == name {
				prettyYaml, err := yaml.Marshal(&stack)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("---\n%s\n\n", string(prettyYaml))
			}
		}
	},
}

func init() {
	getCmd.AddCommand(getStackCmd)
	getStackCmd.Flags().StringVarP(&stackName, "name", "n", "", "Name of bytebot stack")
	viper.BindPFlag("name", getStackCmd.Flags().Lookup("name"))
}
