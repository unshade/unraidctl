[![Go Report Card](https://goreportcard.com/badge/github.com/unshade/unraidctl)](https://goreportcard.com/report/github.com/unshade/unraidctl)
[![License](https://img.shields.io/github/license/unshade/unraidctl)](LICENSE)
[![Release](https://img.shields.io/github/release/unshade/unraidctl.svg)](https://github.com/unshade/unraidctl/releases)

# unraidctl

`unraidctl` is a CLI tool to interact with unraid servers. It allows you to manage various aspects of your unraid server from the command line using the unraid GraphQL API.

It stands for "unraid control" and is inspired by tools like `kubectl` for Kubernetes.

# Usage

To use `unraidctl`, you need either need Go installed or you can download a pre-built binary from the [releases](https://github.com/unshade/unraidctl/releases).
You can build it from source using:

```bash
go install github.com/unshade/unraidctl@latest
```

You will then need to set the following environment variables to connect to your unraid server:

- `API_KEY`: Your unraid API key.
- `API_BASE_URL`: The base URL of your unraid server's GraphQL endpoint.
- `API_SKIP_TLS_VERIFY`: (optional) Set to `true` to skip TLS verification (useful for self-signed certificates).

`unraidctl` also supports a `.env` file for setting these environment variables. You can create a `.env` file in the same directory where you run `unraidctl` with the following format:

```env
API_KEY=your_api_key
API_BASE_URL=https://your-unraid-server:port/graphql
API_SKIP_TLS_VERIFY=true
```

Once installed, you can run commands like:

```bash
unraidctl docker list
```

# Library

`unraidctl` also provides a Go library that can be imported into your own Go projects to interact with unraid servers programmatically. Check out the [pkg](./pkg) directory for more details.

# Local Development

Requirements:

- Go 1.25+

Clone the repository:

```bash
git clone git@github.com:unshade/unraidctl.git
cd unraidctl
```

Set up your environment variables. You can copy the example file and modify it:

```bash
cp .env.example .env
```

You can find the required values in your unraid server settings. Please refer to this documentation if needed: [unraid API Documentation](https://docs.unraid.net/API/).

When generating an API key, make sure it has the necessary permissions for the actions you want to perform with `unraidctl`. Please follow the least privilege principle and only grant the permissions that are absolutely necessary.

# Contributing

If you'd like to contribute to `unraidctl`, please fork the repository and create a pull request with your changes. Contributions are welcome! Please make sure to follow the existing code style and include tests for any new functionality.

Pipeline and code quality checks are automatically run on each pull request to ensure code quality. They MUST pass before your changes can be merged.
