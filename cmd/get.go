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

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
	Long: `Prints a table of the most important information about the specified resources:

Examples:
	# List all users in ps output format
	livesup-cli get users`,
	Run: func(cmd *cobra.Command, args []string) {
		token := ""
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

			fmt.Println(dst.String())
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
