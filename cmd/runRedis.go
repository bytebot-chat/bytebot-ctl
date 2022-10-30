/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// redisCmd represents the redis command
var runRedisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Start a redis server for your Bytebot stacks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("redis called")
	},
}

func init() {
	runCmd.AddCommand(runRedisCmd)
}
