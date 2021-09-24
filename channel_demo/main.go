package main

import (
	"fmt"
	myPool "github.com/stdPkgLearn/channel_demo/pool"
)

func init() {

}

var chanInt = make(chan int, 100)

func main() {
	//chanBaseDemo()
	//chanOperationDemo()
	//chanUseDemo()
	//oneDirectionChanDemo()
	myPool.Pool()
}

func chanBaseDemo() {
	fmt.Printf("chanInt type is: %T\n", chanInt)
	fmt.Printf("chanInt value is: %v\n", chanInt)
	fmt.Printf("chanInt ptr is: %p\n", chanInt)
	fmt.Printf("chanInt buf is: %d\n", len(chanInt))

	fmt.Println("=========================================================")
	// 声明的通道后需要使用make函数初始化之后才能使用。
	chanSlice := make(chan []int, 100)
	fmt.Printf("chanSlice buf is: %d\n", cap(chanSlice))
}

func chanOperationDemo() {
	// 通道有 发送(send)、接收(receive)、关闭(close)
	// 发送和接收都使用 <- 符号
	fmt.Printf("chanInt len is: %v\n", len(chanInt))
	chanInt <- 1
	fmt.Printf("chanInt len is: %v\n", len(chanInt))

	int1 := <-chanInt
	fmt.Printf("int1: %d\n", int1)
	fmt.Printf("chanInt len is: %v\n", len(chanInt))

	// 关闭
	close(chanInt)
}

// 优雅的从通道循环取值
func chanUseDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将数据发送到ch1中
	go func() {
		defer close(ch1)
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		defer close(ch2)
		for i := range ch1 {
			ch2 <- i * i
			// 有两种方式在接收值的时候判断通道是否被关闭，通常使用的是for range的方式

			//i, ok := <-ch1
			//if !ok {
			//	break
			//}
			//ch2 <- i * i
		}
	}()

	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

// 单向通道
func oneDirectionChanDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func counter(out chan<- int) {
	defer close(out)
	for i := 1; i < 100; i++ {
		out <- i
	}
}

func squarer(out chan<- int, in <-chan int) {
	defer close(out)
	for i := range in {
		out <- i * i
	}
}

func printer(in <- chan int){
	for i:= range in{
		fmt.Println(i)
	}
}
