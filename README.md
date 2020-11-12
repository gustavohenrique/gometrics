gometrics
===

[![Coverage Status](https://coveralls.io/repos/github/gustavohenrique/gometrics/badge.svg?branch=main)](https://coveralls.io/github/gustavohenrique/gometrics?branch=main)

Gometrics is a pure Go lib with no third-party dependencies to read CPU and memory info from Linux Kernel via procfs.  
It calculates the CPU and memory usage and exposes metrics from Go runtime.

## Features

- Calculates CPU and memory usage from the current process
- Exposes Go runtime metrics like total times GC run
- Gets memory usage and memory limit when it's run inside a Docker container
- Uses only the Go standard lib
- Works only on Linux x86

## Install

```sh
go get github.com/gustavohenrique/gometrics@
```

## Usage

```go
package main
import (
    "fmt"
    "github.com/gustavohenrique/gometrics"
)

func main() {
    collector := gometrics.New()
    info, _ := collector.GetRuntimeInfo() // or collector.GetDockerInfo()
    bytes, _ := json.Marshal(info)
    fmt.Println(string(bytes))
}
```

## How it works?

What the `GetRuntimeInfo()` does is reads the proc file system. The proc file system (procfs or `/proc`) exposes internal kernel data structures, which is used to obtain information about the system.  
The `GetDockerInfo()` reads the profcs inside a Docker container. Because the way to get metrics from a process running inside and outside a container could be a little different.

Also, both functions get metrics from Go runtime.

## License

Copyright 2020 Gustavo Henrique

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
