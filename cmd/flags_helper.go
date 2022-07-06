package cmd

import "github.com/spf13/cobra"

func addTeamFlags(command *cobra.Command) {
	command.PersistentFlags().StringP("name", "n", "", "Team name")
	command.PersistentFlags().StringP("description", "d", "", "Team description")
	command.PersistentFlags().StringP("avatar_url", "a", "", "Possible values: https://someavatar.com/image.png")
}

func addUserFlags(command *cobra.Command) {
	command.PersistentFlags().String("first_name", "", "First name")
	command.PersistentFlags().String("last_name", "", "Last name")
	command.PersistentFlags().String("email", "", "Email")
}
