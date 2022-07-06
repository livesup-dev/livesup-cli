/*
Copyright © 2022 Emiliano Jankowski <ejankowski@gmail.com>

*/
package cmd

import (
	"github.com/livesup-dev/livesup-cli/internal/api/services"
	"github.com/livesup-dev/livesup-cli/internal/ui"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resource>",
	Short: "Display one or many resources",
	Long: `Prints a table of the most important information about the specified resources:

Examples:
	# List all users in ps output format
	livesup-cli get users`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// getCmd represents the get command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		userList, err := services.NewUserService().All()
		if err != nil {
			return err
		}
		ui.RenderUserTable(userList.Users)

		return nil
	},
}

// getCmd represents the get command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		teamList, err := services.NewTeamService().All()
		if err != nil {
			return err
		}
		ui.RenderTeamTable(teamList.Teams)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.AddCommand(usersCmd)
	getCmd.AddCommand(teamsCmd)
}
