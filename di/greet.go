package di

import (
	"fmt"

	"bytes"
)

// Greet :
func Greet(writer *bytes.Buffer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}