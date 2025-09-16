package main

// 3682. Minimum Index Sum of Common Elements
// You are given two integer arrays nums1 and nums2 of equal length n.

// We define a pair of indices (i, j) as a good pair if nums1[i] == nums2[j].

// Return the minimum index sum i + j among all possible good pairs. If no such pairs exist, return -1.

// Example 1:
// Input: nums1 = [3,2,1], nums2 = [1,3,1]
// Output: 1
// Explanation:
// Common elements between nums1 and nums2 are 1 and 3.
// For 3, [i, j] = [0, 1], giving an index sum of i + j = 1.
// For 1, [i, j] = [2, 0], giving an index sum of i + j = 2.
// The minimum index sum is 1.

// Example 2:
// Input: nums1 = [5,1,2], nums2 = [2,1,3]
// Output: 2
// Explanation:
// Common elements between nums1 and nums2 are 1 and 2.
// For 1, [i, j] = [1, 1], giving an index sum of i + j = 2.
// For 2, [i, j] = [2, 0], giving an index sum of i + j = 2.
// The minimum index sum is 2.

// Example 3:
// Input: nums1 = [6,4], nums2 = [7,8]
// Output: -1
// Explanation:
// Since no common elements between nums1 and nums2, the output is -1.

// Constraints:
//     1 <= nums1.length == nums2.length <= 10^5
//     -10^5 <= nums1[i], nums2[i] <= 10^5

import "fmt"

// 解答错误 991 / 999 
func minimumSum(nums1 []int, nums2 []int) int {
    if len(nums2) > len(nums1) {
        nums1, nums2 = nums2, nums1
    }
    mp := make(map[int]int)
    for i, v := range nums1 {
        if _, ok := mp[v]; ok { continue } // 只记录第一次
        mp[v] = i
    }
    res := 1 << 31
    for i, v := range nums2 {
        if j, ok := mp[v]; ok {
            if i + j < res {
                res = i + j
            }
            if j > res {
                break
            }
        }
    }
    if res == 1 << 31 {
        return -1
    }
    return res
}

func minimumSum1(nums1 []int, nums2 []int) int {
    res, mp := -1, make(map[int]int)
    for i, v := range nums1 {
        if j, ok := mp[v]; !ok || i < j { // 如果元素已存在，保留较小的索引
            mp[v] = i
        }
    }
    for j, v := range nums2 {
        if i, ok := mp[v]; ok {
            sum := i + j
            if res == -1 || sum < res { // 如果是第一个找到的好对，或者当前和更小，则更新最小值
                res = sum
            }
        }
    }
    return res
}

func minimumSum2(nums1 []int, nums2 []int) int {
    res, mp := 1 << 31, make(map[int]int)
    for i, v := range nums2 {
        if _, ok := mp[v]; !ok {
            mp[v] = i
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, v := range nums1 {
        if j, ok := mp[v]; ok {
            res = min(res, i + j)
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [3,2,1], nums2 = [1,3,1]
    // Output: 1
    // Explanation:
    // Common elements between nums1 and nums2 are 1 and 3.
    // For 3, [i, j] = [0, 1], giving an index sum of i + j = 1.
    // For 1, [i, j] = [2, 0], giving an index sum of i + j = 2.
    // The minimum index sum is 1.
    fmt.Println(minimumSum([]int{3,2,1}, []int{1,3,1})) // 1
    // Example 2:
    // Input: nums1 = [5,1,2], nums2 = [2,1,3]
    // Output: 2
    // Explanation:
    // Common elements between nums1 and nums2 are 1 and 2.
    // For 1, [i, j] = [1, 1], giving an index sum of i + j = 2.
    // For 2, [i, j] = [2, 0], giving an index sum of i + j = 2.
    // The minimum index sum is 2.
    fmt.Println(minimumSum([]int{5,1,2}, []int{2,1,3})) // 2
    // Example 3:
    // Input: nums1 = [6,4], nums2 = [7,8]
    // Output: -1
    // Explanation:
    // Since no common elements between nums1 and nums2, the output is -1.
    fmt.Println(minimumSum([]int{6,4}, []int{7,8})) // -1

    fmt.Println(minimumSum([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumSum([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 8
    fmt.Println(minimumSum([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(minimumSum([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(minimumSum1([]int{3,2,1}, []int{1,3,1})) // 1
    fmt.Println(minimumSum1([]int{5,1,2}, []int{2,1,3})) // 2
    fmt.Println(minimumSum1([]int{6,4}, []int{7,8})) // -1
    fmt.Println(minimumSum1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumSum1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 8
    fmt.Println(minimumSum1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(minimumSum1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(minimumSum2([]int{3,2,1}, []int{1,3,1})) // 1
    fmt.Println(minimumSum2([]int{5,1,2}, []int{2,1,3})) // 2
    fmt.Println(minimumSum2([]int{6,4}, []int{7,8})) // -1
    fmt.Println(minimumSum2([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumSum2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 8
    fmt.Println(minimumSum2([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(minimumSum2([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
}