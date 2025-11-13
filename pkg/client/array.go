package client

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/unshade/unraidctl/internal/models"
)

// ArrayClient is used to interract with unraid array system
type ArrayClient interface {
	// ShowArray will return the current state of the unraid array
	ShowArray(ctx context.Context) (*models.ShowArrayModel, error)
	// MutateArray allows to change make a transition to a new state : STARTED -> STOPPED
	MutateArray(ctx context.Context, state ArrayState) (*models.MutateArrayModel, error)
}

// RealArrayClient is a basic implementation for an ArrayClient
type RealArrayClient struct {
	ApiKey        string
	GraphQLClient *graphql.Client
}

var _ ArrayClient = (*RealArrayClient)(nil)

func NewArrayClient(apiKey string, graphqlClient *graphql.Client) ArrayClient {
	return &RealArrayClient{
		ApiKey:        apiKey,
		GraphQLClient: graphqlClient,
	}
}

func (c *RealArrayClient) ShowArray(ctx context.Context) (*models.ShowArrayModel, error) {
	qglQuery := `
	query Query {
		array {
			state
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)

	return query[models.ShowArrayModel](ctx, c.GraphQLClient, req)
}

type ArrayState = string
const (
	ArrayStateStart ArrayState = "START"
	ArrayStateStop  ArrayState = "STOP"
)

func (c *RealArrayClient) MutateArray(ctx context.Context, state ArrayState) (*models.MutateArrayModel, error) {
	qglQuery := `
	mutation Array($input: ArrayStateInput!) {
		array {
			setState(input: $input) {
				id
			}
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)

	req.Var("input", map[string]ArrayState{
		"desiredState": state,
	})

	return query[models.MutateArrayModel](ctx, c.GraphQLClient, req)
}
