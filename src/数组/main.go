package main

import "fmt"

func main() {

	//myArray()
	mySlice()

	//arrraySort()
}

//数组基础
func myArray() {
	//基本使用与c#  java 类似  注意 在go中数组为值类型  而非引用类型
	//先定义在使用 注意go中的数组是固定长度的
	var myarrar1 [3]string
	myarrar1 = [3]string{"22", "33", "444"}

	for i, item := range myarrar1 {
		fmt.Printf("当前值%s 当前索引%d \n", item, i)
	}

	//类型推导初始化
	var myarray = [6]int{1, 1, 2, 3, 4, 5}

	for i, item := range myarray {
		fmt.Printf("当前值%d 当前索引%d \n", item, i)
	}

	// 根据元素的个数，设置数组的大小  ,... 固定写法 表示省略长度
	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	//二维数组
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	fmt.Println(a2)
}

//go 中的切片
func mySlice() {
	//go 中还有一个类型叫切片(slice) 是引用类型，本质是为了解决数组定长问题
	//切片可以自由的增加和减少元素

	//定义方式 1
	var myarrar1 []string
	myarrar1 = []string{"22", "33", "444"}
	fmt.Println(&myarrar1)

	//定义方式2
	var slice1 []int = make([]int, 10)
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)

	fmt.Println(slice1)
}

//排序
func arrraySort() {
	/*
	   数组的排序：
	       让数组中的元素具有一定的顺序。

	       arr :=[5]int{15,23,8,10,7}
	           升序：[7,8,10,15,23]
	           将序：[23,15,10,8,7]

	   排序算法：
	       冒泡排序，插入排序，选择排序，希尔排序，堆排序，快速排序。。。。

	   冒泡排序：（Bubble Sort）
	       依次比较两个相邻的元素，如果他们的顺序（如从大到小）就把他们交换过来。
	*/
	arr := []int{15, 23, 8, 10, 7}
	result := bubbleSort(arr)

	fmt.Print(result)
}

//冒泡排序
func bubbleSort(values []int) []int {
	flag := true
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
	return values
}
