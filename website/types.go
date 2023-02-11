package website

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
	"github.com/taubyte/go-seer"
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
