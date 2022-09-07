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
	var sliceMy []int
	sliceMy = []int{4, 5, 7}
	for _, v := range sliceMy {
		fmt.Println(v)
	}

	myfunc5(sliceMy)

	for _, v := range sliceMy {
		fmt.Println(v)
	}
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
