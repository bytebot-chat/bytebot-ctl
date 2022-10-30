/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stackCmd represents the stack command
var runStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Start or run a configured stack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run stack called")
	},
}

func init() {
	runCmd.AddCommand(runStackCmd)
}
