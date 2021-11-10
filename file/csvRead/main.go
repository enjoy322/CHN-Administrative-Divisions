package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	//	读取csv 转结构体
	s1 := time.Now().UnixNano() / 1e6
	//准备读取文件
	fileName := "./file/v1.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
	}
	fmt.Println("---------------------------")

	fmt.Println(time.Now().UnixNano()/1e6-s1, "ms")

	// 针对小文件，也可以一次性读取所有的文件。注意，r要重新赋值，因为readall是读取剩下的
	//fs1, _ := os.Open(fileName)
	//r1 := csv.NewReader(fs1)
	//content, err := r1.ReadAll()
	//if err != nil {
	//	log.Fatalf("can not readall, err is %+v", err)
	//}
	//for _, row := range content {
	//	fmt.Println(row)
	//}

}
