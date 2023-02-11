package libraries_test

import (
	"testing"

	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"bitbucket.org/taubyte/go-project-schema/libraries"
	"gotest.tools/v3/assert"
)

func TestOpenErrors(t *testing.T) {
	seer, err := internal.NewSeer()
	assert.NilError(t, err)

	_, err = libraries.Open(seer, "", "")
	assert.ErrorContains(t, err, "on library ``; name is empty")

	_, err = libraries.Open(nil, "test_library1", "")
	assert.ErrorContains(t, err, "on library `test_library1`; seer is nil")
}
