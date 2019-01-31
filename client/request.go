package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/giantswarm/credentiald/service/creator"
	"github.com/giantswarm/credentiald/service/lister"
	"github.com/giantswarm/credentiald/service/searcher"
	"github.com/giantswarm/microerror"
)

func (c Client) Create(ctx context.Context, request creator.Request) (*creator.Response, error) {
	var response creator.Response
	{
		req := c.restClient.R()
		req.SetBody(request)

		u, err := c.url.Parse(fmt.Sprintf(CredentialEndpoint, request.Organization))
		if err != nil {
			return nil, microerror.Mask(err)
		}

		resp, err := c.do(ctx, req.Post, u.String())
		if err != nil {
			return nil, microerror.Mask(err)
		}

		if resp.StatusCode() != http.StatusCreated {
			return nil, microerror.Maskf(invalidRequestError, string(resp.Body()))
		}

		err = json.Unmarshal(resp.Body(), &response)
		if err != nil {
			return nil, microerror.Maskf(invalidResponseError, string(resp.Body()))
		}

		locationHeader := resp.Header().Get("Location")
		credentialID := path.Base(locationHeader)
		response.CredentialID = credentialID
	}

	return &response, nil
}

func (c Client) List(ctx context.Context, request lister.Request) ([]lister.Response, error) {
	var response []lister.Response
	{
		req := c.restClient.R()

		u, err := c.url.Parse(fmt.Sprintf(CredentialEndpoint, request.Organization))
		if err != nil {
			return nil, microerror.Mask(err)
		}

		resp, err := c.do(ctx, req.Get, u.String())
		if err != nil {
			return nil, microerror.Mask(err)
		}

		if resp.StatusCode() != http.StatusOK {
			return nil, microerror.Maskf(invalidRequestError, string(resp.Body()))
		}

		err = json.Unmarshal(resp.Body(), &response)
		if err != nil {
			return nil, microerror.Maskf(invalidResponseError, string(resp.Body()))
		}
	}
	return response, nil
}

func (c Client) Search(ctx context.Context, request searcher.Request) (*searcher.Response, error) {
	var response searcher.Response
	{
		req := c.restClient.R()

		u, err := c.url.Parse(fmt.Sprintf(CredentialDetailsEndpoint, request.Organization, request.ID))
		if err != nil {
			return nil, microerror.Mask(err)
		}

		resp, err := c.do(ctx, req.Get, u.String())
		if err != nil {
			return nil, microerror.Mask(err)
		}

		if resp.StatusCode() != http.StatusOK {
			return nil, microerror.Maskf(invalidRequestError, string(resp.Body()))
		}

		err = json.Unmarshal(resp.Body(), &response)
		if err != nil {
			return nil, microerror.Maskf(invalidResponseError, string(resp.Body()))
		}
	}
	return &response, nil

}
