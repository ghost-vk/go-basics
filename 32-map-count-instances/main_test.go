package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		userIDs  []string
		expected map[string]int
	}
	tests := []testCase{
		{
			[]string{"cersei", "jaime", "cersei"},
			map[string]int{"cersei": 2, "jaime": 1},
		},
		{
			[]string{"cersei", "tyrion", "jaime", "tyrion", "tyrion"},
			map[string]int{"cersei": 1, "jaime": 1, "tyrion": 3},
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				[]string{},
				map[string]int{},
			},
			{
				[]string{"cersei", "cersei", "cersei"},
				map[string]int{"cersei": 3},
			},
			{
				[]string{"cersei", "jaime", "cersei", "tyrion", "cersei", "jaime"},
				map[string]int{"cersei": 3, "jaime": 2, "tyrion": 1},
			},
			{
				[]string{"cersei", "cersei", "jaime", "jaime", "cersei", "jaime", "tyrion"},
				map[string]int{"cersei": 3, "jaime": 3, "tyrion": 1},
			},
		}...)
	}

	for _, test := range tests {
		if output := getCounts(test.userIDs); !compareMaps(output, test.expected) {
			t.Errorf(
				"Test Failed: %v ->\n expected: %v\n actual: %v",
				test.userIDs,
				test.expected,
				output,
			)
		} else {
			fmt.Printf("Test Passed: %v ->\n expected: %v\n actual: %v\n",
				test.userIDs,
				test.expected,
				output,
			)
		}
	}
}

func compareMaps(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
