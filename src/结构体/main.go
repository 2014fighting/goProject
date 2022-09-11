package main

import (
	"fmt"
)

func main() {
	//func1()
	func2()
}

// 定义一个结构体
type books struct {
	id      int
	title   string
	author  string
	subject string
}

// 结构体的定义和基本使用
func func1() {

	//根据结构提声明变量 三中方式

	//1 按顺序赋值声明
	mybook1 := books{1, "三体", "张三", "科幻小说"}

	fmt.Println(mybook1)

	//2 按名称声明  类似c#中的具名参数
	mybook2 := books{id: 2, title: "红楼梦", author: "李四", subject: "言情小说"}

	fmt.Println(mybook2)

	//3 这种方式会给结构体每个变量赋上初始值
	mybook3 := new(books)

	fmt.Println(mybook3)

	mybook3.author = "曹雪芹"

	fmt.Println(mybook3)
}

// 结构体指针的使用
func func2() {

	//定义指针类型用星号 ，获取指针用&
	//var p2 *books
	//p2=&books

	p2 := &books{1, "三体", "张三", "科幻小说"}
	fmt.Println(p2.title)

	mybook1 := books{1, "三体", "张三", "科幻小说"}

	fmt.Println(mybook1.title)
}
