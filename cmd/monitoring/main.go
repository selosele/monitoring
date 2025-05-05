package main

import (
	"fmt"
	"monitoring/internal/sysinfo/service"
)

func main() {
	info, err := service.GetSysInfo()
	if err != nil {
		fmt.Println("Error getting system info:", err)
		return
	}
	fmt.Printf("%+v\n", info)
}
