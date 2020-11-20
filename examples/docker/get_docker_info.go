package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/gustavohenrique/gometrics"
	"github.com/gustavohenrique/gometrics/lib/util"
)

func main() {
	collector := gometrics.New()
	interval := 1 * time.Second
	go func() {
		for {
			metrics, _ := collector.Docker()
			fmt.Println(util.PrettyJSON(metrics))
			<-time.After(interval)
			runtime.GC()
		}
	}()
	_ = make([]byte, 1*(1024*1024))
	<-time.After(time.Duration(math.MaxInt64))
}
