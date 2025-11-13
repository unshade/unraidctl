package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unshade/unraidctl/internal/controllers"
)

// vmCmd represents the vm command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Interact with unraid virtual machines",
	Long:  `Interact with unraid virtual machines. This command allows you to manage your VMs, including starting, stopping, and listing them`,
	Run: func(cmd *cobra.Command, args []string) {
		controller := controllers.NewVmController(unraidClient)
		switch args[0] {
		case "list":
			controller.ListVMs(cmd.Context())
		case "start":
			controller.Start(cmd.Context(), args[1])
		case "stop":
			controller.Stop(cmd.Context(), args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(vmCmd)

	dockerCmd.Args = cobra.MinimumNArgs(1)
	dockerCmd.ValidArgs = []string{"force-stop", "pause", "reboot", "reset", "resume", "start", "stop"}
}
