package controllers

import (
	"context"
	"fmt"

	"github.com/unshade/unraidctl/internal"
	"github.com/unshade/unraidctl/pkg/client"
)

type ShareController struct {
	unraidClient *client.UnraidClient
	formater     internal.OutputFormater
}

func NewShareController(unraidClient *client.UnraidClient, formater internal.OutputFormater) *ShareController {
	return &ShareController{
		unraidClient: unraidClient,
		formater:     formater,
	}
}

func (c *ShareController) ListShares(ctx context.Context) {
	response, err := c.unraidClient.Shares.ListShares(ctx)
	if err != nil {
		fmt.Printf("Error listing shares: %v\n", err)
		return
	}
	internal.PrintFormat(c.formater.Format(response))
}
