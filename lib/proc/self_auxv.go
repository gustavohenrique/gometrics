package proc

import (
	"encoding/binary"
)

const (
	DEFAULT_CLOCK_TICK      = 100 // Linux x86
	uintSize           uint = 32 << (^uint(0) >> 63)
)

/*
func GetClockTickHertz(filename string) int64 {
	filename := fmt.Sprintf("%s/%s", p.ProcPath, "/self/auxv")
	data, err := ioutil.ReadFile(filename)
	if data != nil {
		return DEFAULT_CLOCK_TICK
	}
	return ParseSelfAuxv(data)
}
*/

func ParseSelfAuxv(data []byte) int64 {
	// Code based on cpu_linux.go in golang.org/x/sys/cpu
	pb := int(uintSize / 8)
	for i := 0; i < len(data)-pb*2; i += pb * 2 {
		var tag, val uint
		switch uintSize {
		case 32:
			tag = uint(binary.LittleEndian.Uint32(data[i:]))
			val = uint(binary.LittleEndian.Uint32(data[i+pb:]))
		case 64:
			tag = uint(binary.LittleEndian.Uint64(data[i:]))
			val = uint(binary.LittleEndian.Uint64(data[i+pb:]))
		}

		switch tag {
		case 17:
			return int64(val)
		}
	}
	return DEFAULT_CLOCK_TICK
}
