/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unshade/unraidctl/internal"
	"github.com/unshade/unraidctl/internal/controllers"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Interact with unraid Shares",
	Long:  `Interact with unraid Shares`,
	Run: func(cmd *cobra.Command, args []string) {
		formater := internal.OutputFormaterSwitcher(internal.OutputFormat(outputFormat))
		controller := controllers.NewShareController(unraidClient, formater)
		switch args[0] {
		case "list":
			controller.ListShares(cmd.Context())
		}
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)

	shareCmd.Args = cobra.MinimumNArgs(1)
	shareCmd.ValidArgs = []string{"list"}
	shareCmd.Flags().StringVarP(&outputFormat, "output", "o", "text", "Output format: json|yaml|text")
}
