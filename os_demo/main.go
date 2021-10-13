package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	//osOpenFileDemo()
	//osReadDemo()
	//osWriteDemo()
	//osTruncateDemo()
	//osFileInfoDemo()
	//osChtimesDemo()
	//osChmodDemo()
	//osRenameDemo()
	//osMkdirDemo()
	//osRemoveDemo()
	//osReaddirnamesDemo()
	//osReaddirDemo()
	//osEnvDemo()
	osGetpagesizeDemo()
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
Stat()返回有关目标文件的信息，
Lstat()返回有关符号链接(symbolic link)本身的信息。
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
	fmt.Printf("%#v\n", fileInfo)

	fmt.Println("======================")

	fileObj2, _ := os.Lstat("./hello_2.txt")
	fmt.Printf("%#v\n", fileObj2)

	fmt.Println("======================")

	fileObj3, _ := os.Stat("./hello_2.txt")
	fmt.Printf("%#v\n", fileObj3)
}

/*
改变文件时间戳
可以显式改变文件的访问时间和修改时间
*/

func osChtimesDemo() {
	fileInfo, err := os.Stat("./hello_2.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "os Stat failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", fileInfo)
	fmt.Printf("%#v\n", fileInfo.ModTime())
	fmt.Printf("%#v\n", fileInfo.Sys())
	_ = os.Chtimes("./hello_2.txt", time.Now(), time.Now())
	fmt.Println("=================================")
	fmt.Printf("%#v\n", fileInfo)
	fmt.Printf("%#v\n", fileInfo.ModTime())
	fmt.Printf("%#v\n", fileInfo.Sys())
}

/*
每个文件都有一个与之关联的用户 ID（UID）和组 ID（GID），籍此可以判定文件的属主和属组。
系统调用 chown、lchown 和 fchown 可用来改变文件的属主和属组
*/
func osChownDemo() {
	fileObj, err := os.Open("./hello_2.txt")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	fileObj.Chown(12, 12)
	//_ = os.Chown("./hello_2.txt", 12, 12)
	//_ = os.Lchown("./hello_2.txt", 12, 12)
}

/*
权限
*/
func osChmodDemo() {
	f, err := os.Open("./hello_2.txt")
	isPem := os.IsPermission(err)
	if isPem {
		return
	} else {
		fmt.Printf("open file faild, because permission? %v\n", isPem) // false
	}
	fi, _ := f.Stat()
	fmt.Printf("文件权限: %v\n", fi.Mode())

	_ = f.Chmod(0777)
	fi, _ = f.Stat()
	fmt.Printf("修改之后的文件权限: %v\n", fi.Mode())
}

/*
目录与链接
*/

// 更改文件名
func osRenameDemo() {
	_ = os.Rename("./hello.txt", "./hello_new.txt")
}

// 创建和移除目录
func osMkdirDemo() {
	_ = os.Mkdir("./build", 0777)
	_ = os.MkdirAll("./build/dist/jyj", 0777)
}

func osRemoveDemo() {
	_ = os.Remove("./build/dist/jyj")
	_ = os.RemoveAll("./build")
}

/*
读目录
Readdirnames 读取目录 f 的内容，返回一个最多有 n 个成员的[]string，切片成员为目录中文件对象的名字，采用目录顺序。
对本函数的下一次调用会返回上一次调用未读取的内容的信息

如果 n>0，Readdirnames 函数会返回一个最多 n 个成员的切片, 如果到达了目录 f 的结尾，返回值 err 会是 io.EOF
如果 n<=0，Readdirnames 函数返回目录中剩余所有文件对象的名字构成的切片
*/
func osReaddirnamesDemo() {
	f, _ := os.Open("./src")
	n, _ := f.Readdirnames(-1)
	fmt.Printf("readdirname: %v\n", n) // [channel_demo crawler_demo flag_example fmt_example func hello http_example io_demo method os_example runtime_demo stdPkgLearn strconv_example sync_demo time_example]
	fmt.Println("=========================")

	// 对本函数的下一次调用会返回上一次调用未读取的内容的信息
	n, err := f.Readdirnames(100) // []
	if err != nil {
		fmt.Printf("readdirnames failed, err: %v\n", err)
	}
	fmt.Printf("readdirname: %v\n", n)
}

/*
Readdir 内部会调用 Readdirnames，将得到的 names 构造路径，通过 Lstat 构造出 []FileInfo。
*/
func osReaddirDemo() {
	f, _ := os.Open("./src")
	fi, _ := f.Readdir(-1)
	fmt.Printf("%v\n", fi)
	for _, v := range fi {
		fmt.Printf("name: %v, size: %v\n", v.Name(), v.Size())
	}
}

/*
os.Setenv() 设置环境变量。
os.Getenv() 获取环境变量。
os.Unsetenv() 删除环境变量，如果我们尝试使用该环境值来获取该环境值os.Getenv()，则会返回一个空值。
os.LookupEnv() 判断环境变量是否存在。如果系统中不存在该变量，则返回值将为空，并且布尔值将为false。否则，它将返回值（可以为空），并且布尔值为true。
os.Clearenv() 清空所有环境变量
os.ExpandEnv() 根据环境变量的值替换字符串中的 ${var} 或 $var。如果不存在任何环境变量，则将使用空字符串替换它。
os.Environ()：以key = value的形式返回包含所有环境变量的字符串的一部分。
*/
func osEnvDemo() {
	os.Setenv("name", "jyj")
	os.Setenv("age", "23")
	defer os.Unsetenv("name")
	//defer os.Unsetenv("age")
	defer os.Clearenv()
	fmt.Printf("%s is %s years old\n", os.Getenv("name"), os.Getenv("age")) // jyj is 14 years old

	s1, b1 := os.LookupEnv("name")
	s2, b2 := os.LookupEnv("name1")
	fmt.Println(s1, b1) // jyj true
	fmt.Println(s2, b2) // "" false
	fmt.Println(os.ExpandEnv("$name is $age years old"))

	for i, env := range os.Environ() {
		fmt.Println(i, env)
	}
}

/*
os.Getpagesize()
Getpagesize返回底层系统的内存页大小。
*/
func osGetpagesizeDemo() {
	i := os.Getpagesize()
	fmt.Println(i) // 4096
}
