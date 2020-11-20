package util

import "strings"

func IsInsideDockerContainer(filename string) bool {
	if filename == "" {
		filename = "/proc/1/cgroup"
	}
	data, err := ReadFileNoStat(filename)
	if err != nil {
		return false
	}
	return strings.Contains(string(data), "docker")
}
