package proc_test

import (
	"testing"

	"github.com/gustavohenrique/gometrics/lib/proc"
	"github.com/gustavohenrique/gometrics/test"
	"github.com/gustavohenrique/gometrics/test/assert"
)

func TestParsePidStat(ts *testing.T) {

	test.It(ts, "Should parse stats", func(t *testing.T) {
		fake := "123456 (Web Content) S 2548 2034 2034 0 -1 4194560 42489276 0 46 0 3580808 1873803 0 0 20 0 48 0 368750 3881967616 205072 18446744073709551615 94514830016656 94514830431600 140732515325120 0 0 0 0 69634 1082133752 0 0 0 17 6 0 0 1 0 0 94514830443936 94514830444304 94514849296384 140732515333852 140732515334028 140732515334028 140732515336159 0"
		stat, err := proc.ParsePidStat([]byte(fake))

		assert.Nil(t, err)
		assert.Equal(t, stat.State, "S")
		assert.Equal(t, stat.StateName, "sleeping")
		assert.Equal(t, stat.UTime, uint(3580808))
		assert.Equal(t, stat.STime, uint(1873803))
		assert.Equal(t, stat.CUTime, uint(0))
		assert.Equal(t, stat.CSTime, uint(0))
		assert.Equal(t, stat.StartTime, uint64(368750))
		assert.Equal(t, stat.NumThreads, 48)
		assert.Equal(t, stat.RSS, 205072)
		assert.Equal(t, stat.CpuTotalTimeSpent, uint(5454611))
	})

	test.It(ts, "Should return error when the the string format does not contains parentheses", func(t *testing.T) {
		fake := "123456 Web Content S 2548 2034"
		_, err := proc.ParsePidStat([]byte(fake))
		assert.NotNil(t, err, "")
	})
}
