package collectors

import (
	"fmt"
	"time"

	"github.com/gustavohenrique/gometrics/lib/domain"
	"github.com/gustavohenrique/gometrics/lib/util"
)

type DockerCollector struct {
	Cgroup      string
	MemoryUsage string
	MemoryLimit string
	CpuUsage    string
}

func NewDockerCollector() *DockerCollector {
	return &DockerCollector{
		Cgroup:      "/sys/fs/cgroup",
		MemoryUsage: "memory/memory.usage_in_bytes",
		MemoryLimit: "memory/memory.limit_in_bytes",
		CpuUsage:    "cpuacct/cpuacct.usage",
	}
}

func (c *DockerCollector) GetStat(seconds int) (domain.DockerStat, error) {
	var stat domain.DockerStat
	filename := fmt.Sprintf("%s/%s", c.Cgroup, c.MemoryUsage)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return stat, err
	}
	stat.MemoryUsage = util.ParseUint64(string(data))

	filename = fmt.Sprintf("%s/%s", c.Cgroup, c.MemoryLimit)
	data, err = util.ReadFileNoStat(filename)
	if err != nil {
		return stat, err
	}
	stat.MemoryLimit = util.ParseUint64(string(data))

	filename = fmt.Sprintf("%s/%s", c.Cgroup, c.CpuUsage)
	data, err = util.ReadFileNoStat(filename)
	if err != nil {
		return stat, err
	}
	cpuSample1 := util.ParseUint64(string(data))

	interval := time.Duration(seconds) * time.Second
	<-time.After(interval)

	data, _ = util.ReadFileNoStat(filename)
	cpuSample2 := util.ParseUint64(string(data))

	usage := c.calculateCpuUsagePercentage(cpuSample1, cpuSample2, seconds)
	stat.CpuUsagePercentage = usage
	return stat, nil
}

func (c *DockerCollector) calculateCpuUsagePercentage(past, now uint64, seconds int) float64 {
	totalTime := now - past
	usage := (float64(totalTime) / float64(seconds))
	return util.ParseFloat(fmt.Sprintf("%.2f", usage))
}
