package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	info, err := getSysInfo()
	if err != nil {
		fmt.Println("Error getting system info:", err)
		return
	}
	fmt.Printf("%+v\n", info)
}

// 시스템정보를 가져오는 함수
func getSysInfo() (*SysInfo, error) {
	hostStat, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host info:", err)
		return &SysInfo{}, err
	}

	cpuStat, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting cpu info:", err)
		return &SysInfo{}, err
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting memory info:", err)
		return &SysInfo{}, err
	}

	diskStat, err := disk.Usage("//")
	if err != nil {
		fmt.Println("Error getting disk info:", err)
		return &SysInfo{}, err
	}

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	return info, nil
}

// 시스템정보 구조체
type SysInfo struct {
	Hostname string
	Platform string
	CPU      string
	RAM      uint64
	Disk     uint64
}
