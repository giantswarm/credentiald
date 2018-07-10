package creator

import (
	"github.com/giantswarm/credentiald/server/endpoint/creator/aws"
)

type Request struct {
	Provider string  `json:"provider"`
	AWS      aws.AWS `json:"aws"`

	Organization string `json:"-"`
}
