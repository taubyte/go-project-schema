package website

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Website interface {
	Get() Getter
	common.Resource[*structureSpec.Website]
}

type website struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.Website]
	Domains() []string
	Paths() []string
	Branch() string
	Git() (provider, id, fullname string)
}
