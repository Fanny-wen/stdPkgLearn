package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	//pathDirdemo()
	//pathBaseDemo()
	//pathExtDemo()
	pathDemo()
}

/*
path.Dir()
Dir 返回路径中除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。
*/
func pathDirdemo() {
	s1 := path.Dir("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(s1) // src/stdPkgLearn/os_path_demo
}

/*
path.Base()
Base 函数返回路径的最后一个元素。在提取元素前会去掉末尾的斜杠。
*/
func pathBaseDemo() {
	s := path.Base("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(s) // main.go
}

/*
path.Ext()
Ext 函数返回 path 文件扩展名。扩展名是路径中最后一个从 . 开始的部分，包括 .。如果该元素没有 . 会返回空字符串。
*/
func pathExtDemo() {
	s := path.Ext("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(s) // .go
}

/*
path.IsAbs()
IsAbs 返回路径是否是一个绝对路径
Go 好像只支持 Linux 格式的路径

filepath.Abs()
Abs 函数返回 path 代表的绝对路径，如果 path 不是绝对路径，会加入当前工作目录以使之成为绝对路径。
*/
func pathDemo() {
	b1 := path.IsAbs("./src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(b1) //false

	b2 := path.IsAbs("D:/GoPath/src/stdPkgLearn/os_path_demo/main.go")
	fmt.Println(b2) // false

	s, err := filepath.Abs("./src/stdPkgLearn/os_path_demo/main.go")
	if err != nil {
		fmt.Printf("filepath.Abs failed, err:%v\n", err)
		return
	}
	fmt.Println(s) // D:\GoPath\src\stdPkgLearn\os_path_demo\main.go

	fmt.Println(os.Getwd()) // D:\GoPath <nil>
}
