package client

import "github.com/machinebox/graphql"

const (
	ApiKeyHeader = "x-api-key"
)

// UnraidClient is a client wrapper.
// It registers every sub-clients into one unified structure
type UnraidClient struct {
	Docker DockerClient
	Array  ArrayClient
	VM     VMClient
}

type UnraidClientOption func(*UnraidClient)

// NewUnraidClient will return an UnraidClient initialized with standard sub-clients
// You can change initialized sub-clients with your own implementations using functionnal options
func NewUnraidClient(apiKey string, graphqlClient *graphql.Client, options ...UnraidClientOption) *UnraidClient {
	defaultClient := &UnraidClient{
		Docker: NewDockerClient(apiKey, graphqlClient),
		Array:  NewArrayClient(apiKey, graphqlClient),
		VM:     NewVMClient(apiKey, graphqlClient),
	}

	for _, option := range options {
		option(defaultClient)
	}

	return defaultClient
}

// WithDockerClient allow you to custom UnraidClient with your own implementation of DockerClient
func WithDockerClient(dockerClient DockerClient) UnraidClientOption {
	return func(uc *UnraidClient) {
		uc.Docker = dockerClient
	}
}

// WithDockerClient allow you to custom UnraidClient with your own implementation of ArrayClient
func WithArrayClient(arrayClient ArrayClient) UnraidClientOption {
	return func(uc *UnraidClient) {
		uc.Array = arrayClient
	}
}

// WithDockerClient allow you to custom UnraidClient with your own implementation of VMClient
func WithVmClient(vmClient VMClient) UnraidClientOption {
	return func(uc *UnraidClient) {
		uc.VM = vmClient
	}
}
