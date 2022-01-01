package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a int
	var b string
	fmt.Println(a, b)

	// 简写
	d := "hello world"
	fmt.Println(d)
	case1()
}

func case1() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(vcode)
}
