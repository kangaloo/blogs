package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 获取可执行文件的绝对路径
func main() {

	// 获取相对于当前工作目录的相对路径
	root := filepath.Dir(os.Args[0])
	fmt.Println(root)

	// 根据相对路径获取绝对路径
	root, err := filepath.Abs(root)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(root)
}
