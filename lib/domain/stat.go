package domain

type CpuStat struct {
	User           float64 `json:"-"`
	Nice           float64 `json:"-"`
	System         float64 `json:"-"`
	Idle           float64 `json:"-"`
	Iowait         float64 `json:"-"`
	IRQ            float64 `json:"-"`
	SoftIRQ        float64 `json:"-"`
	Steal          float64 `json:"-"`
	Guest          float64 `json:"-"`
	GuestNice      float64 `json:"-"`
	TotalTimeSpent float64 `json:"-"`
	TotalTimeIdle  float64 `json:"-"`
}

type Stat struct {
	BootTime               uint64    `json:"-"`
	CPU                    CpuStat   `json:"-"`
	CPUs                   []CpuStat `json:"-"`
	IRQTotal               uint64    `json:"-"`
	IRQ                    []uint64  `json:"-"`
	ContextSwitches        uint64    `json:"-"`
	ProcessCreated         uint64    `json:"-"`
	ProcessesRunning       uint64    `json:"-"`
	ProcessesBlocked       uint64    `json:"-"`
	SoftIRQTotal           uint64    `json:"-"`
	TotalTimeSpentByAllCPU float64   `json:"-"`
}

type SystemStat struct {
	Stat
	MemoryStat
	Uptime float64 `json:"uptime"`
	NumCPU int     `json:"num_cpu"`
}

type PidStat struct {
	PID                     int     `json:"pid"`
	State                   string  `json:"state"`
	StateName               string  `json:"state_name"`
	PPID                    int     `json:"-"`
	PGRP                    int     `json:"-"`
	Session                 int     `json:"-"`
	TTY                     int     `json:"-"`
	TPGID                   int     `json:"-"`
	Flags                   uint    `json:"-"`
	MinFlt                  uint    `json:"-"`
	CMinFlt                 uint    `json:"-"`
	MajFlt                  uint    `json:"-"`
	CMajFlt                 uint    `json:"-"`
	UTime                   uint    `json:"utime"`
	STime                   uint    `json:"ctime"`
	CUTime                  uint    `json:"cutime"`
	CSTime                  uint    `json:"cstime"`
	Priority                int     `json:"-"`
	Nice                    int     `json:"-"`
	NumThreads              int     `json:"num_threads"`
	StartTime               uint64  `json:"-"`
	VSize                   uint    `json:"vsize"`
	RSS                     int     `json:"rss"`
	CpuTotalTimeSpent       uint    `json:"cpu_total_time_spent"`
	CpuUsagePercentage      float64 `json:"cpu_usage_percentage"`
	ProportionalMemoryUsage uint64  `json:"proportional_memory_usage"`
	MemoryUsage             uint64  `json:"memory_usage"`
}

type DockerStat struct {
	MemoryUsage        uint64  `json:"memory_usage"`
	MemoryLimit        uint64  `json:"memory_limit"`
	CpuUsagePercentage float64 `json:"cpu_usage_percentage"`
}

type RuntimeStat struct {
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

type MemoryStat struct {
	MemTotal          uint64 `json:"mem_total"`
	MemFree           uint64 `json:"mem_free"`
	MemAvailable      uint64 `json:"mem_available"`
	Buffers           uint64 `json:"buffers"`
	Cached            uint64 `json:"cached"`
	SwapCached        uint64 `json:"-"`
	Active            uint64 `json:"-"`
	Inactive          uint64 `json:"-"`
	ActiveAnon        uint64 `json:"-"`
	InactiveAnon      uint64 `json:"-"`
	ActiveFile        uint64 `json:"-"`
	InactiveFile      uint64 `json:"-"`
	Unevictable       uint64 `json:"-"`
	Mlocked           uint64 `json:"-"`
	SwapTotal         uint64 `json:"swap_total"`
	SwapFree          uint64 `json:"swap_free"`
	Dirty             uint64 `json:"-"`
	Writeback         uint64 `json:"-"`
	AnonPages         uint64 `json:"-"`
	Mapped            uint64 `json:"-"`
	Shmem             uint64 `json:"-"`
	Slab              uint64 `json:"-"`
	SReclaimable      uint64 `json:"-"`
	SUnreclaim        uint64 `json:"-"`
	KernelStack       uint64 `json:"-"`
	PageTables        uint64 `json:"-"`
	NFSUnstable       uint64 `json:"-"`
	Bounce            uint64 `json:"-"`
	WritebackTmp      uint64 `json:"-"`
	CommitLimit       uint64 `json:"-"`
	CommittedAS       uint64 `json:"-"`
	VmallocTotal      uint64 `json:"-"`
	VmallocUsed       uint64 `json:"-"`
	VmallocChunk      uint64 `json:"-"`
	HardwareCorrupted uint64 `json:"-"`
	AnonHugePages     uint64 `json:"-"`
	ShmemHugePages    uint64 `json:"-"`
	ShmemPmdMapped    uint64 `json:"-"`
	CmaTotal          uint64 `json:"-"`
	CmaFree           uint64 `json:"-"`
	HugePagesTotal    uint64 `json:"-"`
	HugePagesFree     uint64 `json:"-"`
	HugePagesRsvd     uint64 `json:"-"`
	HugePagesSurp     uint64 `json:"-"`
	Hugepagesize      uint64 `json:"-"`
	DirectMap4k       uint64 `json:"-"`
	DirectMap2M       uint64 `json:"-"`
	DirectMap1G       uint64 `json:"-"`
}
