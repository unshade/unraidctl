package controllers

import (
	"context"
	"fmt"

	"github.com/unshade/unraidctl/internal"
	"github.com/unshade/unraidctl/pkg/client"
)

type ArrayController struct {
	unraidClient *client.UnraidClient
	formater     internal.OutputFormater
}

func NewArrayController(unraidClient *client.UnraidClient, formater internal.OutputFormater) *ArrayController {
	return &ArrayController{
		unraidClient: unraidClient,
		formater:     formater,
	}
}

func (c *ArrayController) ShowArray(ctx context.Context) {
	response, err := c.unraidClient.Array.ShowArray(ctx)
	if err != nil {
		fmt.Printf("Error showing array: %v\n", err)
		return
	}
	internal.PrintFormat(c.formater.Format(response))
}

func (c *ArrayController) StopArray(ctx context.Context) {
	response, err := c.unraidClient.Array.MutateArray(ctx, client.ArrayStateStop)
	if err != nil {
		fmt.Printf("Error stopping array: %v\n", err)
		return
	}
	internal.PrintFormat(c.formater.Format(response))
}

func (c *ArrayController) StartArray(ctx context.Context) {
	response, err := c.unraidClient.Array.MutateArray(ctx, client.ArrayStateStart)
	if err != nil {
		fmt.Printf("Error starting array: %v\n", err)
		return
	}
	internal.PrintFormat(c.formater.Format(response))
}
