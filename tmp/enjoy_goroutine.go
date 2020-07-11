package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("another goroutine!!")
}

func main() {
	fmt.Println("Hi")

	go test()
	go test()

	go func(name string) {
		fmt.Println("Hi " + name)
	}("yeshan")

	time.Sleep(time.Second)
}
