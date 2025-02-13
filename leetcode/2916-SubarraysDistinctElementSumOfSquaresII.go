package main

// 2916. Subarrays Distinct Element Sum of Squares II
// You are given a 0-indexed integer array nums.

// The distinct count of a subarray of nums is defined as:
//     1. Let nums[i..j] be a subarray of nums consisting of all the indices from i to j such that 0 <= i <= j < nums.length. 
//        Then the number of distinct values in nums[i..j] is called the distinct count of nums[i..j].

// Return the sum of the squares of distinct counts of all subarrays of nums.

// Since the answer may be very large, return it modulo 10^9 + 7.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,2,1]
// Output: 15
// Explanation: Six possible subarrays are:
// [1]: 1 distinct value
// [2]: 1 distinct value
// [1]: 1 distinct value
// [1,2]: 2 distinct values
// [2,1]: 2 distinct values
// [1,2,1]: 2 distinct values
// The sum of the squares of the distinct counts in all subarrays is equal to 12 + 12 + 12 + 22 + 22 + 22 = 15.

// Example 2:
// Input: nums = [2,2]
// Output: 3
// Explanation: Three possible subarrays are:
// [2]: 1 distinct value
// [2]: 1 distinct value
// [2,2]: 1 distinct value
// The sum of the squares of the distinct counts in all subarrays is equal to 12 + 12 + 12 = 3.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

// // Time Limit Exceeded 535 / 542
// func sumCounts(nums []int) int {
//     res, n, mod := 0, len(nums), 1_000_000_007
//     for i := 0; i < n; i++ {
//         exist, count := make(map[int]bool), 0
//         for j := i; j < n; j++ {
//             if !exist[nums[j] - 1] {
//                 exist[nums[j] - 1] = true
//                 count++
//             }
//             res += (count * count) % mod
//         }
//     }
//     return res % mod
// }

type LazySegment []struct{ sum, todo int }

func (t LazySegment) do(o, l, r, add int) {
    t[o].sum += add * (r - l + 1)
    t[o].todo += add
}

// o=1  [l,r] 1<=l<=r<=n
// 把 [L,R] 加一，同时返回加一之前的区间和
func (t LazySegment) queryAndAdd1(o, l, r, L, R int) int {
    res := 0
    if L <= l && r <= R {
        res = t[o].sum
        t.do(o, l, r, 1)
        return res
    }
    m := (l + r) >> 1
    if add := t[o].todo; add != 0 {
        t.do(o << 1, l, m, add)
        t.do(o << 1|1, m + 1, r, add)
        t[o].todo = 0
    }
    if L <= m {
        res = t.queryAndAdd1(o << 1, l, m, L, R)
    }
    if m < R {
        res += t.queryAndAdd1(o << 1|1, m+1, r, L, R)
    }
    t[o].sum = t[o << 1].sum + t[o << 1|1].sum
    return res
}

func sumCounts(nums []int) int {
    last := map[int]int{}
    res, sum, n := 0, 0, len(nums)
    t := make(LazySegment, n * 4)
    for i, x := range nums {
        i++
        j := last[x]
        sum += t.queryAndAdd1(1, 1, n, j+1, i) * 2 + i - j
        res = (res + sum) % 1_000_000_007
        last[x] = i
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1]
    // Output: 15
    // Explanation: Six possible subarrays are:
    // [1]: 1 distinct value
    // [2]: 1 distinct value
    // [1]: 1 distinct value
    // [1,2]: 2 distinct values
    // [2,1]: 2 distinct values
    // [1,2,1]: 2 distinct values
    // The sum of the squares of the distinct counts in all subarrays is equal to 12 + 12 + 12 + 22 + 22 + 22 = 15.
    fmt.Println(sumCounts([]int{1,2,1})) // 15
    // Example 2:
    // Input: nums = [1,1]
    // Output: 3
    // Explanation: Three possible subarrays are:
    // [1]: 1 distinct value
    // [1]: 1 distinct value
    // [1,1]: 1 distinct value
    // The sum of the squares of the distinct counts in all subarrays is equal to 12 + 12 + 12 = 3.
    fmt.Println(sumCounts([]int{1,1})) // 3

    fmt.Println(sumCounts([]int{1,2,3,4,5,6,7,8,9})) // 825
    fmt.Println(sumCounts([]int{9,8,7,6,5,4,3,2,1})) // 825

    // fmt.Println(sumCounts1([]int{1,2,1})) // 15
    // fmt.Println(sumCounts1([]int{1,1})) // 3
    // fmt.Println(sumCounts1([]int{1,2,3,4,5,6,7,8,9})) // 825
    // fmt.Println(sumCounts1([]int{9,8,7,6,5,4,3,2,1})) // 825
}