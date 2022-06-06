/*
Copyright © 2022 Emiliano Jankowski <ejankowski@gmail.com>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ApiResponse struct {
	Users []User `json:"data"`
}

type User struct {
	ID        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
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
		URL := "http://host.docker.internal:4000/api/users"

		fmt.Println("Reading users")

		client := &http.Client{}

		// Get the data
		request, _ := http.NewRequest("GET", URL, nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			dst := &bytes.Buffer{}
			data, _ := ioutil.ReadAll(response.Body)

			if err := json.Indent(dst, data, "", "  "); err != nil {
				panic(err)
			}

			var response ApiResponse
			json.Unmarshal(data, &response)

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
		} else {
			fmt.Println(err)
		}
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
