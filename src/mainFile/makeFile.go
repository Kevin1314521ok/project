package main

import (
	"os"
)

func main() {
	name:="C:/Users/Administrator/Desktop/新建Microsoft Excel 工作表.xlsx"

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("电影名称"+"\t"+"评分"+"\t"+"评价人数"+"\t"+"\r\n")
}
