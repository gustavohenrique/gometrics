package gometrics

import (
	"os"

	"github.com/gustavohenrique/gometrics/lib/collectors"
	"github.com/gustavohenrique/gometrics/lib/domain"
)

var (
	dockerCollector  = collectors.NewDockerCollector()
	pidCollector     = collectors.NewPidCollector()
	systemCollector  = collectors.NewSystemCollector()
	runtimeCollector = collectors.NewRuntimeCollector()
)

type Collector struct {
	DurationBetweenCPUSamples int
}

func New() *Collector {
	return &Collector{
		DurationBetweenCPUSamples: 1,
	}
}

func (c *Collector) GetDockerInfo() (domain.DockerInfo, error) {
	var info domain.DockerInfo
	dockerStat, err := dockerCollector.GetStat(c.DurationBetweenCPUSamples)
	if err != nil {
		return info, err
	}
	info.MemoryUsage = dockerStat.MemoryUsage / (1024 * 1024)
	info.MemoryLimit = dockerStat.MemoryLimit / (1024 * 1024)
	info.CpuUsagePercentage = dockerStat.CpuUsagePercentage

	runtimeStat, err := runtimeCollector.GetStat()
	if err != nil {
		return info, err
	}
	info.RuntimeStat = &runtimeStat

	return info, nil
}

func (c *Collector) GetRuntimeInfo() (domain.ProcessInfo, error) {
	info, err := c.GetProcessInfoByPID(os.Getpid())
	if err != nil {
		return info, err
	}

	runtimeStat, err := runtimeCollector.GetStat()
	if err != nil {
		return info, err
	}
	info.RuntimeStat = &runtimeStat

	return info, nil
}

func (c *Collector) GetProcessInfoByPID(pid int) (domain.ProcessInfo, error) {
	var info domain.ProcessInfo
	info.PID = pid

	pidStat, err := pidCollector.GetStat(pid, c.DurationBetweenCPUSamples)
	if err != nil {
		return info, err
	}
	info.PidStat = &pidStat

	systemStat, err := systemCollector.GetStat()
	if err != nil {
		return info, err
	}
	info.SystemStat = &systemStat

	info.MemoryUsage = pidStat.MemoryUsage / 1024
	info.MemoryTotal = systemStat.MemTotal / 1024
	info.CpuUsagePercentage = pidStat.CpuUsagePercentage

	return info, nil
}
