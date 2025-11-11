package client

import (
	"context"
	"fmt"
	"log"

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
