package util_test

import (
	"testing"

	"github.com/gustavohenrique/gometrics/lib/util"
	"github.com/gustavohenrique/gometrics/test"
	"github.com/gustavohenrique/gometrics/test/assert"
	"github.com/gustavohenrique/gometrics/test/fs"
)

func TestIsInsideDockerContainer(ts *testing.T) {
	test.It(ts, "Should return false when running outside a Docker container", func(t *testing.T) {
		filename := fs.GetTestDataPath() + "/1/cgroup"
		result := util.IsInsideDockerContainer(filename)
		assert.Equal(t, false, result)
	})

	test.It(ts, "Should return true when running inside a Docker container", func(t *testing.T) {
		filename := fs.GetTestDataPath() + "/1/cgroup.docker"
		result := util.IsInsideDockerContainer(filename)
		assert.Equal(t, true, result)
	})
}
