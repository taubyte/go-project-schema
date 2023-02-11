package common

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/pretty"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
)

type Mapper []struct {
	Field      string
	IfNotEmpty bool
	Callback   func() error
}

type Resource[T structureSpec.Structure] interface {
	Set(sync bool, ops ...basic.Op) (err error)
	Delete(attributes ...string) (err error)
	Prettify(p pretty.Prettier) map[string]interface{}
	SetWithStruct(sync bool, structure T) error
}
