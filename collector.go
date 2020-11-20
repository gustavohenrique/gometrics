package gometrics

import (
	"os"

	"github.com/gustavohenrique/gometrics/lib/collectors"
	"github.com/gustavohenrique/gometrics/lib/domain"
	"github.com/gustavohenrique/gometrics/lib/util"
)

var (
	dockerCollector  = collectors.NewDockerCollector()
	pidCollector     = collectors.NewPidCollector()
	systemCollector  = collectors.NewSystemCollector()
	runtimeCollector = collectors.NewRuntimeCollector()
	pid              = os.Getpid()
)

type Collector struct {
	DurationBetweenCPUSamples int
}

func New() *Collector {
	return &Collector{
		DurationBetweenCPUSamples: 1,
	}
}

func (c *Collector) Metrics() (domain.Metrics, error) {
	if util.IsInsideDockerContainer() {
		return c.Docker()
	}
	return c.Process()
}

func (c *Collector) Docker() (domain.Metrics, error) {
	var metrics domain.Metrics
	dockerStat, err := dockerCollector.GetStat(c.DurationBetweenCPUSamples)
	if err != nil {
		return metrics, err
	}
	metrics.PID = pid
	metrics.MemoryUsage = dockerStat.MemoryUsage / (1024 * 1024)
	metrics.MemoryTotal = dockerStat.MemoryLimit / (1024 * 1024)
	metrics.CpuUsagePercentage = dockerStat.CpuUsagePercentage / (1024 * 1024)
	metrics.RuntimeStat = c.getGoRuntimeStat()
	metrics.IsDockerContainer = true
	return metrics, nil
}

func (c *Collector) Process() (domain.Metrics, error) {
	metrics, err := c.GetProcessMetricsByPID(pid)
	if err != nil {
		return metrics, err
	}
	metrics.RuntimeStat = c.getGoRuntimeStat()
	metrics.IsDockerContainer = false
	return metrics, nil
}

func (c *Collector) GetProcessMetricsByPID(pid int) (domain.Metrics, error) {
	var metrics domain.Metrics
	metrics.PID = pid

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
