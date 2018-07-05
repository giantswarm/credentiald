package creator

import (
	"github.com/giantswarm/credentiald/service/creator/aws"
	"github.com/giantswarm/credentiald/service/creator/azure"
)

type Request struct {
	Organization string
	Provider     string

	AWS   aws.Request
	Azure azure.Request
}
