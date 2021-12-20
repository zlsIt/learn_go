package main

import "fmt"

// 切片
func main() {
	var s1 []int
	if s1 == nil {
		fmt.Println("s1 为空")
	} else {
		fmt.Println("s1 不为空")
	}
	// 空切片
	s2 := []int{}
	fmt.Println(s2)
	// make
	s3 := make([]int, 0, 10)
	fmt.Println(s3)

	// 初始化slice
	s4 := []int{1, 3, 5, 7, 8}
	fmt.Println(s4[:2])

	fmt.Println("------------------------------------")
	case1()
	case2()
	fmt.Println("------------------------------------")
	case3()
}

func case1() {
	data := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := data[2:4]

	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data)

}

func case2() {
	str := "hello world"
	bytes := []byte(str)
	s := bytes[:8]
	s[6] = 'G'
	fmt.Println(string(s))
}

func case3() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := slice[6:8]
	fmt.Println(s1)
	s2 := slice[:6:8]
	fmt.Println(s2, len(s2), cap(s2))
	s3 := append(s2, 100, 200, 300)
	fmt.Println(s3, len(s3), cap(s3))
}
