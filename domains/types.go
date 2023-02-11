package domains

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
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
