package main

// 2439. Minimize Maximum of Array
// You are given a 0-indexed array nums comprising of n non-negative integers.

// In one operation, you must:
//     Choose an integer i such that 1 <= i < n and nums[i] > 0.
//     Decrease nums[i] by 1.
//     Increase nums[i - 1] by 1.

// Return the minimum possible value of the maximum integer of nums after performing any number of operations.

// Example 1:
// Input: nums = [3,7,1,6]
// Output: 5
// Explanation:
// One set of optimal operations is as follows:
// 1. Choose i = 1, and nums becomes [4,6,1,6].
// 2. Choose i = 3, and nums becomes [4,6,2,5].
// 3. Choose i = 1, and nums becomes [5,5,2,5].
// The maximum integer of nums is 5. It can be shown that the maximum number cannot be less than 5.
// Therefore, we return 5.

// Example 2:
// Input: nums = [10,1]
// Output: 10
// Explanation:
// It is optimal to leave nums as is, and since 10 is the maximum value, we return 10.

// Constraints:
//     n == nums.length
//     2 <= n <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"

func minimizeArrayValue(nums []int) int {
    left, right, sum := 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        right = max(right, v)
        sum += v
    }
    left = sum / len(nums)
    for left < right {
        mid, total := (left + right) / 2, 0
        for i := len(nums) - 1; i > 0; i-- {
            total += (nums[i] - mid)
            if total < 0 {
                total = 0
            }
        }
        if nums[0] + total > mid {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func minimizeArrayValue1(nums []int) int {
    res, sum := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        sum += v
        res = max(res, (sum + i) / (i + 1))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,7,1,6]
    // Output: 5
    // Explanation:
    // One set of optimal operations is as follows:
    // 1. Choose i = 1, and nums becomes [4,6,1,6].
    // 2. Choose i = 3, and nums becomes [4,6,2,5].
    // 3. Choose i = 1, and nums becomes [5,5,2,5].
    // The maximum integer of nums is 5. It can be shown that the maximum number cannot be less than 5.
    // Therefore, we return 5.
    fmt.Println(minimizeArrayValue([]int{3,7,1,6})) // 5
    // Example 2:
    // Input: nums = [10,1]
    // Output: 10
    // Explanation:
    // It is optimal to leave nums as is, and since 10 is the maximum value, we return 10.
    fmt.Println(minimizeArrayValue([]int{10,1})) // 10

    fmt.Println(minimizeArrayValue([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(minimizeArrayValue([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(minimizeArrayValue1([]int{3,7,1,6})) // 5
    fmt.Println(minimizeArrayValue1([]int{10,1})) // 10
    fmt.Println(minimizeArrayValue1([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(minimizeArrayValue1([]int{9,8,7,6,5,4,3,2,1})) // 9
}