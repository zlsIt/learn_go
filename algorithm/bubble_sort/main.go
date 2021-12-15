package main

import "fmt"

func main() {
	arr := []int{18, 30, 2, 5, 6, 4, 7, 10, 8}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		fmt.Println(arr)
	}
	return arr
}
