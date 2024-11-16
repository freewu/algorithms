package main

// 2495. Number of Subarrays Having Even Product
// Given a 0-indexed integer array nums, return the number of subarrays of nums having an even product.

// Example 1:
// Input: nums = [9,6,7,13]
// Output: 6
// Explanation: There are 6 subarrays with an even product:
// - nums[0..1] = 9 * 6 = 54.
// - nums[0..2] = 9 * 6 * 7 = 378.
// - nums[0..3] = 9 * 6 * 7 * 13 = 4914.
// - nums[1..1] = 6.
// - nums[1..2] = 6 * 7 = 42.
// - nums[1..3] = 6 * 7 * 13 = 546.

// Example 2:
// Input: nums = [7,3,5]
// Output: 0
// Explanation: There are no subarrays with an even product.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func evenProduct(nums []int) int64 {
    res, left, right, n := 0, 0, 0, len(nums)
    for right < n {
        v := nums[right]
        right++
        if v & 1 == 0 {
            left = right
        }
        res += left
    }
    return int64(res)
}

func evenProduct1(nums []int) int64 {
    arr := []int{}
    for i, v := range nums { // 获取所有的偶数的位置
        if v & 1 == 0 {
            arr = append(arr, i)
        }
    }
    res, n1, n2 := 0, len(nums), len(arr)
    if n2 == 0 { return 0 } // 没有偶数 直接返回 0
    for i := 0; i < n2 - 1; i++ {
        res += (arr[i] + 1) * (arr[i + 1] - arr[i])
    }
    res += (n1 - arr[n2 - 1]) * (arr[n2 - 1] + 1)
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [9,6,7,13]
    // Output: 6
    // Explanation: There are 6 subarrays with an even product:
    // - nums[0..1] = 9 * 6 = 54.
    // - nums[0..2] = 9 * 6 * 7 = 378.
    // - nums[0..3] = 9 * 6 * 7 * 13 = 4914.
    // - nums[1..1] = 6.
    // - nums[1..2] = 6 * 7 = 42.
    // - nums[1..3] = 6 * 7 * 13 = 546.
    fmt.Println(evenProduct([]int{9,6,7,13})) // 6
    // Example 2:
    // Input: nums = [7,3,5]
    // Output: 0
    // Explanation: There are no subarrays with an even product.
    fmt.Println(evenProduct([]int{7,3,5})) // 0

    fmt.Println(evenProduct1([]int{9,6,7,13})) // 6
    fmt.Println(evenProduct1([]int{7,3,5})) // 0
}