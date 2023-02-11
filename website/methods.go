package website

import (
	"fmt"

	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
)

func (w *website) WrapError(format string, i ...any) error {
	return fmt.Errorf("on website `"+w.name+"`; "+format, i...)
}

func (w *website) Root() *seer.Query {
	return w.Resource.Root()
}

func (w *website) Config() *seer.Query {
	return w.Resource.Config()
}

func (w *website) Name() string {
	return w.name
}

func (w *website) AppName() string {
	return w.application
}

func (*website) Directory() string {
	return common.WebsiteFolder
}

func (s *website) SetName(name string) {
	s.name = name
}
