package main

import "fmt"

// type i interface {}

type Va struct {
	a string
	b int
}

func main() {
	var i interface{}
	v := Va{"zpc", 18}

	fmt.Println(v)

	i = v

	fmt.Println(i)

	d, ok := i.(Va)

	fmt.Println(i.(type))
}
