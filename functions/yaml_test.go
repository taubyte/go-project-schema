package functions_test

import (
	_ "embed"
	"testing"

	"bitbucket.org/taubyte/go-project-schema/functions"
	internal "bitbucket.org/taubyte/go-project-schema/internal/test"
	"gotest.tools/v3/assert"
)

func TestYaml(t *testing.T) {
	getter, err := functions.Yaml("test_function1", "", internal.Function1)
	assert.NilError(t, err)

	assertFunction1_http(t, getter)

	getter, err = functions.Yaml("test_function2", "test_app1", internal.Function2)
	assert.NilError(t, err)

	assertFunction2_pubsub(t, getter)

	getter, err = functions.Yaml("test_function3", "test_app2", internal.Function3)
	assert.NilError(t, err)

	assertFunction3_p2p(t, getter)
}
