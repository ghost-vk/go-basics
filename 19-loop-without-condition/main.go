// Complete the maxMessages function. Given a cost threshold, it should
// calculate the maximum number of messages that can be sent.
//
// Each message costs 100 pennies, plus an additional fee. The fee structure
// is:
//
// 1st message: 100 + 0 2nd message: 100 + 1 3rd message: 100 + 2 4th message:
// 100 + 3

package main

func maxMessages(thresh int) int {
	sum := 0
	for i := 0; ; i++ {
		sum += 100 + i
		if sum >= thresh {
			return i
		}
	}
}
