package mygoroutine

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineTest() {
	//理解：goroutine 类似于线程操作，实现是go内部实现的任务调度器，而并非操作系统

	// 创建WaitGroup变量，可以直接使用，不需要进行其他初始化操作
	var wg sync.WaitGroup
	wg.Add(2)
	// 并发执行程序
	go running("tick")
	//同时执行
	go running("tags")

	wg.Wait()
	// //同时执行
	// running("xxxx")
}
func running(s string) {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println(s, times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
