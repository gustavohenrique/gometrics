package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/gustavohenrique/gometrics"
	"github.com/gustavohenrique/gometrics/lib/util"
)

func main() {
	var pid int
	flag.IntVar(&pid, "pid", os.Getpid(), "PID")
	flag.Parse()

	fmt.Println("PID=", pid)
	collector := gometrics.New()
	interval := 1 * time.Second
	go func() {
		for {
			info, _ := collector.GetProcessInfoByPID(pid)
			fmt.Println(util.PrettyJSON(info))
			<-time.After(interval)
			runtime.GC()
		}
	}()
	_ = make([]byte, 1*(1024*1024))
	<-time.After(time.Duration(math.MaxInt64))
}
