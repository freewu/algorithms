package main

// 4062. Transform Array Using Pair Operations
// You are given two integer arrays source and target.

// In one operation, you may choose two distinct indices i and j in source, along with any integer delta. 
// Then update source as follows:
//     1. source[i] = source[i] + source[j] - delta
//     2. source[j] = delta

// Return true if it is possible to make source equal to target after performing the operation any (including zero) number of times. 
// Otherwise, return false.

// Example 1:
// Input: source = [1,2,3], target = [0,2,4]
// Output: true
// Explanation:
// Choose indices i = 0 and j = 2, and set delta = 4.
// Before operation, source[0] = 1 and source[2] = 3.
// After the operation,
// source[0] = 1 + 3 - 4 = 0
// source[2] = 4
// Hence, source becomes [0, 2, 4], which is equal to target.
// Therefore, the answer is true.

// Example 2:
// Input: source = [-5,-5], target = [-15,5]
// Output: true
// Explanation:
// Choose indices i = 1 and j = 0, and set delta = -15.
// Before operation, source[1] = -5 and source[0] = -5.
// After the operation,
// source[1] = -5 + (-5) - (-15) = 5
// source[0] = -15
// Hence, source becomes [-15, 5], which is equal to target.
// Therefore, the answer is true.

// Example 3:
// Input: source = [1,2,1], target = [0,2,5]
// Output: false
// Explanation:
// It can be shown that no matter what operations are performed, source can never be made equal to target. 
// Therefore, the answer is false.

// Constraints:
//     2 <= source.length == target.length <= 10^5
//     -10^9 <= source[i], target[i] <= 10^9

import "fmt"

func canTransform(source []int, target []int) bool {
    diff, n := 0, len(source)
    for i := range n {
        diff += target[i] - source[i]
    }
    return diff == 0
}

func main() {
    // Example 1:
    // Input: source = [1,2,3], target = [0,2,4]
    // Output: true
    // Explanation:
    // Choose indices i = 0 and j = 2, and set delta = 4.
    // Before operation, source[0] = 1 and source[2] = 3.
    // After the operation,
    // source[0] = 1 + 3 - 4 = 0
    // source[2] = 4
    // Hence, source becomes [0, 2, 4], which is equal to target.
    // Therefore, the answer is true.
    fmt.Println(canTransform([]int{1,2,3}, []int{0,2,4})) // true
    // Example 2:
    // Input: source = [-5,-5], target = [-15,5]
    // Output: true
    // Explanation:
    // Choose indices i = 1 and j = 0, and set delta = -15.
    // Before operation, source[1] = -5 and source[0] = -5.
    // After the operation,
    // source[1] = -5 + (-5) - (-15) = 5
    // source[0] = -15
    // Hence, source becomes [-15, 5], which is equal to target.
    // Therefore, the answer is true.
    fmt.Println(canTransform([]int{-5,-5}, []int{-15,5})) // true
    // Example 3:
    // Input: source = [1,2,1], target = [0,2,5]
    // Output: false
    // Explanation:
    // It can be shown that no matter what operations are performed, source can never be made equal to target. 
    // Therefore, the answer is false.
    fmt.Println(canTransform([]int{1,2,1}, []int{0,2,5})) // false

    fmt.Println(canTransform([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(canTransform([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // true
    fmt.Println(canTransform([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(canTransform([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // true
}