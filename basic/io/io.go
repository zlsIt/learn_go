package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var filename = "abc.txt"
	//createFile(filename)
	readFile(filename)
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file failed. err:", err)
	}
	defer file.Close()

	var buf [128]byte
	var content []byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed. err:", err)
			return
		}
		content = append(content, buf[:n]...)
	}

	fmt.Println(string(content[:]))
}

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("create file failed. err:", err)
		return
	}
	defer file.Close()

	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
	}
}
