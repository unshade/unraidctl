package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/unshade/unraidctl/pkg/client"
)

type DockerController struct {
	unraidClient *client.UnraidClient
}

func NewDockerController(unraidClient *client.UnraidClient) *DockerController {
	return &DockerController{
		unraidClient: unraidClient,
	}
}

func (c *DockerController) ListContainers(ctx context.Context) {
	respData, err := c.unraidClient.Docker.ListContainers(ctx)
	if err != nil {
		fmt.Printf("Error listing containers: %v\n", err)
		return
	}
	for _, container := range respData.Docker.Containers {
		idParts := strings.Split(container.ID, ":")
		compactID := idParts[len(idParts)-1]

		if len(compactID) > 12 {
			compactID = compactID[:12]
		}

		fmt.Printf("ID: %s | Image: %s | State: %s\n", compactID, container.Image, container.State)
	}
}
