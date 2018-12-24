package main

import "fmt"

func main() {
	var a int = 4
	var ptr *int

	ptr = &a

	fmt.Printf("%d", a)
	println()
	fmt.Printf("%d", *ptr)
	*ptr = 5
	println(a)
}
