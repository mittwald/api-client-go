# mittwald API Golang client

This package contains a (mostly auto-generated) Golang client for the mittwald mStudio API.

## License

Copyright (c) 2025 Mittwald CM Service GmbH & Co. KG and contributors

This project is licensed under the MIT License; see the [LICENSE](./LICENSE) file for details.

## Installation

You can install this package using `go get`:

```bash
$ go get github.com/mittwald/api-client-go
```

## API versions

The API client is designed with the support for multiple API versions in mind (although currently, the only available API version is v2).

To use a specific API version, use `github.com/mittwald/api-client-go/mittwaldvX` as import path, with `vX` being the API version you want to use (for example, `v2`):

```go
import "github.com/mittwald/api-client-go/mittwaldv2"
```

## Usage

Use the `mittwaldv2.New` function to instantiate a new client:

```go
ctx := context.Background()
token := os.Getenv("MITTWALD_API_TOKEN")

client, err := mittwaldv2.New(ctx, mittwaldv2.WithAccessToken(token))
if err != nil {
	panic(err)
}
```

You can use different options to configure different kinds of authentication:

1. Omit authentication option for unauthenticated access
2. `mittwaldv2.WithAccessToken` (recommended) to authenticate with an API token
3. `mittwaldv2.WithAccessTokenFromEnv` (recommended) to authenticate with an API token that is automatically retrieved from the process environment variables (`MITTWALD_API_TOKEN`).
3. `mittwaldv2.WithUsernamePassword` to authenticate with username and password; this does not work for users that have 2FA enabled
4. `mittwaldv2.WithAccessTokenRetrievalKey` to authenticate with an [access token retrieval key][atrek] (only relevant for mStudio extensions)
5. `mittwaldv2.WithExtensionSecret` to authenticate as an extension (only relevant for mStudio extensions)

Other options include:

1. `mittwaldv2.WithBaseURL` can be used to override the API base URL (useful for testing purposes).
1. `mittwaldv2.WithEventualConsistency` enables client-side support for the [eventual consistency behaviour][consistency].

Have a look at our [API introduction][api-getting-started] for more information
on how to obtain an API token and how to get started with the API.

## Example

```go
req := project.ListProjectsRequest{
	CustomerId: pointer.To("2ef23459-beb1-4ac2-9a38-d3a9df62bf93"),
	Limit: pointer.To(100),
}

projects, res, err := client.Projects().ListProjects(ctx, req)
for _, project := range projects {
	fmt.Println(project.ShortId)
}
```

## API documentation

The API documentation can be found in our [Developer Portal][api-ref]. You can find the operation ID on the right side of each operation.

[api-getting-started]: https://developer.mittwald.de/docs/v2/api/intro
[api-ref]: https://developer.mittwald.de/reference/v2/
[atrek]: https://developer.mittwald.de/docs/v2/contribution/overview/concepts/authentication/
[consistency]: https://developer.mittwald.de/docs/v2/api/intro/#eventual-consistency
