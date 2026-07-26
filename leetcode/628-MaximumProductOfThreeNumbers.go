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

func maximumProduct1(nums []int) int {
    mx1, mx2, mx3, mn1, mn2 := -1000, -1000, -1000, 1000, 1000
    for i := 0; i < len(nums); i++ {
        v := nums[i]
        if v >= mx1 {
            mx3 = mx2
            mx2 = mx1
            mx1 = v
        } else if v >= mx2 {
            mx3 = mx2
            mx2 = v
        } else if v >= mx3 {
            mx3 = v
        }
        if v <= mn1 {
            mn2 = mn1
            mn1 = v
        } else if v <= mn2 {
            mn2 = v
        }
    }
    res := mx1 * mx2 * mx3
    if mx1 * mn1 * mn2 > res {
        res = mx1 * mn1 * mn2
    }
    return res
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
    fmt.Println(maximumProduct([]int{1,2,3,4,5,6,7,8,9})) // 504
    fmt.Println(maximumProduct([]int{9,8,7,6,5,4,3,2,1})) // 504

    fmt.Println(maximumProduct1([]int{1,2,3})) // 6
    fmt.Println(maximumProduct1([]int{1,2,3,4})) // 24
    fmt.Println(maximumProduct1([]int{-1,-2,-3})) // -6
    fmt.Println(maximumProduct1([]int{-100,-98,-1,2,3,4})) // 39200  4 * -98 * -100
    fmt.Println(maximumProduct1([]int{1,2,3,4,5,6,7,8,9})) // 504
    fmt.Println(maximumProduct1([]int{9,8,7,6,5,4,3,2,1})) // 504   
}