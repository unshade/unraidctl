package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unshade/unraidctl/internal/controllers"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Interact with unraid Docker engine",
	Long:  `Interact with unraid Docker engine`,
	Run: func(cmd *cobra.Command, args []string) {
		controller := controllers.NewDockerController(unraidClient)
		switch args[0] {
		case "list":
			controller.ListContainers(cmd.Context())
		case "stop":
			controller.StopContainer(cmd.Context(), args[1])
		case "start":
			controller.StartContainer(cmd.Context(), args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)

	dockerCmd.Args = cobra.MinimumNArgs(1)
	dockerCmd.ValidArgs = []string{"stop", "restart", "start", "list"}
}
