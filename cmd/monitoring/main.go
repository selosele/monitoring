package main

import (
	"fmt"
	processinfoService "monitoring/internal/processinfo/service"
	sysinfoService "monitoring/internal/sysinfo/service"
)

func main() {
	sysInfo := sysinfoService.GetSysInfo()
	sysInfo.Print()

	fmt.Println("===============================================================")

	processInfo := processinfoService.GetProcessInfo()
	processInfo.Print()
}
