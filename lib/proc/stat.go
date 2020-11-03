package proc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"gometrics/lib/domain"
)

func ParseStat(data []byte) (domain.Stat, error) {
	stat := domain.Stat{}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(scanner.Text())
		if len(parts) < 2 {
			continue
		}
		var err error
		switch {
		case parts[0] == "btime":
			if stat.BootTime, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
				return stat, fmt.Errorf("couldn't parse %s (btime): %s", parts[1], err)
			}
		case parts[0] == "intr":
			if stat.IRQTotal, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
				return stat, fmt.Errorf("couldn't parse %s (intr): %s", parts[1], err)
			}
			numberedIRQs := parts[2:]
			stat.IRQ = make([]uint64, len(numberedIRQs))
			for i, count := range numberedIRQs {
				if stat.IRQ[i], err = strconv.ParseUint(count, 10, 64); err != nil {
					return stat, fmt.Errorf("couldn't parse %s (intr%d): %s", count, i, err)
				}
			}
		case parts[0] == "ctxt":
			if stat.ContextSwitches, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
				return stat, fmt.Errorf("couldn't parse %s (ctxt): %s", parts[1], err)
			}
		case parts[0] == "processes":
			if stat.ProcessCreated, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
				return stat, fmt.Errorf("couldn't parse %s (processes): %s", parts[1], err)
			}
		case parts[0] == "procs_running":
			if stat.ProcessesRunning, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
				return stat, fmt.Errorf("couldn't parse %s (procs_running): %s", parts[1], err)
			}
		case parts[0] == "procs_blocked":
			if stat.ProcessesBlocked, err = strconv.ParseUint(parts[1], 10, 64); err != nil {
				return stat, fmt.Errorf("couldn't parse %s (procs_blocked): %s", parts[1], err)
			}
		case strings.HasPrefix(parts[0], "cpu"):
			cpuStat, cpuID, err := parseCPUStat(line)
			if err != nil {
				return stat, err
			}
			if cpuID == -1 {
				stat.CPU = cpuStat
			} else {
				for int64(len(stat.CPUs)) <= cpuID {
					stat.CPUs = append(stat.CPUs, domain.CPUStat{})
				}
				stat.CPUs[cpuID] = cpuStat
			}
			stat.TotalTimeSpentByAllCPU += cpuStat.TotalTimeSpent
		}
	}

	if err := scanner.Err(); err != nil {
		return stat, fmt.Errorf("couldn't parse stat: %s", err)
	}

	return stat, nil
}

func parseCPUStat(line string) (domain.CPUStat, int64, error) {
	cpuStat := domain.CPUStat{}
	var cpu string

	count, err := fmt.Sscanf(line, "%s %f %f %f %f %f %f %f %f %f %f",
		&cpu,
		&cpuStat.User, &cpuStat.Nice, &cpuStat.System, &cpuStat.Idle,
		&cpuStat.Iowait, &cpuStat.IRQ, &cpuStat.SoftIRQ, &cpuStat.Steal,
		&cpuStat.Guest, &cpuStat.GuestNice)

	if (err != nil && err != io.EOF) || count == 0 {
		return cpuStat, -1, fmt.Errorf("couldn't parse %s (cpu): %s", line, err)
	}

	cpuStat.TotalTimeSpent = cpuStat.User + cpuStat.Nice + cpuStat.System + cpuStat.IRQ + cpuStat.SoftIRQ + cpuStat.Steal
	cpuStat.TotalTimeIdle = cpuStat.Idle + cpuStat.Iowait

	if cpu == "cpu" {
		return cpuStat, -1, nil
	}

	cpuID, err := strconv.ParseInt(cpu[3:], 10, 64)
	if err != nil {
		return cpuStat, -1, fmt.Errorf("couldn't parse %s (cpu/cpuid): %s", line, err)
	}

	return cpuStat, cpuID, nil
}
