package main

// 4118. Minimum Operations to Make Mountain Array I
// You are given an integer array nums of length n.

// In one operation, you may increment or decrement any element of nums by 1.

// All elements of nums must remain positive (at least 1) after every operation.

// An array is called a mountain array if there exists an index peak such that 0 < peak < n - 1 and:
//     1. nums[0] < nums[1] < ... < nums[peak]
//     2. nums[peak] > nums[peak + 1] > ... > nums[n - 1]

// Return the minimum number of operations required to transform nums into a mountain array. 
// If it is impossible, return -1.

// Example 1:
// Input: nums = [2,2,3]
// Output: 2
// Explanation:
// ​​​​​​​One optimal way is to:
//     Increment nums[1] by 1 and decrement nums[2] by 1.
//     The resulting array is [2, 3, 2], which is a mountain array with peak = 1.
//     Therefore, the total number of operations is 1 + 1 = 2.

// Example 2:
// Input: nums = [9,4,9]
// Output: 6
// Explanation:
// One optimal way is to:
//     Increment nums[1] by 1 six times.
//     The resulting array is [9, 10, 9], which is a mountain array with peak = 1.
//     Therefore, the total number of operations is 6.

// Example 3:
// Input: nums = [5,7,7,4]
// Output: 1
// Explanation:
// One optimal way is to:
//     Decrement nums[1] by 1.
//     The resulting array is [5, 6, 7, 4], which is a mountain array with peak = 2.
//     Therefore, the total number of operations is 1.

// Example 4:
// Input: nums = [1,2]
// Output: -1
// Explanation:
// A mountain array must have a peak with at least one element on each side. 
// Since n = 2 and operations do not change the array's length, forming a mountain array is impossible.

// Constraints:
//     1 <= n == nums.length <= 1000
//     1 <= nums[i] <= 1000

import "fmt"
import "slices"

func minOperations(nums []int) int {
    res, n := 1 << 61, len(nums)
    if n < 3 {
        return -1
    }
    mx := slices.Max(nums) + n - 1
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    dp := func() [][]int {
        f := make([][]int, n)
        for i := range f {
            f[i] = make([]int, mx)
        }
        for j := 1; j < mx; j++ {
            f[0][j] = abs(nums[0] - j)
        }
        for i := 1; i < n-1; i++ {
            mn := 1 << 61
            for j := i + 1; j < mx; j++ {
                mn = min(mn, f[i-1][j-1])
                f[i][j] = mn + abs(nums[i] - j)
            }
        }
        return f
    }
    prefix := dp()
    slices.Reverse(nums)
    suffix := dp()
    slices.Reverse(suffix)
    slices.Reverse(nums)
    // 枚举峰顶的位置和值
    for i := 1; i < n - 1; i++ {
        // 从左到右严格递增，j 至少是 i+1
        // 从右到左严格递增，j 至少是 n-i
        for j := max(i+1, n-i); j < mx; j++ {
            res = min(res, prefix[i][j] + suffix[i][j] - abs(nums[i] - j)) // 减去重复算的
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,2,3]
    // Output: 2
    // Explanation:
    // ​​​​​​​One optimal way is to:
    //     Increment nums[1] by 1 and decrement nums[2] by 1.
    //     The resulting array is [2, 3, 2], which is a mountain array with peak = 1.
    //     Therefore, the total number of operations is 1 + 1 = 2.
    fmt.Println(minOperations([]int{2,2,3})) // 2
    // Example 2:
    // Input: nums = [9,4,9]
    // Output: 6
    // Explanation:
    // One optimal way is to:
    //     Increment nums[1] by 1 six times.
    //     The resulting array is [9, 10, 9], which is a mountain array with peak = 1.
    //     Therefore, the total number of operations is 6.
    fmt.Println(minOperations([]int{9,4,9})) // 6
    // Example 3:
    // Input: nums = [5,7,7,4]
    // Output: 1
    // Explanation:
    // One optimal way is to:
    //     Decrement nums[1] by 1.
    //     The resulting array is [5, 6, 7, 4], which is a mountain array with peak = 2.
    //     Therefore, the total number of operations is 1.
    fmt.Println(minOperations([]int{5,7,7,4})) // 1
    // Example 4:
    // Input: nums = [1,2]
    // Output: -1
    // Explanation:
    // A mountain array must have a peak with at least one element on each side. 
    // Since n = 2 and operations do not change the array's length, forming a mountain array is impossible.
    fmt.Println(minOperations([]int{1,2})) // -1

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 2
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 2
}