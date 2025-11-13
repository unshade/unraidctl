package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/unshade/unraidctl/pkg/client"
)

type VmController struct {
	unraidClient *client.UnraidClient
}

func NewVmController(unraidClient *client.UnraidClient) *VmController {
	return &VmController{
		unraidClient: unraidClient,
	}
}

func (c *VmController) ListVMs(ctx context.Context) {
	respData, err := c.unraidClient.VM.ListVMs(ctx)
	if err != nil {
		fmt.Printf("Error listing VMs: %v\n", err)
		return
	}
	for _, vm := range respData.VMs.Domains {
		idParts := strings.Split(vm.Id, ":")
		compactID := idParts[len(idParts)-1]

		fmt.Printf("ID: %s | Name: %s | State: %s\n", compactID, vm.Name, vm.State)
	}
}

func (c *VmController) Start(ctx context.Context, vmId string) {
	_, err := c.unraidClient.VM.Start(ctx, vmId)
	if err != nil {
		fmt.Printf("Could not start vm: %v", err)
	}
} 

func (c *VmController) Stop(ctx context.Context, vmId string) {
	_, err := c.unraidClient.VM.Start(ctx, vmId)
	if err != nil {
		fmt.Printf("Could not start vm: %v", err)
	}
} 
