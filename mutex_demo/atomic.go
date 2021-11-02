package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//addTDemo()
	//CompareAndSwapTDemo()
	//SwapTDemo()
	//LoadTDemo()
	//StoreTDemo()
	ValueDemo()
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
func SwapTDemo() {
	var count1, count2 int64 = int64(123), int64(245)
	old := atomic.SwapInt64(&count1, count2)
	fmt.Printf("count1's new value: %d\n", count1) // count1's new value: 245
	fmt.Printf("count2's value: %d\n", count2)     // count2's value: 245
	fmt.Printf("old: %d\n", old)                   // old: 123
}

/*
	LoadT 系列获取值
*/
func LoadTDemo() {
	var count int64 = int64(123)
	v := atomic.LoadInt64(&count)
	fmt.Printf("v's value: %d\n", v) // v's value: 123
}

/*
	StoreT 系列更新值
*/
func StoreTDemo() {
	var count int64 = int64(123)
	atomic.StoreInt64(&count, 355)
	fmt.Printf("count's new value:%d\n", count) // count's new value:355
}

/*
	Value 存储器，支持Load,Store
*/
func ValueDemo() {
	var v atomic.Value
	v.Store(123)
	x := v.Load()
	fmt.Printf("x's type: %T, x's value: %[1]v\n", x) // x's type: int, x's value: 123
}
