package main

// 3471. Find the Largest Almost Missing Integer
// You are given an integer array nums and an integer k.

// An integer x is almost missing from nums if x appears in exactly one subarray of size k within nums.

// Return the largest almost missing integer from nums. If no such integer exists, return -1.

// A subarray is a contiguous sequence of elements within an array.

// Example 1:
// Input: nums = [3,9,2,1,7], k = 3
// Output: 7
// Explanation:
// 1 appears in 2 subarrays of size 3: [9, 2, 1] and [2, 1, 7].
// 2 appears in 3 subarrays of size 3: [3, 9, 2], [9, 2, 1], [2, 1, 7].
// 3 appears in 1 subarray of size 3: [3, 9, 2].
// 7 appears in 1 subarray of size 3: [2, 1, 7].
// 9 appears in 2 subarrays of size 3: [3, 9, 2], and [9, 2, 1].
// We return 7 since it is the largest integer that appears in exactly one subarray of size k.

// Example 2:
// Input: nums = [3,9,7,2,1,7], k = 4
// Output: 3
// Explanation:
// 1 appears in 2 subarrays of size 4: [9, 7, 2, 1], [7, 2, 1, 7].
// 2 appears in 3 subarrays of size 4: [3, 9, 7, 2], [9, 7, 2, 1], [7, 2, 1, 7].
// 3 appears in 1 subarray of size 4: [3, 9, 7, 2].
// 7 appears in 3 subarrays of size 4: [3, 9, 7, 2], [9, 7, 2, 1], [7, 2, 1, 7].
// 9 appears in 2 subarrays of size 4: [3, 9, 7, 2], [9, 7, 2, 1].
// We return 3 since it is the largest and only integer that appears in exactly one subarray of size k.

// Example 3:
// Input: nums = [0,0], k = 1
// Output: -1
// Explanation:
// There is no integer that appears in only one subarray of size 1.

// Constraints:
//     1 <= nums.length <= 50
//     0 <= nums[i] <= 50
//     1 <= k <= nums.length

import "fmt"

func largestInteger(nums []int, k int) int {
    count := make([]int, 51)
    mx, n := 0, len(nums)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        count[nums[i]] += min(i + 1, min(n - i, k))
        mx = max(mx, nums[i])
    }
    if k == n {
        return mx
    }
    for i := 50 ; i >=0 ; i-- {
        if count[i] == 1 {
            return i
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [3,9,2,1,7], k = 3
    // Output: 7
    // Explanation:
    // 1 appears in 2 subarrays of size 3: [9, 2, 1] and [2, 1, 7].
    // 2 appears in 3 subarrays of size 3: [3, 9, 2], [9, 2, 1], [2, 1, 7].
    // 3 appears in 1 subarray of size 3: [3, 9, 2].
    // 7 appears in 1 subarray of size 3: [2, 1, 7].
    // 9 appears in 2 subarrays of size 3: [3, 9, 2], and [9, 2, 1].
    // We return 7 since it is the largest integer that appears in exactly one subarray of size k.
    fmt.Println(largestInteger([]int{3,9,2,1,7}, 3)) // 7
    // Example 2:
    // Input: nums = [3,9,7,2,1,7], k = 4
    // Output: 3
    // Explanation:
    // 1 appears in 2 subarrays of size 4: [9, 7, 2, 1], [7, 2, 1, 7].
    // 2 appears in 3 subarrays of size 4: [3, 9, 7, 2], [9, 7, 2, 1], [7, 2, 1, 7].
    // 3 appears in 1 subarray of size 4: [3, 9, 7, 2].
    // 7 appears in 3 subarrays of size 4: [3, 9, 7, 2], [9, 7, 2, 1], [7, 2, 1, 7].
    // 9 appears in 2 subarrays of size 4: [3, 9, 7, 2], [9, 7, 2, 1].
    // We return 3 since it is the largest and only integer that appears in exactly one subarray of size k.
    fmt.Println(largestInteger([]int{3,9,7,2,1,7}, 4)) // 3
    // Example 3:
    // Input: nums = [0,0], k = 1
    // Output: -1
    // Explanation:
    // There is no integer that appears in only one subarray of size 1.
    fmt.Println(largestInteger([]int{0,0}, 1)) // -1

    fmt.Println(largestInteger([]int{1,2,3,4,5,6,7,8,9}, 4)) // 9
    fmt.Println(largestInteger([]int{9,8,7,6,5,4,3,2,1}, 4)) // 9
}