package client

import (
	"context"

	"github.com/machinebox/graphql"
)

type ArrayClient interface {
	ShowArray(ctx context.Context) (*ShowArrayModel, error)
	MutateArray(ctx context.Context, state ArrayState) (*MutateArrayModel, error)
}

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

type ShowArrayModel struct {
	Array struct {
		State string `json:"state"`
	} `json:"array"`
}

func (c *RealArrayClient) ShowArray(ctx context.Context) (*ShowArrayModel, error) {
	qglQuery := `
	query Query {
		array {
			state
		}
	}`

	req := graphql.NewRequest(qglQuery)
	auth(req, c.ApiKey)

	return query[ShowArrayModel](ctx, c.GraphQLClient, req)
}

type MutateArrayModel struct {
	Id string `json:"id"`
}

type ArrayState = string

const (
	ArrayStateStart ArrayState = "START"
	ArrayStateStop  ArrayState = "STOP"
)

func (c *RealArrayClient) MutateArray(ctx context.Context, state ArrayState) (*MutateArrayModel, error) {
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

	return query[MutateArrayModel](ctx, c.GraphQLClient, req)
}
