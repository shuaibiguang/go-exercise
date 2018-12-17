package main

import "fmt"

var a = "嘿嘿"
var b string = "asdasd.com"
var c bool

func main () {
	d := []int{1,2,3}
	e := []int{4,5,6}
	// f := d + e
	f := func (e []int, d []int) []int {
		for _,v := range e {
			d = append(d, v)
		}
		return d
	}(e, d)
	// f := make([]int, 0, cap(d) + cap(e))

	// for _,v := range d {
	// 	f = append(f, v)
	// }
	// for _,v := range e {
	// 	f = append(f, v)
	// }

	fmt.Println(f)
}