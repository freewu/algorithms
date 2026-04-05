package main

// 3891. Minimum Increase to Maximize Special Indices
// You are given an integer array nums of length n.

// An index i (0 < i < n - 1) is special if nums[i] > nums[i - 1] and nums[i] > nums[i + 1].

// You may perform operations where you choose any index i and increase nums[i] by 1.

// Your goal is to:
//     1. Maximize the number of special indices.
//     2. Minimize the total number of operations required to achieve that maximum.

// Return an integer denoting the minimum total number of operations required.

// Example 1:
// Input: nums = [1,2,2]
// Output: 1
// Explanation:​​​​​​​
// Start with nums = [1, 2, 2].
// Increase nums[1] by 1, array becomes [1, 3, 2].
// The final array is [1, 3, 2] has 1 special index, which is the maximum achievable.
// It is impossible to achieve this number of special indices with fewer operations. Thus, the answer is 1.

// Example 2:
// Input: nums = [2,1,1,3]
// Output: 2
// Explanation:​​​​​​​
// Start with nums = [2, 1, 1, 3].
// Perform 2 operations at index 1, array becomes [2, 3, 1, 3].
// The final array is [2, 3, 1, 3] has 1 special index, which is the maximum achievable. Thus, the answer is 2.

// Example 3:
// Input: nums = [5,2,1,4,3]
// Output: 4
// Explanation:​​​​​​​​​​​​​​​​​​​​​
// Start with nums = [5, 2, 1, 4, 3].
// Perform 4 operations at index 1, array becomes [5, 6, 1, 4, 3].
// The final array is [5, 6, 1, 4, 3] has 2 special indices, which is the maximum achievable. Thus, the answer is 4.​​​​​​​
 
// Constraints:
//     3 <= n <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func minIncrease(nums []int) int64 {
    prefix, suffix, n := 0, 0, len(nums)
    for i := n - 2; i > 0; i -= 2 {
        suffix += max(max(nums[i-1], nums[i + 1]) - nums[i] + 1, 0)
    }
    if n % 2 > 0 {
        // 修改所有奇数下标
        return int64(suffix)
    }
    res := suffix // 修改 [2,n-2] 中的所有偶数下标
    // 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
    for i := 1; i < n-1; i += 2 {
        prefix += max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0)
        suffix -= max(max(nums[i], nums[i + 2]) - nums[i + 1] + 1, 0) // 撤销 i+1，撤销后 suffix 对应 [i+3,n-2]   
        res = min(res, prefix + suffix)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2]
    // Output: 1
    // Explanation:​​​​​​​
    // Start with nums = [1, 2, 2].
    // Increase nums[1] by 1, array becomes [1, 3, 2].
    // The final array is [1, 3, 2] has 1 special index, which is the maximum achievable.
    // It is impossible to achieve this number of special indices with fewer operations. Thus, the answer is 1.
    fmt.Println(minIncrease([]int{1,2,2})) // 1
    // Example 2:
    // Input: nums = [2,1,1,3]
    // Output: 2
    // Explanation:​​​​​​​
    // Start with nums = [2, 1, 1, 3].
    // Perform 2 operations at index 1, array becomes [2, 3, 1, 3].
    // The final array is [2, 3, 1, 3] has 1 special index, which is the maximum achievable. Thus, the answer is 2.
    fmt.Println(minIncrease([]int{2,1,1,3})) // 2
    // Example 3:
    // Input: nums = [5,2,1,4,3]
    // Output: 4
    // Explanation:​​​​​​​​​​​​​​​​​​​​​
    // Start with nums = [5, 2, 1, 4, 3].
    // Perform 4 operations at index 1, array becomes [5, 6, 1, 4, 3].
    // The final array is [5, 6, 1, 4, 3] has 2 special indices, which is the maximum achievable. Thus, the answer is 4.​​​​​​​
    fmt.Println(minIncrease([]int{5,2,1,4,3}))  // 4

    fmt.Println(minIncrease([]int{1,2,3,4,5,6,7,8,9}))  // 8
    fmt.Println(minIncrease([]int{9,8,7,6,5,4,3,2,1}))  // 8
}