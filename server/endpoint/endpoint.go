package endpoint

import (
	"github.com/giantswarm/microendpoint/endpoint/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/credentiald/server/endpoint/creator"
	"github.com/giantswarm/credentiald/server/endpoint/lister"
	"github.com/giantswarm/credentiald/server/middleware"
	"github.com/giantswarm/credentiald/service"
)

type Config struct {
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *service.Service
}

type Endpoint struct {
	Creator *creator.Endpoint
	Lister  *lister.Endpoint
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

	var creatorEndpoint *creator.Endpoint
	{
		c := creator.Config{
			Logger:  config.Logger,
			Service: config.Service.Creator,
		}

		creatorEndpoint, err = creator.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var listerEndpoint *lister.Endpoint
	{
		c := lister.Config{
			Logger:  config.Logger,
			Service: config.Service.Lister,
		}

		listerEndpoint, err = lister.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionEndpoint *version.Endpoint
	{
		c := version.Config{
			Logger:  config.Logger,
			Service: config.Service.Version,
		}

		versionEndpoint, err = version.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	endpoint := &Endpoint{
		Creator: creatorEndpoint,
		Lister:  listerEndpoint,
		Version: versionEndpoint,
	}

	return endpoint, nil
}
