package databases

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
	"github.com/taubyte/go-seer"
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
