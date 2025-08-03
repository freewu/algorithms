package main

// 3637. Trionic Array I
// You are given an integer array nums of length n.

// An array is trionic if there exist indices 0 < p < q < n − 1 such that:
//     1. nums[0...p] is strictly increasing,
//     2. nums[p...q] is strictly decreasing,
//     3. nums[q...n − 1] is strictly increasing.

// Return true if nums is trionic, otherwise return false.

// Example 1:
// Input: nums = [1,3,5,4,2,6]
// Output: true
// Explanation:
// Pick p = 2, q = 4:
// nums[0...2] = [1, 3, 5] is strictly increasing (1 < 3 < 5).
// nums[2...4] = [5, 4, 2] is strictly decreasing (5 > 4 > 2).
// nums[4...5] = [2, 6] is strictly increasing (2 < 6).

// Example 2:
// Input: nums = [2,1,3]
// Output: false
// Explanation:
// There is no way to pick p and q to form the required three segments.

// Constraints:
//     3 <= n <= 100
//     -1000 <= nums[i] <= 1000

import "fmt"

func isTrionic(nums []int) bool {
    n, i := len(nums), 2
    if nums[1] <= nums[0] { return false } // First increasing sequence must exist
    // First strictly increasing phase
    for i < n && nums[i] > nums[i - 1] {
        i++
    }
    if i == n { return false }
    // Strictly decreasing phase
    for i < n && nums[i] < nums[i - 1] {
        i++
    }
    if i == n { return false }
    // Final strictly increasing phase
    for j := i; j < n; j++ {
        if (nums[j] <= nums[j - 1]) { return false }
    }
    return true;
}

func main() {
    // Example 1:
    // Input: nums = [1,3,5,4,2,6]
    // Output: true
    // Explanation:
    // Pick p = 2, q = 4:
    // nums[0...2] = [1, 3, 5] is strictly increasing (1 < 3 < 5).
    // nums[2...4] = [5, 4, 2] is strictly decreasing (5 > 4 > 2).
    // nums[4...5] = [2, 6] is strictly increasing (2 < 6).
    fmt.Println(isTrionic([]int{1,3,5,4,2,6})) // true
    // Example 2:
    // Input: nums = [2,1,3]
    // Output: false
    // Explanation:
    // There is no way to pick p and q to form the required three segments.
    fmt.Println(isTrionic([]int{2,1,3})) // false

    fmt.Println(isTrionic([]int{1,2,3,4,5,6,7,8,9})) // false
    fmt.Println(isTrionic([]int{9,8,7,6,5,4,3,2,1})) // false
}