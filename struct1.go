package main

import "fmt"

type Man struct {
	age int
	name string
}

func (man *Man) setName (name string) {
	man.name = name
}

func main() {
	var man = Man{age: 18, name: "zpc"}

	man.setName("zxc")

	fmt.Println(man.name)
}