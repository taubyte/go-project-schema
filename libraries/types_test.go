package libraries

import (
	"testing"

	"bitbucket.org/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
	"gotest.tools/v3/assert"
)

func TestTypes(t *testing.T) {
	lib := &library{
		Resource: &basic.Resource{},
		seer:     &seer.Seer{},
		name:     "lib1",
	}

	assert.Equal(t, lib.Name(), "lib1")
	assert.Equal(t, lib.AppName(), "")

	err := lib.WrapError("failed: %s", "test error")
	assert.ErrorContains(t, err, "on library `lib1`; failed: test error")
}
