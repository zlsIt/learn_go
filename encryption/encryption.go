package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	var pwd = "123456"
	data, err := GeneratePwd(pwd)
	if err != nil {
		fmt.Println(err)
	}
	newData := string(data)
	fmt.Println(newData)
	flag, err := ValidatePwd(pwd, newData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(flag)

}

func GeneratePwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func ValidatePwd(pwd string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd)); err != nil {
		return false, errors.New("密码对比错误!")
	}
	return true, nil
}
