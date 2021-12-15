package main

import (
	"fmt"
)

func main() {
	//case1()
	case2()
}

func case2() {
	ch1 := make(chan string, 10)
	go func(ch chan string) {
		i := 0
		for {
			i++
			select {
			case ch <- fmt.Sprintf("hello%d", i):
			default:
				fmt.Println("channel full, cap:", len(ch))
			}
		}
	}(ch1)
	for v := range ch1 {
		fmt.Println("result value:", v)
	}
}

func case1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case n := <-ch1:
				fmt.Printf("ch1:%d\n", n)
			case n := <-ch2:
				fmt.Printf("ch2:%d\n", n)
			}
		}
	}()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
	}
}
