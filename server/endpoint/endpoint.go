package endpoint

import (
	"github.com/giantswarm/microendpoint/endpoint/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/credentiald/server/middleware"
	"github.com/giantswarm/credentiald/service"
)

type Config struct {
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *service.Service
}

type Endpoint struct {
	Version *version.Endpoint
}

func New(config Config) (*Endpoint, error) {
	var err error

	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.Service == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Service must not be empty", config)
	}

	var versionEndpoint *version.Endpoint
	{
		versionConfig := version.DefaultConfig()
		versionConfig.Logger = config.Logger
		versionConfig.Service = config.Service.Version
		versionEndpoint, err = version.New(versionConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	endpoint := &Endpoint{
		Version: versionEndpoint,
	}

	return endpoint, nil
}
