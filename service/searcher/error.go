package searcher

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

var wrongOwnerOrganizationError = &microerror.Error{
	Kind: "wrongOwnerOrganizationError",
}

// IsWrongOwnerOrganizationError asserts wrongOwnerOrganizationError.
func IsWrongOwnerOrganizationError(err error) bool {
	return microerror.Cause(err) == wrongOwnerOrganizationError
}

var secretNotFoundError = &microerror.Error{
	Kind: "secretNotFoundError",
}

// IsSecretNotFoundError asserts secretNotFoundError.
func IsSecretNotFoundError(err error) bool {
	return microerror.Cause(err) == secretNotFoundError
}

var secretInUnexpectedFormatError = &microerror.Error{
	Kind: "secretInUnexpectedFormatError",
	Desc: "the resource we retrieved does not have the expected format",
}

// IsSecretInUnexpectedFormatError asserts secretInUnexpectedFormatError.
func IsSecretInUnexpectedFormatError(err error) bool {
	return microerror.Cause(err) == secretInUnexpectedFormatError
}
