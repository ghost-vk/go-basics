// ASSIGNMENT
// Each time a user receives a message, their username is logged in a slice. We
// want a more efficient way to count how many messages each user received.
//
// Implement the getCounts function. It should return a map of string -> int.
// Each string (key) is a username, and the int (value) should represent the
// number of times that the user has received a message.
//
// So, if "benji" appears in the slice 3 times, the map should have the key
// "benji" with a value 3.

package main

func getCounts(usernames []string) map[string]int {
	counts := map[string]int{}

	for _, u := range usernames {
		counts[u]++
	}

	return counts
}
