package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	ret := 0
	err = conn.Call("Rect.Area", Params{10, 10}, &ret)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("面积为:", ret)

	ret2 := 0
	err2 := conn.Call("Rect.Perimeter", Params{10, 20}, &ret2)
	if err2 != nil {
		log.Println(err)
	}
	fmt.Println("周长为:", ret2)

}
