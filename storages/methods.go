package storages

import (
	"fmt"

	"bitbucket.org/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
)

func (s *storage) WrapError(format string, i ...any) error {
	return fmt.Errorf("on storage `"+s.name+"`; "+format, i...)
}

func (s *storage) Root() *seer.Query {
	return s.Resource.Root()
}

func (s *storage) Config() *seer.Query {
	return s.Resource.Config()
}

func (s *storage) Name() string {
	return s.name
}

func (s *storage) AppName() string {
	return s.application
}

func (*storage) Directory() string {
	return common.StorageFolder
}

func (s *storage) SetName(name string) {
	s.name = name
}
