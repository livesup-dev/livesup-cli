/*
Copyright © 2022 Emiliano Jankowski <ejankowski@gmail.com>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/carlmjohnson/requests"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/livesup-dev/livesup-cli/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ApiResponse struct {
	Users []api.User `json:"data"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
	Long: `Prints a table of the most important information about the specified resources:

Examples:
	# List all users in ps output format
	livesup-cli get users`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile(".env") // optionally look for config in the working directory
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		token := viper.GetString("LIVESUP_TOKEN")
		URL := "http://host.docker.internal:4000/"

		fmt.Println("Reading users")

		var response ApiResponse
		err = requests.
			URL(URL).
			Pathf("api/users").
			ContentType("application/json").
			Bearer(token).
			ToJSON(&response).
			Fetch(context.Background())

		if err != nil {
			panic(fmt.Errorf("fatal error reading API: %w", err))
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Email"})

		for _, user := range response.Users {
			t.AppendRows([]table.Row{
				{user.ID, user.FirstName, user.LastName, user.Email},
			})
			t.AppendSeparator()
		}
		t.Render()
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
