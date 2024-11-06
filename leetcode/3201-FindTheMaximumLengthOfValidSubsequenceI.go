package main

// 3201. Find the Maximum Length of Valid Subsequence I
// You are given an integer array nums.
// A subsequence sub of nums with length x is called valid if it satisfies:
//     (sub[0] + sub[1]) % 2 == (sub[1] + sub[2]) % 2 == ... == (sub[x - 2] + sub[x - 1]) % 2.

// Return the length of the longest valid subsequence of nums.

// A subsequence is an array that can be derived from another array by deleting some 
// or no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 4
// Explanation:
// The longest valid subsequence is [1, 2, 3, 4].

// Example 2:
// Input: nums = [1,2,1,1,2,1,2]
// Output: 6
// Explanation:
// The longest valid subsequence is [1, 2, 1, 2, 1, 2].

// Example 3:
// Input: nums = [1,3]
// Output: 2
// Explanation:
// The longest valid subsequence is [1, 3].

// Constraints:
//     2 <= nums.length <= 2 * 10^5
//     1 <= nums[i] <= 10^7

import "fmt"

func maximumLength(nums []int) int {
    flag := nums[0] % 2
    res, even, odd := 0, 0, 0 // alternating count, even, odd counts
    for _, v := range nums {
        if v % 2 == 0 {
            even++
        } else {
            odd++
        }
        if v % 2 == flag {
            flag ^= 1 // toggle flag value
            res++
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(res,max(even, odd))
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 4
    // Explanation:
    // The longest valid subsequence is [1, 2, 3, 4].
    fmt.Println(maximumLength([]int{1,2,3,4})) // 4
    // Example 2:
    // Input: nums = [1,2,1,1,2,1,2]
    // Output: 6
    // Explanation:
    // The longest valid subsequence is [1, 2, 1, 2, 1, 2].
    fmt.Println(maximumLength([]int{1,2,1,1,2,1,2})) // 6
    // Example 3:
    // Input: nums = [1,3]
    // Output: 2
    // Explanation:
    // The longest valid subsequence is [1, 3].
    fmt.Println(maximumLength([]int{1,3})) // 2
}