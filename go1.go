package main

import "fmt"
import "time"

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(1000000000)
	}
}

func main() {
	go say("hello")
	say("word")
}
