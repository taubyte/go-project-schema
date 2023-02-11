package storages

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
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
