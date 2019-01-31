package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giantswarm/credentiald/service/lister"
	"github.com/giantswarm/microerror"
)

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
