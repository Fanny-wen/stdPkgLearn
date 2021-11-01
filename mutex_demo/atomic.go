package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//addTDemo()
	CompareAndSwapTDemo()
}

/*
	AddT 系列将增量增加到源值上，并返回新值
*/
func addTDemo() {
	var count int64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
		go func() {
			fmt.Println(atomic.LoadInt64(&count))
		}()
	}
	wg.Wait()
	fmt.Printf("count: %d\n", count) // 10
}

/*
	CompareAndSwapT 系列比较两个变量的值，并进行交换
*/
func CompareAndSwapTDemo() {
	var count int64 = 111
	count = int64(124)
	swap := atomic.CompareAndSwapInt64(&count, 111, 999)
	fmt.Printf("count: %d, is swap: %t\n", count, swap) // count: 124, is swap: false

	swap = atomic.CompareAndSwapInt64(&count, 124, 888)
	fmt.Printf("count: %d, is swap: %t\n", count, swap) // count: 888, is swap: true
}

/*
	SwapT系列交换值，并返回旧值
*/
func SwapT() {

}

/*
	LoadT 系列获取值
*/
func LoadTDemo() {

}

/*
	StoreT 系列更新值
*/
func StoreTDemo() {

}

/*
	Value 存储器，支持Load,Store
*/
