package client

import (
	"context"
	"strings"

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

func ignoreNotFound[T any](response *T, err error) (*T, error) {
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return response, nil
		}
		return response, err
	}
	return response, nil
}
