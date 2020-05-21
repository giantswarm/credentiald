// Package searcher provides the endpoint to retrieve one credential.
package searcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/giantswarm/credentiald/server/middleware"
	"github.com/giantswarm/credentiald/service/searcher"
)

const (
	// Method is the HTTP method to act on.
	Method = "GET"
	// Name is the name of our endpoint.
	Name = "searcher"
	// Path is the URI path our endpoint listens to.
	Path = "/v4/organizations/{organization}/credentials/{id}/"
)

// Config defines which configuration our endpoint expects.
type Config struct {
	Logger     micrologger.Logger
	Middleware *middleware.Middleware // nolint: structcheck, unused
	Service    *searcher.Service
}

// Endpoint is the actual endpoint data structure.
type Endpoint struct {
	logger     micrologger.Logger
	middleware *middleware.Middleware // nolint: structcheck, unused
	service    *searcher.Service
}

// New creates a new endpoint with configuration.
func New(config Config) (*Endpoint, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Service == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Service must not be empty", config)
	}

	endpoint := &Endpoint{
		logger:  config.Logger,
		service: config.Service,
	}

	return endpoint, nil
}

// Decoder gets all useful information from a request to the endpoint.
func (e *Endpoint) Decoder() kithttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)

		request := Request{
			Organization: vars["organization"],
			ID:           vars["id"],
		}

		return request, nil
	}
}

// Encoder creates a response in the specified format.
func (e *Endpoint) Encoder() kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		endpointResponse := response.(Response)

		w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusOK)

		return json.NewEncoder(w).Encode(endpointResponse)
	}
}

// Endpoint is where requests are handled and responses are returned.
func (e *Endpoint) Endpoint() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		e.logger.Log("level", "debug", "message", fmt.Sprintf("received endpoint request: %#v", request))

		endpointRequest := request.(Request)

		serviceRequest := searcher.Request{
			Organization: endpointRequest.Organization,
			ID:           endpointRequest.ID,
		}

		searcherResponse, err := e.service.Search(serviceRequest)
		if err != nil {
			if searcher.IsWrongOwnerOrganizationError(err) {
				return nil, microerror.Mask(wrongOwnerOrganizationError)
			} else if searcher.IsSecretNotFoundError(err) {
				return nil, microerror.Mask(secretNotFoundError)
			}
			return nil, microerror.Mask(err)
		}
		e.logger.Log("level", "debug", "message", "received searcher response successfully")

		endpointResponse := Response{
			ID:       searcherResponse.ID,
			Provider: searcherResponse.Provider,
		}

		if searcherResponse.Provider == "aws" {
			endpointResponse.AWS = &ResponseAWS{
				&ResponseAWSRoles{
					Admin:       searcherResponse.AWS.Roles.Admin,
					AWSOperator: searcherResponse.AWS.Roles.AWSOperator,
				},
			}
		} else if searcherResponse.Provider == "azure" {
			endpointResponse.Azure = &ResponseAzure{
				Credential: &ResponseAzureCredential{
					SubscriptionID: searcherResponse.Azure.Credential.SubscriptionID,
					TenantID:       searcherResponse.Azure.Credential.TenantID,
					ClientID:       searcherResponse.Azure.Credential.ClientID,
				},
			}
		}

		return endpointResponse, nil
	}
}

// Method returns the HTTP method supported by this endpoint.
func (e *Endpoint) Method() string {
	return Method
}

// Middlewares returns the middlewares used by this endpoint.
func (e *Endpoint) Middlewares() []kitendpoint.Middleware {
	return []kitendpoint.Middleware{}
}

// Name returns this endpoint's name.
func (e *Endpoint) Name() string {
	return Name
}

// Path returns this endpoint's URI path.
func (e *Endpoint) Path() string {
	return Path
}
