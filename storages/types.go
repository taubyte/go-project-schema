package storages

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
	"github.com/taubyte/go-seer"
)

type Storage interface {
	Get() Getter
	common.Resource[*structureSpec.Storage]
}

type storage struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.Storage]
	Match() string
	Regex() bool
	Type() string

	// if bucket type Object
	Public() bool
	Versioning() bool

	// if bucket type Streaming
	TTL() string

	// independent
	Size() string
	SmartOps() []string
}
