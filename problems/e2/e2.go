package main

import (
	"log"
	"time"

	"github.com/cuken/GoProjectEuler/projectEulerHelper"
)

func main() {

	// Track the time required to execute our function
	defer projectEulerHelper.TrackTime(time.Now(), "Problem 2")

	// Start our returning sum at 0
	sum := 0

	// Define our upper limit
	const Max int = 4000000

	// Create our fibonacci method starting point
	f := projectEulerHelper.Fibonacci()

	//Create our first value
	v := f()

	// Loop until we hit max
	for v < Max {
		//Check if our current fibonacci number is even
		if v%2 == 0 {
			sum += v
		}
		v = f()
	}

	// Log the results to discover output.
	log.Printf("The sum of all fibonacci numbers (not exceeding 1 million) is %d", sum)
}
