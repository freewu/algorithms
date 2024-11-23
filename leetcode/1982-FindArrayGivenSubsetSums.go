package main

// 1982. Find Array Given Subset Sums
// You are given an integer n representing the length of an unknown array that you are trying to recover. 
// You are also given an array sums containing the values of all 2n subset sums of the unknown array (in no particular order).

// Return the array ans of length n representing the unknown array. 
// If multiple answers exist, return any of them.

// An array sub is a subset of an array arr if sub can be obtained from arr by deleting some (possibly zero or all) elements of arr. 
// The sum of the elements in sub is one possible subset sum of arr. 
// The sum of an empty array is considered to be 0.

// Note: Test cases are generated such that there will always be at least one correct answer.

// Example 1:
// Input: n = 3, sums = [-3,-2,-1,0,0,1,2,3]
// Output: [1,2,-3]
// Explanation: [1,2,-3] is able to achieve the given subset sums:
// - []: sum is 0
// - [1]: sum is 1
// - [2]: sum is 2
// - [1,2]: sum is 3
// - [-3]: sum is -3
// - [1,-3]: sum is -2
// - [2,-3]: sum is -1
// - [1,2,-3]: sum is 0
// Note that any permutation of [1,2,-3] and also any permutation of [-1,-2,3] will also be accepted.

// Example 2:
// Input: n = 2, sums = [0,0,0,0]
// Output: [0,0]
// Explanation: The only correct answer is [0,0].

// Example 3:
// Input: n = 4, sums = [0,0,5,5,4,-1,4,9,9,-1,4,3,4,8,3,8]
// Output: [0,-1,4,5]
// Explanation: [0,-1,4,5] is able to achieve the given subset sums.

// Constraints:
//     1 <= n <= 15
//     sums.length == 2n
//     -10^4 <= sums[i] <= 10^4

import "fmt"
import "sort"

func recoverArray(n int, sums []int) []int {
    sort.Ints(sums)
    m := len(sums)
    res, left, right := make([]int, n), make([]int, m / 2), make([]int, m / 2)
    for i := 0; i < n; i++ {
        diff, hasZero, p, q, k := sums[1] - sums[0], 0, -1, -1, 0
        for j := 0; j < m; j++ {
            if k <= q && right[k] == sums[j] {
                k++
            } else {
                if 0 == sums[j] { hasZero = 1 }
                p++
                left[p] = sums[j]
                q++
                right[q] = sums[j] + diff
            }
        }
        if 1 == hasZero {
            res[i], sums = diff, left
        } else {
            res[i], sums = -diff, right
        }
        m /= 2
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, sums = [-3,-2,-1,0,0,1,2,3]
    // Output: [1,2,-3]
    // Explanation: [1,2,-3] is able to achieve the given subset sums:
    // - []: sum is 0
    // - [1]: sum is 1
    // - [2]: sum is 2
    // - [1,2]: sum is 3
    // - [-3]: sum is -3
    // - [1,-3]: sum is -2
    // - [2,-3]: sum is -1
    // - [1,2,-3]: sum is 0
    // Note that any permutation of [1,2,-3] and also any permutation of [-1,-2,3] will also be accepted.
    fmt.Println(recoverArray(3, []int{-3,-2,-1,0,0,1,2,3})) // [1,2,-3]
    // Example 2:
    // Input: n = 2, sums = [0,0,0,0]
    // Output: [0,0]
    // Explanation: The only correct answer is [0,0].
    fmt.Println(recoverArray(2, []int{0,0,0,0})) // [0,0]
    // Example 3:
    // Input: n = 4, sums = [0,0,5,5,4,-1,4,9,9,-1,4,3,4,8,3,8]
    // Output: [0,-1,4,5]
    // Explanation: [0,-1,4,5] is able to achieve the given subset sums.
    fmt.Println(recoverArray(4, []int{0,0,5,5,4,-1,4,9,9,-1,4,3,4,8,3,8})) // [0,-1,4,5]
}