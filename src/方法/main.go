package main

import (
	"fmt"
	"strconv"
)

func main() {

	//go中的方法其实就是包含了一个接收者的函数，所以方法和函数的区别就是是否有一个接收者

	//我理解跟c#中的扩展方式类似

	//演示基本使用
	myPerson := Person{
		name: "张三",
		sex:  "男",
		age:  18,
	}

	fmt.Println(myPerson.SayHi("你好！请做个自我介绍。"))

	//演示继承和重写
	myPerson1 := Student{Person{"Mark", "女", 23}, "MIT"}
	myPerson2 := Employee{Person{"Sam", "女", 32}, "Golang Inc"}

	fmt.Println(myPerson1.SayHi("SayHi"))
	fmt.Println(myPerson2.SayHi("SayHi"))

}

// 先定义一个结构体
type Person struct {
	name string
	sex  string
	age  int
}

//再定义一个方法，这个方法的接收者是这个person的结构体

func (p *Person) SayHi(input string) (output string) {
	fmt.Println(input)
	return "大家好，我叫：" + p.name + "今年：" + strconv.Itoa(p.age)
}

//方法是支持继承的

type Student struct {
	Person //匿名字段
	school string
}
type Employee struct {
	Person  //匿名字段
	company string
}

// 重写
func (p *Student) SayHi(input string) (output string) {
	fmt.Println(input)
	return "大家好我是一个学生，我叫：" + p.name + "今年：" + strconv.Itoa(p.age)
}
