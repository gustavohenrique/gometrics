package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/gustavohenrique/gometrics"
	"github.com/gustavohenrique/gometrics/lib/util"
)

func main() {
	fmt.Println("PID=", os.Getpid())
	collector := gometrics.New()
	interval := 1 * time.Second
	go func() {
		for {
			info, _ := collector.GetInfoFromCurrentProc()
			fmt.Println(util.PrettyJSON(info))
			<-time.After(interval)
			runtime.GC()
		}
	}()
	_ = make([]byte, 1*(1024*1024))
	<-time.After(time.Duration(math.MaxInt64))
}
