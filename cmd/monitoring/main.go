package main

import (
	processinfoService "monitoring/internal/processinfo/service"
	sysinfoService "monitoring/internal/sysinfo/service"
)

func main() {
	sysInfo := sysinfoService.GetInfo()
	sysInfo.Print()

	processInfo := processinfoService.GetInfo()
	processInfo.Print()
}
