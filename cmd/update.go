/*
Copyright © 2022 Emiliano Jankowski

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
	Long: `Update a resource:

	Examples:
		# Update a team
		livesup-cli update team d61f5ae8-5cf3-4290-9c4a-dae8ed91eb60 -d="new description"`,
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

	addTeamFlags(updateTeamCmd)

	updateCmd.AddCommand(updateTeamCmd)

	addUserFlags(updateUserCmd)

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
