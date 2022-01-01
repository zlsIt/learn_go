package main

import "fmt"

func main() {
	sum := add([]int{1, 2}...)
	fmt.Println(sum)
}

func add(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}
