package main

import "fmt"

func main() {
	var numbers1 = []int{1, 2}

	var numbers2 = make([]int, len(numbers1), cap(numbers1)*2)

	copy(numbers2, numbers1)
	fmt.Println(numbers2)

	var numbers3 = []int{2, 3, 4}
	numbers3 = append(numbers3, 1)

	numbers4 := make([]int, len(numbers3), cap(numbers3))
	copy(numbers4, numbers3)

	numbers5 := make([]int, 0)
	numbers5 = append(numbers5, 1)

	fmt.Println(numbers3)
	fmt.Println(numbers4)
	fmt.Println(numbers5)
}
