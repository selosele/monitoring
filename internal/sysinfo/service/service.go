package service

import (
	"fmt"

	"monitoring/internal/sysinfo/types"
	"monitoring/pkg/utils"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// 시스템정보를 가져오는 함수
func GetInfo() *types.SysInfo {
	hostStat, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host info:", err)
		return &types.SysInfo{}
	}

	cpuStat, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting cpu info:", err)
		return &types.SysInfo{}
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting virtual memory info:", err)
		return &types.SysInfo{}
	}

	smStat, err := mem.SwapMemory()
	if err != nil {
		fmt.Println("Error getting swap memory info:", err)
		return &types.SysInfo{}
	}

	diskStat, err := disk.Usage("//")
	if err != nil {
		fmt.Println("Error getting disk info:", err)
		return &types.SysInfo{}
	}

	sysInfo := &types.SysInfo{
		Hostname:    hostStat.Hostname,
		Platform:    hostStat.Platform,
		CPU:         cpuStat[0].ModelName,
		RAM:         utils.BytesToGB(vmStat.Total),
		RAMUsageOf:  fmt.Sprintf("%.2f%%", vmStat.UsedPercent),
		Swap:        utils.BytesToGB(smStat.Total),
		SwapUsageOf: fmt.Sprintf("%.2f%%", smStat.UsedPercent),
		Disk:        utils.BytesToGB(diskStat.Total),
		DiskUsageOf: fmt.Sprintf("%.2f%%", diskStat.UsedPercent),
	}

	return sysInfo
}
