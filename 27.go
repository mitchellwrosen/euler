package main

import (
	"fmt"
	"math"
)

// Returns an isPrime function that accepts values in the range 1-size
func isPrimeGen(size int) func(n int) bool {
	isPrime := make([]bool, size)
	for i := range isPrime {
		isPrime[i] = true
	}

	sqrt := int(math.Sqrt(float64(size)))
	for i := 2; i < sqrt; i++ {
		if isPrime[i] {
			for j := i * i; j < size; j += i {
				isPrime[j] = false
			}
		}
	}

	return func(n int) bool {
		return isPrime[n]
	}
}

type info struct {
	primes int
	a      int
	b      int
}

func ex27() {
	isPrime := isPrimeGen(1000000)

	best := info{0, 0, 0}
	for a := -1000; a <= 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			primes := 0
			for n := 0; ; n++ {
				val := n*n + a*n + b

				if val < 0 {
					continue
				}

				if isPrime(val) {
					primes++
				} else {
					break
				}
			}

			if primes > best.primes {
				best.primes = primes
				best.a = a
				best.b = b
			}
		}
	}

	fmt.Printf("%d consecutive primes: a = %d, b = %d, a*b = %d\n", best.primes,
		best.a, best.b, best.a*best.b)
}
