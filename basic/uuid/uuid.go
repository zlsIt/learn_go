package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func main() {

	v1 := uuid.NewV1()
	fmt.Println(v1.String())

	str := v1.String()
	str = strings.ReplaceAll(uuid.NewV1().String(), "-", "")
	fmt.Println(str)
}
