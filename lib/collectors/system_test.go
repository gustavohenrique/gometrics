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

	test.It(ts, "GetMemoryInfo", func(t *testing.T) {
		info, err := collector.GetMemoryInfo()
		assert.Nil(t, err)
		assert.Equal(t, info.MemTotal, uint64(32767112))
	})
}
