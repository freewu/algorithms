package main

// 454. 4Sum II
// Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, 
// return the number of tuples (i, j, k, l) such that:
//     0 <= i, j, k, l < n
//     nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
 
// Example 1:
// Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
// Output: 2
// Explanation:
// The two tuples are:
// 1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
// 2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

// Example 2:
// Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
// Output: 1
 
// Constraints:
//     n == nums1.length
//     n == nums2.length
//     n == nums3.length
//     n == nums4.length
//     1 <= n <= 200
//     -2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    res, m := 0, map[int]int{}
    for _, num1 := range nums1 {
        for _, num2 := range nums2 {
            m[num1 + num2]++
        }
    }
    for _, num3 := range nums3 {
        for _, num4 := range nums4 {
            if v, ok := m[-(num3 + num4)]; ok {
                res += v
            }
        }
    }
    return res
}

func fourSumCount1(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    res, m := 0, make(map[int]int, len(nums1) * len(nums1)) // val --> 下标出现次数
    for _, v1 := range nums1 {
        for _, v2 := range nums2 {
            m[v1 + v2]++
        }
    }
    for _, v3 := range nums3 {
        for _, v4 := range nums4 {
            res += m[-v3 - v4]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
    // Output: 2
    // Explanation:
    // The two tuples are:
    // 1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
    // 2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
    fmt.Println(fourSumCount([]int{1,2},[]int{-2,-1},[]int{-1,2},[]int{0,2})) // 2
    // Example 2:
    // Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
    // Output: 1
    fmt.Println(fourSumCount([]int{0},[]int{0},[]int{0},[]int{0})) // 1

    fmt.Println(fourSumCount1([]int{1,2},[]int{-2,-1},[]int{-1,2},[]int{0,2})) // 2
    fmt.Println(fourSumCount1([]int{0},[]int{0},[]int{0},[]int{0})) // 1
}