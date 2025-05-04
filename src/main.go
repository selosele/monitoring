package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	info := getSysInfo()
	fmt.Printf("%+v\n", info)
}

// 시스템 정보를 가져오는 함수
func getSysInfo() *SysInfo {
	hostStat, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host info:", err)
		return &SysInfo{}
	}

	cpuStat, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting cpu info:", err)
		return &SysInfo{}
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting memory info:", err)
		return &SysInfo{}
	}

	diskStat, err := disk.Usage("//")
	if err != nil {
		fmt.Println("Error getting disk info:", err)
		return &SysInfo{}
	}

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	return info
}

// 시스템정보 구조체
type SysInfo struct {
	Hostname string `bson:"hostname"`
	Platform string `bson:"platform"`
	CPU      string `bson:"cpu"`
	RAM      uint64 `bson:"ram"`
	Disk     uint64 `bson:"disk"`
}
