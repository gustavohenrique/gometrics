package domain

type SysInfo struct {
	PID                int     `json:"pid"`
	MemoryUsage        uint64  `json:"memory_usage_mib"`
	MemoryTotal        uint64  `json:"memory_total_mib"`
	NumCPU             int     `json:"cpu_quantity"`
	CpuUsagePercentage float64 `json:"cpu_usage_percentage"`
}

type PidInfo struct {
	SysInfo
	State      string `json:"state"`
	NumThreads int    `json:"num_threads"`
}

type GoInfo struct {
	PidInfo
	MemoryMaxRss int64  `json:"memory_max_rss"`
	NumGoroutine int    `json:"num_goroutine"`
	Alloc        uint64 `json:"alloc"`
	TotalAlloc   uint64 `json:"total_alloc"`
	Sys          uint64 `json:"sys"`
	Mallocs      uint64 `json:"mallocs"`
	Frees        uint64 `json:"free"`
	LiveObjects  uint64 `json:"live_objects"`
	PauseTotalNs uint64 `json:"pause_total_ns"`
	NumGC        uint32 `json:"num_gc"`
}

type DockerInfo struct {
	MemoryUsage        uint64  `json:"memory_usage_mib"`
	MemoryLimit        uint64  `json:"memory_limit_mib"`
	CpuUsagePercentage float64 `json:"cpu_usage_percentage"`
}
