package functions

import (
	"testing"

	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
	"gotest.tools/v3/assert"
)

func TestMethodsRoot(t *testing.T) {
	d := &function{
		Resource: &basic.Resource{
			Root: func() *seer.Query { return nil },
		},
	}

	var nilQuery *seer.Query
	assert.Equal(t, d.Root(), nilQuery)
}
