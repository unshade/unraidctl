/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "list":
			unraidClient.Docker.ListContainers(cmd.Context())
		case "stop":
			unraidClient.Docker.StopContainer(cmd.Context(), args[1])
		case "start":
			unraidClient.Docker.StartContainer(cmd.Context(), args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)

	dockerCmd.Args = cobra.MinimumNArgs(1)
	dockerCmd.ValidArgs = []string{"stop", "restart", "start", "list"}
}
