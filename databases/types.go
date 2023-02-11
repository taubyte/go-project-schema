package databases

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Database interface {
	Get() Getter
	common.Resource[*structureSpec.Database]
}

type database struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.Database]
	Match() string
	Regex() bool
	Local() bool
	Secret() bool
	Min() int
	Max() int
	Path() string
	Size() string
	Encryption() (key string, keyType string)
}
