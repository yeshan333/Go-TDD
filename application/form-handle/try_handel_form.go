package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Host: ", r.Host)

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	_, _ = fmt.Fprintf(w, "Hello %s", r.Form.Get("name"))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.tpl")
		log.Println(t.Execute(w, nil))
	} else {
		_ = r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		if pwd := r.Form.Get("password"); pwd == "123456" {
			fmt.Fprintf(w, "欢迎登录, Hello %s", r.Form.Get("username"))
		} else {
			fmt.Fprintf(w, "密码错误")
		}
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
