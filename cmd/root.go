/*
Copyright © 2026 Andy Hayworth <andy.hayworth@outlook.com>
*/

package cmd

import (
	"os"

	"github.com/andrew-hayworth22/dotlocker/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dotlocker",
	Short: "A CLI tool that keeps your config files in sync",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initViper)
}

func initViper() {
	path, err := config.ConfigPath()
	cobra.CheckErr(err)

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
}
