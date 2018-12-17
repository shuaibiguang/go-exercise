package main

func main() {
	var b int = 15
	var a int
	// numbers := [6]int{1,2,3,5}
	var c string = "asdasd"

	for a:=0; a<b ; a++ {
		println(a)
	}

	println ("-------")
	for a<b {
		println (a)
		a++
	}

	for k,v := range c {
		println(k,v)
	}
}