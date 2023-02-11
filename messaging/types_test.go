package messaging

import (
	"testing"

	"bitbucket.org/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
	"gotest.tools/v3/assert"
)

func TestTypes(t *testing.T) {
	msg := &messaging{
		Resource: &basic.Resource{},
		seer:     &seer.Seer{},
		name:     "msg1",
	}

	assert.Equal(t, msg.Name(), "msg1")
	assert.Equal(t, msg.AppName(), "")

	err := msg.WrapError("failed: %s", "test error")
	assert.ErrorContains(t, err, "on messaging `msg1`; failed: test error")
}
