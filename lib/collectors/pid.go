package collectors

import (
	"fmt"
	"runtime"
	"time"

	"gometrics/lib/domain"
	"gometrics/lib/proc"
	"gometrics/lib/util"
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

func (p *PidCollector) GetProportionalMemoryUsage(pid int) (uint64, error) {
	filename := fmt.Sprintf("%s/%d/%s", p.Proc, pid, p.Smaps)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return 0, err
	}
	return proc.ParseSmaps(data)
}

func (p *PidCollector) GetPidStatByInterval(pid, seconds int) (domain.PidStat, error) {
	var stat domain.PidStat
	filename := fmt.Sprintf("%s/%d/%s", p.Proc, pid, p.Stat)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return stat, err
	}
	pidStatSample1, err := proc.ParsePidStat(data)
	if err != nil {
		return stat, err
	}

	time.Sleep(time.Second * time.Duration(seconds))

	data, _ = util.ReadFileNoStat(filename)
	pidStatSample2, err := proc.ParsePidStat(data)
	if err != nil {
		return stat, err
	}
	pidStatSample2.PID = pid

	usage := p.CalculateCpuUsagePercentage(pidStatSample1.CpuTotalTimeSpent, pidStatSample2.CpuTotalTimeSpent, seconds)
	pidStatSample2.CpuUsagePercentage = usage
	return pidStatSample2, nil
}

func (p *PidCollector) CalculateCpuUsagePercentage(past, now uint, seconds int) float64 {
	totalTime := now - past
	usage := (float64(totalTime) / float64(seconds)) / float64(runtime.NumCPU()) // * 100
	return util.ParseFloat(fmt.Sprintf("%.2f", usage))
}
