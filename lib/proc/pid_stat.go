package proc

import (
	"bytes"
	"fmt"

	"gometrics/lib/domain"
)

const (
	RUNNING  = "R"
	SLEEPING = "S"
	WAITING  = "D"
	ZOMBIE   = "Z"
	STOPPED  = "T"
	DEAD     = "X"
)

func ParsePidStat(data []byte) (domain.PidStat, error) {
	pidStat := domain.PidStat{}
	l := bytes.Index(data, []byte("("))
	r := bytes.LastIndex(data, []byte(")"))
	if l < 0 || r < 0 {
		return pidStat, fmt.Errorf("Unexpected format, couldn't extract data: %s", data)
	}
	var ignoredField int
	_, err := fmt.Fscan(
		bytes.NewBuffer(data[r+2:]),
		&pidStat.State,
		&pidStat.PPID,
		&pidStat.PGRP,
		&pidStat.Session,
		&pidStat.TTY,
		&pidStat.TPGID,
		&pidStat.Flags,
		&pidStat.MinFlt,
		&pidStat.CMinFlt,
		&pidStat.MajFlt,
		&pidStat.CMajFlt,
		&pidStat.UTime,
		&pidStat.STime,
		&pidStat.CUTime,
		&pidStat.CSTime,
		&pidStat.Priority,
		&pidStat.Nice,
		&pidStat.NumThreads,
		&ignoredField,
		&pidStat.StartTime,
		&pidStat.VSize,
		&pidStat.RSS,
	)
	switch pidStat.State {
	case RUNNING:
		pidStat.StateName = "running"
	case SLEEPING:
		pidStat.StateName = "sleeping"
	case WAITING:
		pidStat.StateName = "waiting"
	case ZOMBIE:
		pidStat.StateName = "zombie"
	case STOPPED:
		pidStat.StateName = "stopped"
	case DEAD:
		pidStat.StateName = "dead"
	default:
		pidStat.StateName = "unknown"
	}
	pidStat.CpuTotalTimeSpent = pidStat.UTime + pidStat.STime + pidStat.CUTime + pidStat.CSTime
	return pidStat, err
}
