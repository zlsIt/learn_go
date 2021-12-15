package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second * 1)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for {

	}

}
