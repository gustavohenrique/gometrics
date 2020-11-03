package gometrics

import (
	"gometrics/lib/collectors"
	"gometrics/lib/domain"
)

type Collector struct{}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) GetInfoByPid(pid int) (domain.PidInfo, error) {
	info := domain.PidInfo{}
	seconds := 1
	pidCollector := collectors.NewPidCollector()
	pidStat, err := pidCollector.GetPidStatByInterval(pid, seconds)
	if err != nil {
		return info, err
	}
	info.PID = pidStat.PID
	info.CpuUsagePercentage = pidStat.CpuUsagePercentage
	info.NumThreads = pidStat.NumThreads
	info.State = pidStat.StateName

	memory, err := pidCollector.GetProportionalMemoryUsage(pid)
	if err != nil {
		return info, err
	}
	info.MemoryUsage = memory / 1024

	systemCollector := collectors.NewSystemCollector()
	info.NumCPU = systemCollector.GetNumCPU()
	meminfo, err := systemCollector.GetMemoryInfo()
	if err == nil {
		info.MemoryTotal = meminfo.MemTotal / 1024
	}

	return info, nil
}
