package smartops

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
	"github.com/taubyte/go-seer"
)

type SmartOps interface {
	Get() Getter
	common.Resource[*structureSpec.SmartOp]
}

type smartOps struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.SmartOp]
	Source() string
	Timeout() string
	Memory() string
	Call() string
}
