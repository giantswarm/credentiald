package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/giantswarm/microclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"
)

const (
	// API endpoints of the service this client action interacts with.
	CredentialEndpoint        = "/v4/organizations/%s/credentials/"
	CredentialDetailsEndpoint = "/v4/organizations/%s/credentials/%s/"

	// Name is the service name being implemented.
	Name = "credentials"
)

// Config represents the configuration used to create a credentiald Client.
type Config struct {
	Logger     micrologger.Logger
	RestClient *resty.Client

	Address string
}

type Client struct {
	logger     micrologger.Logger
	restClient *resty.Client

	url *url.URL
}

// New creates a new configured credentiald Client.
func New(config Config) (*Client, error) {
	// Dependencies.
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.RestClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.RestClient must not be empty", config)
	}

	u, err := url.Parse(config.Address)
	if err != nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Address %v", config, err)
	}

	client := &Client{
		logger:     config.Logger,
		restClient: config.RestClient,

		url: u,
	}

	return client, nil
}

func (c Client) do(ctx context.Context, requestFunc func(string) (*resty.Response, error), url string) (*resty.Response, error) {
	c.logger.Log("debug", fmt.Sprintf("sending request to %#q", url), "service", Name)

	response, err := microclient.Do(ctx, requestFunc, url)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	c.logger.Log("debug", fmt.Sprintf("received status code %d", response.StatusCode()), "service", Name)

	return response, nil
}
