package smartops_test

import (
	_ "embed"
	"testing"

	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"bitbucket.org/taubyte/go-project-schema/smartops"
	"gotest.tools/v3/assert"
)

func TestYaml(t *testing.T) {
	getter, err := smartops.Yaml("test_smartops1", "", internal.SmartOp1)
	assert.NilError(t, err)

	assertSmartops1(t, getter)

	getter, err = smartops.Yaml("test_smartops2", "test_app1", internal.SmartOp2)
	assert.NilError(t, err)

	assertSmartops2(t, getter)
}
