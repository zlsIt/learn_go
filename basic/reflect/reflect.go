package main

import (
	"fmt"
	"reflect"
)

// 反射
func main() {
	a := make(chan string)
	reflectType(a)
}

func reflectType(v interface{}) {
	tp := reflect.TypeOf(v)
	fmt.Println("类型是:", tp)
	key := tp.Kind()
	fmt.Println(key)
	switch key {
	case reflect.Float64:
		fmt.Println("a is float64")
	case reflect.String:
		fmt.Println("a is string")
	case reflect.Chan:
		fmt.Println("a is chan")
	}
}
