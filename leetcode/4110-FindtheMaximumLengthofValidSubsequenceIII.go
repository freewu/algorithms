package main

// 4110. Find the Maximum Length of Valid Subsequence III
// You are given an integer array nums and a positive integer k.

// A subsequence of nums is valid if the sum of every pair of adjacent elements in the subsequence is divisible by k.

// A subsequence of length 1 is always valid.

// Return the length of the longest valid subsequence of nums.

// Example 1:
// Input: nums = [1,2,3,4,5], k = 3
// Output: 4
// Explanation:
// Pick the subsequence [1, 2, 4, 5].
// The adjacent sums are 1 + 2 = 3, 2 + 4 = 6, and 4 + 5 = 9, all divisible by 3.
// No valid subsequence longer than 4 exists.

// Example 2:
// Input: nums = [5,10,15], k = 5
// Output: 3
// Explanation:
// Pick the whole array [5, 10, 15].
// The adjacent sums are 5 + 10 = 15 and 10 + 15 = 25, both divisible by 5.
// Every element is a multiple of 5, so any two elements can be adjacent in a valid subsequence, and the answer is 3.

// Example 3:
// Input: nums = [1,3,5,7,9], k = 4
// Output: 5
// Explanation:
// Pick the whole array [1, 3, 5, 7, 9].
// The adjacent sums are 1 + 3 = 4, 3 + 5 = 8, 5 + 7 = 12, and 7 + 9 = 16, all divisible by 4.
// The remainders alternate between 1 and 3, which add to 4, so the answer is 5.
 
// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^5

import "fmt"

func maximumLength(nums []int, k int) int {
    res, mp := 0, make(map[int]int)
    for _, v := range nums {
        x, y := v % k, (-v) % k
        // 负数取模修正，保证 y >=0
        if y < 0 {
            y += k
        }
        if mp[y]+1 > mp[x] {
            mp[x] = mp[y] + 1
        }
    }
    for _, v := range mp {
        if v > res {
            res = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], k = 3
    // Output: 4
    // Explanation:
    // Pick the subsequence [1, 2, 4, 5].
    // The adjacent sums are 1 + 2 = 3, 2 + 4 = 6, and 4 + 5 = 9, all divisible by 3.
    // No valid subsequence longer than 4 exists.
    fmt.Println(maximumLength([]int{1,2,3,4,5}, 3)) // 4
    // Example 2:
    // Input: nums = [5,10,15], k = 5
    // Output: 3
    // Explanation:
    // Pick the whole array [5, 10, 15].
    // The adjacent sums are 5 + 10 = 15 and 10 + 15 = 25, both divisible by 5.
    // Every element is a multiple of 5, so any two elements can be adjacent in a valid subsequence, and the answer is 3.
    fmt.Println(maximumLength([]int{5,10,15}, 5)) // 3
    // Example 3:
    // Input: nums = [1,3,5,7,9], k = 4
    // Output: 5
    // Explanation:
    // Pick the whole array [1, 3, 5, 7, 9].
    // The adjacent sums are 1 + 3 = 4, 3 + 5 = 8, 5 + 7 = 12, and 7 + 9 = 16, all divisible by 4.
    // The remainders alternate between 1 and 3, which add to 4, so the answer is 5.
    fmt.Println(maximumLength([]int{1,3,5,7,9}, 4)) // 5

    fmt.Println(maximumLength([]int{1,2,3,4,5,6,7,8,9}, 3)) // 6
    fmt.Println(maximumLength([]int{9,8,7,6,5,4,3,2,1}, 3)) // 6
}