package main

import (
	"fmt"
)

func main() {

	//case1()
	case2()
}

func case1() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	n := <-ch1
	fmt.Println(n)
}

// 需求
// 定义两个chan  往第一个chan里面写值， 在读取第一个chan里面的值，往第二个chan里面写 在把第二个chan的数据读出来
func case2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		//for n := range ch1 {
		//	ch2 <- n + 1
		//}
		for {
			n, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- n * n
		}
		close(ch2)
	}()
	for n := range ch2 {
		fmt.Println(n)
	}
}

//需求：
//计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
//随机生成数字进行计算
func case3() {

}
