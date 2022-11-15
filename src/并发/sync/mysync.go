package mysync

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	count    int32
	wg       sync.WaitGroup
	counter  int64
	shutdown int64
)

// https://blog.csdn.net/weixin_39616056/article/details/110385378
// http://c.biancheng.net/view/4358.html
// 解决这个问题golang中有两种方案
// 1.使用原子函数(c# 中也有类似的 Atomic类的使用，目前支持一些变量的自增自减实用性不强)
// 2 使用互斥锁(类似c#中的lock)
func SyncTest() {

	// wg.Add(2)
	// go incCount()
	// go incCount()
	// wg.Wait()
	// fmt.Println(count)

	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	//fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

// 模拟享资源竞争的问题
func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched() //让当前 goroutine 暂停的意思，退回执行队列，让其他等待的 goroutine 运行
		value++
		count = value
	}
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) //安全的对counter加1
		runtime.Gosched()
	}
}

// 原子函数操作 sync包下面的 Atomic
func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}

	fmt.Printf("666 %s Down\n", name)
}
