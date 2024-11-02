package main

// 2098. Subsequence of Size K With the Largest Even Sum
// You are given an integer array nums and an integer k. 
// Find the largest even sum of any subsequence of nums that has a length of k.

// Return this sum, or -1 if such a sum does not exist.

// A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [4,1,5,3,1], k = 3
// Output: 12
// Explanation:
// The subsequence with the largest possible even sum is [4,5,3]. It has a sum of 4 + 5 + 3 = 12.

// Example 2:
// Input: nums = [4,6,2], k = 3
// Output: 12
// Explanation:
// The subsequence with the largest possible even sum is [4,6,2]. It has a sum of 4 + 6 + 2 = 12.

// Example 3:
// Input: nums = [1,3,5], k = 1
// Output: -1
// Explanation:
// No subsequence of nums with length 1 has an even sum.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5
//     1 <= k <= nums.length

import "fmt"
import "sort"

// 先区分奇偶，再分别对其排序贪心选取最大的两个
func largestEvenSum(nums []int, k int) int64 {
    res, odd, even := 0, []int{}, []int{}
    for _, v := range nums {
        if v & 1 == 1 {
            odd = append(odd, v)
        } else {
            even = append(even, v)
        }
    }
    sort.Slice(odd, func(i, j int) bool { return odd[i] > odd[j] })
    sort.Slice(even, func(i, j int) bool { return even[i] > even[j] })
    i, j := 0, 0
    if k & 1 == 1 {
        k--
        if len(even) == 0 { return -1 }
        res += even[0]
        i++
    }
    for k > 0 {
        e, o := -1, -1
        if i + 1 < len(even) {
            e = even[i] + even[i + 1]
        }
        if j+1 < len(odd) {
            o = odd[j] + odd[j + 1]
        }
        if o == -1 && e == -1 { return -1 }
        if o > e{
            res, j = (res + o), (j + 2)
        } else {
            res, i = (res + e), (i + 2)
        }
        k -= 2
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [4,1,5,3,1], k = 3
    // Output: 12
    // Explanation:
    // The subsequence with the largest possible even sum is [4,5,3]. It has a sum of 4 + 5 + 3 = 12.
    fmt.Println(largestEvenSum([]int{4,1,5,3,1}, 3)) // 12
    // Example 2:
    // Input: nums = [4,6,2], k = 3
    // Output: 12
    // Explanation:
    // The subsequence with the largest possible even sum is [4,6,2]. It has a sum of 4 + 6 + 2 = 12.
    fmt.Println(largestEvenSum([]int{4,6,2}, 3)) // 12
    // Example 3:
    // Input: nums = [1,3,5], k = 1
    // Output: -1
    // Explanation:
    // No subsequence of nums with length 1 has an even sum.
    fmt.Println(largestEvenSum([]int{1,3,5}, 1)) // -1
}