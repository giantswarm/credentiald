package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/giantswarm/microerror"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"

	"github.com/giantswarm/credentiald/server/endpoint"
	"github.com/giantswarm/credentiald/server/middleware"
	"github.com/giantswarm/credentiald/service"
	"github.com/giantswarm/credentiald/service/creator"
)

type Config struct {
	Logger  micrologger.Logger
	Service *service.Service
	Viper   *viper.Viper

	ProjectName string
}

type Server struct {
	logger micrologger.Logger

	bootOnce     sync.Once
	config       microserver.Config
	shutdownOnce sync.Once
}

func New(config Config) (*Server, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Service == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Service must not be empty", config)
	}
	if config.Viper == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Viper must not be empty", config)
	}

	if config.ProjectName == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ProjectName must not be empty", config)
	}

	var err error

	var middlewareCollection *middleware.Middleware
	{
		c := middleware.Config{
			Logger:  config.Logger,
			Service: config.Service,
		}

		middlewareCollection, err = middleware.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var endpointCollection *endpoint.Endpoint
	{
		c := endpoint.Config{
			Logger:     config.Logger,
			Middleware: middlewareCollection,
			Service:    config.Service,
		}

		endpointCollection, err = endpoint.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Server{
		logger:   config.Logger,
		bootOnce: sync.Once{},
		config: microserver.Config{
			Logger:      config.Logger,
			ServiceName: config.ProjectName,
			Viper:       config.Viper,
			Endpoints: []microserver.Endpoint{
				endpointCollection.Creator,
				endpointCollection.Version,
			},
			ErrorEncoder: errorEncoder,
		},
		shutdownOnce: sync.Once{},
	}

	return s, nil
}

func (s *Server) Boot() {
	s.bootOnce.Do(func() {})
}

func (s *Server) Config() microserver.Config {
	return s.config
}

func (s *Server) Shutdown() {
	s.shutdownOnce.Do(func() {})
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	rErr := err.(microserver.ResponseError)
	uErr := rErr.Underlying()

	if creator.IsAlreadyExists(uErr) {
		rErr.SetCode(microserver.CodeResourceAlreadyExists)
		rErr.SetMessage(uErr.Error())
		w.WriteHeader(http.StatusConflict)
	} else {
		rErr.SetCode(microserver.CodeInternalError)
		rErr.SetMessage(uErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
