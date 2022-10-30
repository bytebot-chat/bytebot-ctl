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

type Config struct {
	Stacks []ctl.Stack
}

var C Config

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "bytebot-ctl",
	Short: "Deploy and manage bytebot stacks",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bytebot-ctl.yaml)")
}

func initConfig() {
	var (
		cfgName string = ".bytebot-ctl"
		cfgType string = "yaml" // Format of the file, not the extension
	)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name "bytebot-ctl.yaml"
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType(cfgType)
		viper.SetConfigName(cfgName)

		var defaultStacks = &[]ctl.Stack{}
		viper.SetDefault("stacks", defaultStacks)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// leaving this here because I still don't trust machines
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		// maybe wanna catch this error
		viper.Unmarshal(&C)
	} else {
		fmt.Println("Error reading in config:" + err.Error())
	}

}
