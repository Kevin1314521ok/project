package main

import (
	"time"
	"fmt"
	"utils"
)

func main() {
	t1 := time.Now() // get current time
	utils.Parse()
	elapsed := time.Since(t1)
	fmt.Println("爬虫结束,总共耗时: ", elapsed)

}