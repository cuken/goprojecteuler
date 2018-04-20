package main

import (
	"log"
	"time"

	"github.com/cuken/goprojecteuler/projectEulerHelper"
)

func main() {

	// Track the time required to execute our function
	defer projectEulerHelper.TrackTime(time.Now(), "Problem 3")
	p := 600851475143 // Number from problem
	n := p            // Number to work on
	pl := []int{}     // Create a slice to hold the prime factorizations

	for n%2 == 0 { // Check that the number is divisible by 2
		pl = append(pl, 2) // Add 2 to the list of prime fatorizations
		n = n / 2          // Divide the number by 2
	}

	// N must be odd as we've reached a remainder above
	for i := 3; i*i <= n; i += 2 {
		// Factor using the new i
		for n%i == 0 { // Check for no remainder
			pl = append(pl, i) // Add the prime to the list
			n = n / i          // Divide n by the prime
		}
	}

	// If after we do all this, we still have a remainder greater than 2, then it's prime.
	if n > 2 {
		pl = append(pl, n) // Add the prime to the list
	}

	log.Printf("The largest prime for %d is %d\nThe prime factorization list is: %v", p, n, pl)
}
