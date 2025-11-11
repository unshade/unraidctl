package client

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/machinebox/graphql"
)

type DockerClient struct {
	ApiKey        string
	GraphQLClient *graphql.Client
}

func NewDockerClient(apiKey string, graphqlClient *graphql.Client) *DockerClient {
	return &DockerClient{
		ApiKey:        apiKey,
		GraphQLClient: graphqlClient,
	}
}

func (c *DockerClient) ListContainers(ctx context.Context) {
	query := `
	query Query {
		docker {
			containers {
				id
				image
				state
			}
		}
	}`

	req := graphql.NewRequest(query)
	req.Header.Set(ApiKeyHeader, c.ApiKey)

	var respData struct {
		Docker struct {
			Containers []struct {
				ID    string `json:"id"`
				Image string `json:"image"`
				State string `json:"state"`
			} `json:"containers"`
		} `json:"docker"`
	}

	if err := c.GraphQLClient.Run(ctx, req, &respData); err != nil {
		log.Printf("Error query list containers: %v", err)
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

func (c *DockerClient) StartContainer(ctx context.Context, startId string) {
	mutation := `
	mutation Mutation($startId: PrefixedID!) {
		docker {
			start(id: $startId) {
				id
			}
		}
	}`

	req := graphql.NewRequest(mutation)
	req.Header.Set(ApiKeyHeader, c.ApiKey)

	req.Var("startId", startId)

	var respData struct {
		Docker struct {
			Start struct {
				ID string `json:"id"`
			} `json:"start"`
		} `json:"docker"`
	}

	if err := c.GraphQLClient.Run(ctx, req, &respData); err != nil {
		log.Printf("Error mutation start: %v", err)
		return
	}

	fmt.Printf("Container started with ID: %s\n", respData.Docker.Start.ID)
}

func (c *DockerClient) StopContainer(ctx context.Context, stopId string) {
	mutation := `
	mutation Stop($stopId: PrefixedID!) {
		docker {
			stop(id: $stopId) {
				id
			}
		}
	}`

	req := graphql.NewRequest(mutation)
	req.Header.Set(ApiKeyHeader, c.ApiKey)

	req.Var("stopId", stopId)

	var respData struct {
		Docker struct {
			Stop struct {
				ID string `json:"id"`
			} `json:"stop"`
		} `json:"docker"`
	}

	if err := c.GraphQLClient.Run(ctx, req, &respData); err != nil {
		log.Printf("Error mutation stop: %v", err)
	}

	fmt.Printf("Container stopped with ID: %s\n", respData.Docker.Stop.ID)
}
