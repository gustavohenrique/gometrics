gometrics
===

[![Coverage Status](https://coveralls.io/repos/github/gustavohenrique/gometrics/badge.svg?branch=main)](https://coveralls.io/github/gustavohenrique/gometrics?branch=main)

Gometrics is a pure Go lib with no third-party dependencies to read CPU and memory info from Linux Kernel via procfs.  
It calculates the CPU and memory usage and exposes metrics from Go runtime.

### Features

- Calculates CPU and memory usage from the current process
- Exposes Go runtime metrics like total times GC run
- Gets memory usage and memory limit when it's run inside a Docker container
- Uses only the Go standard lib
- Works only on Linux x86

### Install

```sh
go get github.com/gustavohenrique/gometrics
```

### Usage

```go
import "github.com/gustavohenrique/gometrics"
...
collector := gometrics.New()
info, _ := collector.GetRuntimeInfo() // or collector.GetDockerInfo()
bytes, _ := json.Marshal(info)
fmt.Println(string(bytes))
```

### License

MIT
