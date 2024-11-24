package main

// 1995. Count Special Quadruplets
// Given a 0-indexed integer array nums, return the number of distinct quadruplets (a, b, c, d) such that:
//     nums[a] + nums[b] + nums[c] == nums[d], and
//     a < b < c < d

// Example 1:
// Input: nums = [1,2,3,6]
// Output: 1
// Explanation: The only quadruplet that satisfies the requirement is (0, 1, 2, 3) because 1 + 2 + 3 == 6.

// Example 2:
// Input: nums = [3,3,6,4,5]
// Output: 0
// Explanation: There are no such quadruplets in [3,3,6,4,5].

// Example 3:
// Input: nums = [1,1,1,3,5]
// Output: 4
// Explanation: The 4 quadruplets that satisfy the requirement are:
// - (0, 1, 2, 3): 1 + 1 + 1 == 3
// - (0, 1, 3, 4): 1 + 1 + 3 == 5
// - (0, 2, 3, 4): 1 + 1 + 3 == 5
// - (1, 2, 3, 4): 1 + 1 + 3 == 5

// Constraints:
//     4 <= nums.length <= 50
//     1 <= nums[i] <= 100

import "fmt"

func countQuadruplets(nums []int) int {
    res, n := 0, len(nums)
    for i := 0; i < n - 3; i++ {
        for j := i+1; j < n - 2; j++ {
            for k := j + 1; k < n - 1; k++ {
                for l := k + 1; l < n; l++ {
                    if nums[i] + nums[j] + nums[k] == nums[l] { // nums[a] + nums[b] + nums[c] == nums[d]
                        res++
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,6]
    // Output: 1
    // Explanation: The only quadruplet that satisfies the requirement is (0, 1, 2, 3) because 1 + 2 + 3 == 6.
    fmt.Println(countQuadruplets([]int{1,2,3,6})) // 1
    // Example 2:
    // Input: nums = [3,3,6,4,5]
    // Output: 0
    // Explanation: There are no such quadruplets in [3,3,6,4,5].
    fmt.Println(countQuadruplets([]int{3,3,6,4,5})) // 0
    // Example 3:
    // Input: nums = [1,1,1,3,5]
    // Output: 4
    // Explanation: The 4 quadruplets that satisfy the requirement are:
    // - (0, 1, 2, 3): 1 + 1 + 1 == 3
    // - (0, 1, 3, 4): 1 + 1 + 3 == 5
    // - (0, 2, 3, 4): 1 + 1 + 3 == 5
    // - (1, 2, 3, 4): 1 + 1 + 3 == 5
    fmt.Println(countQuadruplets([]int{1,1,1,3,5})) // 4
}