package main

// 3768. Minimum Inversion Count in Subarrays of Fixed Length
// You are given an integer array nums of length n and an integer k.

// An inversion is a pair of indices (i, j) from nums such that i < j and nums[i] > nums[j].

// The inversion count of a subarray is the number of inversions within it.

// Return the minimum inversion count among all subarrays of nums with length k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [3,1,2,5,4], k = 3
// Output: 0
// Explanation:
// We consider all subarrays of length k = 3 (indices below are relative to each subarray):
// [3, 1, 2] has 2 inversions: (0, 1) and (0, 2).
// [1, 2, 5] has 0 inversions.
// [2, 5, 4] has 1 inversion: (1, 2).
// The minimum inversion count among all subarrays of length 3 is 0, achieved by subarray [1, 2, 5].

// Example 2:
// Input: nums = [5,3,2,1], k = 4
// Output: 6
// Explanation:
// There is only one subarray of length k = 4: [5, 3, 2, 1].
// Within this subarray, the inversions are: (0, 1), (0, 2), (0, 3), (1, 2), (1, 3), and (2, 3).
// Total inversions is 6, so the minimum inversion count is 6.

// Example 3:
// Input: nums = [2,1], k = 1
// Output: 0
// Explanation:
// All subarrays of length k = 1 contain only one element, so no inversions are possible.
// The minimum inversion count is therefore 0.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= n

import "fmt"
import "slices"
import "sort"

type Fenwick []int

func (t Fenwick) update(i, val int) {
    for ; i < len(t); i += i & -i {
        t[i] += val
    }
}

func (t Fenwick) pre(i int) (res int) {
    for ; i > 0; i &= i - 1 {
        res += t[i]
    }
    return
}

func minInversionCount(nums []int, k int) int64 {
    // 离散化
    sorted := slices.Clone(nums)
    slices.Sort(sorted)
    sorted = slices.Compact(sorted)
    for i, x := range nums {
        nums[i] = sort.SearchInts(sorted, x) + 1 // 树状数组下标从 1 开始
    }
    t := make(Fenwick, len(sorted) + 1)
    res, inv := 1 << 61, 0
    for i, in := range nums {
        // 1. 入
        t.update(in, 1)
        inv += min(i+1, k) - t.pre(in) // 窗口大小 - (<=x 的元素个数) = (>x 的元素个数)
        left := i + 1 - k
        if left < 0 { continue } // 尚未形成第一个窗口
        // 2. 更新答案
        res = min(res, inv)
        // 3. 出
        out := nums[left]
        inv -= t.pre(out - 1) // < out 的元素个数
        t.update(out, -1)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,1,2,5,4], k = 3
    // Output: 0
    // Explanation:
    // We consider all subarrays of length k = 3 (indices below are relative to each subarray):
    // [3, 1, 2] has 2 inversions: (0, 1) and (0, 2).
    // [1, 2, 5] has 0 inversions.
    // [2, 5, 4] has 1 inversion: (1, 2).
    // The minimum inversion count among all subarrays of length 3 is 0, achieved by subarray [1, 2, 5].
    fmt.Println(minInversionCount([]int{3,1,2,5,4}, 3)) // 0
    // Example 2:
    // Input: nums = [5,3,2,1], k = 4
    // Output: 6
    // Explanation:
    // There is only one subarray of length k = 4: [5, 3, 2, 1].
    // Within this subarray, the inversions are: (0, 1), (0, 2), (0, 3), (1, 2), (1, 3), and (2, 3).
    // Total inversions is 6, so the minimum inversion count is 6.
    fmt.Println(minInversionCount([]int{5,3,2,1}, 4)) // 6
    // Example 3:
    // Input: nums = [2,1], k = 1
    // Output: 0
    // Explanation:
    // All subarrays of length k = 1 contain only one element, so no inversions are possible.
    // The minimum inversion count is therefore 0.
    fmt.Println(minInversionCount([]int{2,1}, 1)) // 0

    fmt.Println(minInversionCount([]int{1,2,3,4,5,6,7,8,9}, 3)) // 0
    fmt.Println(minInversionCount([]int{9,8,7,6,5,4,3,2,1}, 3)) // 3
}