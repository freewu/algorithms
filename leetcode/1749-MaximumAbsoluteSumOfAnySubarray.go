package main

// 1749. Maximum Absolute Sum of Any Subarray
// You are given an integer array nums. 
// The absolute sum of a subarray [numsl, numsl+1, ..., numsr-1, numsr] is abs(numsl + numsl+1 + ... + numsr-1 + numsr).

// Return the maximum absolute sum of any (possibly empty) subarray of nums.

// Note that abs(x) is defined as follows:
//     If x is a negative integer, then abs(x) = -x.
//     If x is a non-negative integer, then abs(x) = x.

// Example 1:
// Input: nums = [1,-3,2,3,-4]
// Output: 5
// Explanation: The subarray [2,3] has absolute sum = abs(2+3) = abs(5) = 5.

// Example 2:
// Input: nums = [2,-5,1,-4,3,-2]
// Output: 8
// Explanation: The subarray [-5,1,-4] has absolute sum = abs(-5+1-4) = abs(-8) = 8.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^4 <= nums[i] <= 10^4

import "fmt"

func maxAbsoluteSum(nums []int) int {
    pos, posSum, neg, negSum := 0, 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range nums {
        if posSum + nums[i] > 0 {
            posSum += nums[i]
        } else {
            posSum = 0
        }
        pos = max(pos, posSum)
        if negSum + nums[i] < 0 {
            negSum += nums[i]
        } else {
            negSum = 0
        }
        neg = min(neg, negSum)
    }
    return max(pos, -neg)
}

func main() {
    // Example 1:
    // Input: nums = [1,-3,2,3,-4]
    // Output: 5
    // Explanation: The subarray [2,3] has absolute sum = abs(2+3) = abs(5) = 5.
    fmt.Println(maxAbsoluteSum([]int{1,-3,2,3,-4})) // 5
    // Example 2:
    // Input: nums = [2,-5,1,-4,3,-2]
    // Output: 8
    // Explanation: The subarray [-5,1,-4] has absolute sum = abs(-5+1-4) = abs(-8) = 8.
    fmt.Println(maxAbsoluteSum([]int{2,-5,1,-4,3,-2})) // 8
}