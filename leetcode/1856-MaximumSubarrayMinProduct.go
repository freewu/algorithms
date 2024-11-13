package main

// 1856. Maximum Subarray Min-Product
// The min-product of an array is equal to the minimum value in the array multiplied by the array's sum.
//     For example, the array [3,2,5] (minimum value is 2) has a min-product of 2 * (3+2+5) = 2 * 10 = 20.

// Given an array of integers nums, return the maximum min-product of any non-empty subarray of nums. 
// Since the answer may be large, return it modulo 10^9 + 7.

// Note that the min-product should be maximized before performing the modulo operation. 
// Testcases are generated such that the maximum min-product without modulo will fit in a 64-bit signed integer.

// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [1,2,3,2]
// Output: 14
// Explanation: The maximum min-product is achieved with the subarray [2,3,2] (minimum value is 2).
// 2 * (2+3+2) = 2 * 7 = 14.

// Example 2:
// Input: nums = [2,3,3,1,2]
// Output: 18
// Explanation: The maximum min-product is achieved with the subarray [3,3] (minimum value is 3).
// 3 * (3+3) = 3 * 6 = 18.

// Example 3:
// Input: nums = [3,1,5,6,4,2]
// Output: 60
// Explanation: The maximum min-product is achieved with the subarray [5,6,4] (minimum value is 4).
// 4 * (5+6+4) = 4 * 15 = 60.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^7

import "fmt"

func maxSumMinProduct(nums []int) int {
    res, n := 0, len(nums)
    stack, prefix := [][2]int{}, make([]int, n + 1)
    for i := 0; i < n; i++ {
        prefix[i + 1] = prefix[i] + nums[i]
    }
    for i := 0; i < n; i++ {
        start := i
        for len(stack) > 0 && stack[len(stack) - 1][1] > nums[i] {
            v := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1] // pop
            res, start = max(res, (prefix[i] - prefix[v[0]]) * v[1]), v[0]
        }
        stack = append(stack, [2]int{ start, nums[i] })
    }
    for _, v := range stack {
        res = max(res, (prefix[n] - prefix[v[0]]) * v[1])
    }
    return res % 1_000_000_007
}

func maxSumMinProduct1(nums []int) int {
    n := len(nums)
    res, prefix, stack := 0, make([]int, n + 1), []int{}
    for i := 1; i <= n; i++ {
        prefix[i] = prefix[i - 1] + nums[i - 1]
    }
    left, right := make([]int, n), make([]int, n) // 以 nums[i] 为最小值 左右边界
    for i := 0; i < n; i++ {
        right[i], left[i] = n - 1, -1
    }
    for i, v := range nums {
        for len(stack) > 0 && nums[stack[len(stack) - 1]] >= v { // 右边界记录的是小于
            right[stack[len(stack) - 1]] = i - 1
            stack = stack[:len(stack) - 1]
        }
        if len(stack) > 0 {
            left[i] = stack[len(stack) - 1]
        }
        stack = append(stack, i)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        res = max(res, (nums[i] * (prefix[right[i] + 1] - prefix[left[i] + 1])))
    }
    return res % 1000000007
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,2]
    // Output: 14
    // Explanation: The maximum min-product is achieved with the subarray [2,3,2] (minimum value is 2).
    // 2 * (2+3+2) = 2 * 7 = 14.
    fmt.Println(maxSumMinProduct([]int{1,2,3,2})) // 14
    // Example 2:
    // Input: nums = [2,3,3,1,2]
    // Output: 18
    // Explanation: The maximum min-product is achieved with the subarray [3,3] (minimum value is 3).
    // 3 * (3+3) = 3 * 6 = 18.
    fmt.Println(maxSumMinProduct([]int{2,3,3,1,2})) // 18
    // Example 3:
    // Input: nums = [3,1,5,6,4,2]
    // Output: 60
    // Explanation: The maximum min-product is achieved with the subarray [5,6,4] (minimum value is 4).
    // 4 * (5+6+4) = 4 * 15 = 60.
    fmt.Println(maxSumMinProduct([]int{3,1,5,6,4,2})) // 60

    fmt.Println(maxSumMinProduct1([]int{1,2,3,2})) // 14
    fmt.Println(maxSumMinProduct1([]int{2,3,3,1,2})) // 18
    fmt.Println(maxSumMinProduct1([]int{3,1,5,6,4,2})) // 60
}