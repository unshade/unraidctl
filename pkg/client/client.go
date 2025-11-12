package client

import "github.com/machinebox/graphql"

const (
	ApiKeyHeader = "x-api-key"
)

type UnraidClient struct {
	Docker DockerClient
	Array  ArrayClient
}

func NewUnraidClient(apiKey string, graphqlClient *graphql.Client) *UnraidClient {
	return &UnraidClient{
		Docker: NewDockerClient(apiKey, graphqlClient),
		Array:  NewArrayClient(apiKey, graphqlClient),
	}
}
