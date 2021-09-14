package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//stringsCompareDemo()
	//stringsEqualFoldDemo()
	//stringsContains_ContainsAny_ContainsRuneDemo()
	//stringsCountDemo()
	//stringsFieldsDemo()
	//stringsSplitDemo()
	//stringsHasPrefix_HasSuffixDemo()
	//stringsIndexDemo()
	stringsJoinDemo()
}

/*
字符串比较
*/
func stringsCompareDemo() {
	// 用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。
	// 不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观
	a := "gopher"
	b := "hello world"
	fmt.Println(strings.Compare(a, b)) // -1
	fmt.Println(strings.Compare(a, a)) // 0
	fmt.Println(strings.Compare(b, a)) // 1
}

func stringsEqualFoldDemo() {
	// 计算 s 与 t 忽略字母大小写后是否相等
	s1 := "GO"
	t1 := "go"
	s2 := "一"
	t2 := "壹"
	fmt.Println(strings.EqualFold(s1, t1)) // true
	fmt.Println(strings.EqualFold(s2, t2)) // false
}

/*
是否存在某个字符或子串
*/
func stringsContains_ContainsAny_ContainsRuneDemo() {
	// 子串 substr 在 s 中，返回 true
	fmt.Println(strings.Contains("hello world", "hell"))       // true
	fmt.Println(strings.Contains("hello world", "hell world")) // false
	//	chars 中任何一个 Unicode 代码点在 s 中，返回 true
	fmt.Println(strings.ContainsAny("hello world", "nihao")) // true
	fmt.Println(strings.ContainsAny("hello world", "bu"))    // false
	// Unicode 代码点 r 在 s 中，返回 true
	fmt.Println(strings.ContainsRune("hello world", 111)) // true
	fmt.Println(strings.ContainsRune("hello world", 112)) // false
}

/*
子串出现次数 ( 字符串匹配 )
*/
func stringsCountDemo() {
	// 在 Go 中，查找子串出现次数即字符串模式匹配，实现的是 Rabin-Karp 算法
	bar := strings.Count("hello world, hi girl, hello golang", "h")
	fmt.Println(bar)
	// 特别说明一下的是当 sep 为空时，Count 的返回值是：utf8.RuneCountInString(s) + 1
	fmt.Println("Golang中文社区")
	fmt.Println(strings.Count("Golang中文社区", ""))
}

/*
 字符串分割为[]string
*/
func stringsFieldsDemo() {
	// Fields 用一个或多个连续的空格分隔字符串 s，返回子字符串的数组（slice）
	// 如果字符串 s 只包含空格，则返回空列表 ([]string 的长度为 0）
	bar1 := strings.Fields("hello world")
	fmt.Println(bar1)                                                       // [hello world]
	fmt.Println(strings.Fields(" hello world \n hi girl\t hi \vgolang \f")) // [hello world hi girl hi golang]
	fmt.Println("=======================================================")
	bar2 := strings.FieldsFunc("ABC123PQR456XYZ789", func(r rune) bool {
		return unicode.IsNumber(r)
	})
	fmt.Println(bar2) // [ABC PQR XYZ]
}

func stringsSplitDemo() {
	bar1 := strings.Split("hello,world,hi,girl,h", ",")
	fmt.Println(bar1) // [hello  world  hi  girl h]
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Println("=======================================================")

	// 带 N 的方法可以通过最后一个参数 n 控制返回的结果中的 slice 中的元素个数，当 n < 0 时，返回所有的子字符串；当 n == 0 时，返回的结果是 nil；
	// 当 n > 0 时，表示返回的 slice 中最多只有 n 个元素，其中，最后一个元素不会分割
	bar2 := strings.SplitN("hello, world, hi, girl,h", ",", 2)
	fmt.Println(bar2) // hello  world, hi, girl,h]

	fmt.Println("=======================================================")

	bar3 := strings.SplitAfter("a,b,c", ",")
	fmt.Printf("%q\n", bar3) // ["a," "b," "c"]
	bar4 := strings.SplitAfterN("a,b,c,d,e,foo,j,k", ",", 7)
	fmt.Printf("%q\n", bar4) // ["a," "b," "c," "d," "e," "foo," "j,k"]
}

/*
字符串是否有某个前缀或后缀
*/
func stringsHasPrefix_HasSuffixDemo() {
	// s 中是否以 prefix 开始
	fmt.Println(strings.HasPrefix("hello world", "hel"))   // true
	fmt.Println(strings.HasPrefix("hello world", "heloo")) // false

	fmt.Println("=======================================================")

	// s 中是否以 suffix 结尾
	fmt.Println(strings.HasSuffix("hello world", "old"))  // false
	fmt.Println(strings.HasSuffix("hello world", "oold")) // false
}

/*
字符或子串在字符串中出现的位置
*/
func stringsIndexDemo() {
	// 在 s 中查找 sep 的第一次出现，返回第一次出现的索引, -1表示未出现过
	bar1 := strings.Index("hello world", "lo")
	fmt.Printf("bar1: %d\n", bar1) // 3

	// 在 s 中查找字节 c 的第一次出现，返回第一次出现的索引
	bar2 := strings.IndexByte("hello world", 'l')
	fmt.Printf("bar2: %d\n", bar2) // 3

	// chars 中任何一个 Unicode 代码点在 s 中首次出现的位置
	bar3 := strings.IndexAny("hello world, 你好, 世界", "世界")
	fmt.Printf("bar3: %d\n", bar3) // 21

	// Unicode 代码点 r 在 s 中第一次出现的位置
	bar4 := strings.IndexRune("hello world", 101) // 101: e
	fmt.Printf("bar4: %d\n", bar4)                // 1

	// 查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true
	han := func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 汉字
	}
	fmt.Println(strings.IndexFunc("Hello, world", han))
	fmt.Println(strings.IndexFunc("Hello, 世界", han))
}

/*
字符串 JOIN 操作
*/
func stringsJoinDemo() {
	bar1 := strings.Split("a,b,c,d", ",")
	fmt.Printf("%q\n", bar1) // ["a" "b" "c" "d"]
	bar2 := strings.Join(bar1, "_")
	fmt.Printf("%q\n", bar2) // "a_b_c_d"
}
