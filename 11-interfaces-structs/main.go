package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type birthdayMsg struct {
	bTime         time.Time
	recipientName string
}

// birthdayMsg не явно имплементирует message
func (bm birthdayMsg) getMessage() string {
	return fmt.Sprintf("Hi, %s, it is your birthday!", bm.recipientName)
}

type reportMsg struct {
	text string
}

// reportMsg не явно имплементирует message
func (rMsg reportMsg) getMessage() string {
	return fmt.Sprintf("Got report: %s", rMsg.text)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("================================")

	// Type Assertion
	_, ok := m.(reportMsg)
	if ok {
		fmt.Println("That is report msg!")
	}
}

func main() {
	test(birthdayMsg{
		recipientName: "Bob",
		bTime:         time.Date(1999, 04, 10, 0, 0, 0, 0, time.UTC),
	})
	test(reportMsg{
		text: "Alert! Production is down!",
	})
}
