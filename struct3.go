package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v, %v, hahahah", p.name, p.age)
}

func main () {
	p1 := Person{"zpc", 18}
	str := p1.String()
	fmt.Println(str)
}