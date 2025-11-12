package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/unshade/unraidctl/pkg/client"
)

type ArrayController struct {
	unraidClient *client.UnraidClient
}

func NewArrayController(unraidClient *client.UnraidClient) *ArrayController {
	return &ArrayController{
		unraidClient: unraidClient,
	}
}

func (c *ArrayController) ShowArray(ctx context.Context) {
	respData, err := c.unraidClient.Array.ShowArray(ctx)
	if err != nil {
		fmt.Printf("Error showing array: %v\n", err)
		return
	}
	fmt.Printf("Array State: %s\n", strings.ToLower(respData.Array.State))
}

func (c *ArrayController) StopArray(ctx context.Context) {
	_, err := c.unraidClient.Array.MutateArray(ctx, client.ArrayStateStop)
	if err != nil {
		fmt.Printf("Error stopping array: %v\n", err)
		return
	}
}

func (c *ArrayController) StartArray(ctx context.Context) {
	_, err := c.unraidClient.Array.MutateArray(ctx, client.ArrayStateStart)
	if err != nil {
		fmt.Printf("Error starting array: %v\n", err)
		return
	}
}
