package main

// 3944. Minimum Operations to Make Array Modulo Alternating II
// You are given an integer array nums and an integer k.

// In one operation, you can increase or decrease any element of nums by 1.

// An array is called modulo alternating if there exist two distinct integers x and y (0 <= x, y < k) such that:
//     1. For every even index i, nums[i] % k == x
//     2. For every odd index i, nums[i] % k == y

// Return the minimum number of operations required to make nums modulo alternating.

// Example 1:
// Input: nums = [1,4,2,8], k = 3
// Output: 2
// Explanation:
// Let's choose x = 1 for even indices and y = 2 for odd indices.
// Perform the following operations:
// Increment nums[1] = 4 by 1, giving nums = [1, 5, 2, 8].
// Decrement nums[2] = 2 by 1, giving nums = [1, 5, 1, 8].
// Now, for even indices, nums[i] % k = 1, and for odd indices, nums[i] % k = 2.
// Thus, the total number of operations required is 2.

// Example 2:
// Input: nums = [1,1,1], k = 3
// Output: 1
// Explanation:
// Incrementing nums[1] by 1 gives nums = [1, 2, 1], which satisfies the condition with x = 1 and y = 2.
// Thus, the total number of operations required is 1.

// Example 3:
// Input: nums = [6,7,8], k = 2
// Output: 0
// Explanation:
// The array already satisfies the condition with x = 0 and y = 1. Thus, no operations are required.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     2 <= k <= 10^5

import "fmt"

func minOperations(nums []int, k int) int64 {
    n := len(nums)
    count := func(n, m, k int) int {
        m %= k
        diff := (m - n + k) % k
        return min(diff, k-diff)
    }
    // 初始化两个长度为k的切片，初始值都是0
    se1, se2 := make([]int, k),  make([]int, k)
    for i := 0; i < n; i += 2 { // 遍历偶数下标（0,2,4...）
        for j := 0; j < k; j++ {
            se1[j] += count(j, nums[i], k)
        }
    }
    for i := 1; i < n; i += 2 { // 遍历奇数下标（1,3,5...）
        for j := 0; j < k; j++ {
            se2[j] += count(j, nums[i], k)
        }
    }
    res := 1 << 61
    for j := 0; j < k; j++ { // 双重循环找最小值
        for i := 0; i < k; i++ {
            if i == j { continue }
            res = min(res, se1[j] + se2[i])
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,4,2,8], k = 3
    // Output: 2
    // Explanation:
    // Let's choose x = 1 for even indices and y = 2 for odd indices.
    // Perform the following operations:
    // Increment nums[1] = 4 by 1, giving nums = [1, 5, 2, 8].
    // Decrement nums[2] = 2 by 1, giving nums = [1, 5, 1, 8].
    // Now, for even indices, nums[i] % k = 1, and for odd indices, nums[i] % k = 2.
    // Thus, the total number of operations required is 2.
    fmt.Println(minOperations([]int{1,4,2,8}, 3)) // 2
    // Example 2:
    // Input: nums = [1,1,1], k = 3
    // Output: 1
    // Explanation:
    // Incrementing nums[1] by 1 gives nums = [1, 2, 1], which satisfies the condition with x = 1 and y = 2.
    // Thus, the total number of operations required is 1.
    fmt.Println(minOperations([]int{1,1,1}, 3)) // 1
    // Example 3:
    // Input: nums = [6,7,8], k = 2
    // Output: 0
    // Explanation:
    // The array already satisfies the condition with x = 0 and y = 1. Thus, no operations are required.
    fmt.Println(minOperations([]int{6,7,8}, 2)) // 0

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, 3)) // 5
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, 3)) // 5
}