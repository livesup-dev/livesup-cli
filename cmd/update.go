/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/livesup-dev/livesup-cli/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <resource> <resource-id>",
	Short: "Update a resource",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("the resource id is required")
		}

		fmt.Println(args)

		switch args[0] {
		case "user":
			api.UpdateTeam(buildTeam(args[1], cmd.Flags()))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.PersistentFlags().StringP("name", "n", "", "Team name")
	updateCmd.PersistentFlags().StringP("description", "d", "", "Team description")
	updateCmd.PersistentFlags().StringP("avatar_url", "a", "", "Possible values: https://someavatar.com/image.png")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func buildTeam(id string, flags *pflag.FlagSet) api.Team {
	var team api.Team

	team.ID = id
	team.Name, _ = flags.GetString("name")
	team.Description, _ = flags.GetString("description")
	team.AvatarUrl, _ = flags.GetString("avatar_url")
	return team
}
