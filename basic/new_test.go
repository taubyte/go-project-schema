package basic_test

import (
	"testing"

	"bitbucket.org/taubyte/go-project-schema/basic"
	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"gotest.tools/v3/assert"
)

func TestNewError(t *testing.T) {
	project, err := internal.NewProjectReadOnly()
	assert.NilError(t, err)

	db, err := project.Database("test_database1", "")
	assert.NilError(t, err)

	_, err = basic.NewNoName(nil, db.(basic.ResourceIface))
	assert.ErrorContains(t, err, "on database `test_database1`; seer is nil")
}
