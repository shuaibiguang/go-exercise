package main

import "fmt"

func test1(sum int) (a, b int) {
	a = sum + 2
	b = a + 2
	return
}

func main() {
	fmt.Println(test1(1))
}
