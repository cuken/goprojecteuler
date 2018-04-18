package main

import (
	"log"
	"time"

	"github.com/cuken/GoProjectEuler/projectEulerHelper"
)

func main() {

	// Track the time required to execute our function
	defer projectEulerHelper.TrackTime(time.Now(), "Problem 1")

	// Start at 0
	sum := 0

	// Start i at 1 as checking something mod 0 is not the best idea =)
	for i := 1; i < 1000; i++ {

		// Check if i (from our for loop) when divided by 3 or 5's remainder = 0.
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	// PrintF helpful output: https://golang.org/pkg/fmt/
	log.Printf("The sum of all digits: %d\n", sum)
}
