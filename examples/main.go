package main

import (
	"flag"
	"fmt"

	"gometrics"
	"gometrics/lib/util"
)

func main() {
	// pid := os.Getpid()
	var pid int
	flag.IntVar(&pid, "pid", 1234, "PID")
	flag.Parse()
	collector := gometrics.NewCollector()
	pidInfo := collector.GetSysInfoBy(pid)

	fmt.Println(util.PrettyJSON(pidInfo))
}
