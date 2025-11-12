package client

import (
	"context"

	"github.com/machinebox/graphql"
)

type VMClient interface {
	ListVMs(ctx context.Context) (*ListVMsModel, error)
	Start(ctx context.Context, id string) (*StartVMModel, error)
	Stop(ctx context.Context, id string) (*StopVMModel, error)
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

type ListVMsModel struct {
	VMs struct {
		ID      string `json:"id"`
		Domains []struct {
			Name  string `json:"name"`
			State string `json:"state"`
			Id    string `json:"id"`
		} `json:"domains"`
	} `json:"vms"`
}

func (c *RealVMClient) ListVMs(ctx context.Context) (*ListVMsModel, error) {
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

	return query[ListVMsModel](ctx, c.GraphQLClient, req)
}

type StartVMModel struct {
	Id string `json:"id"`
}

func (c *RealVMClient) Start(ctx context.Context, id string) (*StartVMModel, error) {
	qglQuery := `
	mutation Start($startId: PrefixedID!) {
		vm {
			start(id: $startId)
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)
	req.Var("startId", id)

	return query[StartVMModel](ctx, c.GraphQLClient, req)
}

type StopVMModel struct {
	Id string `json:"id"`
}

func (c *RealVMClient) Stop(ctx context.Context, id string) (*StopVMModel, error) {
	qglQuery := `
	mutation Stop($stopId: PrefixedID!) {
		vm {
			stop(id: $stopId)
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)
	req.Var("stopId", id)

	return ignoreNotFound(query[StopVMModel](ctx, c.GraphQLClient, req))
}
