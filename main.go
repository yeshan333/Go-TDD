package main

import (
	"fmt"

	"github.com/yeshan333/Go-TDD/mypackage"
)

func main() {
	fmt.Println("Using package")
	// sayHello() // go run main.go hello.go
	mypackage.A()
}
