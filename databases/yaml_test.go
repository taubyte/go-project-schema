package databases_test

import (
	_ "embed"
	"testing"

	"bitbucket.org/taubyte/go-project-schema/databases"
	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"gotest.tools/v3/assert"
)

func TestYaml(t *testing.T) {
	getter, err := databases.Yaml("test_database1", "", internal.Database1)
	assert.NilError(t, err)

	assertDatabase1(t, getter)

	getter, err = databases.Yaml("test_database2", "test_app1", internal.Database2)
	assert.NilError(t, err)

	assertDatabase2(t, getter)
}
