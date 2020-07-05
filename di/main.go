package di

import (
	"fmt"
	"io"
	"net/http"
)

// Greeter : 注入剥离
func Greeter(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler : 依赖注入
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greeter(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
