package client

import (
	"github.com/giantswarm/microerror"
)

var invalidConfigError = microerror.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidRequestError = microerror.New("invalid request")

// IsInvalidRequest asserts invalidRequestError.
func IsInvalidRequest(err error) bool {
	return microerror.Cause(err) == invalidRequestError
}

var invalidResponseError = microerror.New("invalid response")

// IsInvalidResponse asserts invalidResponseError.
func IsInvalidResponse(err error) bool {
	return microerror.Cause(err) == invalidResponseError
}
