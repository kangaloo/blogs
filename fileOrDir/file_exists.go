package main

import (
	"log"
	"os"
)

// 判断文件/目录是否存在
func main() {

	file := "/root/data/testFile.txt"

	_, err := os.Stat(file)

	if err != nil {
		if os.IsNotExist(err) {
			// 用 os.IsNotExist() 检查 err ，返回 true 则文件/目录不存在
			log.Printf("file %s not exist .", file)
			log.Fatalln(err)
		}
	}

}
