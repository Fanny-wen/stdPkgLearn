package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
type stringStruct struct {
	str unsafe.Pointer
	len int
}
GO 中的 string类型一般是指向字符串字面量

字符串字面量存储位置是在虚拟内存分区的只读段上面，而不是堆或栈上

因此，GO 的 string 类型不可修改的
*/

func main() {
	//stringsCompareDemo()
	//stringsEqualFoldDemo()
	//stringsContains_ContainsAny_ContainsRuneDemo()
	//stringsCountDemo()
	//stringsFieldsDemo()
	//stringsSplitDemo()
	//stringsHasPrefix_HasSuffixDemo()
	//stringsIndexDemo()
	//stringsJoinDemo()
	//stringsRepeatDemo()
	//stringsMapDemo()
	//stringsReplaceDemo()
	//stringsLoweUpperDemo()
	//stringsTitleDemo()
	//stringsTrimDemo()
	//stringsReplacerDemo()
	//stringsReaderDemo()
	stringsBuilderDemo()
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
	fmt.Println(bar2) // [hello  world, hi, girl,h]

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

/*
字符串重复几次
*/
func stringsRepeatDemo() {
	bar := strings.Repeat("hello world", 3)
	fmt.Println("ni hao" + bar) // ni haohello worldhello worldhello world
}

/*
字符替换
*/
func stringsMapDemo() {
	// Map 函数，将 s 的每一个字符按照 mapping 的规则做映射替换，如果 mapping 返回值 <0 ，则舍弃该字符。
	//该方法只能对每一个字符做处理，但处理方式很灵活，可以方便过滤，筛选汉字等
	mapping := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z': // 大写字母转小写
			return r + 32
		case r >= 'a' && r <= 'z': // 小写字母不处理
			return r
		case unicode.Is(unicode.Han, r): // 汉字换行
			return '\n'
		}
		return -1 // 过滤所有非字母、汉字的字符
	}
	fmt.Println(strings.Map(mapping, "Hello WORLD, 你好 世界, Hello Golang"))
}

/*
字符串子串替换
*/
func stringsReplaceDemo() {
	// 用 new 替换 s 中的 old，一共替换 n 个。
	// 如果 n < 0，则不限制替换次数，即全部替换
	bar1 := strings.Replace("hello world, hello golang, ni hao", "h", "H", 2)
	fmt.Printf("%q\n", bar1) // "Hello world, Hello golang, ni hao"

	// ReplaceAll内部直接调用了函数 Replace(s, old, new , -1)
	bar2 := strings.ReplaceAll("hello world", "llo", "LLO")
	fmt.Printf("%q\n", bar2) // "heLLO world"
}

/*
大小写转换
*/
func stringsLoweUpperDemo() {
	bar1 := strings.ToUpper("hello world")
	fmt.Printf("hello world ToUpper -> %s\n", bar1) // hello world ToUpper -> HELLO WORLD
	bar2 := strings.ToLower(bar1)
	fmt.Printf("HELLO WORLD ToLower -> %s\n", bar2) // HELLO WORLD ToLower -> hello world

	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "hello world")) // HELLO WORLD
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "HELLO WORLD")) // hello world
}

/*
标题处理
*/
func stringsTitleDemo() {
	// 其中 Title 会将 s 每个单词的首字母大写，不处理该单词的后续字符
	bar1 := strings.Title("hello world")
	fmt.Println(bar1) // Hello World

	// ToTitle 将 s 的每个字母大写
	bar2 := strings.ToTitle("hello world")
	fmt.Println(bar2) // HELLO WORLD

	//ToTitleSpecial 将 s 的每个字母大写，并且会将一些特殊字母转换为其对应的特殊大写字母
	bar3 := strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir")
	fmt.Println(bar3)
}

/*
修剪
*/
func stringsTrimDemo() {
	// 将 s 左侧和右侧中匹配 cutset 中的任一字符的字符去掉
	bar1 := strings.Trim("hello world hel", "helxv")
	fmt.Printf("%q\n", bar1) // "o world "

	// 将 s 左侧的匹配 cutset 中的任一字符的字符去掉
	bar2 := strings.TrimLeft("hell1 1o world hel", "helxv")
	fmt.Printf("%q\n", bar2) // "1 1o world hel"

	// 将 s 右侧的匹配 cutset 中的任一字符的字符去掉
	bar3 := strings.TrimRight("hello world hel", "helxv")
	fmt.Printf("%q\n", bar3) // "hello world "

	// 如果 s 的前缀为 prefix 则返回去掉前缀后的 string , 否则 s 没有变化。
	bar4 := strings.TrimPrefix("hello world hel", "hel")
	fmt.Printf("%q\n", bar4) // "lo world hel"

	// 如果 s 的后缀为 suffix 则返回去掉后缀后的 string , 否则 s 没有变化。
	bar5 := strings.TrimSuffix("hello world hel", "hel")
	fmt.Printf("%q\n", bar5) // "hello world "

	// 将 s 左侧和右侧的间隔符去掉。常见间隔符包括：'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL)
	bar6 := strings.TrimSpace("\n\thello world\n\t")
	fmt.Printf("%q\n", bar6) // "hello world"
}

/*
Replacer 类型
*/
func stringsReplacerDemo() {
	// 不定参数 oldnew 是 old-new 对，即进行多个替换。如果 oldnew 长度与奇数，会导致 panic.
	replacer := strings.NewReplacer("<", "&lt", ">", "&gt")
	bar1 := replacer.Replace("This is <b>HTML</b>!")
	fmt.Printf("%q\n", bar1)

	// WriteString在替换之后将结果写入 io.Writer 中
	n, err := replacer.WriteString(os.Stdout, "<div>hello</div>\n")
	if err != nil {
		fmt.Fprintf(os.Stdout, "err: %v\n", err)
	}
	fmt.Println(n)
}

/*
Reader 类型
*/
func stringsReaderDemo() {
	reader := strings.NewReader("hello world")
	fmt.Println(reader.Size())
	_, _ = reader.WriteTo(os.Stdout)
	fmt.Println()
	fmt.Println(reader.Len())
}

/*
Builder 类型
*/
func stringsBuilderDemo() {
	builder := &strings.Builder{}
	_ = builder.WriteByte('7')
	n, _ := builder.WriteRune('夕')
	fmt.Println(n)
	n, _ = builder.Write([]byte("Hello, World"))
	fmt.Println(n)
	n, _ = builder.WriteString("你好，世界")
	fmt.Println(n)
	fmt.Println(builder.Len())
	fmt.Println(builder.Cap())
	builder.Grow(100)
	fmt.Println(builder.Len())
	fmt.Println(builder.Cap())
	fmt.Println(builder.String())
	builder.Reset()
	fmt.Println(builder.String())
}
