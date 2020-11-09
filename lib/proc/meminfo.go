package proc

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/gustavohenrique/gometrics/lib/domain"
)

func ParseMemoryInfo(data []byte) (domain.MemoryInfo, error) {
	var m domain.MemoryInfo
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			return m, fmt.Errorf("malformed meminfo line: %q", scanner.Text())
		}
		v, err := strconv.ParseUint(fields[1], 0, 64)
		if err != nil {
			return m, err
		}

		switch fields[0] {
		case "MemTotal:":
			m.MemTotal = v
		case "MemFree:":
			m.MemFree = v
		case "MemAvailable:":
			m.MemAvailable = v
		case "Buffers:":
			m.Buffers = v
		case "Cached:":
			m.Cached = v
		case "SwapCached:":
			m.SwapCached = v
		case "Active:":
			m.Active = v
		case "Inactive:":
			m.Inactive = v
		case "Active(anon):":
			m.ActiveAnon = v
		case "Inactive(anon):":
			m.InactiveAnon = v
		case "Active(file):":
			m.ActiveFile = v
		case "Inactive(file):":
			m.InactiveFile = v
		case "Unevictable:":
			m.Unevictable = v
		case "Mlocked:":
			m.Mlocked = v
		case "SwapTotal:":
			m.SwapTotal = v
		case "SwapFree:":
			m.SwapFree = v
		case "Dirty:":
			m.Dirty = v
		case "Writeback:":
			m.Writeback = v
		case "AnonPages:":
			m.AnonPages = v
		case "Mapped:":
			m.Mapped = v
		case "Shmem:":
			m.Shmem = v
		case "Slab:":
			m.Slab = v
		case "SReclaimable:":
			m.SReclaimable = v
		case "SUnreclaim:":
			m.SUnreclaim = v
		case "KernelStack:":
			m.KernelStack = v
		case "PageTables:":
			m.PageTables = v
		case "NFS_Unstable:":
			m.NFSUnstable = v
		case "Bounce:":
			m.Bounce = v
		case "WritebackTmp:":
			m.WritebackTmp = v
		case "CommitLimit:":
			m.CommitLimit = v
		case "Committed_AS:":
			m.CommittedAS = v
		case "VmallocTotal:":
			m.VmallocTotal = v
		case "VmallocUsed:":
			m.VmallocUsed = v
		case "VmallocChunk:":
			m.VmallocChunk = v
		case "HardwareCorrupted:":
			m.HardwareCorrupted = v
		case "AnonHugePages:":
			m.AnonHugePages = v
		case "ShmemHugePages:":
			m.ShmemHugePages = v
		case "ShmemPmdMapped:":
			m.ShmemPmdMapped = v
		case "CmaTotal:":
			m.CmaTotal = v
		case "CmaFree:":
			m.CmaFree = v
		case "HugePages_Total:":
			m.HugePagesTotal = v
		case "HugePages_Free:":
			m.HugePagesFree = v
		case "HugePages_Rsvd:":
			m.HugePagesRsvd = v
		case "HugePages_Surp:":
			m.HugePagesSurp = v
		case "Hugepagesize:":
			m.Hugepagesize = v
		case "DirectMap4k:":
			m.DirectMap4k = v
		case "DirectMap2M:":
			m.DirectMap2M = v
		case "DirectMap1G:":
			m.DirectMap1G = v
		}
	}
	return m, nil
}
