package main

import "fmt"

type Sayer interface {
	say()
	eat()
}

type Dog struct{}

type Cat struct{}

func (d Dog) say() {
	fmt.Println("汪汪汪")
}

func (d Dog) eat() {
	fmt.Println("吃骨头")
}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}

func (c Cat) eat() {
	fmt.Println("吃鱼")
}

func main() {
	dog := &Dog{}
	dog.say()
	dog.eat()

	cat := &Cat{}
	cat.say()

	var peo People = &Student{}
	think := "sb"
	fmt.Println(peo.Speak(think))
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
