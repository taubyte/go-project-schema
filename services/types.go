package services

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
	"github.com/taubyte/go-seer"
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
