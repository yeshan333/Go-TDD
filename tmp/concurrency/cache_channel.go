package main

import (
	"fmt"
	"time"
)

func consumer(ch chan int) {
	time.Sleep(time.Second * 10)
	val := <-ch
	fmt.Println(val)
}

func main() {
	ch := make(chan int, 2)

	go consumer(ch)

	ch <- 0
	ch <- 1
	// ch <- 2

	fmt.Println("I am free")

	ch <- 2 // 继续发会被阻塞

	fmt.Printf("I can not go there within 10 seconds")

	time.Sleep(time.Second)
}
