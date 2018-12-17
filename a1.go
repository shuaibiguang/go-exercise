package main

func main () {
	var arr [10]int

	for k,_ := range arr {
		arr[k] = k + 100
	}

	for k,v := range arr {
		println(k,v)
	}
}