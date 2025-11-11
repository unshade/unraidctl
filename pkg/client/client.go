package client

import "github.com/machinebox/graphql"

const (
	ApiKeyHeader = "x-api-key"
)

type UnraidClient struct {
	Docker DockerClient
}

func NewUnraidClient(apiKey string, graphqlClient *graphql.Client) *UnraidClient {
	return &UnraidClient{
		Docker: NewDockerClient(apiKey, graphqlClient),
	}
}
