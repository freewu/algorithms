package main

// 3132. Find the Integer Added to Array II
// You are given two integer arrays nums1 and nums2.

// From nums1 two elements have been removed, and all other elements have been increased (or decreased in the case of negative) by an integer, represented by the variable x.

// As a result, nums1 becomes equal to nums2. 
// Two arrays are considered equal when they contain the same integers with the same frequencies.

// Return the minimum possible integer x that achieves this equivalence.

// Example 1:
// Input: nums1 = [4,20,16,12,8], nums2 = [14,18,10]
// Output: -2
// Explanation:
// After removing elements at indices [0,4] and adding -2, nums1 becomes [18,14,10].

// Example 2:
// Input: nums1 = [3,5,5,3], nums2 = [7,7]
// Output: 2
// Explanation:
// After removing elements at indices [0,3] and adding 2, nums1 becomes [7,7].

// Constraints:
//     3 <= nums1.length <= 200
//     nums2.length == nums1.length - 2
//     0 <= nums1[i], nums2[i] <= 1000
//     The test cases are generated in a way that there is an integer x such that nums1 can become equal to nums2 by removing two elements and adding x to each element of nums1.

import "fmt"
import "sort"

func minimumAddedInteger(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    res := 1 << 31
    f := func(x int) bool {
        i, j, cnt := 0, 0, 0
        for i < len(nums1) && j < len(nums2) {
            if nums2[j]-nums1[i] != x {
                cnt++
            } else {
                j++
            }
            i++
        }
        return cnt <= 2
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range nums1[:3] {
        x := nums2[0] - v
        if f(x) {
            res = min(res, x)
        }
    }
    return res
}

func minimumAddedInteger1(nums1, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    // 枚举保留 nums1[2] 或者 nums1[1] 或者 nums1[0]
    // 倒着枚举是因为 nums1[i] 越大答案越小，第一个满足的就是答案
    for i := 2; i > 0; i-- {
        x := nums2[0] - nums1[i]
        // 在 {nums1[i] + x} 中找子序列 nums2
        j := 0
        for _, v := range nums1[i:] {
            if nums2[j] == v+x {
                j++
                // nums2 是 {nums1[i] + x} 的子序列
                if j == len(nums2) {
                    return x
                }
            }
        }
    }
    // 题目保证答案一定存在
    return nums2[0] - nums1[0]
}

func main() {
    // Example 1:
    // Input: nums1 = [4,20,16,12,8], nums2 = [14,18,10]
    // Output: -2
    // Explanation:
    // After removing elements at indices [0,4] and adding -2, nums1 becomes [18,14,10].
    fmt.Println(minimumAddedInteger([]int{4,20,16,12,8},[]int{14,18,10})) // -2
    // Example 2:
    // Input: nums1 = [3,5,5,3], nums2 = [7,7]
    // Output: 2
    // Explanation:
    // After removing elements at indices [0,3] and adding 2, nums1 becomes [7,7].
    fmt.Println(minimumAddedInteger([]int{3,5,5,3},[]int{7,7})) // 2

    fmt.Println(minimumAddedInteger1([]int{4,20,16,12,8},[]int{14,18,10})) // -2
    fmt.Println(minimumAddedInteger1([]int{3,5,5,3},[]int{7,7})) // 2
}