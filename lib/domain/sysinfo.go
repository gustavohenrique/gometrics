package domain

type SysInfo struct {
	PID                     int     `json:"pid"`
	NumCPUs                 int     `json:"num_cpus"`
	NumGoroutines           int     `json:"num_goroutines"`
	MemoryAllocatedBySystem int64   `json:"memory_allocated_by_system"`
	MemoryUsageByPID        uint64  `json:"memory_usage_by_pid"`
	CpuUsagePercentage      float64 `json:"cpu_usage_percentage"`
}
