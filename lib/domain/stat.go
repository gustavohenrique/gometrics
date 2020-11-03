package domain

type CPUStat struct {
	User           float64
	Nice           float64
	System         float64
	Idle           float64
	Iowait         float64
	IRQ            float64
	SoftIRQ        float64
	Steal          float64
	Guest          float64
	GuestNice      float64
	TotalTimeSpent float64
	TotalTimeIdle  float64
}

type Stat struct {
	BootTime               uint64
	CPU                    CPUStat
	CPUs                   []CPUStat
	IRQTotal               uint64
	IRQ                    []uint64
	ContextSwitches        uint64
	ProcessCreated         uint64
	ProcessesRunning       uint64
	ProcessesBlocked       uint64
	SoftIRQTotal           uint64
	TotalTimeSpentByAllCPU float64
}

type PidStat struct {
	PID                int
	State              string
	StateName          string
	PPID               int
	PGRP               int
	Session            int
	TTY                int
	TPGID              int
	Flags              uint
	MinFlt             uint
	CMinFlt            uint
	MajFlt             uint
	CMajFlt            uint
	UTime              uint
	STime              uint
	CUTime             uint
	CSTime             uint
	Priority           int
	Nice               int
	NumThreads         int
	StartTime          uint64
	VSize              uint
	RSS                int
	CpuTotalTimeSpent  uint
	CpuUsagePercentage float64
}
