package creator

import (
	"github.com/giantswarm/microerror"
)

var alreadyExistsError = &microerror.Error{
	Kind: "alreadyExistsError",
}

// IsAlreadyExists asserts alreadyExistsError.
func IsAlreadyExists(err error) bool {
	return microerror.Cause(err) == alreadyExistsError
}

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidProviderError = &microerror.Error{
	Kind: "invalidProviderError",
}

// IsInvalidProvider asserts invalidProviderError.
func IsInvalidProvider(err error) bool {
	return microerror.Cause(err) == invalidProviderError
}
