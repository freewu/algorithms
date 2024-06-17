package main

// 628. Maximum Product of Three Numbers
// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

// Example 1:
// Input: nums = [1,2,3]
// Output: 6

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 24

// Example 3:
// Input: nums = [-1,-2,-3]
// Output: -6

// Constraints:
//     3 <= nums.length <= 10^4
//     -1000 <= nums[i] <= 1000

import "fmt"
import "sort"

// func maximumProduct(nums []int) int {
//     sort.Ints(nums)
//     res, arr := 1, nums[len(nums) - 3:]
//     for _, v := range arr {
//         res *= v
//     }
//     return res
// }

func maximumProduct(nums []int) int {
    sort.Ints(nums)
    sum1 := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1] // 后三位的乘积
    if nums[0] < 0 && nums[1] < 0 && nums[len(nums)-1] > 0 { // 如果前面两位都为负数
        sum2 := nums[0] * nums[1] * nums[len(nums)-1]
        if sum2 > sum1 {
            return sum2
        }
    }
	return sum1
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 6
    fmt.Println(maximumProduct([]int{1,2,3})) // 6
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 24
    fmt.Println(maximumProduct([]int{1,2,3,4})) // 24
    // Example 3:
    // Input: nums = [-1,-2,-3]
    // Output: -6
    fmt.Println(maximumProduct([]int{-1,-2,-3})) // -6

    fmt.Println(maximumProduct([]int{-100,-98,-1,2,3,4})) // 39200  4 * -98 * -100
}