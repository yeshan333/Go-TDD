package main

import "fmt"

const prefix string = "Hello, "

// Hello 哈喽我咯
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("yeshan"))
}
