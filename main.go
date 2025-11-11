package main

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/machinebox/graphql"
	"github.com/unshade/unraidctl/cmd"
	"github.com/unshade/unraidctl/internal"
)

func main() {
	config, err := internal.GetConfig()
	if err != nil {
		os.Exit(1)
	}

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.Api.SkipTlsVerify,
		},
	}

	httpClient := &http.Client{
		Transport: customTransport,
	}

	graphqlClient := graphql.NewClient(config.Api.BaseUrl, graphql.WithHTTPClient(httpClient))

	cmd.Execute(config, graphqlClient)
}
