package controllers

import (
	"context"
	"fmt"

	"github.com/unshade/unraidctl/internal"
	"github.com/unshade/unraidctl/pkg/client"
)

type DockerController struct {
	unraidClient *client.UnraidClient
	formater     internal.OutputFormater
}

func NewDockerController(unraidClient *client.UnraidClient, formater internal.OutputFormater) *DockerController {
	return &DockerController{
		unraidClient: unraidClient,
		formater:     formater,
	}
}

func (c *DockerController) ListContainers(ctx context.Context) {
	response, err := c.unraidClient.Docker.ListContainers(ctx)
	if err != nil {
		fmt.Printf("Error listing containers: %v\n", err)
		return
	}
	internal.PrintFormat(c.formater.Format(response))
}

func (c *DockerController) StartContainer(ctx context.Context, containerid string) {
	response, err := c.unraidClient.Docker.StartContainer(ctx, containerid)
	if err != nil {
		fmt.Printf("Could not start container: %v", err)
	}
	internal.PrintFormat(c.formater.Format(response))
}

func (c *DockerController) StopContainer(ctx context.Context, containerid string) {
	response, err := c.unraidClient.Docker.StopContainer(ctx, containerid)
	if err != nil {
		fmt.Printf("Could not stop container: %v", err)
	}
	internal.PrintFormat(c.formater.Format(response))
}
