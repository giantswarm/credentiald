package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/credentiald/service/creator"
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
