package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/credentiald/v2/service/searcher"
)

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
