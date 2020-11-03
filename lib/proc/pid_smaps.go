package proc

import (
	"bufio"
	"bytes"
	"fmt"
)

// https://en.wikipedia.org/wiki/Proportional_set_size
func ParseSmaps(data []byte) (uint64, error) {
	res := uint64(0)
	pfx := []byte("Pss:")
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.HasPrefix(line, pfx) {
			var size uint64
			_, err := fmt.Sscanf(string(line[4:]), "%d", &size)
			if err != nil {
				return 0, err
			}
			res += size
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return res, nil
}
