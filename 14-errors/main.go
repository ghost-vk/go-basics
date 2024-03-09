package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	t1, err1 := sendSMS(msgToCustomer)
	if err1 != nil {
		return 0, err1
	}

	t2, err2 := sendSMS(msgToSpouse)
	if err2 != nil {
		return 0, err1
	}

	return t1 + t2, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
