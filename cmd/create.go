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

// createCmd represents the update command
var createCmd = &cobra.Command{
	Use:   "create <resource>",
	Short: "Create a resource",
	Long: `Create a resource:

	Examples:
		# create a team
		livesup-cli create team -n="My cool team"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// updateCmd represents the update command
var createTeamCmd = &cobra.Command{
	Use:   "team",
	Short: "Create a team",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		team, err := services.NewTeamService().Create(buildNewTeam(cmd.Flags()))
		if err != nil {
			return err
		}

		ui.RenderTeamTable([]models.Team{*team})

		return nil
	},
}

// updateCmd represents the update command
var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Create a user",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := services.NewUserService().Create(buildNewUser(cmd.Flags()))
		if err != nil {
			return err
		}
		// TODO: We need to have a RenderSingleRow function
		ui.RenderUserTable([]models.User{*user})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Create a function to add the flags on both commands create and update
	createTeamCmd.PersistentFlags().StringP("name", "n", "", "Team name")
	createTeamCmd.PersistentFlags().StringP("description", "d", "", "Team description")
	createTeamCmd.PersistentFlags().StringP("avatar_url", "a", "", "Possible values: https://someavatar.com/image.png")

	createCmd.AddCommand(createTeamCmd)

	createUserCmd.PersistentFlags().String("first_name", "", "First name")
	createUserCmd.PersistentFlags().String("last_name", "", "Last name")
	createUserCmd.PersistentFlags().String("email", "", "Email")

	createCmd.AddCommand(createUserCmd)
}

func buildNewTeam(flags *pflag.FlagSet) *models.Team {
	var team models.Team

	team.Name, _ = flags.GetString("name")
	team.Description, _ = flags.GetString("description")
	team.AvatarUrl, _ = flags.GetString("avatar_url")
	return &team
}

func buildNewUser(flags *pflag.FlagSet) *models.User {
	var user models.User

	user.FirstName, _ = flags.GetString("first_name")
	user.LastName, _ = flags.GetString("last_name")
	user.Email, _ = flags.GetString("email")
	return &user
}
