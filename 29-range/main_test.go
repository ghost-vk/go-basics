package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		msg      []string
		badWords []string
		expected int
	}
	tests := []testCase{
		{[]string{"hey", "there", "john"}, []string{"crap", "shoot", "dang", "frick"}, -1},
		{[]string{"ugh", "oh", "my", "frick"}, []string{"crap", "shoot", "dang", "frick"}, 3},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{[]string{"what", "the", "shoot", "I", "hate", "that", "crap"}, []string{"crap", "shoot", "dang", "frick"}, 2},
			{[]string{"crap", "shoot", "dang", "frick"}, []string{""}, -1},
			{[]string{""}, nil, -1},
		}...)
	}

	for _, test := range tests {
		if output := indexOfFirstBadWord(test.msg, test.badWords); output != test.expected {
			t.Errorf(`Test [<73;109;21MFailed:
message:
%v
bad words:
%v
=>
expected:
  %v
actual:
  %v
===========================
`, sliceWithBullets(test.msg), sliceWithBullets(test.badWords), test.expected, output)
		} else {
			fmt.Printf(`Test Passed:
message:
%v
bad words:
%v
=>
expected:
  %v
actual:
  %v
===========================
`, sliceWithBullets(test.msg), sliceWithBullets(test.badWords), test.expected, output)
		}
	}
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending
// on which b[<73;109;21Mutton is used to run the tests
var withSubmit = true
