package main

import "fmt"

// 算法步骤
// 将数据根据一个值按照大小分成左右两边，左边小于此值，右边大于
// 将两边数据进行递归调用步骤1
// 将所有数据合并
func main() {
	arr := []int{1, 3, 5, 7, 2, 4, 8, 7, 3}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	first := arr[0]
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, first)
	for i := 1; i < len(arr); i++ {
		if arr[i] < first {
			low = append(low, arr[i])
		} else if arr[i] > first {
			high = append(high, arr[i])
		} else if arr[i] == first {
			mid = append(mid, arr[i])
		}
	}
	low, high = quickSort(low), quickSort(high)
	arr = append(append(low, mid...), high...)
	return arr
}
