package main

// 2761. Prime Pairs With Target Sum
// You are given an integer n. 
// We say that two integers x and y form a prime number pair if:
//     1 <= x <= y <= n
//     x + y == n
//     x and y are prime numbers

// Return the 2D sorted list of prime number pairs [xi, yi]. 
// The list should be sorted in increasing order of xi. 
// If there are no prime number pairs at all, return an empty array.

// Note: A prime number is a natural number greater than 1 with only two factors, itself and 1.

// Example 1:
// Input: n = 10
// Output: [[3,7],[5,5]]
// Explanation: In this example, there are two prime pairs that satisfy the criteria. 
// These pairs are [3,7] and [5,5], and we return them in the sorted order as described in the problem statement.

// Example 2:
// Input: n = 2
// Output: []
// Explanation: We can show that there is no prime number pair that gives a sum of 2, so we return an empty array. 

// Constraints:
//     1 <= n <= 10^6

import "fmt"

func findPrimePairs(n int) [][]int {
    getPrimes := func(num int) []bool {
        res := make([]bool, num + 1)
        for i := 1; i <= num; i++ {
            res[i] = true
        }
        for i := 2; i * i <= num; i++ {
            if res[i] {
                for j := i * i; j < num; j += i {
                    res[j] = false
                }
            }
        }
        return res
    }
    primes := getPrimes(n)
    res := [][]int{}
    for i := 2; i <= n / 2; i++ {
        if primes[i] && primes[n - i] {
            res = append(res, []int{ i, n - i})
        }
    }
    return res
}

const mx = 1e6
var primes []int
var isPrime = [mx + 1]bool{}

func init() {
    for i := 2; i <= mx; i++ {
        isPrime[i] = true
    }
    for i := 2; i <= mx; i++ {
        if isPrime[i] {
            primes = append(primes, i)
            for j := i * i; j <= mx; j += i {
                isPrime[j] = false
            }
        }
    }
}

func findPrimePairs1(n int) [][]int {
    res := [][]int{}
    if n % 2 > 0 {
        if n > 4 && isPrime[n - 2] {
            return [][]int{{ 2, n - 2 }}
        }
        return res
    }
    for _, x := range primes {
        y := n - x
        if y < x { break }
        if isPrime[y] {
            res = append(res, []int{ x, y })
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 10
    // Output: [[3,7],[5,5]]
    // Explanation: In this example, there are two prime pairs that satisfy the criteria. 
    // These pairs are [3,7] and [5,5], and we return them in the sorted order as described in the problem statement.
    fmt.Println(findPrimePairs(10)) // [[3,7],[5,5]]
    // Example 2:
    // Input: n = 2
    // Output: []
    // Explanation: We can show that there is no prime number pair that gives a sum of 2, so we return an empty array. 
    fmt.Println(findPrimePairs(2)) // []

    fmt.Println(findPrimePairs(1)) // []
    fmt.Println(findPrimePairs(3)) // []
    fmt.Println(findPrimePairs(8)) // [[3 5]]
    fmt.Println(findPrimePairs(64)) // [[3 61] [5 59] [11 53] [17 47] [23 41]]
    fmt.Println(findPrimePairs(1024)) // [[3 1021] [5 1019] [11 1013] [41 983] [47 977] [53 971] [71 953] [83 941] [113 911] [137 887] [167 857] [197 827] [227 797] [251 773] [263 761] [281 743] [347 677] [383 641] [431 593] [461 563] [467 557] [503 521]]
    //fmt.Println(findPrimePairs(99999)) //
    //fmt.Println(findPrimePairs(100000)) //

    fmt.Println(findPrimePairs1(10)) // [[3,7],[5,5]]
    fmt.Println(findPrimePairs1(2)) // []
    fmt.Println(findPrimePairs1(1)) // []
    fmt.Println(findPrimePairs1(3)) // []
    fmt.Println(findPrimePairs1(8)) // [[3 5]]
    fmt.Println(findPrimePairs1(64)) // [[3 61] [5 59] [11 53] [17 47] [23 41]]
    fmt.Println(findPrimePairs1(1024)) // [[3 1021] [5 1019] [11 1013] [41 983] [47 977] [53 971] [71 953] [83 941] [113 911] [137 887] [167 857] [197 827] [227 797] [251 773] [263 761] [281 743] [347 677] [383 641] [431 593] [461 563] [467 557] [503 521]]
    //fmt.Println(findPrimePairs(99999)) //
    //fmt.Println(findPrimePairs(100000)) // 
}