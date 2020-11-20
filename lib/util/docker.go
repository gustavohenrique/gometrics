package util

import (
	"strings"
)

func IsInsideDockerContainer(file ...string) bool {
	filename := "/proc/1/cgroup"
	if len(file) > 0 && file[0] != "" {
		filename = file[0]
	}
	data, err := ReadFileNoStat(filename)
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "docker")
}
