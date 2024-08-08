package main

// 793. Preimage Size of Factorial Zeroes Function
// Let f(x) be the number of zeroes at the end of x!. 
// Recall that x! = 1 * 2 * 3 * ... * x and by convention, 0! = 1.
//     For example, f(3) = 0 because 3! = 6 has no zeroes at the end, while f(11) = 2 because 11! = 39916800 has two zeroes at the end.

// Given an integer k, return the number of non-negative integers x have the property that f(x) = k.

// Example 1:
// Input: k = 0
// Output: 5
// Explanation: 0!, 1!, 2!, 3!, and 4! end with k = 0 zeroes.

// Example 2:
// Input: k = 5
// Output: 0
// Explanation: There is no x such that x! ends in k = 5 zeroes.

// Example 3:
// Input: k = 3
// Output: 5

// Constraints:
//     0 <= k <= 10^9

import "fmt"

func preimageSizeFZF(k int) int {
    trailingZeroes := func(n int) int {
        res := 0
        for n > 0 {
            res += n / 5
            n = n / 5
        }
        return res
    }
     // trailingZeroes(1<<32) > 1e9
    low, high := 0, 1 << 32
    for low < high {
        mid := low + (high - low) / 2
        if trailingZeroes(mid) < k {
            low = mid + 1
        } else {
            high = mid
        }
    }
    if trailingZeroes(low) == k { // the result should be either 0 or 5
        return 5
    }
    return 0
}

func main() {
    // Example 1:
    // Input: k = 0
    // Output: 5
    // Explanation: 0!, 1!, 2!, 3!, and 4! end with k = 0 zeroes.
    fmt.Println(preimageSizeFZF(0)) // 5
    // Example 2:
    // Input: k = 5
    // Output: 0
    // Explanation: There is no x such that x! ends in k = 5 zeroes.
    fmt.Println(preimageSizeFZF(5)) // 0
    // Example 3:
    // Input: k = 3
    // Output: 5
    fmt.Println(preimageSizeFZF(3)) // 5

    fmt.Println(preimageSizeFZF(1000000000)) // 5
}