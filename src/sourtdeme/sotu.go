package main

import (
	"sort"
	"fmt"
)

//简单排序
func main() {
	//create slice of int  创建切片
	a:=[]int{3,4,6}
	sort.Ints(a)

	for i,v:= range a{
		fmt.Println(i," ",v)
	}
}
