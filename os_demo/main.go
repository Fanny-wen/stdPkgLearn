package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//osOpenFileDemo()
	//osReadDemo()
	//osWriteDemo()
	//osTruncateDemo()
	osFileInfoDemo()
}

func osOpenFileDemo() {
	//fileObj, err := os.OpenFile("./hello.txt", os.O_CREATE|os.O_RDWR, 0666)
	// openFile 是一个更一般性的文件打开函数，大多数调用者都应用 Open 或 Create 代替本函数。

	//    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	//    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	//    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	//    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	//    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	//    O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
	//    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
	//    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	fileObj, err := os.Create("./hello.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "open file faild, err: %v\n", err)
	}
	defer fileObj.Close()
	fmt.Println(fileObj)
}

/*
ReadAt 从指定的位置（相对于文件开始位置）读取长度为 len(b) 个字节数据并写入 b。
它返回读取的字节数和可能遇到的任何错误。当 n<len(b) 时，本方法总是会返回错误；如果是因为到达文件结尾，返回值 err 会是 io.EOF。
它对应的系统调用是 pread。

Read 和 ReadAt 的区别：前者从文件当前偏移量处读，且会改变文件当前的偏移量；
而后者从 off 指定的位置开始读，且不会改变文件当前偏移量。
*/
func osReadDemo() {
	b := make([]byte, 128)
	fileObj, err := os.Open("./hello.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file faild, err: %v\n", err)
	}
	defer fileObj.Close()
	for {
		n, err := fileObj.Read(b)
		if err == io.EOF {
			fmt.Println("read fileObj end")
			//return
			break
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "read fileObj faild, err: %v\n", err)
		}
		// 获取文件对象当前偏移量
		ret, err := fileObj.Seek(0, os.SEEK_CUR)
		fmt.Println(ret)
		fmt.Printf("read fileObj %d byte\n", n)
		fmt.Println(string(b))
	}

	n, err := fileObj.ReadAt(b, 3)
	if err != nil {
		fmt.Fprintf(os.Stdout, "ReadAt file faild, err: %v\n", err)
	}
	fmt.Printf("file readAt method read %d byte\n", n)
	fmt.Println(b[:n])
	fmt.Println(string(b[:n]))
	ret, err := fileObj.Seek(0, os.SEEK_CUR)
	fmt.Println(ret)
}

func osWriteDemo() {
	b := make([]byte, 128)
	fileObj1, err := os.Open("./hello.txt")
	n, err := fileObj1.Read(b)
	if err == io.EOF {
		fmt.Println("read fileObj end")
		return
	}
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "read fileObj faild, err: %v\n", err)
		return
	}
	defer fileObj1.Close()

	fileObj2, err := os.Create("./hello_2.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file faild, err: %v\n", err)
		return
	}
	defer fileObj2.Close()
	n, err = fileObj2.Write(b)
	if err != nil {
		fmt.Fprintf(os.Stdout, "write file faild, err: %v\n", err)
		return
	}
	fmt.Printf("write file success, write %d byte\n", n)

	n, err = fileObj2.WriteAt([]byte("\nhello world"), 128)
	if err != nil {
		fmt.Fprintf(os.Stdout, "writeAt faild, err: %v\n", err)
		return
	}
	fmt.Printf("writeAt success, write %d byte\n", n)
}

/*
osTruncateDemo: 截断文件
调用 `File.Truncate` 前，需要先以  可写方式  打开操作文件，该方法不会修改文件偏移量。
*/
func osTruncateDemo() {
	fileObj, err := os.OpenFile("./hello_2.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer fileObj.Close()
	err = fileObj.Truncate(148)
	if err != nil {
		fmt.Printf("truncate file failed, err:%v\n", err)
		return
	}
	ret, _ := fileObj.Seek(0, os.SEEK_CUR)
	fmt.Println(ret)
}

/*
	type FileInfo interface {
		Name() string       // 文件的名字（不含扩展名）
		Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
		Mode() FileMode     // 文件的模式位
		ModTime() time.Time // 文件的修改时间
		IsDir() bool        // 等价于 Mode().IsDir()
		Sys() interface{}   // 底层数据来源（可以返回 nil）
	}
*/
func osFileInfoDemo() {
	fileObj, err := os.Open("./hello_2.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	fileInfo, _ := fileObj.Stat()
	fmt.Println(fileInfo.Name())
	sys := fileInfo.Sys()
	fmt.Println(sys)
	//stat := sys.(*syscall.Stat_t)
	//fmt.Println(time.Unix(stat.Atimespec.Unix()))
}
