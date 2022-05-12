package main

import "fmt"

/**
204. Count Primes
Given an integer n, return the number of prime numbers that are strictly less than n.

Example 1:

	Input: n = 10
	Output: 4
	Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

Example 2:

	Input: n = 0
	Output: 0

Example 3:

	Input: n = 1
	Output: 0


Constraints:

	0 <= n <= 5 * 10^6

统计所有小于非负整数 n 的质数的数量。
 */

func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	// 先循环跑所有小于 n 的非素数 i ^ 2 < n
	for i := 2; i*i < n; i++ {
		fmt.Printf("i = %v\n",i)
		if isNotPrime[i] { // 如果已存在了 就不需要再跑跑了 4 能被 2 整除,在i = 2 时,已经跑出来了
			continue
		}
		// 把可以被自己整除的数据都先跑出来
		for j := i * i; j < n; j = j + i { // j = j + i 这个很重要 注意不是 + 1
			fmt.Printf("j = %v\n",j)
			isNotPrime[j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ { // 统计 isNotPrime[x] = false 的数量
		if !isNotPrime[i] {
			count++
		}
	}
	return count
}

// best solution
func countPrimesBest(n int) int {
	// eliminate an odd prime 1 and only even prime 2.
	if n < 3 {
		return 0
	}
	nonPrimes := make([]bool, n) // all are false
	count := n / 2 // we will decrement it when we found odd non-prime
	// sieve of Eratosthenes
	for i := 3; i*i < n; i += 2 {
		if nonPrimes[i] {
			continue
		}
		// mark all multiples of i as non-prime
		for j := i * i; j < n; j += 2 * i {
			if !nonPrimes[j] {
				//nonPrimes[j] = true
				count--
			}
			nonPrimes[j] = true
		}
	}
	return count
}

func main() {
	fmt.Printf("countPrimes(10) = %v\n",countPrimes(10)) // 4 There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
	fmt.Printf("countPrimes(0) = %v\n",countPrimes(0)) // 0
	fmt.Printf("countPrimes(1) = %v\n",countPrimes(1)) // 0

	fmt.Printf("countPrimesBest(10) = %v\n",countPrimesBest(10)) // 4 There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
	fmt.Printf("countPrimesBest(0) = %v\n",countPrimesBest(0)) // 0
	fmt.Printf("countPrimesBest(1) = %v\n",countPrimesBest(1)) // 0
}