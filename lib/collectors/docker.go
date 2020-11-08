package collectors

import (
	"fmt"
	"strconv"
	"time"

	"gometrics/lib/domain"
	"gometrics/lib/util"
)

type DockerCollector struct {
	Proc        string
	MemoryUsage string
	MemoryLimit string
	CpuUsage    string
}

func NewDockerCollector() *DockerCollector {
	return &DockerCollector{
		Proc:        "/sys/fs/cgroup",
		MemoryUsage: "memory/memory.usage_in_bytes",
		MemoryLimit: "memory/memory.limit_in_bytes",
		CpuUsage:    "cpuacct/cpuacct.usage",
	}
}

func (p *DockerCollector) GetDockerStatByInterval(seconds int) (domain.DockerStat, error) {
	var stat domain.DockerStat
	filename := fmt.Sprintf("%s/%s", p.Proc, p.MemoryUsage)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return stat, err
	}
	memoryUsageInBytes, _ := strconv.ParseUint(string(data), 10, 64)
	stat.MemoryUsage = memoryUsageInBytes / (1024 * 1024)

	filename = fmt.Sprintf("%s/%s", p.Proc, p.MemoryLimit)
	data, err = util.ReadFileNoStat(filename)
	if err != nil {
		return stat, err
	}
	memoryLimitInBytes, _ := strconv.ParseUint(string(data), 10, 64)
	stat.MemoryLimit = memoryLimitInBytes / (1024 * 1024)

	filename = fmt.Sprintf("%s/%s", p.Proc, p.CpuUsage)
	data, err = util.ReadFileNoStat(filename)
	if err != nil {
		return stat, err
	}
	cpuSample1, _ := strconv.ParseUint(string(data), 10, 64)

	interval := time.Duration(seconds) * time.Second
	<-time.After(interval)

	data, _ = util.ReadFileNoStat(filename)
	cpuSample2, _ := strconv.ParseUint(string(data), 10, 64)

	usage := p.CalculateCpuUsagePercentage(cpuSample1, cpuSample2, seconds)
	stat.CpuUsagePercentage = usage
	return stat, nil
}

func (p *DockerCollector) CalculateCpuUsagePercentage(past, now uint64, seconds int) float64 {
	totalTime := now - past
	usage := (float64(totalTime) / float64(seconds))
	return util.ParseFloat(fmt.Sprintf("%.2f", usage))
}
