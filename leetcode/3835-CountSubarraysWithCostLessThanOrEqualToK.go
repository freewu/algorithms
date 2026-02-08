package main

// 3835. Count Subarrays With Cost Less Than or Equal to K
// You are given an integer array nums, and an integer k.

// For any subarray nums[l..r], define its cost as:

//     cost = (max(nums[l..r]) - min(nums[l..r])) * (r - l + 1).

// Return an integer denoting the number of subarrays of nums whose cost is less than or equal to k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,3,2], k = 4
// Output: 5
// Explanation:
// We consider all subarrays of nums:
// nums[0..0]: cost = (1 - 1) * 1 = 0
// nums[0..1]: cost = (3 - 1) * 2 = 4
// nums[0..2]: cost = (3 - 1) * 3 = 6
// nums[1..1]: cost = (3 - 3) * 1 = 0
// nums[1..2]: cost = (3 - 2) * 2 = 2
// nums[2..2]: cost = (2 - 2) * 1 = 0
// There are 5 subarrays whose cost is less than or equal to 4.

// Example 2:
// Input: nums = [5,5,5,5], k = 0
// Output: 10
// Explanation:
// For any subarray of nums, the maximum and minimum values are the same, so the cost is always 0.
// As a result, every subarray of nums has cost less than or equal to 0.
// For an array of length 4, the total number of subarrays is (4 * 5) / 2 = 10.

// Example 3:
// Input: nums = [1,2,3], k = 0
// Output: 3
// Explanation:
// The only subarrays of nums with cost 0 are the single-element subarrays, and there are 3 of them.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= k <= 10^15

import "fmt"

func countSubarrays(nums []int, k int64) int64 {
    minQ, maxQ := []int{}, []int{}
    res, left := int64(0), 0
    for right, x := range nums {
        // 元素进入窗口
        for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
            minQ = minQ[:len(minQ)-1]
        }
        minQ = append(minQ, right)
        for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
            maxQ = maxQ[:len(maxQ)-1]
        }
        maxQ = append(maxQ, right)
        // 如果窗口不满足要求，左端点离开窗口
        for int64(nums[maxQ[0]]-nums[minQ[0]]) * int64(right-left+1) > k {
            left++
            if minQ[0] < left {
                minQ = minQ[1:]
            }
            if maxQ[0] < left {
                maxQ = maxQ[1:]
            }
        }
        res += int64(right - left + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2], k = 4
    // Output: 5
    // Explanation:
    // We consider all subarrays of nums:
    // nums[0..0]: cost = (1 - 1) * 1 = 0
    // nums[0..1]: cost = (3 - 1) * 2 = 4
    // nums[0..2]: cost = (3 - 1) * 3 = 6
    // nums[1..1]: cost = (3 - 3) * 1 = 0
    // nums[1..2]: cost = (3 - 2) * 2 = 2
    // nums[2..2]: cost = (2 - 2) * 1 = 0
    // There are 5 subarrays whose cost is less than or equal to 4.
    fmt.Println(countSubarrays([]int{1,3,2}, 4)) // 5
    // Example 2:
    // Input: nums = [5,5,5,5], k = 0
    // Output: 10
    // Explanation:
    // For any subarray of nums, the maximum and minimum values are the same, so the cost is always 0.
    // As a result, every subarray of nums has cost less than or equal to 0.
    // For an array of length 4, the total number of subarrays is (4 * 5) / 2 = 10.
    fmt.Println(countSubarrays([]int{5,5,5,5}, 0)) // 10
    // Example 3:
    // Input: nums = [1,2,3], k = 0
    // Output: 3
    // Explanation:
    // The only subarrays of nums with cost 0 are the single-element subarrays, and there are 3 of them. 
    fmt.Println(countSubarrays([]int{1,2,3}, 0)) // 3

    fmt.Println(countSubarrays([]int{1,2,3,4,5,6,7,8,9}, 4)) // 17
    fmt.Println(countSubarrays([]int{9,8,7,6,5,4,3,2,1}, 4)) // 17
}