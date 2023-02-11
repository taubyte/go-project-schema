package domains

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
	"github.com/taubyte/go-seer"
)

type Domain interface {
	Get() Getter
	common.Resource[*structureSpec.Domain]
}

type domain struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.Domain]
	FQDN() string
	UseCertificate() bool
	Type() string
	Cert() string
	Key() string
}
