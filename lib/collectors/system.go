package collectors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gustavohenrique/gometrics/lib/domain"
	"github.com/gustavohenrique/gometrics/lib/proc"
	"github.com/gustavohenrique/gometrics/lib/util"
)

type SystemCollector struct {
	Proc    string
	Uptime  string
	Meminfo string
}

func NewSystemCollector() *SystemCollector {
	return &SystemCollector{
		Proc:    "/proc",
		Uptime:  "uptime",
		Meminfo: "meminfo",
	}
}

func (c *SystemCollector) GetStat() (domain.SystemStat, error) {
	var systemStat domain.SystemStat
	memoryStat, err := c.GetMemoryStat()
	if err != nil {
		return systemStat, err
	}
	systemStat.MemoryStat = memoryStat
	systemStat.NumCPU = c.GetNumCPU()
	systemStat.Uptime = c.GetUptime()
	return systemStat, nil
}

func (c *SystemCollector) GetNumCPU() int {
	return runtime.NumCPU()
}

func (c *SystemCollector) GetUptime() float64 {
	filename := fmt.Sprintf("%s/%s", c.Proc, c.Uptime)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return 0
	}
	return util.ParseFloat(strings.Split(string(data), " ")[0])
}

func (c *SystemCollector) GetMemoryStat() (domain.MemoryStat, error) {
	filename := fmt.Sprintf("%s/%s", c.Proc, c.Meminfo)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return domain.MemoryStat{}, err
	}
	return proc.ParseMemoryStat(data)
}
