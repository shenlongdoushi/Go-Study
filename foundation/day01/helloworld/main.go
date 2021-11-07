package main

import "fmt"

const (
	_ = iota
	KB = 1 << (10*iota) // 2^10 = 1024
	MB = 1 << (10*iota)
	GB = 1 << (10*iota)
	TB = 1 << (10*iota)
	PB = 1 << (10*iota)
)

// Go is UTF-8 native, so supports more international characters.
// 1 byte = 8 bits
// 1 English world like 'A' = 1 byte
// 1 UTF-8 Chinese character like '中' usually takes 3 bytes

func main()  {
	name1 := "hello"
	fmt.Println(name1)
	fmt.Println(KB)
}