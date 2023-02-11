package project_test

import (
	"errors"
	"testing"

	"github.com/taubyte/go-project-schema/basic"
	internal "github.com/taubyte/go-project-schema/internal/test"
	"github.com/taubyte/go-project-schema/project"
	"gotest.tools/v3/assert"
)

func TestMethods(t *testing.T) {
	_project, err := internal.NewProjectEmpty()
	assert.NilError(t, err)

	_project.Set(true, project.Name("test_project"))

	iface, ok := _project.(basic.ResourceIface)
	assert.Equal(t, ok, true)

	err = iface.WrapError("failed with: %s", errors.New("test error"))
	assert.ErrorContains(t, err, "on project `test_project`; failed with: test error")

	assert.Equal(t, iface.Directory(), "")
	assert.Equal(t, iface.AppName(), "")

	iface.SetName("other_project")

	assert.Equal(t, iface.Name(), "other_project")
	assert.Equal(t, _project.Get().Name(), "other_project")
}
