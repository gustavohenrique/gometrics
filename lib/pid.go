package lib

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"gometrics/lib/domain"
	"gometrics/lib/util"
)

const (
	userHZ        = 100
	uintSize uint = 32 << (^uint(0) >> 63)
	RUNNING       = "R"
	SLEEPING      = "S"
	WAITING       = "D"
	ZOMBIE        = "Z"
	STOPPED       = "T"
	DEAD          = "X"
)

type PIDStat struct {
	ProcPath        string
	ProcPIDFilename string
}

func NewPIDStat() *PIDStat {
	return &PIDStat{
		ProcPath:        "/proc",
		ProcPIDFilename: "stat",
	}
}

func (p *PIDStat) GetSysInfoByPID(pid int) domain.SysInfo {
	cpus := runtime.NumCPU()
	goroutines := runtime.NumGoroutine()
	var mem syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &mem)
	memoryUsage, _ := p.getPrivateMemory(pid)
	return domain.SysInfo{
		PID:                     pid,
		NumCPUs:                 cpus,
		NumGoroutines:           goroutines,
		MemoryAllocatedBySystem: mem.Maxrss,
		MemoryUsageByPID:        memoryUsage,
		CpuUsagePercentage:      p.GetCpuUsage(pid),
	}
}

func (p *PIDStat) getPrivateMemory(pid int) (uint64, error) {
	// https://en.wikipedia.org/wiki/Proportional_set_size
	f, err := os.Open(fmt.Sprintf("%s/%d/smaps", p.ProcPath, pid))
	if err != nil {
		return 0, err
	}
	defer f.Close()

	res := uint64(0)
	pfx := []byte("Pss:")
	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Bytes()
		if bytes.HasPrefix(line, pfx) {
			var size uint64
			_, err := fmt.Sscanf(string(line[4:]), "%d", &size)
			if err != nil {
				return 0, err
			}
			res += size
		}
	}
	if err := r.Err(); err != nil {
		return 0, err
	}

	return res, nil
}

func (p *PIDStat) getClkTck() int64 {
	// Code based on cpu_linux.go in golang.org/x/sys/cpu
	filename := fmt.Sprintf("%s/%s", p.ProcPath, "/self/auxv")
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return userHZ
	}
	pb := int(uintSize / 8)
	for i := 0; i < len(buf)-pb*2; i += pb * 2 {
		var tag, val uint
		switch uintSize {
		case 32:
			tag = uint(binary.LittleEndian.Uint32(buf[i:]))
			val = uint(binary.LittleEndian.Uint32(buf[i+pb:]))
		case 64:
			tag = uint(binary.LittleEndian.Uint64(buf[i:]))
			val = uint(binary.LittleEndian.Uint64(buf[i+pb:]))
		}

		switch tag {
		case 17:
			return int64(val)
		}
	}
	return userHZ
}

func (p *PIDStat) getStatFromProcPidStat(pid int) (domain.ProcStat, error) {
	s := domain.ProcStat{}
	filename := fmt.Sprintf("%s/%d/%s", p.ProcPath, pid, p.ProcPIDFilename)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return s, err
	}
	l := bytes.Index(data, []byte("("))
	r := bytes.LastIndex(data, []byte(")"))
	if l < 0 || r < 0 {
		return s, fmt.Errorf("unexpected format, couldn't extract comm: %s", data)
	}
	var ignore int
	_, err = fmt.Fscan(
		bytes.NewBuffer(data[r+2:]),
		&s.State,
		&s.PPID,
		&s.PGRP,
		&s.Session,
		&s.TTY,
		&s.TPGID,
		&s.Flags,
		&s.MinFlt,
		&s.CMinFlt,
		&s.MajFlt,
		&s.CMajFlt,
		&s.UTime,
		&s.STime,
		&s.CUTime,
		&s.CSTime,
		&s.Priority,
		&s.Nice,
		&s.NumThreads,
		&ignore,
		&s.Starttime,
		&s.VSize,
		&s.RSS,
	)
	s.TotalTime = s.UTime + s.STime + s.CUTime + s.CSTime
	s.TotalTimeInSecs = s.TotalTime / userHZ
	return s, nil
}

func (p *PIDStat) GetCpuUsage(pid int) float64 {
	// clkTck := float64(getClkTck())
	s, err := p.getStatFromProcPidStat(pid)
	if err != nil {
		log.Println("getStatFromProcPidStat:", err)
		return 0
	}
	totalTime1 := s.TotalTime

	time.Sleep(time.Second * 1)

	s2, _ := p.getStatFromProcPidStat(pid)
	totalTime2 := s2.TotalTime

	totalTime := totalTime2 - totalTime1
	fmt.Println("totaltime1=", totalTime1)
	fmt.Println("totaltime2=", totalTime2)

	uptime := p.getFromProcUptime()
	fmt.Println("uptime=", uptime)
	seconds := 1.0                                                      //math.Abs(uptime - float64(s.Starttime/100))
	usage := (float64(totalTime) / seconds) / float64(runtime.NumCPU()) // * 100

	// stat, err := p.getFromProcStat()
	// if err != nil {
	// log.Println("getFromProcStat:", err)
	// return 0
	// }
	// uptime := float64(stat.BootTime) - float64(s.Starttime / userHZ)
	// usage := ((stat.TotalTimeByAllCPU / totalTime) / float64(uptime)) * 100
	// usage := ((totalTime / stat.TotalTimeByAllCPU) / float64(uptime)) * 100
	fmt.Println("usage=", usage)
	return util.ParseFloat(fmt.Sprintf("%.2f", usage))
}

func (p *PIDStat) getFromProcUptime() float64 {
	filename := fmt.Sprintf("%s/uptime", p.ProcPath)
	data, err := util.ReadFileNoStat(filename)
	if err != nil {
		return 0
	}
	return util.ParseFloat(strings.Split(string(data), " ")[0])
}

func (p *PIDStat) getFromProcStat() (domain.Stat, error) {
	fileName := fmt.Sprintf("%s/stat", p.ProcPath)
	data, err := util.ReadFileNoStat(fileName)
	if err != nil {
		return domain.Stat{}, err
	}

	stat := domain.Stat{}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(scanner.Text())
		if len(parts) < 2 {
			continue
		}
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
			cpuStat, cpuID, err := p.parseCPUStat(line)
			if err != nil {
				return stat, err
			}
			if cpuID == -1 {
				stat.CPUTotal = cpuStat
			} else {
				for int64(len(stat.CPU)) <= cpuID {
					stat.CPU = append(stat.CPU, domain.CPUStat{})
				}
				stat.CPU[cpuID] = cpuStat
			}
			stat.TotalTimeByAllCPU += cpuStat.TotalTimeInSecs
		}
	}

	if err := scanner.Err(); err != nil {
		return stat, fmt.Errorf("couldn't parse %s: %s", fileName, err)
	}

	return stat, nil
}

// Parse a cpu statistics line and returns the domain.CPUStat struct plus the cpu id (or -1 for the overall sum).
func (p *PIDStat) parseCPUStat(line string) (domain.CPUStat, int64, error) {
	cpuStat := domain.CPUStat{}
	var cpu string

	count, err := fmt.Sscanf(line, "%s %f %f %f %f %f %f %f %f %f %f",
		&cpu,
		&cpuStat.User, &cpuStat.Nice, &cpuStat.System, &cpuStat.Idle,
		&cpuStat.Iowait, &cpuStat.IRQ, &cpuStat.SoftIRQ, &cpuStat.Steal,
		&cpuStat.Guest, &cpuStat.GuestNice)

	if err != nil && err != io.EOF {
		return cpuStat, -1, fmt.Errorf("couldn't parse %s (cpu): %s", line, err)
	}
	if count == 0 {
		return cpuStat, -1, fmt.Errorf("couldn't parse %s (cpu): 0 elements parsed", line)
	}

	cpuStat.TotalTime = cpuStat.User + cpuStat.Nice + cpuStat.System + cpuStat.IRQ + cpuStat.SoftIRQ + cpuStat.Steal
	cpuStat.TotalTimeIdle = cpuStat.Idle + cpuStat.Iowait
	// cpuStat.TotalTime = cpuStat.User + cpuStat.Nice + cpuStat.System + cpuStat.Idle + cpuStat.Iowait + cpuStat.IRQ + cpuStat.SoftIRQ + cpuStat.Steal
	cpuStat.User /= userHZ
	cpuStat.Nice /= userHZ
	cpuStat.System /= userHZ
	cpuStat.Idle /= userHZ
	cpuStat.Iowait /= userHZ
	cpuStat.IRQ /= userHZ
	cpuStat.SoftIRQ /= userHZ
	cpuStat.Steal /= userHZ
	cpuStat.Guest /= userHZ
	cpuStat.GuestNice /= userHZ
	cpuStat.TotalTimeInSecs = cpuStat.TotalTime / userHZ

	if cpu == "cpu" {
		return cpuStat, -1, nil
	}

	cpuID, err := strconv.ParseInt(cpu[3:], 10, 64)
	if err != nil {
		return cpuStat, -1, fmt.Errorf("couldn't parse %s (cpu/cpuid): %s", line, err)
	}

	return cpuStat, cpuID, nil
}
