package pretty

import (
	commonSpec "github.com/taubyte/go-specs/common"
)

type Object interface {
	Interface() interface{}
}

type Prettier interface {
	Fetch(path *commonSpec.TnsPath) (Object, error)
	Project() string
	Branch() string
}

type PrettyResource interface {
	Prettify(p Prettier) map[string]interface{}
}

type PrettyResourceIface struct {
	Type string
	Get  func(name, application string) (PrettyResource, error)
	List func(application string) (local []string, global []string)
}
