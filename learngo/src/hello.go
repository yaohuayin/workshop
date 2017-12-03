package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}

type PhoneConnect struct {
	name string
}

func (pc PhoneConnect) Name() string {
	return pc.name
}

func (pc PhoneConnect) Connect() {
	fmt.Println("Connect: ", pc.name)
}

func main() {
	var a USB
	a = PhoneConnect{"PhoneConnect"}
	a.Connect()
}
