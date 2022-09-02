package main

import "fmt"

func main() {
	//内置输出
	print("测试打印")
	println("println测试打印")
	println("println测试打印")

	//fmt 包输出
	fmt.Print("helloworld,go语言基础学习记录\n\n")
	fmt.Println("xxx", "yyyy", "zzz")

	fmt.Printf("这是测试格式化字符串%s\n", "参数")

	fmt.Printf("这是一个整数%d\n", 100)

	fmt.Printf("这是一个整数%f \n", 100.444)

}
