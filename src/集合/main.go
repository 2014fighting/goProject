package main

import "fmt"

func main() {
	myMap()
}

func myMap() {
	//go 语言中的map 也是一个引用类型 类似与c#中 Dictionary  但是由于是非泛型所以更像是c#中的Hashtable
	// go map 的底层实现也是 哈希表

	// 声明变量，默认 map 是 nil 空对象会引发异常
	var map_variable map[int]string

	/* 使用 make 函数 */
	map_variable = make(map[int]string)

	map_variable[1] = "张三"
	map_variable[2] = "李四"
	map_variable[3] = "王五"
	map_variable[4] = "麻六"
	fmt.Println(map_variable)
	for _, v := range map_variable {
		fmt.Println(v)
	}

	//判断元素是否存在
	name, ok := map_variable[1]
	if ok {
		fmt.Printf("%s 元素存在map中\n", name)
	} else {
		fmt.Println("不存在")
	}

	//打印map长度
	fmt.Println(len(map_variable))

	//删除map元素
	delete(map_variable, 2)

	for _, v := range map_variable {
		fmt.Println(v)
	}
}
