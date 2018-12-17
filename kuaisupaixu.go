package main

import "fmt"
import "time"
import "math/rand"

func main()  {
	var array []int
	for i := 0; i<100000; i++ {
		array = append(array, rand.Intn(100))
	}
	fmt.Println(time.Now())
	array = quickSort(array)
	fmt.Println(time.Now())
}

func quickSort(array []int) []int {
	if (len(array) <= 1) {return array}

	base_value := array[0]
	array = array[1:]

	var left_array []int
	var right_array []int

	for _,v := range array {
		if (v > base_value) {
			right_array = append(right_array, v)
		} else {
			left_array = append(left_array, v)
		}
	}

	left_array = quickSort(left_array)
	right_array = quickSort(right_array)

	return func (left, right []int, base_value int) []int {
		left = append(left, base_value)
		for _, v := range right {
			left = append(left, v)
		}
		return left
	}(left_array, right_array, base_value)
}
