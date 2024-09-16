package main

// 1708. Largest Subarray Length K
// An array A is larger than some array B if for the first index i where A[i] != B[i], A[i] > B[i].

// For example, consider 0-indexing:
//     [1,3,2,4] > [1,2,2,4], since at index 1, 3 > 2.
//     [1,4,4,4] < [2,1,1,1], since at index 0, 1 < 2.

// A subarray is a contiguous subsequence of the array.

// Given an integer array nums of distinct integers, return the largest subarray of nums of length k.

// Example 1:
// Input: nums = [1,4,5,2,3], k = 3
// Output: [5,2,3]
// Explanation: The subarrays of size 3 are: [1,4,5], [4,5,2], and [5,2,3].
// Of these, [5,2,3] is the largest.

// Example 2:
// Input: nums = [1,4,5,2,3], k = 4
// Output: [4,5,2,3]
// Explanation: The subarrays of size 4 are: [1,4,5,2], and [4,5,2,3].
// Of these, [4,5,2,3] is the largest.

// Example 3:
// Input: nums = [1,4,5,2,3], k = 1
// Output: [5]

// Constraints:
//     1 <= k <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     All the integers of nums are unique.

// Follow up: What if the integers in nums are not distinct?

import "fmt"

// 在 0 ~ n-k 范围内，找最大值
func largestSubarray(nums []int, k int) []int {
    left, n := 0, len(nums)
    for right := 0; right <= n - k; right++ {
        if nums[right] > nums[left] {
            left = right
        }
    }
    return nums[left:left + k]
}

func largestSubarray1(nums []int, k int) []int {
    n := len(nums) 
    if n == k { return nums }
    // 枚举每个子数组的首元素，然后返回以该元素开始的子数组:
    mx, j := nums[0], 0 
    for i := 1; i <= (n - k); i++ {
        if nums[i] > mx {
            mx = nums[i]
            j = i 
        }
    }
    return nums[j:(j + k)] 
}

func main() {
    // Example 1:
    // Input: nums = [1,4,5,2,3], k = 3
    // Output: [5,2,3]
    // Explanation: The subarrays of size 3 are: [1,4,5], [4,5,2], and [5,2,3].
    // Of these, [5,2,3] is the largest.
    fmt.Println(largestSubarray([]int{1,4,5,2,3}, 3)) // [5,2,3]
    // Example 2:
    // Input: nums = [1,4,5,2,3], k = 4
    // Output: [4,5,2,3]
    // Explanation: The subarrays of size 4 are: [1,4,5,2], and [4,5,2,3].
    // Of these, [4,5,2,3] is the largest.
    fmt.Println(largestSubarray([]int{1,4,5,2,3}, 4)) // [4,5,2,3]
    // Example 3:
    // Input: nums = [1,4,5,2,3], k = 1
    // Output: [5]
    fmt.Println(largestSubarray([]int{1,4,5,2,3}, 1)) // [5]

    fmt.Println(largestSubarray1([]int{1,4,5,2,3}, 3)) // [5,2,3]
    fmt.Println(largestSubarray1([]int{1,4,5,2,3}, 4)) // [4,5,2,3]
    fmt.Println(largestSubarray1([]int{1,4,5,2,3}, 1)) // [5]
}