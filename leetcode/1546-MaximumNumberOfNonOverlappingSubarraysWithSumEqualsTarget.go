package main

// 1546. Maximum Number of Non-Overlapping Subarrays With Sum Equals Target
// Given an array nums and an integer target, 
// return the maximum number of non-empty non-overlapping subarrays such that the sum of values in each subarray is equal to target.

// Example 1:
// Input: nums = [1,1,1,1,1], target = 2
// Output: 2
// Explanation: There are 2 non-overlapping subarrays [1,1,1,1,1] with sum equals to target(2).

// Example 2:
// Input: nums = [-1,3,5,1,4,2,-9], target = 6
// Output: 2
// Explanation: There are 3 subarrays with sum equal to 6.
// ([5,1], [4,2], [3,5,1,4,2,-9]) but only the first 2 are non-overlapping.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^4 <= nums[i] <= 10^4
//     0 <= target <= 10^6

import "fmt"

func maxNonOverlapping(nums []int, target int) int {
    mp := map[int]int{ 0: -1 }
    sum, res, pre := 0, 0, -1
    for i := range nums { // Prefix sum array
        sum += nums[i]
        if index, ok := mp[sum - target]; ok {
            if pre <= index {
                res++
                pre = i
            }
        }
        mp[sum] = i
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1,1,1], target = 2
    // Output: 2
    // Explanation: There are 2 non-overlapping subarrays [1,1,1,1,1] with sum equals to target(2).
    fmt.Println(maxNonOverlapping([]int{1,1,1,1,1}, 2)) // 2
    // Example 2:
    // Input: nums = [-1,3,5,1,4,2,-9], target = 6
    // Output: 2
    // Explanation: There are 3 subarrays with sum equal to 6.
    // ([5,1], [4,2], [3,5,1,4,2,-9]) but only the first 2 are non-overlapping.
    fmt.Println(maxNonOverlapping([]int{-1,3,5,1,4,2,-9}, 6)) // 2
}