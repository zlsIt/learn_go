package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":19990", nil)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	fmt.Println(number)
	if number == 1 {
		time.Sleep(time.Second * 3) // 耗时10秒的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}
