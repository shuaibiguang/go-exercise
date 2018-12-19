package main

var c string;
var d = 4;

const (
	ABC = 1
	DEF = "heihei"
	GHT = len(DEF)
)

func main () {
	var a = 1
	var b = a
	a = 3
	println(a, b)
	println(&a)
	println(ABC, DEF, GHT)
	c := 10
	d := c<<3
	println(d)
}
