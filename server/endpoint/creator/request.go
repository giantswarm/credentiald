package creator

import (
	"github.com/giantswarm/credentiald/v2/server/endpoint/creator/aws"
	"github.com/giantswarm/credentiald/v2/server/endpoint/creator/azure"
)

type Request struct {
	Organization string `json:"-"`
	Provider     string `json:"provider"`

	AWS   aws.AWS     `json:"aws"`
	Azure azure.Azure `json:"azure"`
}
