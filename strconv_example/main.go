package main

import (
	"fmt"
	"strconv"
)

func main() {
	//stringToIntDemo()
	//intToStringDemo()
	//parseDemo()
	//formatDemo()
	//quoteDemo()
	appendDemo()
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
