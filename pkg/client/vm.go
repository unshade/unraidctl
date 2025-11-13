package client

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/unshade/unraidctl/internal/models"
)

type VMClient interface {
	ListVMs(ctx context.Context) (*models.ListVMsModel, error)
	Start(ctx context.Context, id string) (*models.StartVMModel, error)
	Stop(ctx context.Context, id string) (*models.StopVMModel, error)
}

type RealVMClient struct {
	ApiKey        string
	GraphQLClient *graphql.Client
}

var _ VMClient = (*RealVMClient)(nil)

func NewVMClient(apiKey string, graphqlClient *graphql.Client) VMClient {
	return &RealVMClient{
		ApiKey:        apiKey,
		GraphQLClient: graphqlClient,
	}
}

func (c *RealVMClient) ListVMs(ctx context.Context) (*models.ListVMsModel, error) {
	qglQuery := `
	query Query {
		vms {
			id
			domains {
				id
				name
				state
			}
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)

	return query[models.ListVMsModel](ctx, c.GraphQLClient, req)
}

func (c *RealVMClient) Start(ctx context.Context, id string) (*models.StartVMModel, error) {
	qglQuery := `
	mutation Start($startId: PrefixedID!) {
		vm {
			start(id: $startId)
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)
	req.Var("startId", id)

	return query[models.StartVMModel](ctx, c.GraphQLClient, req)
}

func (c *RealVMClient) Stop(ctx context.Context, id string) (*models.StopVMModel, error) {
	qglQuery := `
	mutation Stop($stopId: PrefixedID!) {
		vm {
			stop(id: $stopId)
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)
	req.Var("stopId", id)

	return ignoreNotFound(query[models.StopVMModel](ctx, c.GraphQLClient, req))
}
