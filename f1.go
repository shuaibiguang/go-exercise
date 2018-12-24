package main

func main() {
	var a int = 100
	var b int = 200

	var ret int = max(a, b)
	println(ret)

	var x string = "a"
	var y string = "b"

	var n, m = swap(x, y)
	println(n, m)

	swap2(&a, &b)
	println(a, b)
}
func swap2(num1, num2 *int) {
	*num1++
	*num2++
}

func swap(x, y string) (string, string) {
	return y, x
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
