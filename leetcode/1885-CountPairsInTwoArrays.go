package main

// 1885. Count Pairs in Two Arrays
// Given two integer arrays nums1 and nums2 of length n, 
// count the pairs of indices (i, j) such that i < j and nums1[i] + nums1[j] > nums2[i] + nums2[j].

// Return the number of pairs satisfying the condition.

// Example 1:
// Input: nums1 = [2,1,2,1], nums2 = [1,2,1,2]
// Output: 1
// Explanation: The pairs satisfying the condition are:
// - (0, 2) where 2 + 2 > 1 + 1.

// Example 2:
// Input: nums1 = [1,10,6,2], nums2 = [1,4,1,5]
// Output: 5
// Explanation: The pairs satisfying the condition are:
// - (0, 1) where 1 + 10 > 1 + 4.
// - (0, 2) where 1 + 6 > 1 + 1.
// - (1, 2) where 10 + 6 > 4 + 1.
// - (1, 3) where 10 + 2 > 4 + 5.
// - (2, 3) where 6 + 2 > 1 + 5.

// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     1 <= nums1[i], nums2[i] <= 10^5

import "fmt"
import "sort"

func countPairs(nums1 []int, nums2 []int) int64 {
    res, n := 0, len(nums1)
    //nums1[i] + nums1[j] > nums2[i] + nums2[j]
    //转化为  nums1[i] - nums2[i] >   nums2[j] - nums1[j]
    //differences[i] > −differences[j]
    //differences[i] + differences[j] > 0
    for i := 0; i < n; i++ {
        nums1[i] -= nums2[i]
    }
    sort.Ints(nums1)
    binarySearch := func(diffs []int,  n,  target, startIndex int) int {
        low, high := startIndex, n - 1
        if diffs[high] < target { return n }
        for low < high {
            mid := (low + high) >> 1
            if diffs[mid] < target {
                low = mid + 1
            } else {
                high = mid
            }
        }
        return low
    }
    for i := 0; i < n - 1; i++ {
        target := -nums1[i] + 1
        index := binarySearch(nums1, n, target, i + 1)
        res += n - index
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums1 = [2,1,2,1], nums2 = [1,2,1,2]
    // Output: 1
    // Explanation: The pairs satisfying the condition are:
    // - (0, 2) where 2 + 2 > 1 + 1.
    fmt.Println(countPairs([]int{2,1,2,1}, []int{1,2,1,2})) // 1
    // Example 2:
    // Input: nums1 = [1,10,6,2], nums2 = [1,4,1,5]
    // Output: 5
    // Explanation: The pairs satisfying the condition are:
    // - (0, 1) where 1 + 10 > 1 + 4.
    // - (0, 2) where 1 + 6 > 1 + 1.
    // - (1, 2) where 10 + 6 > 4 + 1.
    // - (1, 3) where 10 + 2 > 4 + 5.
    // - (2, 3) where 6 + 2 > 1 + 5.
    fmt.Println(countPairs([]int{1,10,6,2}, []int{1,4,1,5})) // 5
}