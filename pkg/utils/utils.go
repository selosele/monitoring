package utils

// 바이트 단위를 GB 단위로 변환하는 함수
func BytesToGB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}
