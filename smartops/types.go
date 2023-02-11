package smartops

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
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
