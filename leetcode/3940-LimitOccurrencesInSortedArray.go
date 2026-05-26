package main

// 3940. Limit Occurrences in Sorted Array
// You are given a sorted integer array nums and an integer k.

// Return an array such that each distinct element appears at most k times, while preserving the relative order of the elements in nums.

// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,1,2,2,3]
// Explanation:
// Each element can appear at most 2 times.
// The element 1 appears 3 times, so only 2 occurrences are kept.
// The element 2 appears 2 times, so both occurrences are kept.
// The element 3 appears 1 time, so it is kept.
// Thus, the resulting array is [1, 1, 2, 2, 3].

// Example 2:
// Input: nums = [1,2,3], k = 1
// Output: [1,2,3]
// Explanation:
// All elements are distinct and already appear at most once, so the array remains unchanged.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100
//     nums is sorted in non-decreasing order.
//     1 <= k <= nums.length

import "fmt"

func limitOccurrences(nums []int, k int) []int {
    res, mp := []int{}, map[int]int{}
    for _, v := range nums {
        if mp[v] < k {
            res = append(res, v)
            mp[v]++
        }
    }
    return res
}

func limitOccurrences1(nums []int, k int) []int {
    res, count := []int{}, 0
    for i, v := range nums {
        if i == 0 || v != nums[i-1] {
            count = 1
        } else {
            count++
        }
        if count <= k {
            res = append(res, v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1,2,2,3], k = 2
    // Output: [1,1,2,2,3]
    // Explanation:
    // Each element can appear at most 2 times.
    // The element 1 appears 3 times, so only 2 occurrences are kept.
    // The element 2 appears 2 times, so both occurrences are kept.
    // The element 3 appears 1 time, so it is kept.
    // Thus, the resulting array is [1, 1, 2, 2, 3].
    fmt.Println(limitOccurrences([]int{1,1,1,2,2,3}, 2)) // [1,1,2,2,3]
    // Example 2:
    // Input: nums = [1,2,3], k = 1
    // Output: [1,2,3]
    // Explanation:
    // All elements are distinct and already appear at most once, so the array remains unchanged.
    fmt.Println(limitOccurrences([]int{1,2,3}, 1)) // [1,2,3]

    fmt.Println(limitOccurrences([]int{1,2,3,4,5,6,7,8,9}, 2)) // [1,2,3,4,5,6,7,8,9]
    fmt.Println(limitOccurrences([]int{9,8,7,6,5,4,3,2,1}, 2)) // [9,8,7,6,5,4,3,2,1]

    fmt.Println(limitOccurrences1([]int{1,1,1,2,2,3}, 2)) // [1,1,2,2,3]
    fmt.Println(limitOccurrences1([]int{1,2,3}, 1)) // [1,2,3]
    fmt.Println(limitOccurrences1([]int{1,2,3,4,5,6,7,8,9}, 2)) // [1,2,3,4,5,6,7,8,9]
    fmt.Println(limitOccurrences1([]int{9,8,7,6,5,4,3,2,1}, 2)) // [9,8,7,6,5,4,3,2,1]
}