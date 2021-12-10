package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

func main() {
	u := &User{
		Name: "jack",
		Age:  18,
		Sex:  1,
	}
	// 序列化
	bt, err := json.Marshal(u)
	if err != nil {
		fmt.Println("json marshal failed. err:", err)
		return
	}
	fmt.Println(string(bt))

	// 反序列化
	u2 := &User{}
	err = json.Unmarshal(bt, u2)
	if err != nil {
		fmt.Println("json unMarshal failed. err:", err)
		return
	}
	fmt.Printf("%+v\n", u2)

}
