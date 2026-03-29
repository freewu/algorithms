package main

// 3880. Minimum Absolute Difference Between Two Values
// You are given an integer array nums consisting only of 0, 1, and 2.

// A pair of indices (i, j) is called valid if nums[i] == 1 and nums[j] == 2.

// Return the minimum absolute difference between i and j among all valid pairs. If no valid pair exists, return -1.

// The absolute difference between indices i and j is defined as abs(i - j).

// Example 1:
// Input: nums = [1,0,0,2,0,1]
// Output: 2
// Explanation:
// The valid pairs are:
// (0, 3) which has absolute difference of abs(0 - 3) = 3.
// (5, 3) which has absolute difference of abs(5 - 3) = 2.
// Thus, the answer is 2.

// Example 2:
// Input: nums = [1,0,1,0]
// Output: -1
// Explanation:
// There are no valid pairs in the array, thus the answer is -1.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 2

import "fmt"

func minAbsoluteDifference(nums []int) int {
    res, one, two := 1 << 61, -1, -1
    for i, v := range nums {
        if v == 1 {
            one = i
            if two > -1 {
                res = min(res, one - two)
            }
        }
        if v == 2 {
            two = i
            if one > -1 {
                res = min(res, two - one)
            }
        }
    }
    if res ==  1 << 61 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,0,0,2,0,1]
    // Output: 2
    // Explanation:
    // The valid pairs are:
    // (0, 3) which has absolute difference of abs(0 - 3) = 3.
    // (5, 3) which has absolute difference of abs(5 - 3) = 2.
    // Thus, the answer is 2.
    fmt.Println(minAbsoluteDifference([]int{1,0,0,2,0,1})) // 2
    // Example 2:
    // Input: nums = [1,0,1,0]
    // Output: -1
    // Explanation:
    // There are no valid pairs in the array, thus the answer is -1.
    fmt.Println(minAbsoluteDifference([]int{1,0,1,0})) // -1

    fmt.Println(minAbsoluteDifference([]int{0,0,0,0,0,0,0,0,0})) // -1
    fmt.Println(minAbsoluteDifference([]int{1,1,1,1,1,1,1,1,1})) // -1
    fmt.Println(minAbsoluteDifference([]int{2,2,2,2,2,2,2,2,2})) // -1
    fmt.Println(minAbsoluteDifference([]int{0,0,0,1,1,1,2,2,2})) // 1
    fmt.Println(minAbsoluteDifference([]int{1,1,1,2,2,2,0,0,0})) // 1
    fmt.Println(minAbsoluteDifference([]int{0})) // -1
    fmt.Println(minAbsoluteDifference([]int{1})) // -1
    fmt.Println(minAbsoluteDifference([]int{2})) // -1
}