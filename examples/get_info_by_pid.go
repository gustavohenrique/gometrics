package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gustavohenrique/gometrics"
	"github.com/gustavohenrique/gometrics/lib/util"
)

func main() {
	// pid := os.Getpid()
	var pid int
	flag.IntVar(&pid, "pid", 1, "PID")
	flag.Parse()
	collector := gometrics.New()
	info, err := collector.GetInfoByPid(pid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(util.PrettyJSON(info))
}
