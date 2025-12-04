package client

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/unshade/unraidctl/internal/models"
)

type ShareClient interface {
	ListShares(ctx context.Context) (*models.ListSharesModel, error)
}

type RealShareClient struct {
	ApiKey        string
	GraphQLClient *graphql.Client
}

var _ ShareClient = (*RealShareClient)(nil)

func NewShareClient(apiKey string, graphqlClient *graphql.Client) ShareClient {
	return &RealShareClient{
		ApiKey:        apiKey,
		GraphQLClient: graphqlClient,
	}
}

func (c *RealShareClient) ListShares(ctx context.Context) (*models.ListSharesModel, error) {
	qglQuery := `
	query Query {
		shares {
			name
			free
			used
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)

	return query[models.ListSharesModel](ctx, c.GraphQLClient, req)
}
