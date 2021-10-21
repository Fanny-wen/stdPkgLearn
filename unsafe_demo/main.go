package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//unsafeArbitraryTypeDemo()
	unsafePointerDemo()
}

/*
Arbitrary 类型
官方导出这个类型只是出于完善文档的考虑，在其他的库和任何项目中都没有使用价值，除非程序员故意使用它。
*/
func unsafeArbitraryTypeDemo() {}

/*
Pointer 类型
实现定位欲读写的内存的基础
（1）任何类型的指针都可以被转化为 Pointer
（2）Pointer 可以被转化为任何类型的指针
（3）uintptr 可以被转化为 Pointer
（4）Pointer 可以被转化为 uintptr
*/
func unsafePointerDemo() {
	i := 100
	fmt.Println(i) // 100
	p := (*int)(unsafe.Pointer(&i))
	fmt.Println(*p) // 100
	*p = 0
	fmt.Println(i)  // 0
	fmt.Println(*p) // 0

	fmt.Println(unsafe.Sizeof(p)) // 8
}
