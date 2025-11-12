package cmd

import (
	"os"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
	"github.com/unshade/unraidctl/internal"
	"github.com/unshade/unraidctl/pkg/client"
)

var unraidClient *client.UnraidClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "unraidctl",
	Short: "unraidctl is a CLI tool to interact with unraid servers",
	Long: `unraidctl is a CLI tool to interact with unraid servers.
	It stands for "unraid control" and allows you to manage various aspects of your unraid server from the command line.
	It uses the unraid GraphQL API to perform actions and retrieve information.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(config *internal.Config, graphqlClient *graphql.Client) {
	unraidClient = client.NewUnraidClient(config.Api.ApiKey, graphqlClient)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
