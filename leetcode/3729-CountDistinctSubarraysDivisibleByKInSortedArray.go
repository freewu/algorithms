package main

// 3729. Count Distinct Subarrays Divisible by K in Sorted Array
// You are given an integer array nums sorted in non-descending order and a positive integer k.

// A subarray of nums is good if the sum of its elements is divisible by k.

// Return an integer denoting the number of distinct good subarrays of nums.

// Subarrays are distinct if their sequences of values are. For example, there are 3 distinct subarrays in [1, 1, 1], namely [1], [1, 1], and [1, 1, 1].

// Example 1:
// Inpu
// Output: 3
// Explanation:
// The good subarrays are [1, 2], [3], and [1, 2, 3]. For example, [1, 2, 3] is good because the sum of its elements is 1 + 2 + 3 = 6, and 6 % k = 6 % 3 = 0.

// Example 2:
// Input: nums = [2,2,2,2,2,2], k = 6
// Output: 2
// Explanation:
// The good subarrays are [2, 2, 2] and [2, 2, 2, 2, 2, 2]. For example, [2, 2, 2] is good because the sum of its elements is 2 + 2 + 2 = 6, and 6 % k = 6 % 6 = 0.
// Note that [2, 2, 2] is counted only once.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     nums is sorted in non-descending order.
//     1 <= k <= 10^9

import "fmt"

func numGoodSubarrays(nums []int, k int) int64 {
    res, s, p, l := 0, 0, nums[0] - 1, 0
    mp := map[int]int{ 0: 1 }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for _, v := range nums {
        s = (s + v) % k
        res += mp[s]
        mp[s]++
        if v == p {
            l++
        } else {
            for m, d := 0, k / gcd(k, p); m + d <= l; m += d {
                res -= (l - m - d)
            }
            p, l = v, 1
        }
    }
    for m, d := 0, k / gcd(k, p); m + d <= l; m += d {
        res -= (l - m - d)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Inpu
    // Output: 3
    // Explanation:
    // The good subarrays are [1, 2], [3], and [1, 2, 3]. For example, [1, 2, 3] is good because the sum of its elements is 1 + 2 + 3 = 6, and 6 % k = 6 % 3 = 0.
    fmt.Println(numGoodSubarrays([]int{1,2,3}, 3)) // 3
    // Example 2:
    // Input: nums = [2,2,2,2,2,2], k = 6
    // Output: 2
    // Explanation:
    // The good subarrays are [2, 2, 2] and [2, 2, 2, 2, 2, 2]. For example, [2, 2, 2] is good because the sum of its elements is 2 + 2 + 2 = 6, and 6 % k = 6 % 6 = 0.
    // Note that [2, 2, 2] is counted only once.
    fmt.Println(numGoodSubarrays([]int{2,2,2,2,2,2}, 6)) // 2

    fmt.Println(numGoodSubarrays([]int{1,2,3,4,5,6,7,8,9}, 3)) // 24
    fmt.Println(numGoodSubarrays([]int{9,8,7,6,5,4,3,2,1}, 3)) // 24
}