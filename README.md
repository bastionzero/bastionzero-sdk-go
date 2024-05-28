# bastionzero-sdk-go

[![bastionzero-sdk-go release (latest SemVer)](https://img.shields.io/github/v/release/bastionzero/bastionzero-sdk-go?sort=semver)](https://github.com/bastionzero/bastionzero-sdk-go/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/bastionzero/bastionzero-sdk-go)
[![Test Status](https://github.com/bastionzero/bastionzero-sdk-go/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/bastionzero/bastionzero-sdk-go/actions?query=workflow%3Aci%20branch:master)

`bastionzero-sdk-go` is a Go client library for accessing the **BastionZero API
v2**. 

You can view the client API docs here: [https://pkg.go.dev/github.com/bastionzero/bastionzero-sdk-go](http://godoc.org/github.com/bastionzero/bastionzero-sdk-go)

You can view BastionZero API docs here: [https://cloud.bastionzero.com/api/](https://cloud.bastionzero.com/api/)

## Install

```sh
go get github.com/bastionzero/bastionzero-sdk-go@vX.Y.Z
```

where X.Y.Z is the [version](https://github.com/bastionzero/bastionzero-sdk-go/releases) you need.

Alternatively,

```sh
go get github.com/bastionzero/bastionzero-sdk-go
```

to get the latest version.

## Usage

```go
import "github.com/bastionzero/bastionzero-sdk-go"
```

Currently, using an API key is the only method of authentication that is
supported by this library. You can manage your API keys at the BastionZero API
key panel found [here](https://cloud.bastionzero.com/admin/apikeys). See the
[admin
guide](https://docs.bastionzero.com/docs/admin-guide/authorization#creating-an-api-key)
for more information.

After an API key has been created, you can use its secret to create a new
client:

```go
client, err := bastionzero.NewFromAPISecret(http.DefaultClient, "bastionzero-api-secret")
if err != nil {
    return err
}

// list all target connect policies in your organization
policies, _, err := client.Policies.ListTargetConnectPolicies(context.Background(), nil)
```

The services of a [`bastionzero.Client`](https://pkg.go.dev/github.com/bastionzero/bastionzero-sdk-go/bastionzero#Client) divide the API into logical chunks and correspond to
the structure of the BastionZero API documentation.

NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancellation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.

## Examples

To create a new environment:

```go
createRequest := &environments.CreateEnvironmentRequest{
    Name: "my-new-env",
    Description: "An environment created using the BastionZero Go SDK.",
    OfflineCleanupTimeoutHours: uint(24),
}

createResponse, _, err := client.Environments.CreateEnvironment(context.TODO(), createRequest)
if err != nil {
    return err
}

fmt.Printf("Created new environment with ID: %s", createResponse.ID)
```

## Versioning

Each version of the client is tagged and the version is updated accordingly.

In general, `bastionzero-sdk-go` follows [semver](https://semver.org/) as
closely as we can for tagging releases of the package.

## License

This library is distributed under the Apache License 2.0 license found in the
[LICENSE](./LICENSE) file and includes [open source software](https://docs.bastionzero.com/docs/credits/go-sdk)
under a variety of other licenses.