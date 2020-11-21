package proc_test

import (
	"testing"

	"github.com/gustavohenrique/gometrics/lib/proc"
	"github.com/gustavohenrique/gometrics/test"
	"github.com/gustavohenrique/gometrics/test/assert"
)

func TestParseCgroup(ts *testing.T) {
	test.It(ts, "Should return init.scope when running outside a Docker container", func(t *testing.T) {
		data := []byte(`
12:devices:/
11:freezer:/
10:perf_event:/
9:rdma:/
8:pids:/
7:hugetlb:/
6:memory:/
5:net_cls,net_prio:/
4:cpu,cpuacct:/
3:blkio:/
2:cpuset:/
1:name=systemd:/init.scope
0::/init.scope`)
		platform, err := proc.ParseCgroup(data)
		assert.Nil(t, err)
		assert.Equal(t, "init.scope", platform)
	})

	test.It(ts, "Should return docker when running inside a Docker container", func(t *testing.T) {
		data := []byte(`
12:devices:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
11:freezer:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
10:perf_event:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
9:rdma:/
8:pids:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
7:hugetlb:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
6:memory:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
5:net_cls,net_prio:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
4:cpu,cpuacct:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
3:blkio:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
2:cpuset:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
1:name=systemd:/docker/0e91dfebc23d7e71032702770de92774d7852ff01f873f9c7d32c7ace8d8027d
0::/system.slice/containerd.service`)
		platform, err := proc.ParseCgroup(data)
		assert.Nil(t, err)
		assert.Equal(t, "docker", platform)
	})
}
