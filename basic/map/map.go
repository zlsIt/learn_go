package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	m := map[string]string{
		"jack": "美国",
		"tom":  "英国",
	}
	fmt.Println(m)
	// 判断map是否有该元素
	v, ok := m["jojo"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("没有对应的值")
	}
	// 添加值
	m["jojo"] = "中国"
	fmt.Println(m)
	// 删除值
	delete(m, "jojo")
	fmt.Println(m)
	// 遍历map
	for k, v := range m {
		fmt.Printf("姓名:%s, 国家:%s\n", k, v)
	}
	fmt.Println("----------------------------------")
	case1()
}

// 按照指定顺序遍历map
func case1() {
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("同学%02d", i)
		v := rand.Intn(100)
		scoreMap[key] = v
	}
	var keys = make([]string, 0, 200)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
