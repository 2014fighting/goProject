package utils

import "fmt"

//测试包引用
func AtestDependent(param string) {
	fmt.Printf("调用utils中的方法%v,传入%v", "AtestDependent", param)
}
