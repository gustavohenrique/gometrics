package domain

type Base struct {
	MemoryUsage        uint64       `json:"memory_usage_in_mib"`
	CpuUsagePercentage float64      `json:"cpu_usage_percentage"`
	RuntimeStat        *RuntimeStat `json:"runtime,omitempty"`
}

type ProcessInfo struct {
	Base
	PID         int         `json:"pid"`
	MemoryTotal uint64      `json:"memory_total_in_mib"`
	SystemStat  *SystemStat `json:"system,omitempty"`
	PidStat     *PidStat    `json:"process,omitempty"`
}

type DockerInfo struct {
	Base
	MemoryLimit uint64 `json:"memory_limit_in_mib"`
}
