package main

// 3914. Minimum Operations to Make Array Non Decreasing
// You are given an integer array nums of length n.

// In one operation, you may choose any subarray nums[l..r] and increase each element in that subarray by x, where x is any positive integer.

// Return the minimum possible sum of the values of x across all operations required to make the array non-decreasing.

// An array is non-decreasing if nums[i] <= nums[i + 1] for all 0 <= i < n - 1.

// Example 1:
// Input: nums = [3,3,2,1]
// Output: 2
// Explanation:
// One optimal set of operations:
// Choose subarray [2..3] and add x = 1 resulting in [3, 3, 3, 2]
// Choose subarray [3..3] and add x = 1 resulting in [3, 3, 3, 3]
// The array becomes non-decreasing, and the total sum of chosen x values is 1 + 1 = 2.

// Example 2:
// Input: nums = [5,1,2,3]
// Output: 4
// Explanation:
// One optimal set of operations:
// Choose subarray [1..3] and add x = 4 resulting in [5, 5, 6, 7]
// The array becomes non-decreasing, and the total sum of chosen x values is 4.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func minOperations(nums []int) int64 {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
        res += max(nums[i-1] - nums[i], 0)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,3,2,1]
    // Output: 2
    // Explanation:
    // One optimal set of operations:
    // Choose subarray [2..3] and add x = 1 resulting in [3, 3, 3, 2]
    // Choose subarray [3..3] and add x = 1 resulting in [3, 3, 3, 3]
    // The array becomes non-decreasing, and the total sum of chosen x values is 1 + 1 = 2.
    fmt.Println(minOperations([]int{3,3,2,1})) // 2
    // Example 2:
    // Input: nums = [5,1,2,3]
    // Output: 4
    // Explanation:
    // One optimal set of operations:
    // Choose subarray [1..3] and add x = 4 resulting in [5, 5, 6, 7]
    // The array becomes non-decreasing, and the total sum of chosen x values is 4.
    fmt.Println(minOperations([]int{5,1,2,3})) // 4

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 8
}