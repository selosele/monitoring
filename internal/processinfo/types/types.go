package types

import (
	"fmt"
	"time"
)

// 프로세스 정보 구조체
type ProcessInfo struct {
	Pid        int32     // 프로세스 ID
	Name       string    // 프로세스 이름
	Status     []string  // 프로세스 상태
	CreateTime time.Time // 프로세스 생성 시간
	NumThreads int32     // 프로세스 스레드 수
}

// 프로세스 정보 리스트 타입
type ProcessList []*ProcessInfo

// 프로세스 정보를 출력하는 함수
func (p ProcessList) Print() {
	fmt.Println("\n■ PROCESS INFO:")
	fmt.Println("pid name numThreads createTime status")
	for _, i := range p {
		fmt.Printf("%+v %+v %+v %+v %+v\n", i.Pid, i.Name, i.NumThreads, i.CreateTime, i.Status)
	}
}
