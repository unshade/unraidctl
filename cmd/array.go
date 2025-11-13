package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unshade/unraidctl/internal"
	"github.com/unshade/unraidctl/internal/controllers"
)

// arrayCmd represents the array command
var arrayCmd = &cobra.Command{
	Use:   "array",
	Short: "Interact with unraid array",
	Long:  `Interact with unraid array`,
	Run: func(cmd *cobra.Command, args []string) {
		formater := internal.OutputFormaterSwitcher(internal.OutputFormat(outputFormat))
		controller := controllers.NewArrayController(unraidClient, formater)
		switch args[0] {
		case "stop":
			controller.StopArray(cmd.Context())
		case "start":
			controller.StartArray(cmd.Context())
		case "show":
			controller.ShowArray(cmd.Context())
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(arrayCmd)
	arrayCmd.Args = cobra.MinimumNArgs(1)
	arrayCmd.ValidArgs = []string{"stop", "start", "show"}
	arrayCmd.Flags().StringVarP(&outputFormat, "output", "o", "text", "Output format: json|yaml|text")
}
