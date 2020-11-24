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
	pid              = os.Getpid()
	cgroup           string
)

type Collector struct {
	DurationBetweenCPUSamples int
	EcsCpuTaskUnit            int
}

func init() {
	cgroup = systemCollector.GetCgroup()
}

func New() *Collector {
	return &Collector{
		DurationBetweenCPUSamples: 1,
		EcsCpuTaskUnit:            2048,
	}
}

func (c *Collector) Metrics() (domain.Metrics, error) {
	if cgroup == domain.LINUX {
		return c.Process()
	}
	metrics, err := c.Docker()
	if err != nil {
		return metrics, err
	}
	if cgroup == domain.ECS {
		metrics.CpuUsagePercentage = (metrics.CpuUsagePercentage * 100) / float64(c.EcsCpuTaskUnit)
	}
	return metrics, nil
}

func (c *Collector) Docker() (domain.Metrics, error) {
	var metrics domain.Metrics
	dockerStat, err := dockerCollector.GetStat(c.DurationBetweenCPUSamples)
	if err != nil {
		return metrics, err
	}
	metrics.PID = pid
	metrics.Cgroup = cgroup
	metrics.MemoryUsage = dockerStat.MemoryUsage / (1024 * 1024)
	metrics.MemoryTotal = dockerStat.MemoryLimit / (1024 * 1024)
	metrics.CpuUsagePercentage = dockerStat.CpuUsagePercentage / (1024 * 1024)
	metrics.RuntimeStat = c.getGoRuntimeStat()
	return metrics, nil
}

func (c *Collector) Process() (domain.Metrics, error) {
	metrics, err := c.GetProcessMetricsByPID(pid)
	if err != nil {
		return metrics, err
	}
	metrics.RuntimeStat = c.getGoRuntimeStat()
	return metrics, nil
}

func (c *Collector) GetProcessMetricsByPID(pid int) (domain.Metrics, error) {
	var metrics domain.Metrics
	metrics.PID = pid
	metrics.Cgroup = cgroup

	pidStat, err := pidCollector.GetStat(pid, c.DurationBetweenCPUSamples)
	if err != nil {
		return metrics, err
	}
	metrics.PidStat = &pidStat

	systemStat, err := systemCollector.GetStat()
	if err != nil {
		return metrics, err
	}
	metrics.SystemStat = &systemStat

	metrics.MemoryUsage = pidStat.MemoryUsage / 1024
	metrics.MemoryTotal = systemStat.MemTotal / 1024
	metrics.CpuUsagePercentage = pidStat.CpuUsagePercentage

	return metrics, nil
}

func (c *Collector) getGoRuntimeStat() *domain.RuntimeStat {
	runtimeStat, err := runtimeCollector.GetStat()
	if err != nil {
		return nil
	}
	return &runtimeStat
}
