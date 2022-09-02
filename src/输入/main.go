package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("请输入名称....")
	count, err := fmt.Scan(&name) //等待输入后才会继续代码往下走   Scanln 等待回车后代码往下走
	if err != nil {
		fmt.Print("没有错误信息")
	}
	fmt.Println(count, err)
	fmt.Print(name)
}
