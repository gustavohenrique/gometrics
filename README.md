# Gometrics (Deprecated)

[![Screenshot](https://i.imgur.com/QHcXPIQ.png)](https://gustavohenrique.github.io/gometrics)


[![Coverage Status](https://coveralls.io/repos/github/gustavohenrique/gometrics/badge.svg?branch=main)](https://coveralls.io/github/gustavohenrique/gometrics?branch=main)

> Important Notice: This project is no longer maintained and has been officially discontinued. A more up-to-date and robust version of this library is now available under the name syspector. Please consider migrating to the new library for continued support and enhanced features.

## Why was it discontinued?

Gometrics served as a great starting point for reading system metrics in Go, but it has been superseded by a more comprehensive 
and actively maintained library, [syspector](https://github.com/ravoni4devs/syspector). Syspector offers enhanced features, improved compatibility, and support for multiple
operating systems (Linux, macOS, Windows), as well as Docker environments.

## Features (in Syspector)

- More accurate and updated system resource monitoring (CPU, memory, disk usage)
- Cross-platform support for Linux, Windows, and macOS (Darwin)
- Docker container stats for memory and CPU usage
- Better support for the Go runtime and application-level metrics
- Actively maintained with regular updates
- We highly encourage all users of Gometrics to transition to syspector for continued development and new features.

## Install (for Syspector)

```sh
go get github.com/ravoni4devs/syspector
```

## Why Gometrics?

Metrics are a standard for measurement. They play an important role in understanding why your application is working in a certain way. You will need some information to find out what is happening with your application to keep it stable or help you when something goes wrong.

## About

Gometrics does not persist data or give you a dashboard. You're free to use any technology to show and persist metrics.  
There is a dashboard example written in Javascript inside the `/docs` directory.

Features:

- Calculates CPU and memory usage from the current process
- Exposes Go runtime metrics like total times GC run and total of goroutines
- Gets memory usage and limit when it's run inside a Docker container
- Uses only the Go standard lib
- Works on Linux x86_64 and Docker containers

Supported metrics' types:

- **Counter**: is a metric value which can only increase or reset, e.g., total times the Garbage Collector was run
- **Gauge**: is a number which can either go up or down, e.g., CPU usage

## Install

```sh
go get github.com/gustavohenrique/gometrics
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
    metrics, _ := collector.Metrics()  // collector.Process() or collector.Docker()
    bytes, _ := json.Marshal(metrics)
    fmt.Println(string(bytes))
}
```

## How it works?

What the `Process()` does is reads the proc file system. The proc file system (procfs or `/proc`) exposes internal kernel data structures, which is used to obtain information about the system.  
The `Docker()` reads the profcs inside a Docker container. Because the way to get metrics from a process running inside and outside a container could be a little different. Both functions can be called by `Metrics()` which auto-detects if Gometrics are running inside a container or not.

Also, both functions exposes metrics from Go runtime.

The calculation of CPU utilization is based on the total available utilization of all cores. Gometrics picking sample the total process time, every second by default, and find the difference.

Memory usage is got from procfs. When Gometrics is running outside of a container, it gives [Proportional set size](https://en.wikipedia.org/wiki/Proportional_set_size).

**Amazon ECS**

Gometrics read data from `/proc/self/cgroup` and it can be different in AWS ECS. You can disable the ECS hierarchy by setting `ECS_ENABLE_TASK_CPU_MEM_LIMIT=false` to revert the `/proc/self/cgroup` output to use the "normal" Docker output.  
Also, you can set the ECS CPU units before call `Metrics()`:

```go
collector := gometrics.New()
collector.EcsCpuTaskUnit = 2048  // 2 vCPUs
metrics, _ := collector.Metrics()
```

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
