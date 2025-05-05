package service

import (
	"fmt"

	"monitoring/internal/sysinfo/types"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// 시스템정보를 가져오는 함수
func GetSysInfo() (*types.SysInfo, error) {
	hostStat, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host info:", err)
		return &types.SysInfo{}, err
	}

	cpuStat, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting cpu info:", err)
		return &types.SysInfo{}, err
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting memory info:", err)
		return &types.SysInfo{}, err
	}

	diskStat, err := disk.Usage("//")
	if err != nil {
		fmt.Println("Error getting disk info:", err)
		return &types.SysInfo{}, err
	}

	return &types.SysInfo{
		Hostname: hostStat.Hostname,
		Platform: hostStat.Platform,
		CPU:      cpuStat[0].ModelName,
		RAM:      vmStat.Total / 1024 / 1024,
		Disk:     diskStat.Total / 1024 / 1024,
	}, nil
}
