/*
Copyright © 2022 Bren "fraq" Briggs <code@fraq.io>

*/

package cmd

import (
	"fmt"
	"strings"

	"github.com/bytebot-chat/bytebot-ctl/ctl"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Add a service to your Bytebot stack",
	Run: func(cmd *cobra.Command, args []string) {
		stackNamePrompt := promptui.Select{
			Label: "Which stack to update?",
			Items: C.GetStackNames(),
		}

		_, stackName, err := stackNamePrompt.Run()
		fmt.Println("Modifying stack: " + stackName)
		if err != nil {
			fmt.Println(err)
		}

		appConfig, err := addServicePrompt()
		if err != nil {
			fmt.Println(err)
		}

		// Update the global in-memory config first
		fmt.Println("Updating config...")
		for i, stack := range C.Stacks {
			if stack.Name == stackName {
				C.Stacks[i].Apps = append(C.Stacks[i].Apps, appConfig)
			}
		}

		// Then write it to viper
		viper.Set("stacks", C.Stacks)
		fmt.Println("Writing config to disk...")
		// Then dump the config to disk
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Done!")
	},
}

func init() {
	addCmd.AddCommand(serviceCmd)
}

func addServicePrompt() (ctl.AppConfig, error) {
	var app ctl.AppConfig

	serviceNamePrompt := promptui.Prompt{
		Label: "Name of the service to add",
	}
	id, err := serviceNamePrompt.Run()
	if err != nil {
		return app, err
	}
	app.ID = id

	serviceImageNamePrompt := promptui.Prompt{
		Label: "Docker image name (without tag)",
	}
	app.Image.Name, err = serviceImageNamePrompt.Run()
	if err != nil {
		return app, err
	}

	serviceImageTagPrompt := promptui.Prompt{
		Label: "Docker image tag (whithout \":\")",
	}
	app.Image.Tag, err = serviceImageTagPrompt.Run()
	if err != nil {
		return app, err
	}

	inboundTopicsPrompt := promptui.Prompt{
		Label: "Comma separated list of topics to listen on",
	}
	inboundTopics, err := inboundTopicsPrompt.Run()
	if err != nil {
		return app, err
	}
	app.InboundTopics = strings.Split(inboundTopics, ",")

	outboundTopicsPrompt := promptui.Prompt{
		Label: "Comma separated list of topics to publish on",
	}
	outboundTopics, err := outboundTopicsPrompt.Run()
	if err != nil {
		return app, err
	}
	app.OutboundTopics = strings.Split(outboundTopics, ",")

	return app, nil
}
