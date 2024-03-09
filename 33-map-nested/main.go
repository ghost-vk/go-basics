// ASSIGNMENT
// Because Textio is a glorified customer database, we have a lot of internal
// logic for sorting and dealing with customer names.
//
// Complete the getNameCounts function. It takes a slice of strings (names) and
// returns a nested map where the first key is all the unique first characters
// (runes) of the names, the second key is all the names themselves, and the
// value is the count of each name.
//
// For example:
//
// billy
// billy
// bob
// joe
//
// Creates the following nested map:
//
// b: {
//     billy: 2,
//     bob: 1
// },
// j: {
//     joe: 1
// }
//
// Note: the test suite is not printing the map you're returning directly, but
// instead checking some specific keys.

package main

func getNameCounts(names []string) map[rune]map[string]int {
	d := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		c := rune(name[0])
		_, exists := d[c]
		if !exists {
			d[c] = make(map[string]int)
		}
		d[c][name]++
	}
	return d
}
