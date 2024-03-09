// ASSIGNMENT
//
// Complete the removeProfanity function.
//
// It should use the strings.ReplaceAll function to replace all instances of
// the following words in the input message with asterisks.
//
// "fubb" -> "****"
// "shiz" -> "****"
// "witch" -> "*****"
//
// It should mutate the value in the pointer and return nothing.
//
// Do not alter the function signature.

package main

import (
	"strings"
)

func removeProfanity(message *string) {
	ref := *message
	ref = strings.ReplaceAll(ref, "fubb", "****")
	ref = strings.ReplaceAll(ref, "shiz", "****")
	ref = strings.ReplaceAll(ref, "witch", "*****")
	*message = ref
}
