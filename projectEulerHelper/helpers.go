package projectEulerHelper

import (
	"log"
	"math"
	"time"
)

// Helper functions pulled from https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/

// IsPrimeSqrt Returns true or false if a number is prime
func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// ListOfPrimeFactors returns a []int slice of all prime factors within the number n
func ListOfPrimeFactors(n int) (pl []int) {
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

	return pl
}

// SieveOfEratosthenes returns a []int slice that has a list of primes up to (not including) n.
func SieveOfEratosthenes(n int) []int {
	f := make([]bool, n)
	var r []int
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if f[i] == false {
			for j := i * i; j < n; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if f[i] == false {
			r = append(r, i)
		}
	}
	return r
}

// Fibonacci Recursive Fibonacci sequence that calls a function that returns an int
func Fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

// TrackTime outputs the total time taken to execute the function
func TrackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s to complete.", name, elapsed)
}
