package main

import "unsafe"

func main() {
	var b byte = 11
	var r rune = 22

	println(unsafe.Sizeof(b))
	println(unsafe.Sizeof(r))
}
