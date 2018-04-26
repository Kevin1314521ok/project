package main

import (
	"utils"
	"fmt"
)

//外部排序  归并


func main() {
	cc := utils.Sourt(1,2,3,4,5,6,66,777,888,99999,333,332354)
	mes := utils.InMes(cc)
	for  {
		if msg,ok:=<- mes;ok {
			fmt.Println(msg)
		}else {
			break
		}

	}
}
