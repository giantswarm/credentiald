package client

import (
	"github.com/giantswarm/microerror"
)

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidRequestError = &microerror.Error{
	Kind: "invalidRequestError",
}

// IsInvalidRequest asserts invalidRequestError.
func IsInvalidRequest(err error) bool {
	return microerror.Cause(err) == invalidRequestError
}

var invalidResponseError = &microerror.Error{
	Kind: "invalidResponseError",
}

// IsInvalidResponse asserts invalidResponseError.
func IsInvalidResponse(err error) bool {
	return microerror.Cause(err) == invalidResponseError
}
