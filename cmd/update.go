/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
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
		return nil
	},
}

// updateCmd represents the update command
var updateTeamCmd = &cobra.Command{
	Use:   "team <team-id>",
	Short: "Update a team",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		team, err := services.NewTeamService().Update(buildTeam(args[0], cmd.Flags()))
		if err != nil {
			return err
		}

		ui.RenderTeamTable([]models.Team{*team})

		return nil
	},
}

// updateCmd represents the update command
var updateUserCmd = &cobra.Command{
	Use:   "user <user-id>",
	Short: "Update a user",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := services.NewUserService().Update(buildUser(args[0], cmd.Flags()))
		if err != nil {
			return err
		}
		// TODO: We need to have a RenderSingleRow function
		ui.RenderUserTable([]models.User{*user})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateTeamCmd.PersistentFlags().StringP("name", "n", "", "Team name")
	updateTeamCmd.PersistentFlags().StringP("description", "d", "", "Team description")
	updateTeamCmd.PersistentFlags().StringP("avatar_url", "a", "", "Possible values: https://someavatar.com/image.png")

	updateCmd.AddCommand(updateTeamCmd)

	updateUserCmd.PersistentFlags().String("first_name", "", "First name")
	updateUserCmd.PersistentFlags().String("last_name", "", "Last name")
	updateUserCmd.PersistentFlags().String("email", "", "Email")

	updateCmd.AddCommand(updateUserCmd)
}

func buildTeam(id string, flags *pflag.FlagSet) *models.Team {
	var team models.Team

	team.ID = id
	team.Name, _ = flags.GetString("name")
	team.Description, _ = flags.GetString("description")
	team.AvatarUrl, _ = flags.GetString("avatar_url")
	return &team
}

func buildUser(id string, flags *pflag.FlagSet) *models.User {
	var user models.User

	user.ID = id
	user.FirstName, _ = flags.GetString("first_name")
	user.LastName, _ = flags.GetString("last_name")
	user.Email, _ = flags.GetString("email")
	return &user
}
