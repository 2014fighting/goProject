package main

import (
	"fmt"
	"reflect"
)

func main() {
	//myfunc1(1, "xxx")
	//myfunc2(1, 3, 4)
	//myfunc2(1, 2)

	// x := 3
	// fmt.Println("x = ", x) // 应该输出 "x = 3"
	// x1 := myfunc4(x)       // 调用 add1(&x) 传x的地址

	// fmt.Println("x1 = ", x1) // 应该输出 "x1 = 4"
	// fmt.Println("x = ", x)   // 应该输出 "x = 3"

	// var sliceMy []int
	// sliceMy = []int{4, 5, 7}
	// for _, v := range sliceMy {
	// 	fmt.Println(v)
	// }

	// myfunc5(sliceMy)

	// for _, v := range sliceMy {
	// 	fmt.Println(v)
	// }

	//boolRes := myfunc6()
	//fmt.Println(boolRes)
	//myfunc7()

	//myfunc8()
	res := myfunc9(fun1)
	fmt.Println(res)
}

// 函数基本使用
func myfunc1(parame1 int, parame2 string) (output1 int, output2 string) {

	fmt.Println(parame1)

	fmt.Println(parame2)
	//，定义返回多个值的函数  用法类似与c#中的元组类型
	//这里是处理逻辑代码
	//返回多个值
	return 1, "ok"
}

// 定义一个全是int 类型的可变参数
func myfunc2(arg ...int) {
	fmt.Println(reflect.TypeOf(arg))
	for _, v := range arg {
		fmt.Println(v)
	}

}

// 简单的一个函数，实现了参数+1的操作  引用传递，类似c#中的 ref的使用
func myfunc3(a *int) int { // 请注意，
	*a = *a + 1 // 修改了a的值
	return *a   // 返回新值
}

// 对比上面的 myfunc3  a  在外部不会发生变化 值传递
func myfunc4(a int) int { // 请注意，
	a = a + 1 // 修改了a的值
	return a  // 返回新值
}

// 验证切片作为引用类型传递的结果会发生改变
func myfunc5(arg []int) []int { // 请注意，
	arg[0] = 2 // 修改了a的值
	return arg // 返回新值
}

// defer 使用 ,后进先出原则，应该打印1 4 2   3
func myfunc6() bool {
	a := 1
	b := 2
	c := 3
	d := 4
	defer fmt.Println(c)
	fmt.Println(a)
	defer fmt.Println(b)

	fmt.Println(d)
	return true
}

// defer 堆栈打印
func myfunc7() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
	fmt.Print("\n")
	for _, v := range []rune("abckd") {
		defer fmt.Printf("%c\n", v)
	}
}

// 匿名函数
func myfunc8() {
	//匿名函数
	func() {
		fmt.Println("我是一个匿名函数。。")
	}()

	//匿名函数
	fun3 := func() {
		fmt.Println("我也是一个匿名函数。。")
	}

	fun3()

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//定义一个带参数 并且带返回值的匿名函数
	fmt.Println(func(a bool, b int) int {
		fmt.Println(a, b)
		return b
	}(false, 4))
}

/*
Go语言是支持函数式编程：
1.将匿名函数作为另一个函数的参数，回调函数
2.将匿名函数作为另一个函数的返回值，可以形成闭包结构。
*/

// 回调函数 定义了一个  你匿名函数作为参数的函数，这个参数就是回调函数
// 这点使用跟js 类似，在c#中实现是用的委托
func myfunc9(callback func(a int) int) int {
	return callback(1)
}

func fun1(a int) int {
	fmt.Printf("我是函数 fun1,传入参数%d \n", a)
	return a + 1
}
