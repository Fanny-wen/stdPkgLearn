package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	//bytesCountDemo()
	//bytesContainsDemo()
	//bytesIndexDemo()
	//bytesJoinDemo()
	//bytesRunesDemo()
	//bytesReaderDemo()
	bytesBufferDemo()
}

func bytesCountDemo() {
	bar := []byte("hello world")
	n := bytes.Count(bar, []byte("o"))
	fmt.Printf("%d\n", n) // 2
}

func bytesContainsDemo() {
	bar := []byte("hello world")
	fmt.Printf("%t\n", bytes.Contains(bar, []byte("hel"))) // true
	fmt.Printf("%t\n", bytes.ContainsAny(bar, "abcdefg"))  // true
	fmt.Printf("%t\n", bytes.ContainsRune(bar, 123))       // false
}

func bytesIndexDemo() {
	bar := []byte("hello world")
	fmt.Printf("%d\n", bytes.Index(bar, []byte("o"))) // 4
	fmt.Printf("%d\n", bytes.IndexByte(bar, 'w'))     // 6
	fmt.Printf("%d\n", bytes.IndexByte(bar, 123))     // -1
	fmt.Printf("%d\n", bytes.IndexAny(bar, "r"))      // 8
	fmt.Printf("%d\n", bytes.IndexRune(bar, 101))     // 1
	fmt.Printf("%d\n", bytes.IndexRune(bar, 'd'))     // 10
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Printf("%d\n", bytes.IndexFunc([]byte("hi 你好"), f)) // 3
}

func bytesJoinDemo() {
	bar := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	s := bytes.Join(bar, []byte{119, 120})
	fmt.Printf("%s\n", s) // foowxbarwxbaz
	fmt.Printf("%T\n", s) // []uint8
}

func bytesRunesDemo() {
	bar := []byte("你好, 世界")
	fmt.Println(bar)
	for k, v := range bar {
		fmt.Printf("%d : %v -- %s -- %U\n", k, v, string(v), v)
	}
	baz := bytes.Runes(bar)
	fmt.Println(baz)
	for k, v := range baz {
		fmt.Printf("%d : %v -- %s -- %U\n", k, v, string(v), v)
	}
	fmt.Printf("%s\n", string(baz))
}

func bytesReaderDemo() {
	bar := []byte("hello world, 你好, 世界")
	r1 := bytes.NewReader(bar)
	d1 := make([]byte, len(bar))
	_, _ = r1.Read(d1)
	fmt.Printf("%v\n", string(d1))
	fmt.Println(r1)
	r1.Reset(bar)
	fmt.Println(r1)
}

func bytesBufferDemo() {
	bar := []byte("hello world")
	b := bytes.NewBuffer(bar)
	line, _ := b.ReadBytes('o')
	fmt.Printf("%v--%s\n", line, string(line))                 // [104 101 108 108 111]--hello
	fmt.Printf("%q--%d--%v\n", b.String(), b.Len(), b.Bytes()) // " world"--6--[32 119 111 114 108 100]
	b.Truncate(3)                                              // 截断
	fmt.Printf("%q--%d--%v\n", b.String(), b.Len(), b.Bytes()) //  " wo"--3--[32 119 111]
}
