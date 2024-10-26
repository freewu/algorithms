package main

// 1983. Widest Pair of Indices With Equal Range Sum
// You are given two 0-indexed binary arrays nums1 and nums2.
// Find the widest pair of indices (i, j) such that i <= j and nums1[i] + nums1[i+1] + ... + nums1[j] == nums2[i] + nums2[i+1] + ... + nums2[j].

// The widest pair of indices is the pair with the largest distance between i and j. 
// The distance between a pair of indices is defined as j - i + 1.

// Return the distance of the widest pair of indices. 
// If no pair of indices meets the conditions, return 0.

// Example 1:
// Input: nums1 = [1,1,0,1], nums2 = [0,1,1,0]
// Output: 3
// Explanation:
// If i = 1 and j = 3:
// nums1[1] + nums1[2] + nums1[3] = 1 + 0 + 1 = 2.
// nums2[1] + nums2[2] + nums2[3] = 1 + 1 + 0 = 2.
// The distance between i and j is j - i + 1 = 3 - 1 + 1 = 3.

// Example 2:
// Input: nums1 = [0,1], nums2 = [1,1]
// Output: 1
// Explanation:
// If i = 1 and j = 1:
// nums1[1] = 1.
// nums2[1] = 1.
// The distance between i and j is j - i + 1 = 1 - 1 + 1 = 1.

// Example 3:
// Input: nums1 = [0], nums2 = [1]
// Output: 0
// Explanation:
// There are no pairs of indices that meet the requirements.

// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     nums1[i] is either 0 or 1.
//     nums2[i] is either 0 or 1.

import "fmt"

func widestPairOfIndices(nums1 []int, nums2 []int) int {
    res, sum1, sum2 := 0, 0, 0
    mp := make(map[int]int)
    mp[0] = -1
    for i := 0; i < len(nums1); i++ {
        sum1 += nums1[i]
        sum2 += nums2[i]
        if v, ok := mp[sum1 - sum2]; ok {
            if i - v > res {
                res = i - v
            }
        } else {
            mp[sum1 - sum2] = i // 前缀和记录每个下标nums1,nums2和的差值
        }
    }
    return res
}

func widestPairOfIndices1(nums1, nums2 []int) int {
    res, sum1, sum2, n := 0, 0, 0, len(nums1) 
    mp1, mp2 := make([]int, n + 1), make([]int, n + 1) 
    for i := range mp1 { mp1[i], mp2[i] = -1, -1 } // fill -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i += 1 {
        sum1 += nums1[i]
        sum2 += nums2[i] 
        if sum1 == sum2 {
            res = max(res, i + 1)
        } else {
            diff := sum1 - sum2 
            if diff > 0 {
                if mp1[diff] >= 0 {
                    res = max(res, i - mp1[diff]) 
                } else {
                    mp1[diff] = i 
                }
            } else {
                if mp2[-diff] >= 0 {
                    res = max(res, i - mp2[-diff]) 
                } else {
                    mp2[-diff] = i 
                }
            }
        }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: nums1 = [1,1,0,1], nums2 = [0,1,1,0]
    // Output: 3
    // Explanation:
    // If i = 1 and j = 3:
    // nums1[1] + nums1[2] + nums1[3] = 1 + 0 + 1 = 2.
    // nums2[1] + nums2[2] + nums2[3] = 1 + 1 + 0 = 2.
    // The distance between i and j is j - i + 1 = 3 - 1 + 1 = 3.
    fmt.Println(widestPairOfIndices([]int{1,1,0,1}, []int{0,1,1,0})) // 3
    // Example 2:
    // Input: nums1 = [0,1], nums2 = [1,1]
    // Output: 1
    // Explanation:
    // If i = 1 and j = 1:
    // nums1[1] = 1.
    // nums2[1] = 1.
    // The distance between i and j is j - i + 1 = 1 - 1 + 1 = 1.
    fmt.Println(widestPairOfIndices([]int{0,1}, []int{1,1})) // 1
    // Example 3:
    // Input: nums1 = [0], nums2 = [1]
    // Output: 0
    // Explanation:
    // There are no pairs of indices that meet the requirements.
    fmt.Println(widestPairOfIndices([]int{0}, []int{1})) // 0

    fmt.Println(widestPairOfIndices1([]int{1,1,0,1}, []int{0,1,1,0})) // 3
    fmt.Println(widestPairOfIndices1([]int{0,1}, []int{1,1})) // 1
    fmt.Println(widestPairOfIndices1([]int{0}, []int{1})) // 0
}