package main

// Import multiple packages
import (
	"fmt"
	"unsafe"
)

func main()  {
	// Variable has default value:
	// int = 0, string = ""
	var i int
	var j string
	var q float64 = 3.14
	fmt.Println("i=",i,"j=",j,"q=",q)
	// i= 0 j=

	// Let Go determine variable type
	num := 3.14
	fmt.Printf("num is %f\n", num)
	fmt.Printf("num type is %T\n", num)
	// num is 3.140000
	// num type is float64

	// Define multiple variables together.
	var (
		a = 100
		b = "Holiday"
		c = 3.14
		d byte
	)
	fmt.Printf("%d %s %f\n", a,b,c)
	fmt.Printf("d takes %d bytes\n", unsafe.Sizeof(d))
	// 100 Holiday 3.140000
	// d takes 1 bytes
}