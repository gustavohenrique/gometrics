package collectors

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gustavohenrique/gometrics/lib/domain"
	"github.com/gustavohenrique/gometrics/lib/proc"
	"github.com/gustavohenrique/gometrics/lib/util"
)

type PidCollector struct {
	Proc  string
	Stat  string
	Smaps string
}

func NewPidCollector() *PidCollector {
	return &PidCollector{
		Proc:  "/proc",
		Stat:  "stat",
		Smaps: "smaps",
	}
}

func (c *PidCollector) GetStat(pid, seconds int) (domain.PidStat, error) {
	stat, err := c.GetPidStatByInterval(pid, seconds)
	if err != nil {
		return stat, err
	}
	usage, err := c.GetMemoryUsageProportional(pid)
	if err != nil {
		return stat, err
	}
	stat.MemoryUsage = usage
	return stat, nil
}

func (c *PidCollector) GetMemoryUsageProportional(pid int) (uint64, error) {
	filename := fmt.Sprintf("%s/%d/%s", c.Proc, pid, c.Smaps)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return 0, err
	}
	return proc.ParseSmaps(data)
}

func (c *PidCollector) GetPidStatByInterval(pid, seconds int) (domain.PidStat, error) {
	var stat domain.PidStat
	filename := fmt.Sprintf("%s/%d/%s", c.Proc, pid, c.Stat)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return stat, err
	}
	pidStatSample1, err := proc.ParsePidStat(data)
	if err != nil {
		return stat, err
	}

	interval := time.Duration(seconds) * time.Second
	<-time.After(interval)

	data, _ = util.ReadFileNoStat(filename)
	pidStatSample2, err := proc.ParsePidStat(data)
	if err != nil {
		return stat, err
	}
	pidStatSample2.PID = pid

	usage := c.calculateCpuUsagePercentage(pidStatSample1.CpuTotalTimeSpent, pidStatSample2.CpuTotalTimeSpent, seconds)
	pidStatSample2.CpuUsagePercentage = usage
	return pidStatSample2, nil
}

func (c *PidCollector) calculateCpuUsagePercentage(past, now uint, seconds int) float64 {
	totalTime := now - past
	usage := (float64(totalTime) / float64(seconds)) / float64(runtime.NumCPU()) // * 100
	return util.ParseFloat(fmt.Sprintf("%.2f", usage))
}
