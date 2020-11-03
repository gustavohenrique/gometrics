package lib_test

import (
	"os"
	"path/filepath"
	"testing"

	"gometrics/lib"
	"gometrics/test"
	"gometrics/test/assert"
)

func getTestDataPath() string {
	current, _ := os.Getwd()
	parent := filepath.Dir(current)
	return filepath.Join(parent, "test", "testdata")
}

func TestGetCpuUsage(ts *testing.T) {
	pid := 123456

	test.It(ts, "Return", func(t *testing.T) {
		pidStat := lib.NewPIDStat()
		pidStat.ProcPath = getTestDataPath() + "/proc"
		usage := pidStat.GetCpuUsage(pid)
		assert.Equal(t, 5.12, usage, "")
	})
}
