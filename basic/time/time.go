package main

import (
	"fmt"
	"time"
)

func main() {
	// 时间格式化
	now := time.Now()
	fmtTime := now.Format("2006-01-02 15:04:05")
	fmt.Println(fmtTime)

}
