package collectors_test

import (
	"testing"

	"github.com/gustavohenrique/gometrics/lib/collectors"
	"github.com/gustavohenrique/gometrics/test"
	"github.com/gustavohenrique/gometrics/test/assert"
	"github.com/gustavohenrique/gometrics/test/fs"
)

func TestSystemCollector(ts *testing.T) {
	collector := collectors.NewSystemCollector()
	collector.Proc = fs.GetTestDataPath()

	test.It(ts, "GetUptime", func(t *testing.T) {
		uptime := collector.GetUptime()
		assert.Equal(t, uptime, 534235.93)
	})

	test.It(ts, "GetMemoryStat", func(t *testing.T) {
		stat, err := collector.GetMemoryStat()
		assert.Nil(t, err)
		assert.Equal(t, stat.MemTotal, uint64(32767112))
	})
}
