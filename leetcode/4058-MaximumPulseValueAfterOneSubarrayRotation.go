package main

// 4058. Maximum Pulse Value After One Subarray Rotation
// You are given an integer array nums of length n.

// Define the pulse value of an integer array arr as the alternating sum starting at index 0: 
//     pulse(arr) = arr[0] - arr[1] + arr[2] - arr[3] + ....

// You may perform at most one operation on nums:
//     1. Choose two indices l and r such that 0 <= l < r < n.
//     2. Left-rotate the subarray nums[l..r] by exactly one position. For example, [a, b, c, d] becomes [b, c, d, a].

// Return the maximum pulse value that can be obtained after performing at most one such operation.

// Example 1:
// Input: nums = [1,5,2]
// Output: 6
// Explanation:
// The original pulse value is 1 - 5 + 2 = -2.
// Rotate the subarray nums[0..1] from [1, 5] to [5, 1].
// The resulting array is [5, 1, 2] and its pulse value is 5 - 1 + 2 = 6, which is the maximum possible.

// Example 2:
// Input: nums = [6,4,3]
// Output: 7
// Explanation:
// The original pulse value is 6 - 4 + 3 = 5.
// Rotate the subarray nums[1..2] from [4, 3] to [3, 4].
// The resulting array is [6, 3, 4] and its pulse value is 6 - 3 + 4 = 7, which is the maximum possible.

// Example 3:
// Input: nums = [9,7]
// Output: 2
// Explanation:
// The original pulse value is 9 - 7 = 2, which is already maximum. Thus, no rotation is required.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func maxValue(nums []int) int64 {
    f0, f1, mx,sum := 0, 0, 0, nums[0]
    for i := 1; i < len(nums); i++ {
        sum += nums[i] * (1 - i % 2 * 2)
        d := (nums[i] - nums[i-1]) * (i % 2 * 2 - 1)
        f0, f1 = f1, max(f0, 0)+d
        mx = max(mx, f1)
    }
    return int64(sum + mx * 2)
}

func main() {
    // Example 1:
    // Input: nums = [1,5,2]
    // Output: 6
    // Explanation:
    // The original pulse value is 1 - 5 + 2 = -2.
    // Rotate the subarray nums[0..1] from [1, 5] to [5, 1].
    // The resulting array is [5, 1, 2] and its pulse value is 5 - 1 + 2 = 6, which is the maximum possible.
    fmt.Println(maxValue([]int{1,5,2})) // 6
    // Example 2:
    // Input: nums = [6,4,3]
    // Output: 7
    // Explanation:
    // The original pulse value is 6 - 4 + 3 = 5.
    // Rotate the subarray nums[1..2] from [4, 3] to [3, 4].
    // The resulting array is [6, 3, 4] and its pulse value is 6 - 3 + 4 = 7, which is the maximum possible.
    fmt.Println(maxValue([]int{6,4,3})) // 7
    // Example 3:
    // Input: nums = [9,7]
    // Output: 2
    // Explanation:
    // The original pulse value is 9 - 7 = 2, which is already maximum. Thus, no rotation is required.
    fmt.Println(maxValue([]int{9,7})) // 2

    fmt.Println(maxValue([]int{1,2,3,4,5,6,7,8,9})) // 13
    fmt.Println(maxValue([]int{9,8,7,6,5,4,3,2,1})) // 13
}