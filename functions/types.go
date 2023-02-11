package functions

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
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
