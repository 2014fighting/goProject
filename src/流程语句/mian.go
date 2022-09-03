package main

import (
	"fmt"
)

func main() {
	//condinction()
	//switchCase()

	forSentence()
}

func condinction() {
	//基本用法跟java  c# 差不多
	var length = 1000

	if length < 100 {
		fmt.Println("长度小于100")
	} else {
		fmt.Println("大于等于100")
	}
	fmt.Println(&length)
	//if变体
	if length := 10; length < 100 {

		//这里注意在if变体中 新定义了一个length，所以只能在if作用域中生效，
		//跟外层的 length不是同一个，通过打印内存地址可以看出
		fmt.Println(length)
		fmt.Println(&length)
	}
	fmt.Println(length)
}

func switchCase() {

	//基本用法跟 c# java 类似
	age := 30
	switch age {
	case 18:
		fmt.Println("刚刚成年")
	case 30:
		fmt.Println("已经而立之年")
	default:
		fmt.Printf("当前年龄是 %d", age)
	}
}

func forSentence() {

	//基本用法 和java  c#  类似  其中  break  和 continue 以及goto 的用法也是一样的
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)

		//time.Sleep(time.Second * 1)
	}

	//go中 没有while  这点和c# java 不同
	number := 10
	for number < 14 {
		fmt.Printf("%d", number)
		number++
	}
	fmt.Println("for结束")

	//go中 没有 forearch  使用 range实现
	numbers := [6]int{1, 2, 3, 5}

	for i, v := range numbers {
		fmt.Printf("当前值为%d 索引为%d \n", v, i)
	}

}
