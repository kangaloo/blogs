package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 获取可执行文件的绝对路径
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(root)
}
