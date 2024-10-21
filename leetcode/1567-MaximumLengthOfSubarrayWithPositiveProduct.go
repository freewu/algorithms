package main

// 1567. Maximum Length of Subarray With Positive Product
// Given an array of integers nums, find the maximum length of a subarray where the product of all its elements is positive.

// A subarray of an array is a consecutive sequence of zero or more values taken out of that array.

// Return the maximum length of a subarray with positive product.

// Example 1:
// Input: nums = [1,-2,-3,4]
// Output: 4
// Explanation: The array nums already has a positive product of 24.

// Example 2:
// Input: nums = [0,1,-2,-3,-4]
// Output: 3
// Explanation: The longest subarray with positive product is [1,-2,-3] which has a product of 6.
// Notice that we cannot include 0 in the subarray since that'll make the product 0 which is not positive.

// Example 3:
// Input: nums = [-1,-2,-3,0,1]
// Output: 2
// Explanation: The longest subarray with positive product is [-1,-2] or [-2,-3].

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func getMaxLen(nums []int) int {
    // dp[i] holds the max length of subarray with positive product
    // We start with the end of the array and work towards the beginning.
    res, lastneg, dp := 0, -1, make([]int, len(nums) + 1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] > 0 {
            dp[i] = dp[i + 1] + 1
        } else if nums[i] == 0 {
            dp[i] = 0
            lastneg = -1 // reset last negative value since we encountered a zero
        } else {
            if lastneg != -1 {
                dp[i] = dp[lastneg + 1] + lastneg - i + 1
                lastneg = i
            } else {
                dp[i] = 0
                lastneg = i
            }
        }
        res = max(res, dp[i])
    }
    return res
}

func getMaxLen1(nums []int) int {
    // dp[i][0] 以i结尾乘积为负数的子数组长度
    // dp[i][1] 以i结尾乘积为正数的子数组长度
    // nums[i] < 0 dp[i][0] = dp[i-1][1] + 1
    //             dp[i][1] = dp[i-1][0] == 0 ? 0 : dp[i-1][0] + 1
    // nums[i] > 0 dp[i][0] = dp[i-1][0] == 0 ? 0 : dp[i-1][0] + 1
    //             dp[i][1] = dp[i-1][1] + 1
    res, a, b, aa, bb := 0, 0, 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        i = i + 1
        if v < 0 {
            a, b = bb + 1, 0
            if aa > 0 {
                b = aa + 1
            }
        } else if v > 0 {
            a, b = 0, bb + 1
            if aa > 0 {
                a = aa + 1
            }
        } else {
            a, b = 0, 0
        }
        aa, bb = a, b
        res = max(res, b)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,-2,-3,4]
    // Output: 4
    // Explanation: The array nums already has a positive product of 24.
    fmt.Println(getMaxLen([]int{1,-2,-3,4})) // 4
    // Example 2:
    // Input: nums = [0,1,-2,-3,-4]
    // Output: 3
    // Explanation: The longest subarray with positive product is [1,-2,-3] which has a product of 6.
    // Notice that we cannot include 0 in the subarray since that'll make the product 0 which is not positive.
    fmt.Println(getMaxLen([]int{0,1,-2,-3,-4})) // 3
    // Example 3:
    // Input: nums = [-1,-2,-3,0,1]
    // Output: 2
    // Explanation: The longest subarray with positive product is [-1,-2] or [-2,-3].
    fmt.Println(getMaxLen([]int{-1,-2,-3,0,1})) // 2

    fmt.Println(getMaxLen1([]int{1,-2,-3,4})) // 4
    fmt.Println(getMaxLen1([]int{0,1,-2,-3,-4})) // 3
    fmt.Println(getMaxLen1([]int{-1,-2,-3,0,1})) // 2
}