package types

import "fmt"

// 시스템 정보 구조체
type SysInfo struct {
	Hostname    string  // 호스트 이름
	Platform    string  // 플랫폼
	CPU         string  // CPU 모델 이름
	RAM         float64 // RAM 용량 (GB 단위)
	RAMUsageOf  string  // RAM 사용량 비율 (퍼센트 단위)
	Swap        float64 // 스왑 용량 (GB 단위)
	SwapUsageOf string  // 스왑 사용량 비율 (퍼센트 단위)
	Disk        float64 // 디스크 용량 (GB 단위)
	DiskUsageOf string  // 디스크 사용량 비율 (퍼센트 단위)
}

// 시스템 정보를 출력하는 함수
func (s SysInfo) Print() {
	fmt.Println("\n■ SYSTEM INFO:")
	fmt.Printf("> hostname: %+v\n", s.Hostname)
	fmt.Printf("> platform: %+v\n", s.Platform)
	fmt.Printf("> cpu: %+v\n", s.CPU)
	fmt.Printf("> ram: %+vGB\n", s.RAM)
	fmt.Printf("> ram usage of: %+v\n", s.RAMUsageOf)
	fmt.Printf("> swap: %+vGB\n", s.Swap)
	fmt.Printf("> swap usage of: %+v\n", s.SwapUsageOf)
	fmt.Printf("> disk: %+vGB\n", s.Disk)
	fmt.Printf("> disk usage of: %+v\n", s.DiskUsageOf)
}
