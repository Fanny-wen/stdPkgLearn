package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	//pathDirDemo()
	//pathBaseDemo()
	//pathExtDemo()
	//pathAbsDemo()
	//pathSplitDemo()
	//pathJoinDemo()
	//pathSplitListDemo()
	//pathCleanDemo()
	//pathGlobDemo()
	//pathWalkDemo()
	//pathVolumeNameDemo()
	//pathEvalSymlinkDemo()
	pathSlashDemo()
}

/*
path.Dir()
Dir 返回路径中除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。

filepath包提供了兼容操作系统的路径操作。它使用正斜杠还是反斜杠取决于操作系统
filepath.Dir()

path包只能用来处理正斜杠(/)分割的路径, 不能处理带有驱动符或反斜杠(\)的Windows路径
在处理路径时, 应经量使用filepath包, 处理url时, 使用path包
*/
func pathDirDemo() {
	s1 := path.Dir("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(s1) // src/stdPkgLearn/os_path_demo

	s2 := filepath.Dir("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s2) // D:\GoPath\src\stdPkgLearn\os_path_demo

	s3 := path.Dir("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s3) // .
}

/*
path.Base()
Base 函数返回路径的最后一个元素。在提取元素前会去掉末尾的斜杠。
*/
func pathBaseDemo() {
	s1 := path.Base("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(s1) // main.go

	s2 := filepath.Base("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s2) // main.go

	s3 := path.Base("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	// 没法解析非正斜杠路径
	fmt.Println(s3) // D:\GoPath\src\stdPkgLearn\os_path_demo\main.go
}

/*
path.Ext()
Ext 函数返回 path 文件扩展名。扩展名是路径中最后一个从 . 开始的部分，包括 .。如果该元素没有 . 会返回空字符串。
*/
func pathExtDemo() {
	s1 := path.Ext("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(s1) // .go

	s2 := filepath.Ext("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s2) // .go

	s3 := path.Ext("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s3) // .go
}

/*
path.IsAbs()
IsAbs 返回路径是否是一个绝对路径
Go 好像只支持 Linux 格式的路径

filepath.Abs()
Abs 函数返回 path 代表的绝对路径，如果 path 不是绝对路径，会加入当前工作目录以使之成为绝对路径。
*/
func pathAbsDemo() {
	b1 := path.IsAbs("/src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(b1) // true

	b2 := path.IsAbs("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(b2) // false

	b3 := filepath.IsAbs("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(b3) // true

	s, err := filepath.Abs("./src/stdPkgLearn/os_path_demo/main.go")
	if err != nil {
		fmt.Printf("filepath.Abs failed, err:%v\n", err)
		return
	}
	fmt.Println(s)          // D:\GoPath\src\stdPkgLearn\os_path_demo\main.go
	fmt.Println(os.Getwd()) // D:\GoPath <nil>
}

/*
路径的切分
path.Split()
filepath.Split()
*/
func pathSplitDemo() {
	d1, f1 := path.Split("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Printf("dir: %v, file: %v\n", d1, f1) // dir: ./src/stdPkgLearn/os_path_demo/, file: main.go

	d2, f2 := filepath.Split("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Printf("dir: %v, file: %v\n", d2, f2) // dir: D:\GoPath\src\stdPkgLearn\os_path_demo\, file: main.go

	d3, f3 := path.Split("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Printf("dir: %v, file: %v\n", d3, f3) // dir: , file: D:\GoPath\src\stdPkgLearn\os_path_demo\main.go

	d4, f4 := path.Split("./src/stdPkgLearn/os_path_demo/")
	fmt.Printf("dir: %v, file: %v\n", d4, f4) // dir: ./src/stdPkgLearn/os_path_demo/, file:

	d5, f5 := path.Split("/src")
	fmt.Printf("dir: %v, file: %v\n", d5, f5) // dir: /, file: src
}

/*
路径的拼接
path.Join()
相对路径到绝对路径的转变，需要经过路径的拼接。Join 用于将多个路径拼接起来，会根据情况添加路径分隔符。
*/
func pathJoinDemo() {
	s1 := path.Join("./", "src", "stdPkgLearn", "os_path_demo")
	fmt.Println(s1) // src/stdPkgLearn/os_path_demo

	s2 := filepath.Join("D:", "GoPath", "src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s2) // D:GoPath\src\stdPkgLearn\os_path_demo\main.go

	s3 := path.Join("D:", "GoPath", "src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s3) // D:/GoPath/src\stdPkgLearn\os_path_demo\main.go
}

/*
filepath.SplitList()
有时，我们需要分割 PATH 或 GOPATH 之类的环境变量（这些路径被特定于 OS 的列表分隔符连接起来）
功能：按os.PathListSeparator即(;)將路徑進行分割
*/
func pathSplitListDemo() {
	s := os.Getenv("GOPATH")
	ss1 := filepath.SplitList(s)
	fmt.Println(ss1) // [D:\GoPath C:\Users\W1998\go]
}

/*
规整化路径
path.Clean()
Clean 函数通过单纯的词法操作返回和 path 代表同一地址的最短路径
*/
func pathCleanDemo() {
	s1 := path.Clean("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s1) // D:\GoPath\src\stdPkgLearn\os_path_demo\main.go

	s2 := filepath.Clean("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s2) // D:\GoPath\src\stdPkgLearn\os_path_demo\main.go

	s3 := path.Clean("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(s3) // src/stdPkgLearn/os_path_demo/main.go

	s4 := filepath.Clean("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(s4) // src\stdPkgLearn\os_path_demo\main.go
}

/*
文件路径匹配
path.Glob()
*/
func pathGlobDemo() {
	m, err := filepath.Glob("./src/stdPkgLearn/io_demo/*")
	if err != nil {
		fmt.Printf("Glob failed, err: %v\n", err)
		return
	}
	fmt.Println(m) // [src\stdPkgLearn\io_demo\main.go]

	m2, err := filepath.Glob("./src/stdPkgLearn/io_demo/[a-z][a-z][a-z][a-z].go")
	if err != nil {
		fmt.Printf("Glob failed, err: %v\n", err)
		return
	}
	fmt.Println(m2) // [src\stdPkgLearn\io_demo\main.go]
}

/*
遍历目录
*/
func pathWalkDemo() {
	pwd, _ := os.Getwd()
	_ = filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if match, err := filepath.Match("*", filepath.Base(path)); match {
			fmt.Println("path:", path)
			fmt.Println("info:", info)
			fmt.Println()
			return err
		}
		return nil
	})
}

/*
filepath.VolumeName()
返回路徑字符串中的卷名
*/
func pathVolumeNameDemo() {
	s := filepath.VolumeName("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s) // D:
}

/*
filepath.EvalSymlinks()
功能：返回软链指向的路径
*/
func pathEvalSymlinkDemo() {
	s, err := filepath.EvalSymlinks("./temp_db.json.lnk")
	if err != nil {
		fmt.Printf("filepath EvalSymlinks failed, err: %v\n", err)
	}
	fmt.Println(s)
}

/*
filepath.FromSlash() 將 path 中的 ‘/’ 轉換為系統相關的路徑分隔符
filepath.ToSlash()	將path中平臺相關的路徑分隔符轉換成’/’
*/
func pathSlashDemo() {
	s1 := filepath.FromSlash("/var/log/celery/celery_err.log")
	fmt.Println(s1) // \var\log\celery\celery_err.log

	s2 := filepath.ToSlash("D:\\GoPath\\src\\stdPkgLearn\\os_path_demo\\main.go")
	fmt.Println(s2) // D:/GoPath/src/stdPkgLearn/os_path_demo/main.go
}
