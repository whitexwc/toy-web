package main

import "fmt"

func main() {
	b := []byte{34, 15}
	println(printBytes(b))
	pi := 3.1415926
	println(printNumWith2(pi))
}

// 输出两位小数
func printNumWith2(float642 float64) string {
	s := fmt.Sprintf("%.2f", float642)
	return s
}

func printBytes(data []byte) string {
	s := fmt.Sprintf("%X", data)
	return s
}
