package main

// 3400. Maximum Number of Matching Indices After Right Shifts
// You are given two integer arrays, nums1 and nums2, of the same length.

// An index i is considered matching if nums1[i] == nums2[i].

// Return the maximum number of matching indices after performing any number of right shifts on nums1.

// A right shift is defined as shifting the element at index i to index (i + 1) % n, for all indices.

// Example 1:
// Input: nums1 = [3,1,2,3,1,2], nums2 = [1,2,3,1,2,3]
// Output: 6
// Explanation:
// If we right shift nums1 2 times, it becomes [1, 2, 3, 1, 2, 3]. Every index matches, so the output is 6.

// Example 2:
// Input: nums1 = [1,4,2,5,3,1], nums2 = [2,3,1,2,4,6]
// Output: 3
// Explanation:
// If we right shift nums1 3 times, it becomes [5, 3, 1, 1, 4, 2]. Indices 1, 2, and 4 match, so the output is 3.

// Constraints:
//     nums1.length == nums2.length
//     1 <= nums1.length, nums2.length <= 3000
//     1 <= nums1[i], nums2[i] <= 10^9

import "fmt"

func maximumMatchingIndices(nums1 []int, nums2 []int) int {
    res, n := -1, len(nums1)
    check := func(arr1, arr2 []int) int {
        res := 0
        for i := 0; i < n; i++ {
            if arr1[i] == arr2[i] {
                res++
            }
        }
        return res
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        tmp := nums1[n - i :] 
        tmp = append(tmp, nums1[: n - i]...)
        res = max(res, check(tmp, nums2))
    }
    return res
}

func maximumMatchingIndices1(nums1 []int, nums2 []int) int {
    res, n := 0, len(nums1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n && res < n; i++ {
        count := 0
        for j := 0; j < n; j++ {
            if nums1[j] == nums2[(j + i) % n] {
                count++
            }
        }
        res = max(res, count)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [3,1,2,3,1,2], nums2 = [1,2,3,1,2,3]
    // Output: 6
    // Explanation:
    // If we right shift nums1 2 times, it becomes [1, 2, 3, 1, 2, 3]. Every index matches, so the output is 6.
    fmt.Println(maximumMatchingIndices([]int{3,1,2,3,1,2}, []int{1,2,3,1,2,3})) // 6
    // Example 2:
    // Input: nums1 = [1,4,2,5,3,1], nums2 = [2,3,1,2,4,6]
    // Output: 3
    // Explanation:
    // If we right shift nums1 3 times, it becomes [5, 3, 1, 1, 4, 2]. Indices 1, 2, and 4 match, so the output is 3.
    fmt.Println(maximumMatchingIndices([]int{1,4,2,5,3,1}, []int{2,3,1,2,4,6})) // 3

    fmt.Println(maximumMatchingIndices1([]int{3,1,2,3,1,2}, []int{1,2,3,1,2,3})) // 6
    fmt.Println(maximumMatchingIndices1([]int{1,4,2,5,3,1}, []int{2,3,1,2,4,6})) // 3
}