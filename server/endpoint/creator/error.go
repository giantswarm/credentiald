package creator

import (
	"github.com/giantswarm/microerror"
)

var aleradyExistsError = microerror.New("alerady exists")

// IsAleradyExists asserts aleradyExistsError.
func IsAleradyExists(err error) bool {
	return microerror.Cause(err) == aleradyExistsError
}

var invalidConfigError = microerror.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}
