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
	rwMutexDemo()
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
