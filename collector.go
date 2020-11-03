package gometrics

import (
	"gometrics/lib"
	"gometrics/lib/domain"
)

type Collector struct{}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) GetSysInfoBy(pid int) domain.SysInfo {
	pidStat := lib.NewPIDStat()
	return pidStat.GetSysInfoByPID(pid)
}
