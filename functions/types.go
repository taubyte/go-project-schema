package functions

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
	"github.com/taubyte/go-seer"
)

type Function interface {
	Get() Getter
	common.Resource[*structureSpec.Function]
}

type function struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.Function]
	Type() string
	Method() string
	Paths() []string
	Local() bool
	Command() string
	Channel() string
	Source() string
	Domains() []string
	Timeout() string
	Memory() string
	Call() string
	Protocol() string
}
