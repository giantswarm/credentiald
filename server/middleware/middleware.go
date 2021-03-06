package middleware

import (
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/credentiald/v2/service"
)

type Config struct {
	Logger  micrologger.Logger
	Service *service.Service
}

type Middleware struct{}

func New(config Config) (*Middleware, error) {
	return &Middleware{}, nil
}
