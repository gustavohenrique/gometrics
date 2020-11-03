package proc_test

import (
	"testing"

	"gometrics/lib/proc"
	"gometrics/test"
	"gometrics/test/assert"
)

func TestParseMemInfo(ts *testing.T) {
	const fake = `MemTotal:       32767112 kB
MemFree:         1434680 kB
MemAvailable:   14707732 kB
Buffers:         1111964 kB
Cached:         13726008 kB
SwapCached:            0 kB
Active:          6581532 kB
Inactive:       22941528 kB
Active(anon):     119668 kB
Inactive(anon): 16414056 kB
Active(file):    6461864 kB
Inactive(file):  6527472 kB
Unevictable:        1056 kB
Mlocked:              96 kB
SwapTotal:             0 kB
SwapFree:              0 kB
Dirty:              2440 kB
Writeback:             0 kB
AnonPages:      14682324 kB
Mapped:          3348776 kB
Shmem:           1873264 kB
KReclaimable:     754724 kB
Slab:            1097188 kB
SReclaimable:     754724 kB
SUnreclaim:       342464 kB
KernelStack:       43936 kB
PageTables:       158468 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:    16383556 kB
Committed_AS:   47961536 kB
VmallocTotal:   34359738367 kB
VmallocUsed:      111940 kB
VmallocChunk:          0 kB
Percpu:            16064 kB
HardwareCorrupted:     0 kB
AnonHugePages:         0 kB
ShmemHugePages:        0 kB
ShmemPmdMapped:        0 kB
FileHugePages:         0 kB
FilePmdMapped:         0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
Hugetlb:               0 kB
DirectMap4k:     1724364 kB
DirectMap2M:    29616128 kB
DirectMap1G:     3145728 kB`

	test.It(ts, "Should parse stats", func(t *testing.T) {
		info, err := proc.ParseMemoryInfo([]byte(fake))
		assert.Nil(t, err, "")
		assert.Equal(t, info.MemTotal, uint64(32767112))
		assert.Equal(t, info.Hugepagesize, uint64(2048))
	})
}
