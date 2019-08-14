package client

import (
	"context"
	"fmt"

	"github.com/giantswarm/credentiald/service/creator"
	"github.com/giantswarm/credentiald/service/lister"
	"github.com/giantswarm/credentiald/service/searcher"
	"github.com/giantswarm/micrologger"
	"gopkg.in/resty.v1"
)

// ExampleClient shows how to create a Client.
func ExampleClient() Client {
	var (
		apiURL   = "http://api.g8s.example"
		apiToken = "my-token"
	)

	var logger micrologger.Logger
	{
		config := micrologger.Config{
			Caller:             micrologger.DefaultCaller,
			IOWriter:           micrologger.DefaultIOWriter,
			TimestampFormatter: micrologger.DefaultTimestampFormatter,
		}

		logger, _ = micrologger.New(config)
	}

	var restyClient *resty.Client
	{
		restyClient = resty.New()
		restyClient.SetAuthToken(apiToken)
	}

	var config = Config{
		Address:    apiURL,
		Logger:     logger,
		RestClient: restyClient,
	}

	var c *Client
	{
		c, _ = New(config)
	}

	return *c
}

// ExampleCreate shows how to create a set of credentials for an organization.
func ExampleCreate() {
	var (
		organizationID = "my-organization"
		provider       = "azure"
	)

	var c Client
	{
		c = ExampleClient()
	}

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var response *creator.Response
	{
		request := creator.Request{
			Organization: organizationID,
			Provider:     provider,
		}

		response, _ = c.Create(ctx, request)
	}

	fmt.Printf("response: \n%#v\n", response)
}

// ExampleList shows how to list credentials which belong to an organization.
func ExampleList() {
	var (
		organizationID = "my-organization"
	)

	var c Client
	{
		c = ExampleClient()
	}

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var response []lister.Response
	{
		request := lister.Request{
			Organization: organizationID,
		}

		response, _ = c.List(ctx, request)
	}

	fmt.Printf("response: \n%#v\n", response)

}

// ExampleSearch shows how to get a set of credentials which belong to an organzation.
func ExampleSearch() {
	var (
		organizationID = "my-organization"
		credentialID   = "credential-a1b2c3"
	)

	var c Client
	{
		c = ExampleClient()
	}

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var response *searcher.Response
	{
		request := searcher.Request{
			Organization: organizationID,
			ID:           credentialID,
		}

		response, _ = c.Search(ctx, request)
	}

	fmt.Printf("response: \n%#v\n", response)
}
