package main

import "fmt"

func main() {
	fmt.Println("指针学习")
	func1()
	//func2()
}

//打印变量地址
func func1() {

	var xx = 7
	fmt.Println(&xx) //先定义一个

	var ip *int /* 声明指针变量 */

	fmt.Println(ip) //声明了未使用就是空指针nil

	ip = &xx //把xx变量的指针赋值给ip
	fmt.Println(ip)

}

//指针变量声明与常规变量声明对比
func func2() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}
