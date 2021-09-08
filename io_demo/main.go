package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

func main() {
	//ioPipeDemo()
	//ioCopyDemo()
	//ioCopyNDemo()
	//ioMultiReaderDemo()
	//ioReadAtWriteAtDemo()
	//ioReadFromDemo()
	//ioWriteToDemo()
	//ioSeekerDemo()
	ioByteReaderByteWriterDemo()
}

func ioReadAtWriteAtDemo() {
	// Hello world [72 101 108 108 111 32 119 111 114 108 100]

	//byteBuffer := bytes.NewBuffer(make([]byte, 128))
	//byteReader := bytes.NewReader([]byte("Hello world"))
	stringReader := strings.NewReader("Hello world")
	stringsReaderSlice := make([]byte, 128)
	stringsReaderN, err := stringReader.ReadAt(stringsReaderSlice, 0)
	if err != nil {
		// 当 ReadAt 返回的 n < len(p) 时，它就会返回一个 非nil 的错误来解释 为什么没有返回更多的字节(返回EOF)
		fmt.Println(err)
	}
	fmt.Printf("ReadAt: %d, %v\n", stringsReaderN, stringsReaderSlice[:stringsReaderN])
	fmt.Println("=======================================================================")

	file, _ := os.Create("hello.txt")
	defer file.Close()
	file.WriteString("Hello World")
	n, err := file.WriteAt([]byte("Hello Go"), 17)
	if err != nil {
		panic(err)
	}
	fmt.Printf("WriteAt: %d\n", n)
}

func ioReadFromDemo() {
	file, err := os.OpenFile("hello.txt", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	n, err := file.ReadFrom(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ReadFrom: %d\n", n)
	//writer := bufio.NewWriter(file)
	//writer.ReadFrom(file)
	//writer.Flush()
}

func ioWriteToDemo() {
	reader := strings.NewReader("Hello World")
	reader.WriteTo(os.Stdout)
}

func ioSeekerDemo() {
	//b := make([]byte, 10)
	reader := strings.NewReader("Hello world")
	reader.Seek(6, 1)
	//n, err := reader.Read(b)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(b[:n])

	r, size, _ := reader.ReadRune() //ReadRune读取单个UTF-8编码的Unicode字符，并返回以字节为单位的rune及其大小。如果没有可用的字符，则设置err。
	fmt.Println(r, size)
}

func ioByteReaderByteWriterDemo() {
	var ch byte
	fmt.Scanf("%c\n", &ch)
	buffer := new(bytes.Buffer)
	err := buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("成功写入一个字节, 正在读取该字节")
		newch, err := buffer.ReadByte()
		if err != nil {
			panic(err)
		}
		fmt.Printf("成功读取字节: %c\n", newch)
	} else {
		fmt.Println("写入错误")
	}
	fmt.Println("=======================================")
	var ch2 byte
	reader := bytes.NewReader([]byte("Hello world"))
	ch2, _ = reader.ReadByte()
	fmt.Println(ch2)

}

func Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:index])
}

func ioPipeDemo() {
	pipeReader, pipeWriter := io.Pipe()
	fmt.Println(pipeReader)
	fmt.Println(pipeWriter)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(writer *io.PipeWriter) {
		defer wg.Done()
		data := []byte("Go语言中文网")
		for i := 0; i < 3; i++ {
			n, err := writer.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("写入字节: %d\n", n)
		}
		writer.CloseWithError(errors.New("写入段已关闭"))
	}(pipeWriter)
	go func(reader *io.PipeReader) {
		defer wg.Done()
		buf := make([]byte, 128)
		for {
			fmt.Println("接口端开始阻塞5秒钟...")
			time.Sleep(time.Second * 5)
			fmt.Println("接口端开始接收")
			n, err := reader.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("收到字节: %d, 内容: %v\n", n, buf)
		}
	}(pipeReader)
	wg.Wait()
}

func ioCopyDemo() {
	written, err := io.Copy(os.Stdout, strings.NewReader("Go语言中文网"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println(written)
}

func ioCopyNDemo() {
	written, err := io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println(written)
}

func ioMultiReaderDemo() {
	readers := []io.Reader{
		strings.NewReader("from strings reader"),
		bytes.NewBufferString("from bytes buffer"),
	}
	reader := io.MultiReader(readers...)
	data := make([]byte, 0, 128)
	buf := make([]byte, 10)

	for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
		if err != nil {
			panic(err)
		}
		data = append(data, buf[:n]...)
	}
	fmt.Printf("%s\n", data)
}
