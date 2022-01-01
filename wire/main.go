package main

import "fmt"

type Message struct {
	msg string
}
type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

func NewGreeter(m Message) Greeter {
	return Greeter{
		Message: m,
	}
}

func NewEvent(g Greeter) Event {
	return Event{
		Greeter: g,
	}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func (g Greeter) Greet() Message {
	return g.Message
}

//func main() {
//	msg := NewMessage("hello world")
//	greeter := NewGreeter(msg)
//	event := NewEvent(greeter)
//	event.Start()
//}

func main() {
	event := InitializeEvent("hello wire")
	event.Start()
}
