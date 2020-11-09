package collectors_test

import (
	"testing"

	"github.com/gustavohenrique/gometrics/lib/collectors"
	"github.com/gustavohenrique/gometrics/test"
	"github.com/gustavohenrique/gometrics/test/assert"
	"github.com/gustavohenrique/gometrics/test/fs"
)

func TestDockerCollector(ts *testing.T) {
	collector := collectors.NewDockerCollector()
	collector.Cgroup = fs.GetTestDataPath()

	test.It(ts, "GetStat", func(t *testing.T) {
		stat, err := collector.GetStat(1)
		assert.Nil(t, err)
		assert.Equal(t, stat.MemoryUsage, uint64(8810496))
		assert.Equal(t, stat.MemoryLimit, uint64(10485760))
		assert.Equal(t, stat.CpuUsagePercentage, 0.0)
	})
}
