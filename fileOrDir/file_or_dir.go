package main

import (
	"log"
	"os"
)

func main() {

	file := "/root/data/testFile.txt"
	fi, err := os.Stat(file)

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln(err)
		}
	}

	if fi.IsDir() {
		log.Printf("%s is a file. ", file)
	} else {
		log.Printf("%s is a directory, ", file)
	}
}
