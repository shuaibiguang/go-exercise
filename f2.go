package main

type Circle struct {
	radius int
}

func (c Circle) getArea() int {
	return 3 * c.radius * c.radius
}

type Man struct {
	age  int
	name string
}

func setName(man Man, name string) Man {
	man.name = name
	return man
}

func (m Man) getAge() int {
	return m.age
}

func setAge(man Man, age int) Man {
	man.age = age
	return man
}

func (m Man) getName() string {
	return m.name
}

func main() {
	var man Man
	man = setAge(man, 18)
	man = setName(man, "zxc")
	// man.setName("zxc")

	println(man.getName())
	println(man.getAge())
}
