package services

import (
	"testing"

	"bitbucket.org/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
	"gotest.tools/v3/assert"
)

func TestTypes(t *testing.T) {
	srv := &service{
		Resource: &basic.Resource{},
		seer:     &seer.Seer{},
		name:     "srv1",
	}

	assert.Equal(t, srv.Name(), "srv1")
	assert.Equal(t, srv.AppName(), "")

	err := srv.WrapError("failed: %s", "test error")
	assert.ErrorContains(t, err, "on service `srv1`; failed: test error")
}
