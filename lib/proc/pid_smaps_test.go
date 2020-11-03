package proc_test

import (
	"testing"

	"gometrics/lib/proc"
	"gometrics/test"
	"gometrics/test/assert"
	"gometrics/test/fs"
)

func TestParseSmaps(ts *testing.T) {
	fake := fs.ReadFromTestData("smaps")

	test.It(ts, "Should parse smaps", func(t *testing.T) {
		pss, err := proc.ParseSmaps(fake)
		assert.Nil(t, err, "")
		assert.Equal(t, pss, uint64(545531))
	})
}
