package controllers

import (
	"context"
	"fmt"

	"github.com/unshade/unraidctl/internal"
	"github.com/unshade/unraidctl/pkg/client"
)

type VmController struct {
	unraidClient *client.UnraidClient
	formater     internal.OutputFormater
}

func NewVmController(unraidClient *client.UnraidClient, formater internal.OutputFormater) *VmController {
	return &VmController{
		unraidClient: unraidClient,
		formater:     formater,
	}
}

func (c *VmController) ListVMs(ctx context.Context) {
	response, err := c.unraidClient.VM.ListVMs(ctx)
	if err != nil {
		fmt.Printf("Error listing VMs: %v\n", err)
		return
	}
	internal.PrintFormat(c.formater.Format(response))
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
