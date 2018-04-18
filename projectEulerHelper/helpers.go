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

// TrackTime outputs the total time taken to execute the function
func TrackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s to complete.", name, elapsed)
}
