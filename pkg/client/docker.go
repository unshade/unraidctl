package client

import (
	"context"

	"github.com/machinebox/graphql"
)

type DockerClient interface {
	ListContainers(ctx context.Context) (*ListContainersModel, error)
	StartContainer(ctx context.Context, startId string) (*StartContainerModel, error)
	StopContainer(ctx context.Context, stopId string) (*StopContainerModel, error)
}

type RealDockerClient struct {
	ApiKey        string
	GraphQLClient *graphql.Client
}

var _ DockerClient = (*RealDockerClient)(nil)

func NewDockerClient(apiKey string, graphqlClient *graphql.Client) DockerClient {
	return &RealDockerClient{
		ApiKey:        apiKey,
		GraphQLClient: graphqlClient,
	}
}

type ListContainersModel struct {
	Docker struct {
		Containers []struct {
			ID    string `json:"id"`
			Image string `json:"image"`
			State string `json:"state"`
		} `json:"containers"`
	} `json:"docker"`
}

func (c *RealDockerClient) ListContainers(ctx context.Context) (*ListContainersModel, error) {
	qglQuery := `
	query Query {
		docker {
			containers {
				id
				image
				state
			}
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)

	return query[ListContainersModel](ctx, c.GraphQLClient, req)
}

type StartContainerModel struct {
	Docker struct {
		Start struct {
			ID string `json:"id"`
		} `json:"start"`
	} `json:"docker"`
}

func (c *RealDockerClient) StartContainer(ctx context.Context, startId string) (*StartContainerModel, error) {
	mutation := `
	mutation Mutation($startId: PrefixedID!) {
		docker {
			start(id: $startId) {
				id
			}
		}
	}`

	req := graphql.NewRequest(mutation)
	auth(req, c.ApiKey)

	req.Var("startId", startId)

	return query[StartContainerModel](ctx, c.GraphQLClient, req)
}

type StopContainerModel struct {
	Docker struct {
		Stop struct {
			ID string `json:"id"`
		} `json:"stop"`
	} `json:"docker"`
}

func (c *RealDockerClient) StopContainer(ctx context.Context, stopId string) (*StopContainerModel, error) {
	mutation := `
	mutation Stop($stopId: PrefixedID!) {
		docker {
			stop(id: $stopId) {
				id
			}
		}
	}`

	req := graphql.NewRequest(mutation)
	auth(req, c.ApiKey)

	req.Var("stopId", stopId)

	return query[StopContainerModel](ctx, c.GraphQLClient, req)
}
