package main

import (
	"flag"
	"fmt"
)

func commandlineArgs() {
	firstName := flag.String("firstname", "shan", "名")

	var lastName string

	flag.StringVar(&lastName, "lastname", "ye", "姓")

	flag.Parse() // 解析命令参数

	fmt.Printf("姓：%v，名：%v", *firstName, lastName)
}
