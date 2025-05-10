package service

import (
	"fmt"
	"time"

	"monitoring/internal/processinfo/types"

	"github.com/shirou/gopsutil/v3/process"
)

// 프로세스 정보를 가져오는 함수
func GetInfo() types.ProcessList {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error getting process info:", err)
		return types.ProcessList{}
	}

	processList := types.ProcessList{}

	for _, proc := range processes {
		name, err := proc.Name()
		if err != nil {
			fmt.Println("Error getting process name:", err)
			continue
		}

		status, err := proc.Status()
		if err != nil {
			fmt.Println("Error getting process status:", err)
			continue
		}

		createTimeInt, err := proc.CreateTime()
		if err != nil {
			fmt.Println("Error getting process createTime:", err)
			continue
		}

		loc, err := time.LoadLocation("Asia/Seoul")
		if err != nil {
			fmt.Println("Error loading location:", err)
			continue
		}

		createTime := time.UnixMilli(createTimeInt).In(loc)

		processList = append(processList, &types.ProcessInfo{
			Pid:        proc.Pid,
			Name:       name,
			Status:     status,
			CreateTime: createTime,
		})
	}

	return processList
}
