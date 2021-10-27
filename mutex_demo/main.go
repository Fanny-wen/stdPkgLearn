package main

import (
	"fmt"
	"sync"
)

var a = 0
var wg sync.WaitGroup
var lock = &sync.Mutex{} // 互斥锁
var rwlock = new(sync.RWMutex)

/*
当声明了一个结构体指针变量var conn *MConn , 但是没有初始化 , 直接调用属性时候 , 就会出现
panic: runtime error: invalid memory address or nil pointer dereference
*/

func main() {
	//mutexDemo()
	//rwMutexDemo()
	//onceDemo()
	condDemo()
}

/*
sync.Mutex{}
互斥锁代表着当数据被加锁了之后，除了加锁的程序，其他程序不能对数据进行读操作和写操作
*/
func mutexDemo() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(a)
}
func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock() // 加锁
		a += 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

/*
sync.RWMutex{}
当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
*/
func rwMutexDemo() {
	wg.Add(3)
	go read(1)
	go read(1)
	go write()
	wg.Wait()
}
func read(i int) {
	//time.Sleep(time.Duration(i) * time.Second)
	rwlock.RLock()
	fmt.Printf("a: %d\n", a)
	//fmt.Printf("get lock, but can't read at first time.\n")
	defer rwlock.RUnlock()
	defer wg.Done()

}
func write() {
	rwlock.RLock()
	defer rwlock.RUnlock()
	defer wg.Done()
	//time.Sleep(3 * time.Second)
	a = 100
}

/*
sync.once{}
sync.Once 被用于控制变量的初始化，这个变量的读写满足如下三个条件：
1. 当且仅当第一次访问某个变量时，进行初始化（写）；
2. 变量初始化过程中，所有读都被阻塞，直到初始化完成；
3. 变量仅初始化一次，初始化完成后驻留在内存里。
*/
func onceDemo() {
	once := new(sync.Once)
	defWg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		defWg.Add(1)
		go func(i int) {
			defer defWg.Done()
			once.Do(func() {
				fmt.Printf("sync once do %d times, i: %[1]d\n", i)
			})
		}(i)
	}
	defWg.Wait()
}

/*
sync.cond{}
*/
func condDemo() {
	var clocker = new(sync.Mutex)
	var cond = sync.NewCond(clocker)
	//var cwp = sync.WaitGroup{}
	var cwp = new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		go func(x int) {
			cwp.Add(1)
			cond.L.Lock() // 获取锁
			cond.Wait()   // 等待通知  暂时阻塞
			defer cwp.Done()
			defer cond.L.Unlock() // 释放锁，不释放的话将只会有一次输出
			fmt.Println(x)
		}(i)
	}
	fmt.Println("start all")
	cond.Broadcast() // 下发给所有等待的 goroutine
	cwp.Wait()
}
