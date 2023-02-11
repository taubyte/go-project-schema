package project_test

import (
	"testing"

	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"bitbucket.org/taubyte/go-project-schema/project"
	"github.com/spf13/afero"
	"gotest.tools/v3/assert"
)

func TestOpenError(t *testing.T) {
	_, err := project.Open(project.VirtualFS(afero.NewMemMapFs(), "invalid"))
	assert.ErrorContains(t, err, "Opening repository failed with open invalid: file does not exist")
}

func TestOpenSystemFS(t *testing.T) {
	project, err := internal.NewProjectSystemFS()
	assert.NilError(t, err)

	assert.Equal(t, project.Get().Id(), "projectID1")
}
