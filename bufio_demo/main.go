package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//bufioReadSliceDemo()
	//bufioReadBytesDmeo()
	//bufioReadStringDemo()
	//bufioReadLineDemo()
	//bufioPeekDemo()
	//bufioBufferedDemo()
	//bufioDiscardDemo()
	//bufioScannerDemo()
	//bufioSplitDemo()
	bufioWriteToDemo()
}

func bufioReadSliceDemo() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	// ReadSlice 返回的 []byte 是指向 Reader 中的 buffer ，而不是 copy 一份返回
	line, err := reader.ReadSlice('\n')
	if err != nil {
		fmt.Printf("readSlice failed, err: %v\n", err)
		return
	}
	fmt.Printf("the line: %s\n", line)
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line) // 和第一次的line不同
	fmt.Println(string(n))
}

func bufioReadBytesDmeo() {
	// 返回的 []byte 是指向 Reader 中的 buffer，而不是 copy 一份返回，也正因为如此，通常我们会使用 ReadBytes 或 ReadString
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "readBytes failed, err: %v\n", err)
		return
	}
	fmt.Printf("the line: %v\n", line)
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the line: %v\n", line)
	fmt.Printf(string(n))
}

func bufioReadStringDemo() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "readString failed, err: %v\n", err)
		return
	}
	fmt.Printf("the line: %v\n", line)
	n, _ := reader.ReadString('\n')
	fmt.Printf("the line: %v\n", line)
	fmt.Print(n)
}

func bufioReadLineDemo() {
	//reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"), 16)
	line, isPrefix, err := reader.ReadLine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadLine failed, err: %v\n", err)
		return
	}
	fmt.Printf("the line: %v\nisPrefix: %v\n", string(line), isPrefix)
}

// Peek返回下一个n个字节，而不推进读取器
func bufioPeekDemo() {
	// 同ReadSlice一样，返回的 []byte 只是 buffer 中的引用，在下次IO操作后会无效
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"), 16)
	line, err := reader.Peek(16)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Peek failed, err: %v\n", err)
		return
	}
	fmt.Printf("the line: %v\n", line)
	fmt.Printf("the line: %v\n", string(line))

	fmt.Println("===================================================")

	// 对多 goroutine 是不安全的，在多并发环境下，不能依赖其结果
	reader = bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	go func(reader *bufio.Reader) {
		line, _ := reader.Peek(14)
		fmt.Printf("%s\n", line)
		time.Sleep(1)
		fmt.Printf("%s\n", line)
	}(reader)
	go reader.ReadBytes('\t')
	time.Sleep(1e8)
}

// 返回可从缓冲区读取的字节数
func bufioBufferedDemo() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"), 16)
	n := reader.Buffered()
	fmt.Println(n)
}

// 跳过n个字节
func bufioDiscardDemo() {
	reader := bufio.NewReaderSize(strings.NewReader("hello world"), 6)
	n, err := reader.Discard(7)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Discard failed, err: %v\n", err)
		return
	}
	fmt.Printf("Discard %d bytes\n", n)

	ch, err := reader.ReadByte()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadByte failed, err: %v\n", err)
	}
	fmt.Println(string(ch))
}

func bufioWriteToDemo() {
	buffer := bytes.NewBuffer([]byte("hello"))
	buffer2 := bytes.NewBuffer(make([]byte, 6))
	reader := bufio.NewReader(buffer)
	_, _ = reader.ReadByte()
	n, err := reader.WriteTo(buffer2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "reader WriteTo failed, err:%v\n", err)
	}
	fmt.Printf("reader WriteTo success, n:%d\n", n)
	fmt.Println(buffer.Bytes())
	fmt.Println(buffer2.Bytes())
}

func bufioScannerDemo() {
	// 默认的split(分词行为) 是 ScanLines
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("scanner Text: %v\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input: %v\n", err)
	}
}

func bufioSplitDemo() {
	// 可以通过Split方法为Scanner实例设置分词行为
	const input = "This is The Golang Standard Library.\nWelcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}
