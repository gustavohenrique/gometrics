package proc

import (
	"bufio"
	"bytes"
	"strings"
)

func ParseCgroup(data []byte) (string, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	pfx := []byte("1:")
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.HasPrefix(line, pfx) {
			l := string(line)
			parts := strings.Split(l, "/")
			if len(parts) >= 2 {
				return parts[1], nil
			}
			return l, nil
		}
	}
	err := scanner.Err()
	return "", err
}
