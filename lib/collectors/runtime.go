package collectors

import (
	"runtime"
	"syscall"

	"github.com/gustavohenrique/gometrics/lib/domain"
)

var (
	ru  syscall.Rusage
	rtm runtime.MemStats
)

type RuntimeCollector struct{}

func NewRuntimeCollector() *RuntimeCollector {
	return &RuntimeCollector{}
}

func (c *RuntimeCollector) GetStat() (domain.RuntimeStat, error) {
	syscall.Getrusage(syscall.RUSAGE_SELF, &ru)
	maxRss := ru.Maxrss

	runtime.ReadMemStats(&rtm)

	// https://golang.org/pkg/runtime/
	stat := domain.RuntimeStat{}
	stat.NumGoroutine = runtime.NumGoroutine()
	stat.MemoryMaxRss = maxRss
	stat.Alloc = rtm.Alloc
	stat.TotalAlloc = rtm.TotalAlloc
	stat.Sys = rtm.Sys
	stat.Mallocs = rtm.Mallocs
	stat.Frees = rtm.Frees
	stat.LiveObjects = stat.Mallocs - stat.Frees
	stat.PauseTotalNs = rtm.PauseTotalNs
	stat.NumGC = rtm.NumGC
	return stat, nil
}
