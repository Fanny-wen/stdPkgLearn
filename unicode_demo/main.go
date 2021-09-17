package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

func main() {
	//unicodeDemo()
	//unicodeUtf8Demo()
	//utf8DecodeDemo()
	//utf8FullRuneDemo()
	//utf8RuneStartDemo()
	utf16Demo()
}

func unicodeDemo() {
	// 是否为控制字符
	single := '\u0015'
	fmt.Println("IsControl", unicode.IsControl(single)) // true
	single = '\ufe35'
	fmt.Println("IsControl", unicode.IsControl(single)) // false
	fmt.Println("---------------")
	// 是否为阿拉伯数字字符，即 0-9
	fmt.Println("IsDigit", unicode.IsDigit('1')) // true
	fmt.Println("---------------")
	// 是否数字字符，比如罗马数字Ⅷ也是数字字符
	fmt.Println("IsNumber", unicode.IsNumber('Ⅰ')) // true
	fmt.Println("IsNumber", unicode.IsNumber('一')) // false
	fmt.Println("IsNumber", unicode.IsNumber('1')) // true
	fmt.Println("---------------")
	// 是否图形字符
	fmt.Println("IsGraphic", unicode.IsGraphic('😀')) // true
	fmt.Println("IsGraphic", unicode.IsGraphic('j')) // true
	fmt.Println("---------------")

	// 是否字母
	fmt.Println("IsLetter", unicode.IsLetter('R')) // true
	fmt.Println("IsLetter", unicode.IsLetter('r')) // true
	fmt.Println("IsLetter", unicode.IsLetter('1')) // false
	// 是否大写字符
	fmt.Println("IsUpper", unicode.IsUpper('R')) // true
	fmt.Println("IsUpper", unicode.IsUpper('r')) // false
	fmt.Println("IsUpper", unicode.IsUpper('1')) // false
	// 是否小写字母
	fmt.Println("IsLower", unicode.IsLower('R')) // false
	fmt.Println("IsLower", unicode.IsLower('r')) // true
	fmt.Println("IsLower", unicode.IsLower('1')) // false
	fmt.Println("---------------")

	// 是否符号字符
	fmt.Println("IsMark", unicode.IsMark('R')) // false
	fmt.Println("IsMark", unicode.IsMark('r')) // false
	fmt.Println("IsMark", unicode.IsMark('.')) // false
	fmt.Println("IsMark", unicode.IsMark('>')) // false
	fmt.Println("IsMark", unicode.IsMark('?')) // false
	fmt.Println("IsMark", unicode.IsMark('#')) // false
	fmt.Println("---------------")
	fmt.Println("IsSymbol", unicode.IsSymbol('R')) // false
	fmt.Println("IsSymbol", unicode.IsSymbol('r')) // false
	fmt.Println("IsSymbol", unicode.IsSymbol('.')) // false
	fmt.Println("IsSymbol", unicode.IsSymbol('>')) // true
	fmt.Println("IsSymbol", unicode.IsSymbol('?')) // false
	fmt.Println("IsSymbol", unicode.IsSymbol('#')) // false
	fmt.Println("---------------")
	// 是否标点符号
	fmt.Println("IsPunct", unicode.IsPunct('R')) // false
	fmt.Println("IsPunct", unicode.IsPunct('r')) // false
	fmt.Println("IsPunct", unicode.IsPunct('.')) // true
	fmt.Println("IsPunct", unicode.IsPunct('?')) // true
	fmt.Println("IsPunct", unicode.IsPunct('>')) // false
}

func unicodeUtf8Demo() {
	// 判断是否符合utf8编码
	word := []byte("界")
	fmt.Println("Valid", utf8.Valid(word[:2]))                                     // false
	fmt.Println("Valid", utf8.Valid(word))                                         // true
	fmt.Println("ValidRune", utf8.ValidRune('界'))                                  // true
	fmt.Println("ValidString", utf8.ValidString("世界"))                             // true
	fmt.Println("ValidString", utf8.ValidString(string([]byte{0xff, 0xfe, 0xfd}))) // false
	fmt.Println(string([]byte{0xff, 0xfe, 0xfd}))

	fmt.Println("---------------")

	// 判断rune所占字节数
	fmt.Println("RuneLen", utf8.RuneLen('h')) // 1
	fmt.Println("RuneLen", utf8.RuneLen('界')) // 3

	fmt.Println("---------------")

	// 判断字节串或者字符串的 rune 数
	fmt.Println("RuneCount", utf8.RuneCount([]byte("hello world, 你好, 世界")))         // 19
	fmt.Println("RuneCount", len([]byte("hello world, 你好, 世界")))                    // 27
	fmt.Println("RuneCountInString", utf8.RuneCountInString("hello world, 你好, 世界")) // 19
	fmt.Println("RuneCountInString", len("hello world, 你好, 世界"))                    // 27

	fmt.Println("---------------")

	// 编码和解码到 rune
	p := make([]byte, 3)
	n := utf8.EncodeRune(p, '界')
	fmt.Println(n, p) // n: 3, p: [231 149 140]
	//file, _ := os.OpenFile("./hello.txt", os.O_RDWR, 0666)
	//utf8EncodeDemo(file)
}

func utf8EncodeDemo(file *os.File) {
	w := bufio.NewWriter(file)
	buf := make([]byte, utf8.UTFMax)
	for i := 0; i < 1000; i++ {
		size := utf8.EncodeRune(buf, rune(i))
		nbytes, _ := w.WriteRune(rune(i))
		if nbytes != size {
			fmt.Fprintf(os.Stdout, "err")
			return
		}
	}
	w.Flush()
}

func utf8DecodeDemo() {
	r, size := utf8.DecodeRune([]byte("Hello world"))
	fmt.Println("DecodeRune", r, size) // 72, 1

	r, size = utf8.DecodeLastRune([]byte("Hello world"))
	fmt.Println("DecodeLastRune", r, size) // 100, 1

	s1 := "hello world, 你好世界"
	for len(s1) > 0 {
		r, size := utf8.DecodeRuneInString(s1)
		fmt.Printf("%T, %c, %d\n", r, r, size)
		s1 = s1[size:]
	}
	fmt.Println("---------------")
	s2 := "hello world"
	for len(s2) > 0 {
		r, size := utf8.DecodeLastRuneInString(s2)
		fmt.Printf("%T, %c, %d\n", r, r, size)
		s2 = s2[:len(s2)-size]
	}
}

func utf8FullRuneDemo() {
	buf := []byte("世界")
	fmt.Printf("%t\n", utf8.FullRune(buf))     // true
	fmt.Printf("%t\n", utf8.FullRune(buf[:2])) // false
	fmt.Printf("%t\n", utf8.FullRune(buf[:3])) // ture

	fmt.Println(utf8.FullRuneInString("你好"))     // true
	fmt.Println(utf8.FullRuneInString("世界"[:2])) // false
}

func utf8RuneStartDemo() {
	fmt.Println(utf8.RuneStart([]byte("a世界")[0])) // true
	fmt.Println(utf8.RuneStart([]byte("a世界")[1])) // true
	fmt.Println(utf8.RuneStart([]byte("a世界")[2])) // false
}

func utf16Demo() {
	u16 := utf16.Encode([]rune("hello world, 你好世界"))
	fmt.Println(u16)                   // [104 101 108 108 111 32 119 111 114 108 100 44 32 20320 22909 19990 30028]
	fmt.Println(utf16.EncodeRune('界')) // 65533 65533

	fmt.Println(utf16.Decode(u16))
	fmt.Println(utf16.DecodeRune(65533, 65533)) // 65533
}
