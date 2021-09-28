package main

import "fmt"

const (
	_ = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 2^10 = 1024
	MB // 2^10*2^10 = 2^20 = 1048576
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isOperation
	isDevOps
	canSeeNA
	canSeeSA
	canSeeAsia
)

func main()  {
	// Example 1 - Disk File Size
	fmt.Printf("KB is %v, type is %T\n", KB, KB)
	fmt.Printf("MB is %v, type is %T\n", MB, MB)
	// KB is 1024, type is int
	// MB is 1048576, type is int

	fileSize := 4000000000.
	fmt.Printf("File Size is %.2f GB\n", fileSize/GB)
	// File Size is 3.73 GB

	// Example 2 - Admin Privilege
	var roles byte = isAdmin | canSeeNA | canSeeAsia
	fmt.Printf("int value is %v, binary is %b\n",roles, roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is canSeeSA? %v\n", canSeeSA & roles == canSeeSA)
	// int value is 41, binary is 101001
	// Is Admin? true
	// Is canSeeSA? false


}