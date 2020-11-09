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

func (c *SystemCollector) GetMemoryInfo() (domain.MemoryInfo, error) {
	filename := fmt.Sprintf("%s/%s", c.Proc, c.Meminfo)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return domain.MemoryInfo{}, err
	}
	return proc.ParseMemoryInfo(data)
}
