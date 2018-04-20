package projectEulerHelper

import "testing"

func Test_IsPrimeSqrt(t *testing.T) {
	tests := []struct {
		i int  //input
		o bool //output
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
	}
	for _, test := range tests {
		r := IsPrimeSqrt(test.i)
		if r != test.o {
			t.Errorf("Checking prime for %d was incorrect, got: %v, want: %v", test.i, r, test.o)
		}
	}
}

func Test_ListOfPrimeFactors(t *testing.T) {
	tests := []struct {
		i int   //input
		o []int //output
	}{
		{1, nil},
		{2, []int{2}},
		{4, []int{2, 2}},
		{12, []int{2, 2, 3}},
		{25, []int{5, 5}},
	}
	for _, test := range tests {
		r := ListOfPrimeFactors(test.i)
		if len(r) != len(test.o) {
			t.Errorf("The number of primes returned for %d was incorrect, got: %v, want: %v", test.i, len(r), len(test.o))
		} else {
			for i := range r {
				if r[i] != test.o[i] {
					t.Errorf("The returned prime was incorrect, got: %v, want: %v", r[i], test.o[i])
				}
			}
		}
	}
}

func Test_SieveOfEratosthenes(t *testing.T) {
	tests := []struct {
		i int   // input
		o []int // output
	}{
		{1, nil},
		{2, nil},
		{3, []int{2}},
		{4, []int{2, 3}},
		{5, []int{2, 3}},
		{6, []int{2, 3, 5}},
		{72, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71}},
	}
	for _, test := range tests {
		r := SieveOfEratosthenes(test.i)
		if len(r) != len(test.o) {
			t.Errorf("The number of primes returned for %d was incorrect, got: %v, want: %v", test.i, len(r), len(test.o))
		} else {
			for i := range r {
				if r[i] != test.o[i] {
					t.Errorf("The returned prime was incorrect, got: %v, want: %v", r[i], test.o[i])
				}
			}
		}
	}
}
