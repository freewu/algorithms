package main

// 3641. Longest Semi-Repeating Subarray
// You are given an integer array nums of length n and an integer k.

// A semi‑repeating subarray is a contiguous subarray in which at most k elements repeat (i.e., appear more than once).

// Return the length of the longest semi‑repeating subarray in nums.

// Example 1:
// Input: nums = [1,2,3,1,2,3,4], k = 2
// Output: 6
// Explanation:
// The longest semi-repeating subarray is [2, 3, 1, 2, 3, 4], which has two repeating elements (2 and 3).

// Example 2:
// Input: nums = [1,1,1,1,1], k = 4
// Output: 5
// Explanation:
// The longest semi-repeating subarray is [1, 1, 1, 1, 1], which has only one repeating element (1).

// Example 3:
// Input: nums = [1,1,1,1,1], k = 0
// Output: 1
// Explanation:
// The longest semi-repeating subarray is [1], which has no repeating elements.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     0 <= k <= nums.length

import "fmt"

func longestSubarray(nums []int, k int) int {
    res, left, count, n := 0, 0, 0, len(nums)
    freq := make(map[int]int)  // 记录窗口中每个元素的出现频率
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        v := nums[i]
        freq[v]++
        // 当元素从出现1次变为出现2次时，重复元素数量增加1
        if freq[v] == 2 {
            count++
        }
        // 当重复元素数量超过k时，移动左指针缩小窗口
        for count > k {
            freq[nums[left]]--
            // 当元素从出现2次变为出现1次时，重复元素数量减少1
            if freq[nums[left]] == 1 {
                count--
            }
            left++
        }
        res = max(res,  i - left + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,1,2,3,4], k = 2
    // Output: 6
    // Explanation:
    // The longest semi-repeating subarray is [2, 3, 1, 2, 3, 4], which has two repeating elements (2 and 3).
    fmt.Println(longestSubarray([]int{1,2,3,1,2,3,4}, 2)) // 6
    // Example 2:
    // Input: nums = [1,1,1,1,1], k = 4
    // Output: 5
    // Explanation:
    // The longest semi-repeating subarray is [1, 1, 1, 1, 1], which has only one repeating element (1).
    fmt.Println(longestSubarray([]int{1,1,1,1,1}, 4)) // 5
    // Example 3:
    // Input: nums = [1,1,1,1,1], k = 0
    // Output: 1
    // Explanation:
    // The longest semi-repeating subarray is [1], which has no repeating elements.
    fmt.Println(longestSubarray([]int{1,1,1,1,1}, 0)) // 1

    fmt.Println(longestSubarray([]int{1,2,3,4,5,6,7,8,9}, 2)) // 9
    fmt.Println(longestSubarray([]int{9,8,7,6,5,4,3,2,1}, 2)) // 9
    fmt.Println(longestSubarray([]int{1,2,3,4,5,6,7,8,9}, 0)) // 9
    fmt.Println(longestSubarray([]int{9,8,7,6,5,4,3,2,1}, 0)) // 9
}