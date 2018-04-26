package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(
		writer http.ResponseWriter, request *http.Request) {
			fmt.Println(request.FormValue("name"))
			fmt.Fprint(writer,request.FormValue("name"))
	})
	//监听端口号
	http.ListenAndServe(":8081",nil)
}
