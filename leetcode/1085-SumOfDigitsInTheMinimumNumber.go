package main

// 1085. Sum of Digits in the Minimum Number
// Given an integer array nums, return 0 if the sum of the digits of the minimum integer in nums is odd, or 1 otherwise.

// Example 1:
// Input: nums = [34,23,1,24,75,33,54,8]
// Output: 0
// Explanation: The minimal element is 1, and the sum of those digits is 1 which is odd, so the answer is 0.

// Example 2:
// Input: nums = [99,77,33,66,55]
// Output: 1
// Explanation: The minimal element is 33, and the sum of those digits is 3 + 3 = 6 which is even, so the answer is 1.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func sumOfDigits(nums []int) int {
    mn := 1 << 32 - 1
    for _, v := range nums { // 找到最小值 
        if v < mn {
            mn = v
        }
    }
    res := 0
    for mn > 0 { // 计算 各个数位上的数字
        res += mn % 10
        mn /= 10
    }
    if res % 2 == 0 {
        return 1
    }
    return 0
}

func main() {
    // Example 1:
    // Input: nums = [34,23,1,24,75,33,54,8]
    // Output: 0
    // Explanation: The minimal element is 1, and the sum of those digits is 1 which is odd, so the answer is 0.
    fmt.Println(sumOfDigits([]int{34,23,1,24,75,33,54,8})) // 0
    // Example 2:
    // Input: nums = [99,77,33,66,55]
    // Output: 1
    // Explanation: The minimal element is 33, and the sum of those digits is 3 + 3 = 6 which is even, so the answer is 1.
    fmt.Println(sumOfDigits([]int{99,77,33,66,55})) // 1
}