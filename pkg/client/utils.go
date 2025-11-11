package client

import (
	"context"

	"github.com/machinebox/graphql"
)

func query[T any](ctx context.Context, client *graphql.Client, req *graphql.Request) (*T, error) {
	var respData T
	if err := client.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func auth(req *graphql.Request, apiKey string) {
	req.Header.Set(ApiKeyHeader, apiKey)
}
