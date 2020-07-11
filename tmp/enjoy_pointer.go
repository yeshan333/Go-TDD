package main

// NewFunc : 感受下 new 关键字
func NewFunc() *string {
	str := new(string)

	*str = "Golang is Good!"

	return str
}
