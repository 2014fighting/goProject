package main

import "fmt"

func main() {
	fmt.Println("指针学习")
	//指针其实就是一种数据类型，用于表示数据内存地址
	//存在的意义：创建了一个地址的引用 便于后期通过这个应用去获取地址对应的值
	//func1()
	//func2()
	func3()
}

//指针的基本使用
func func1() {

	var xx = 7
	fmt.Println(&xx) //先定义一个

	var ip *int /* 声明指针变量 */

	fmt.Println(ip) //声明了未使用就是空指针nil

	ip = &xx //把xx变量的地址赋值给ip指针
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

//指针的应用场景
func func3() {
	// v1 := "张山"

	// v2 := v1
	// v1 = "李四"

	// fmt.Println(v1, v2) //结果：李四 张三

	v1 := "张三"

	v2 := &v1

	v1 = "李四"

	fmt.Println(v1, *v2) //结果：李四 李四

	//两个结果对比可以看出指针的作用就是保存变量的地址
}
