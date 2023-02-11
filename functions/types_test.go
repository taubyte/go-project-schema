package functions

import (
	"testing"

	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
	"gotest.tools/v3/assert"
)

func TestTypes(t *testing.T) {
	fun := &function{
		Resource: &basic.Resource{},
		seer:     &seer.Seer{},
		name:     "fun1",
	}

	assert.Equal(t, fun.Name(), "fun1")
	assert.Equal(t, fun.AppName(), "")

	err := fun.WrapError("failed: %s", "test error")
	assert.ErrorContains(t, err, "on function `fun1`; failed: test error")
}
