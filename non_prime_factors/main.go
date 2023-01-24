package main

func NPF(n int) int {
	primes := sieve(n)
	return sumSlice(primes)
}

func sieve(n int) []int {
	// Create a slice to hold the prime numbers
	primes := make([]int, 0)

	// Create a boolean array to mark composite numbers
	composite := make([]bool, n+1)

	// Initialize all numbers as potentially prime
	for i := 2; i <= n; i++ {
		if !composite[i] {
			// i is prime, so add it to the list of primes
			primes = append(primes, i)
		}
		// Mark all multiples of i as composite
		for j := 2 * i; j <= n; j += i {
			composite[j] = true
		}
	}

	return primes
}

func sumSlice(a []int) int {
	var total int
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	return total
}
