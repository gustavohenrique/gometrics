package domain

type SysInfo struct {
	MemoryUsage        uint64  `json:"memory_usage_mib"`
	MemoryTotal        uint64  `json:"memory_total_mib"`
	NumCPU             int     `json:"cpu_quantity"`
	CpuUsagePercentage float64 `json:"cpu_usage_percentage"`
}

type GoInfo struct {
	SysInfo
	PID                     *int   `json:"pid,omitempty"`
	NumGoroutines           *int   `json:"num_goroutines,omitempty"`
	MemoryAllocatedBySystem *int64 `json:"memory_allocated_by_system,omitempty"`
}

type PidInfo struct {
	SysInfo
	PID        int    `json:"pid"`
	State      string `json:"state"`
	NumThreads int    `json:"num_threads"`
}
