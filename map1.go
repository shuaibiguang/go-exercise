package main

import "fmt"

func main() {
	list1 := map[string]string{"a":"abc", "b": "abc"}
	list1["c"] = "zpc"

	list2 := make(map[string]string)

	list2["a"] = "a"
	list2["b"] = "b"
	// list2["C"] = "C"

	fmt.Println(list2)

	data, ok := list2["b"]

	if (ok) {
		fmt.Println(data)
	} else {
		fmt.Println("没有找到")
	}

	for k,v := range list1 {
		fmt.Println(k,v)
	}
	delete(list2, "a")
	
	fmt.Println(list2)
}