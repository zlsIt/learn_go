package main

import "fmt"

// 鸟
type Bird interface {
	Fry()
	Type() string
}

type Canary struct {
	Name string
}

func (c *Canary) Fry() {
	fmt.Printf("我是%s, 用黄色的翅膀飞\n", c.Name)
}

func (c *Canary) Type() string {
	return c.Name
}

type Crow struct {
	Name string
}

func (c *Crow) Fry() {
	fmt.Printf("我是%s, 用黑色的翅膀飞\n", c.Name)
}

func (c *Crow) Type() string {
	return c.Name
}

func LetItFry(bird Bird) {
	bird.Fry()
}

func main() {
	LetItFry(&Canary{"金丝雀"})
	LetItFry(&Crow{"乌鸦"})
}
