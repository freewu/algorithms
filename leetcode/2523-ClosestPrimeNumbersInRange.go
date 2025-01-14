package main

// 2523. Closest Prime Numbers in Range
// Given two positive integers left and right, find the two integers num1 and num2 such that:
//     1. left <= num1 < num2 <= right .
//     2. num1 and num2 are both prime numbers.
//     3. num2 - num1 is the minimum amongst all other pairs satisfying the above conditions.

// Return the positive integer array ans = [num1, num2]. 
// If there are multiple pairs satisfying these conditions, 
// return the one with the minimum num1 value or [-1, -1] if such numbers do not exist.

// A number greater than 1 is called prime if it is only divisible by 1 and itself.

// Example 1:
// Input: left = 10, right = 19
// Output: [11,13]
// Explanation: The prime numbers between 10 and 19 are 11, 13, 17, and 19.
// The closest gap between any pair is 2, which can be achieved by [11,13] or [17,19].
// Since 11 is smaller than 17, we return the first pair.

// Example 2:
// Input: left = 4, right = 6
// Output: [-1,-1]
// Explanation: There exists only one prime number in the given range, so the conditions cannot be satisfied.

// Constraints:
//     1 <= left <= right <= 10^6

import "fmt"
import "sort"

// Sieve Of Eratosthenes Algorithms
func closestPrimes(left int, right int) []int {
    primes := make([]bool, right+1) // array for 0 to right+1
    if left < 2 { left = 2 } // if left < 2 left equal to 2 
    // use Sieve Of Eratosthenes algorithms
    for p := 2; p*p <= right; p++ {  // for loop from 2 to p*p < right , increment ++
        if !primes[p] { // if p index in primes is false
            for i := p * p; i <= right; i += p { // for loop p*p < right to right , increment i+p  
                primes[i] = true //  p index in primes equal to true
            }
        }
    }
    p1, p2, last, diff := -1, -1, -1, 1 << 31
    for p := left; p <= right; p++ { // for loop from left to  right , increment ++
        if !primes[p] { // if p index in primes is false
            if p - last < diff { //  p - last to min diff 
                diff, p1, p2 = p - last, last,  p
            }
            last = p // equal last to p 
        }
    }
    if p1 < 0 || p2 < 0 { // if p1 or p2 low 0 , assign p1 , p2 to -1
        p1, p2 = -1, -1
    }
    return []int{ p1, p2 }
}

func closestPrimes1(left int, right int) []int {
    const mx = 1e6
    var primes []int
    init := func() {
        st := [mx + 1]bool{}
        for i := 2; i <= mx; i ++ {
            if !st[i] {
                primes = append(primes, i)
                for j := i * i; j <= mx; j += i {
                    st[j] = true
                }
            }
        }
    }
    init()
    l, r := -1, -1
    i := sort.SearchInts(primes, left)
    if i >= len(primes) - 1 {
        return []int{ l, r }
    }
    for ; i + 1 < len(primes) && primes[i + 1] <= right; i ++ {
        if l < 0 || primes[i + 1] - primes[i] < r - l {
            l, r = primes[i], primes[i + 1]
        }
    }
    return []int{ l, r }
}

func main() {
    // Example 1:
    // Input: left = 10, right = 19
    // Output: [11,13]
    // Explanation: The prime numbers between 10 and 19 are 11, 13, 17, and 19.
    // The closest gap between any pair is 2, which can be achieved by [11,13] or [17,19].
    // Since 11 is smaller than 17, we return the first pair.
    fmt.Println(closestPrimes(10, 19)) // [11,13]
    // Example 2:
    // Input: left = 4, right = 6
    // Output: [-1,-1]
    // Explanation: There exists only one prime number in the given range, so the conditions cannot be satisfied.
    fmt.Println(closestPrimes(4, 6)) // [-1,-1]

    fmt.Println(closestPrimes(1, 1)) // [-1,-1]
    fmt.Println(closestPrimes(1, 1_000_000)) // [2,3]
    fmt.Println(closestPrimes(1_000_000, 1_000_000)) // [-1,-1]

    fmt.Println(closestPrimes(10, 19)) // [11,13]
    fmt.Println(closestPrimes(4, 6)) // [-1,-1]
    fmt.Println(closestPrimes(1, 1)) // [-1,-1]
    fmt.Println(closestPrimes(1, 1_000_000)) // [2,3]
    fmt.Println(closestPrimes(1_000_000, 1_000_000)) // [-1,-1]
}