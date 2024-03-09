// ASSIGNMENT As an easter egg, we decided to reward our users with a free text
// message if they send a prime number of text messages this year.
//
// Complete the printPrimes function. It should print all of the prime numbers
// up to and including max. It should skip any numbers that are not prime.
//
// Here's the pseudocode:
//
// printPrimes(max):
//   for n in range(2, max+1):
//     if n is 2:
//       n is prime, print it
//     if n is even:
//       n is not prime, skip to next n
//     for i in range (3, sqrt(n) + 1):
//       if i can be multiplied into n:
//         n is not prime, skip to next n
//     n is prime, print it

package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := 2; n <= max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}

		isPrime := true

		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if !isPrime {
			continue
		}

		fmt.Println(n)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
