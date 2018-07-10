package creator

import (
	"github.com/giantswarm/credentiald/service/creator/aws"
)

type Request struct {
	Organization string

	AWS aws.Request
}
