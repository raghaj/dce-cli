package cmd

import (
	"github.com/spf13/cobra"
)

var authUrl string

func init() {
	RootCmd.AddCommand(authCmd)
	authCmd.Flags().StringVarP(&authUrl, "url-override", "u", "", "Override the DCE login url")
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login to dce",
	Run: func(cmd *cobra.Command, args []string) {
		service.Authenticate(authUrl)
	},
}