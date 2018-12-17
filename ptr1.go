package main

const MAX int = 3

func main() {
	var a = []int{1,2,3}
	var ptr [MAX]*int
	
	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	// for i := 0; i < MAX; i++ {
	// 	println(ptr[i])
	// }

	for _,v := range ptr {
		println(*v)
	}
}