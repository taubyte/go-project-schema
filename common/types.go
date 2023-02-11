package common

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/pretty"
	structureSpec "github.com/taubyte/go-specs/structure"
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
