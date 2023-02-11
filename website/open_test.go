package website_test

import (
	"testing"

	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"bitbucket.org/taubyte/go-project-schema/website"
	"gotest.tools/v3/assert"
)

func TestOpenErrors(t *testing.T) {
	seer, err := internal.NewSeer()
	assert.NilError(t, err)

	_, err = website.Open(seer, "", "")
	assert.ErrorContains(t, err, "on website ``; name is empty")

	_, err = website.Open(nil, "test_website1", "")
	assert.ErrorContains(t, err, "on website `test_website1`; seer is nil")
}
