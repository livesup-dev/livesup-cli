/*
Copyright © 2022 Emiliano Jankowski <ejankowski@gmail.com>

*/
package cmd

import (
	"errors"

	"github.com/livesup-dev/livesup-cli/pkg/api"
	"github.com/livesup-dev/livesup-cli/pkg/ui"
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
		if len(args) < 1 {
			return errors.New("the resource argument is required")
		}

		resource := args[0]

		var resourceTable ui.ModelTable

		switch resource {
		case "users":
			resourceTable = getUsersTable()
		case "teams":
			resourceTable = getTeamsTable()
		default:
			return errors.New("the resource does not exists")
		}

		ui.DrawTable(resourceTable)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getUsersTable() ui.ModelTable {
	response := api.GetAllUsers()
	return ui.BuildUserTable(response.Users)
}

func getTeamsTable() ui.ModelTable {
	response := api.GetAllTeams()
	return ui.BuildTeamTable(response.Teams)
}
