package handle

import "fmt"

// 错误处理
func HandleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
