package main

import (
	"fmt"
	"time"
)

func main() {
	//go关键字  开启gorountin 协程
	ch :=make(chan  string)
	for i:=0;i<5000 ;i++  {
		go say(i,ch)
	}

	for  {
		msg:=<- ch
		fmt.Println(msg)
	}

	time.Sleep(5*time.Second)
}

func say(i int,ch chan string)  {
	for  {
		ch<- fmt.Sprint(i)
	}
}