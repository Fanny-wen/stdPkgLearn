package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Datas struct {
	c0 byte
	c1 int
	c2 string
	c3 int
}

func main() {
	//unsafeArbitraryTypeDemo()
	//unsafePointerDemo()
	//unsafeSizeofDemo()
	//unsafeOffsetofDemo()
	unsafeAlignofDemo()
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

/*
string 与 []byte 的强转换
*/
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

/*
unsafe.Sizeof()
*/
func unsafeSizeofDemo() {
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int(10)))
	fmt.Println(unsafe.Sizeof(int32(190)))
	fmt.Println(unsafe.Sizeof("asong"))
	fmt.Println(unsafe.Sizeof([]int{1, 3, 4}))
}

/*
unsafe.Offsetof()
该函数返回由 v 所指示的某结构体中的字段在该结构体中的位置偏移字节数，注意，v 的表达方式必须是“ struct.filed ”形式。
*/
func unsafeOffsetofDemo() {
	d := Datas{}
	d.c3 = 13
	p := unsafe.Pointer(&d)
	fmt.Printf("%T, %v\n", p, p)
	offset := unsafe.Offsetof(d.c3)
	q := (*int)(unsafe.Pointer(uintptr(p) + offset))
	fmt.Printf("%T, %v\n", q, q)
	fmt.Println(*q) // 13
	*q = 1013
	fmt.Println(d.c3) // 1013
}

/*
unsafe.Offsetof()
获取变量的对齐值
*/
func unsafeAlignofDemo() {
	var b bool
	var i8 int8
	var i16 int16
	var i64 int64
	var f32 float32
	var s string
	var m map[string]string
	var p *int32

	fmt.Printf("bool: %d\n", unsafe.Alignof(b))
	fmt.Printf("int8: %d\n", unsafe.Alignof(i8))
	fmt.Printf("int16: %d\n", unsafe.Alignof(i16))
	fmt.Printf("int64: %d\n", unsafe.Alignof(i64))
	fmt.Printf("int32: %d\n", unsafe.Alignof(f32))
	fmt.Printf("string: %d\n", unsafe.Alignof(s))
	fmt.Printf("map[string]string: %d\n", unsafe.Alignof(m))
	fmt.Printf("*int32: %d\n", unsafe.Alignof(p))
}
