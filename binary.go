package main

import "fmt"

const (
	_  = iota             //0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB
	GB
	TB
)

func main() {
	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}

// %#U, %#x
