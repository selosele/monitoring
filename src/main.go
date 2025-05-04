package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type SysInfo struct {
	Hostname string `bson:"hostname"`
	Platform string `bson:"platform"`
	CPU      string `bson:"cpu"`
	RAM      uint64 `bson:"ram"`
	Disk     uint64 `bson:"disk"`
}

func main() {
	hostStat, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host info:", err)
		return
	}

	cpuStat, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting cpu info:", err)
		return
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting memory info:", err)
		return
	}

	diskStat, err := disk.Usage("//")
	if err != nil {
		fmt.Println("Error getting disk info:", err)
		return
	}

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	fmt.Printf("%+v\n", info)
}
