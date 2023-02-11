package services_test

import (
	"testing"

	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"bitbucket.org/taubyte/go-project-schema/services"
	"gotest.tools/v3/assert"
)

func TestOpenErrors(t *testing.T) {
	seer, err := internal.NewSeer()
	assert.NilError(t, err)

	_, err = services.Open(seer, "", "")
	assert.ErrorContains(t, err, "on service ``; name is empty")

	_, err = services.Open(nil, "test_service1", "")
	assert.ErrorContains(t, err, "on service `test_service1`; seer is nil")
}
