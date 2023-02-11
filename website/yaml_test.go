package website_test

import (
	_ "embed"
	"testing"

	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"bitbucket.org/taubyte/go-project-schema/website"
	"gotest.tools/v3/assert"
)

func TestYaml(t *testing.T) {
	getter, err := website.Yaml("test_website1", "", internal.Website1)
	assert.NilError(t, err)

	assertWebsite1(t, getter)

	getter, err = website.Yaml("test_website2", "test_app1", internal.Website2)
	assert.NilError(t, err)

	assertWebsite2(t, getter)
}
