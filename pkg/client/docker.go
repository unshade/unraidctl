package client

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/unshade/unraidctl/internal/models"
)

type DockerClient interface {
	ListContainers(ctx context.Context) (*models.ListContainersModel, error)
	StartContainer(ctx context.Context, startId string) (*models.StartContainerModel, error)
	StopContainer(ctx context.Context, stopId string) (*models.StopContainerModel, error)
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

func (c *RealDockerClient) ListContainers(ctx context.Context) (*models.ListContainersModel, error) {
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

	return query[models.ListContainersModel](ctx, c.GraphQLClient, req)
}

func (c *RealDockerClient) StartContainer(ctx context.Context, startId string) (*models.StartContainerModel, error) {
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

	return ignoreNotFound(query[models.StartContainerModel](ctx, c.GraphQLClient, req))
}

func (c *RealDockerClient) StopContainer(ctx context.Context, stopId string) (*models.StopContainerModel, error) {
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

	return ignoreNotFound(query[models.StopContainerModel](ctx, c.GraphQLClient, req))
}
