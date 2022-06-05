/*
Copyright © 2022 Emiliano Jankowski <ejankowski@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "livesup-cli",
	Short: "Simplifies the information you need when you need it.",
	Long: `We live in the Cloud era, today's applications are about connecting APIs.
To manage, monitor, control your app, teams and business you probably need to log in into 10
different services (or more!). It's even harder if you just joined the organization or
switched to a new team. LiveSup tries to add transparency to the services you use and it
creates a layer that organizes and simplifies the information you need when you need it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.livesup-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
