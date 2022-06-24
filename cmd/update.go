/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/livesup-dev/livesup-cli/internal/api/models"
	"github.com/livesup-dev/livesup-cli/internal/api/services"
	"github.com/livesup-dev/livesup-cli/internal/ui"
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

		resource := args[0]

		switch resource {
		case "team":
			team := services.UpdateTeam(buildTeam(args[1], cmd.Flags()))
			ui.RenderTeamTable([]models.Team{team})
		default:
			return fmt.Errorf("resource <%s> not found", resource)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.PersistentFlags().StringP("name", "n", "", "Team name")
	updateCmd.PersistentFlags().StringP("description", "d", "", "Team description")
	updateCmd.PersistentFlags().StringP("avatar_url", "a", "", "Possible values: https://someavatar.com/image.png")
}

func buildTeam(id string, flags *pflag.FlagSet) models.Team {
	var team models.Team

	team.ID = id
	team.Name, _ = flags.GetString("name")
	team.Description, _ = flags.GetString("description")
	team.AvatarUrl, _ = flags.GetString("avatar_url")
	return team
}
