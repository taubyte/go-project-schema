package services

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
)

type Service interface {
	Get() Getter
	common.Resource[*structureSpec.Service]
}

type service struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.Service]
	Protocol() string
}
