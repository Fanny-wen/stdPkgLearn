package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 函数只能与 nil 比较
// 可变函数是指针传递。
// 不能使用短变量声明设置结构体字段值。

func main() {
	var wg = sync.WaitGroup{}
	var dCount uint64
	var fCount uint64
	var cCount uint64
	var dogCh = make(chan struct{}, 1)
	var fishCh = make(chan struct{}, 1)
	var catCh = make(chan struct{}, 1)
	wg.Add(3)
	go dog(&wg, dogCh, fishCh, dCount)
	go fish(&wg, fishCh, catCh, fCount)
	go cat(&wg, catCh, dogCh, cCount)
	dogCh <- struct{}{}
	wg.Wait()
}

func dog(wg *sync.WaitGroup, dogCh, fishCh chan struct{}, count uint64) {
	for {
		if count < uint64(100) {
			<-dogCh
			fmt.Println("dog", count)
			atomic.AddUint64(&count, 1)
			fishCh <- struct{}{}
		} else {
			wg.Done()
			return
		}
	}
}
func fish(wg *sync.WaitGroup, fishCh, catCh chan struct{}, count uint64) {
	for {
		if count < uint64(100) {
			<-fishCh
			fmt.Println("fish", count)
			catCh <- struct{}{}
			atomic.AddUint64(&count, 1)
		} else {
			wg.Done()
			return
		}
	}
}
func cat(wg *sync.WaitGroup, catCh, dogCh chan struct{}, count uint64) {
	for {
		if count < uint64(100) {
			<-catCh
			fmt.Println("cat", count)
			dogCh <- struct{}{}
			atomic.AddUint64(&count, 1)
		} else {
			wg.Done()
			return
		}
	}
}

func contextWithDeadline() {
	//d := time.Now().Add(50 * time.Millisecond)
	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func test() {
	rand.Seed(time.Now().UnixNano())
	const Max = 10
	const NumSenders = 1
	dataCh := make(chan int, 10)
	stopCh := make(chan struct{})
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case v, ok := <-stopCh:
					fmt.Printf("stop chan closed %v, %v\n", v, ok)
					fmt.Printf("len of dataCh: %d\n", len(dataCh))
					return
				case dataCh <- rand.Intn(Max):
					fmt.Printf("cap: %d--%d\n", cap(dataCh), len(dataCh))
				}
			}
		}()
	}
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}
			fmt.Println(value)
		}
	}()
	select {
	case <-time.After(time.Second):
	}
}
func FallthroughDemo(char string) bool {
	switch char {
	case " ":
		// 强制执行下一个 case 代码块
		fallthrough
	case "\t":
		return true
	}
	return false
}

func JsonDemo() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	var j, _ = json.Marshal(struct {
		Name string
		Age  int
	}{"六", 123})
	_ = json.Unmarshal(data, &result)
	fmt.Println(string(j))
	fmt.Println(result)
}

func diffNewMake() {
	// make 的作用是为 slice map chan 初始化并返回引用
	// make 返回的是 T 类型的引用, 只适用于 slice map channel
	m := make(map[string]interface{})
	fmt.Printf("%+v\n", m) // map[]
	// new 的作用是初始化一个指向类型的指针
	// new 返回的是 一个指针, 指针指向的是 新分配的、类型为T 的 零值
	n := new(map[string]interface{})
	fmt.Printf("%+v\n", n) // &map[]
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func Demo() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		fmt.Printf("%p, %+v\n", &val, val)
		m[key] = &val
		fmt.Printf("%+v\n", m)
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

// Demo2 类型别名和类型自定
func Demo2() {
	type int1 int   // 类型自定
	type int2 = int // 类型别名
	var i1 int1
	var i2 int2
	fmt.Printf("%T, %T\n", i1, i2)
}

// Demo3 字符串拼接
func Demo3() {
	a := "hello" + "world"
	b := bytes.NewBufferString("hello")
	n, _ := b.WriteString("world")
	fmt.Printf("n: %d\n", n)
	c := fmt.Sprintf("%s%s", "hello", "world")
	fmt.Println(a, b, c)
}

// Demo4 数组和切片的截取操作, 截取得到的切片长度和容量计算方法是 j-i、l-i
func Demo4() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Printf("len: %d------ cap: %d\n", len(a), cap(a)) // len: 0------ cap: 3
	fmt.Printf("len: %d------ cap: %d\n", len(b), cap(b)) // len: 2------ cap: 3
	fmt.Printf("len: %d------ cap: %d\n", len(c), cap(c)) // len: 1------ cap: 2
	fmt.Printf("len: %d------ cap: %d\n", len(s), cap(s)) // len: 3------ cap: 3
}

// Demo5 map 的value是 不可寻址的
func Demo5() {
	type Math struct {
		x, y int
	}

	var m = map[string]Math{
		"foo": Math{2, 3},
	}
	fmt.Println(m)
	//m["foo"].x = 4
}
