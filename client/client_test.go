package client

import (
	"context"
	"fmt"

	"github.com/giantswarm/micrologger"
	"gopkg.in/resty.v1"

	"github.com/giantswarm/credentiald/v2/service/creator"
	"github.com/giantswarm/credentiald/v2/service/lister"
	"github.com/giantswarm/credentiald/v2/service/searcher"
)

// ClientExample shows how to create a Client.
// nolint: unused
func ClientExample() Client {
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

// CreateExample shows how to create a set of credentials for an organization.
// nolint: unused
func CreateExample() {
	var (
		organizationID = "my-organization"
		provider       = "azure"
	)

	var c Client
	{
		c = ClientExample()
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

// ListExample shows how to list credentials which belong to an organization.
// nolint: unused
func ListExample() {
	var (
		organizationID = "my-organization"
	)

	var c Client
	{
		c = ClientExample()
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

// SearchExample shows how to get a set of credentials which belong to an organzation.
// nolint: unused
func SearchExample() {
	var (
		organizationID = "my-organization"
		credentialID   = "credential-a1b2c3" //nolint:gosec
	)

	var c Client
	{
		c = ClientExample()
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
