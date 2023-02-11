package libraries

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Library interface {
	Get() Getter
	common.Resource[*structureSpec.Library]
}

type library struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.Library]
	Path() string
	Branch() string
	Git() (provider, id, fullname string)
}
