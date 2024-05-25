package main

// 634. Find the Derangement of An Array
// In combinatorial mathematics, a derangement is a permutation of the elements of a set, such that no element appears in its original position.

// You are given an integer n. There is originally an array consisting of n integers from 1 to n in ascending order, 
// return the number of derangements it can generate. Since the answer may be huge, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3
// Output: 2
// Explanation: The original array is [1,2,3]. The two derangements are [2,3,1] and [3,1,2].

// Example 2:
// Input: n = 2
// Output: 1
 
// Constraints:
//     1 <= n <= 10^6

import "fmt"

func findDerangement(n int) int {
    res := 1
    for i := 1; i <= n; i++ {
        res = (i * res - ((i & 1) << 1) + 1) % 1_000_000_007
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 2
    // Explanation: The original array is [1,2,3]. The two derangements are [2,3,1] and [3,1,2].
    fmt.Println(findDerangement(3)) // 2
    // Example 2:
    // Input: n = 2
    // Output: 1
    fmt.Println(findDerangement(2)) // 1

    fmt.Println(findDerangement(1024)) // 691143516
}