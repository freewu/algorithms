package main

// 1673. Find the Most Competitive Subsequence
// Given an integer array nums and a positive integer k, return the most competitive subsequence of nums of size k.
// An array's subsequence is a resulting sequence obtained by erasing some (possibly zero) elements from the array.

// We define that a subsequence a is more competitive than a subsequence b (of the same length) if in the first position where a and b differ, 
// subsequence a has a number less than the corresponding number in b. 
// For example, [1,3,4] is more competitive than [1,3,5] because the first position they differ is at the final number, and 4 is less than 5.

// Example 1:
// Input: nums = [3,5,2,6], k = 2
// Output: [2,6]
// Explanation: Among the set of every possible subsequence: {[3,5], [3,2], [3,6], [5,2], [5,6], [2,6]}, [2,6] is the most competitive.

// Example 2:
// Input: nums = [2,4,3,3,5,4,9,6], k = 4
// Output: [2,3,3,4]
 
// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9
//     1 <= k <= nums.length

import "fmt"

// stack
func mostCompetitive(nums []int, k int) []int {
    n, res := len(nums), make([]int, 0, k)
    for i := 0; i < n; i++ {
        for len(res) > 0 && k-len(res) < n-i && nums[i] < res[len(res)-1] {
            res = res[:len(res)-1]
        }
        if len(res) < k {
            res = append(res, nums[i])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,5,2,6], k = 2
    // Output: [2,6]
    // Explanation: Among the set of every possible subsequence: {[3,5], [3,2], [3,6], [5,2], [5,6], [2,6]}, [2,6] is the most competitive.
    fmt.Println(mostCompetitive([]int{3,5,2,6}, 2)) // [2,6]
    // Example 2:
    // Input: nums = [2,4,3,3,5,4,9,6], k = 4
    // Output: [2,3,3,4]
    fmt.Println(mostCompetitive([]int{2,4,3,3,5,4,9,6}, 4)) // [2,3,3,4]
}