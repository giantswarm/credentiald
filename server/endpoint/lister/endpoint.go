// Package lister provides the endpoint to retrieve multiple credentials
package lister

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
	"github.com/giantswarm/credentiald/service/lister"
)

const (
	// Method is the HTTP method to act on.
	Method = "GET"
	// Name is the name of our endpoint.
	Name = "lister"
	// Path is the URI path our endpoint listens to.
	Path = "/v4/organizations/{organization}/credentials/"
)

// Config defines which configuration our endpoint expects.
type Config struct {
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *lister.Service
}

// Endpoint is the actual endpoint data structure.
type Endpoint struct {
	logger     micrologger.Logger
	middleware *middleware.Middleware
	service    *lister.Service
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

		var request Request
		request.Organization = vars["organization"]

		return request, nil
	}
}

// Encoder creates a response in the specified format.
func (e *Endpoint) Encoder() kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		endpointResponse := response.([]*Response)

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

		serviceRequest := lister.Request{Organization: endpointRequest.Organization}
		listerResponse, err := e.service.List(serviceRequest)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		e.logger.Log("level", "debug", "message", fmt.Sprintf("received service response: %#v", listerResponse))

		endpointResponse := []*Response{}

		for _, credential := range listerResponse {
			responseItem := &Response{
				ID:       credential.ID,
				Provider: credential.Provider,
			}

			if credential.Provider == "aws" {
				responseItem.AWS = &ResponseAWS{
					&ResponseAWSRoles{
						Admin:       credential.AWS.Roles.Admin,
						AWSOperator: credential.AWS.Roles.AWSOperator,
					},
				}
			} else if credential.Provider == "azure" {
				responseItem.Azure = &ResponseAzure{
					SubscriptionID: credential.Azure.SubscriptionID,
					TenantID:       credential.Azure.TenantID,
					ClientID:       credential.Azure.ClientID,
				}
			}

			endpointResponse = append(endpointResponse, responseItem)
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
