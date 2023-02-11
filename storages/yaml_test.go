package storages_test

import (
	_ "embed"
	"testing"

	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"bitbucket.org/taubyte/go-project-schema/storages"
	"gotest.tools/v3/assert"
)

func TestYaml(t *testing.T) {
	getter, err := storages.Yaml("test_storage1", "", internal.Storage1)
	assert.NilError(t, err)

	assertStorage1(t, getter)

	getter, err = storages.Yaml("test_storage2", "test_app1", internal.Storage2)
	assert.NilError(t, err)

	assertStorage2(t, getter)
}
