package main

import "fmt"

func main() {

	//变量定义方式1
	var firstName string
	firstName = "小王"
	fmt.Print(firstName)

	//变量定义方式2
	lastName := "张三"
	fmt.Print(lastName)

	//变量定义方式3
	var myName = "李四"
	fmt.Print(myName)

	//go 语言内存存储原理，只类型
	var name = "勒布朗"
	var nickName = name //只要是声明变量就会新建内存地址进行保存

	fmt.Println(name, &name)
	fmt.Println(nickName, &nickName)
	name = "勒布朗.詹姆斯" //赋值是不会改边内存存储地址的
	fmt.Println(name, &name)
	fmt.Println(nickName, &nickName)

	//验证 赋值和声明，注意number的变化，作用域内是新声明，number2 作用域内是赋值

	var number = 99
	//var number2 = 22
	fmt.Println("number", number)
	fmt.Println("number", &number)
	if true {
		number := 33
		//number2 = 88
		fmt.Println("number", number)
		fmt.Println("number", &number)
		//fmt.Println("number2", number2)
	}

	fmt.Println("number", number)
	fmt.Println("number", &number)
	//fmt.Println("number2", number2)

	const age = 33
}
