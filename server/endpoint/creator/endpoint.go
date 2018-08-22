package creator

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
	"github.com/giantswarm/credentiald/service/creator"
	"github.com/giantswarm/credentiald/service/creator/aws"
	"github.com/giantswarm/credentiald/service/creator/azure"
)

const (
	Method = "POST"
	Name   = "creator"
	Path   = "/v4/organizations/{organization}/credentials/"
)

type Config struct {
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *creator.Service
}

type Endpoint struct {
	logger     micrologger.Logger
	middleware *middleware.Middleware
	service    *creator.Service
}

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

func (e *Endpoint) Decoder() kithttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		organization := vars["organization"]

		var request Request

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		request.Organization = organization

		return request, nil
	}
}

func (e *Endpoint) Encoder() kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		endpointResponse := response.(Response)

		w.Header().Set(
			"Location",
			fmt.Sprintf("/v4/organizations/%s/credentials/%s/",
				endpointResponse.Organization,
				endpointResponse.CredentialID,
			),
		)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		return json.NewEncoder(w).Encode(endpointResponse)
	}
}

func (e *Endpoint) Endpoint() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		endpointRequest := request.(Request)
		e.logger.Log("level", "debug", "message", fmt.Sprintf("received endpoint request: %#v", endpointRequest))

		creatorRequest := creator.Request{
			Organization: endpointRequest.Organization,
			Provider:     endpointRequest.Provider,

			AWS: aws.Request{
				AdminARN:       endpointRequest.AWS.Roles.Admin,
				AwsOperatorARN: endpointRequest.AWS.Roles.AwsOperator,
			},
			Azure: azure.Request{
				ClientID:       endpointRequest.Azure.Credential.ClientID,
				SecretID:       endpointRequest.Azure.Credential.SecretID,
				SubscriptionID: endpointRequest.Azure.Credential.SubscriptionID,
				TenantID:       endpointRequest.Azure.Credential.TenantID,
			},
		}

		creatorResponse, err := e.service.Create(ctx, creatorRequest)
		if creator.IsAlreadyExists(err) {
			return nil, microerror.Maskf(alreadyExistsError, err.Error())
		} else if err != nil {
			return nil, microerror.Mask(err)
		}
		e.logger.Log("level", "debug", "message", fmt.Sprintf("received service response: %#v", creatorResponse))

		endpointResponse := Response{
			Code:    creatorResponse.Code,
			Message: creatorResponse.Message,

			CredentialID: creatorResponse.CredentialID,
			Organization: creatorResponse.Organization,
		}

		return endpointResponse, nil
	}
}

func (e *Endpoint) Method() string {
	return Method
}

func (e *Endpoint) Middlewares() []kitendpoint.Middleware {
	return []kitendpoint.Middleware{}
}

func (e *Endpoint) Name() string {
	return Name
}

func (e *Endpoint) Path() string {
	return Path
}
