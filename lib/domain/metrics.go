package domain

const LINUX = "init.scope"
const ECS = "ecs"

type Base struct {
	MemoryUsage        uint64       `json:"memory_usage_in_mib"`
	MemoryTotal        uint64       `json:"memory_total_in_mib"`
	CpuUsagePercentage float64      `json:"cpu_usage_percentage"`
	RuntimeStat        *RuntimeStat `json:"runtime,omitempty"`
}

type Process struct {
	Base
	PID        int         `json:"pid"`
	SystemStat *SystemStat `json:"system,omitempty"`
	PidStat    *PidStat    `json:"process,omitempty"`
}

type Metrics struct {
	Process
	Cgroup string `json:"cgroup"`
}
