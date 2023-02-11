package domains_test

import (
	"testing"

	"bitbucket.org/taubyte/go-project-schema/domains"
	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"gotest.tools/v3/assert"
)

func TestOpenErrors(t *testing.T) {
	seer, err := internal.NewSeer()
	assert.NilError(t, err)

	_, err = domains.Open(seer, "", "")
	assert.ErrorContains(t, err, "on domain ``; name is empty")

	_, err = domains.Open(nil, "test_domain1", "")
	assert.ErrorContains(t, err, "on domain `test_domain1`; seer is nil")
}
