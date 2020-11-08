package gometrics

import (
	"os"
	"runtime"
	"syscall"

	"gometrics/lib/collectors"
	"gometrics/lib/domain"
)

var (
	seconds = 1
	dockerCollector = collectors.NewDockerCollector()
	pidCollector = collectors.NewPidCollector()
)

type Collector struct{}

func New() *Collector {
	return &Collector{}
}

func (c *Collector) GetInfoFromCurrentDockerContainer() (domain.DockerInfo, error) {
	info := domain.DockerInfo{}
	dockerStat, err := dockerCollector.GetDockerStatByInterval(seconds)
	if err != nil {
		return info, err
	}
	info.MemoryUsage = dockerStat.MemoryUsage
	info.MemoryLimit = dockerStat.MemoryLimit
	info.CpuUsagePercentage = dockerStat.CpuUsagePercentage
	return info, nil
}

func (c *Collector) GetInfoFromCurrentProc() (domain.GoInfo, error) {
	var ru syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &ru)
	maxRss := ru.Maxrss

	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	pid := os.Getpid()
	pidInfo, err := c.GetInfoByPid(pid)
	info := domain.GoInfo{}
	info.PidInfo = pidInfo
	info.NumGoroutine = runtime.NumGoroutine()

	// https://golang.org/pkg/runtime/
	info.MemoryMaxRss = maxRss
	info.Alloc = rtm.Alloc
	info.TotalAlloc = rtm.TotalAlloc
	info.Sys = rtm.Sys
	info.Mallocs = rtm.Mallocs
	info.Frees = rtm.Frees
	info.LiveObjects = info.Mallocs - info.Frees
	info.PauseTotalNs = rtm.PauseTotalNs
	info.NumGC = rtm.NumGC
	return info, err
}

func (c *Collector) GetInfoByPid(pid int) (domain.PidInfo, error) {
	info := domain.PidInfo{}
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
