package util

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// ReadFileNoStat uses ioutil.ReadAll to read contents of entire file.
// This is similar to ioutil.ReadFile but without the call to os.Stat, because
// many files in /proc and /sys report incorrect file sizes (either 0 or 4096).
// Reads a max file size of 512kB.  For files larger than this, a scanner
// should be used.
func ReadFileNoStat(filename string) ([]byte, error) {
	const maxBufferSize = 1024 * 512

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := io.LimitReader(f, maxBufferSize)
	return ioutil.ReadAll(reader)
}

func ParseFloat(val string) float64 {
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}
	return floatVal
}
