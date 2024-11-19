package main

// 2461. Maximum Sum of Distinct Subarrays With Length K
// You are given an integer array nums and an integer k. 
// Find the maximum subarray sum of all the subarrays of nums that meet the following conditions:
//     The length of the subarray is k, and
//     All the elements of the subarray are distinct.

// Return the maximum subarray sum of all the subarrays that meet the conditions. 
// If no subarray meets the conditions, return 0.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,5,4,2,9,9,9], k = 3
// Output: 15
// Explanation: The subarrays of nums with length 3 are:
// - [1,5,4] which meets the requirements and has a sum of 10.
// - [5,4,2] which meets the requirements and has a sum of 11.
// - [4,2,9] which meets the requirements and has a sum of 15.
// - [2,9,9] which does not meet the requirements because the element 9 is repeated.
// - [9,9,9] which does not meet the requirements because the element 9 is repeated.
// We return 15 because it is the maximum subarray sum of all the subarrays that meet the conditions

// Example 2:
// Input: nums = [4,4,4], k = 3
// Output: 0
// Explanation: The subarrays of nums with length 3 are:
// - [4,4,4] which does not meet the requirements because the element 4 is repeated.
// We return 0 because no subarrays meet the conditions.

// Constraints:
//     1 <= k <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func maximumSubarraySum(nums []int, k int) int64 {
    mp := make(map[int]int)
    res, sum,  l, r, n := 0, 0, 0, 0, len(nums)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for ; r < n; r++ {
        sum += nums[r]
        mp[nums[r]]++
        for mp[nums[r]] > 1 { // 当前有元素重复, 及时收缩左窗口
            sum -= nums[l]
            mp[nums[l]]--
            l++
        }
        for r > l + k - 1 { 
            sum -= nums[l]
            mp[nums[l]]--
            l++
        }
        if r == l + k - 1 { // 若当前元素为k个, 更新答案
            res = max(res, sum)
        }
    }
    return int64(res)
}


func main() {
    // Example 1:
    // Input: nums = [1,5,4,2,9,9,9], k = 3
    // Output: 15
    // Explanation: The subarrays of nums with length 3 are:
    // - [1,5,4] which meets the requirements and has a sum of 10.
    // - [5,4,2] which meets the requirements and has a sum of 11.
    // - [4,2,9] which meets the requirements and has a sum of 15.
    // - [2,9,9] which does not meet the requirements because the element 9 is repeated.
    // - [9,9,9] which does not meet the requirements because the element 9 is repeated.
    // We return 15 because it is the maximum subarray sum of all the subarrays that meet the conditions
    fmt.Println(maximumSubarraySum([]int{1,5,4,2,9,9,9}, 3)) // 15
    // Example 2:
    // Input: nums = [4,4,4], k = 3
    // Output: 0
    // Explanation: The subarrays of nums with length 3 are:
    // - [4,4,4] which does not meet the requirements because the element 4 is repeated.
    // We return 0 because no subarrays meet the conditions.
    fmt.Println(maximumSubarraySum([]int{4,4,4}, 3)) // 0
}