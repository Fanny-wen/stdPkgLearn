package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strToIntDemo()
	//intToStrDemo()
	strconvQuoteDemo()
}

func strToIntDemo() {
	bar1, _ := strconv.Atoi("123")
	fmt.Println(bar1) // 123

	// base 的取值为 2~36，如果 base 的值为 0，则会根据字符串的前缀来确定 base 的值
	// 参数 bitSize 表示的是整数取值范围，或者说整数的具体类型。取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64。
	bar2, _ := strconv.ParseInt("-1010", 2, 64) // -10
	fmt.Println(bar2)

	bar3, _ := strconv.ParseUint("1010", 2, 0) //10
	fmt.Println(bar3)
}

func intToStrDemo() {
	bar1 := strconv.Itoa(123)
	fmt.Println(bar1) // 123

	bar2 := strconv.FormatInt(-123, 10)
	fmt.Println(bar2) // -123

	bar3 := strconv.FormatUint(10, 2)
	fmt.Println(bar3) // 1010
}

func strconvQuoteDemo() {
	// Quote返回一个双引号的Go字符串字面量，表示s
	s := strconv.Quote(`I'm your father`)
	fmt.Printf("%s\n", s) // "I'm your father"
}

func stringToIntDemo() {
	// Atoi() 用于将字符串类型的整数转换为int类型

	var s1 = "1"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Printf("strconv.Atoi failed, err: %v\n", err)
		return
	}
	fmt.Printf("type: %T, value: %#v\n", i1, i1)
}

func intToStringDemo() {
	// Itoa() 用于将int类型的数据转换为对应的字符串表示

	var i1 = 123
	s1 := strconv.Itoa(i1)
	fmt.Printf("type: %T, value: %#v\n", s1, s1)
}

func parseDemo() {
	//	Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。

	b, err := strconv.ParseBool("1")
	fmt.Printf("strconv.ParseBool: %v, err: %v\n", b, err)
	i, err := strconv.ParseInt("0xf4", 0, 16) // strconv.ParseInt: -128, err: strconv.ParseInt: parsing "-1203": value out of range
	fmt.Printf("strconv.ParseInt: %v, err: %v\n", i, err)
	u, err := strconv.ParseUint("345", 10, 64)
	fmt.Printf("strconv.ParseUint: %v, err: %v\n", u, err)
	f, err := strconv.ParseFloat("3.14", 64)
	fmt.Printf("strconc.ParseFloat: %v, err: %v\n", f, err)
}

func formatDemo() {
	//	Format系列函数实现了将给定类型数据格式化为string类型数据的功能。

	b := strconv.FormatBool(true)
	fmt.Printf("strconv.FormatBool: %v, type: %T\n", b, b)
	i := strconv.FormatInt(1123, 10)
	fmt.Printf("strconv.FormatInt: %v, type: %T\n", i, i)
	u := strconv.FormatUint(1123, 10)
	fmt.Printf("strconv.FormatUint: %v, type: %T\n", u, u)
	f := strconv.FormatFloat(3.1415926, 'f', -1, 32)
	fmt.Printf("strconv.FormatFloat: %v, type: %T\n", f, f)
}

func quoteDemo() {
	b := strconv.IsPrint('\u263a')
	fmt.Printf("strconv.IsPrint: %v\n", b)
	s1 := strconv.Quote("123123\t")
	fmt.Printf("strconv.Quote: %v\n", s1)

	s2 := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	fmt.Printf("strconv.QuoteToASCII: %v\n", s2)

	s3 := strconv.QuoteRune('☺')
	fmt.Printf("strconv.QuoteRune: %v\n", s3)
}

func appendDemo() {
	b := []byte("bool:")
	fmt.Println(string(strconv.AppendBool(b, true)))
	fmt.Println(append(b, "qwedfh"...))

	fmt.Println(strconv.AppendInt(b, 1, 10))
}
