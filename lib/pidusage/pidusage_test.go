package pidusage_test

import (
	"os"
	"path/filepath"
	"testing"

	"gometrics/lib/pidusage"
	"gometrics/test"
	"gometrics/test/assert"
)

func getTestDataPath() string {
	current, _ := os.Getwd()
	parent := filepath.Dir(current)
	root := filepath.Dir(parent)
	return filepath.Join(root, "test", "testdata")
}

func TestGetPidUsage(ts *testing.T) {
	pid := 123456

	test.It(ts, "Return", func(t *testing.T) {
		procPath := getTestDataPath() + "/proc"
		stat, err := pidusage.GetStat(pid, procPath)
		assert.Nil(t, err, "error is not null")
		assert.Equal(t, float64(0.001205755629506568), stat.CPU, "")
		assert.Equal(t, float64(6140), stat.Memory, "")
	})
}
