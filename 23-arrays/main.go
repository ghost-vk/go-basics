// ASSIGNMENT
// When our clients don't respond to a message, they can be reminded
// with up to 2 additional messages. getMessageWithRetries returns:

// 1. An array of 3 strings
// 2. An array of 3 integers

// The strings are just the original messages structured as an array. The first
// is the primary message, the second is the first reminder, and the third is
// the last reminder. The integers represent the cost of sending each message.
// The cost of a message is equal to the length of the message, plus the cost
// of any previous messages. For example:

// 1. "hello" costs 5
// 2. "world" costs 10
// 3. "!" costs 11

package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	primaryCount := len(primary)
	secondaryCount := primaryCount + len(secondary)
	tertiaryCount := secondaryCount + len(tertiary)

	return [3]string{
		primary,
		secondary,
		tertiary,
	}, [3]int{primaryCount, secondaryCount, tertiaryCount}
}
