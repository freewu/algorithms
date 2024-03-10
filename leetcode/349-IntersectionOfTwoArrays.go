package main

// 349. Intersection of Two Arrays
// Given two integer arrays nums1 and nums2, return an array of their intersection. 
// Each element in the result must be unique and you may return the result in any order.

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.
 
// Constraints:
//         1 <= nums1.length, nums2.length <= 1000
//         0 <= nums1[i], nums2[i] <= 1000

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
    // 长的放前面
    if len(nums1) < len(nums2) {
        intersection(nums2, nums1)
    }
    // 存在 map 中
    m1 := make(map[int]int)
    m2 := make(map[int]int)
    for _,v := range nums1 {
        m1[v]++
    }
    for _,v := range nums2 {
        m2[v]++
    }
    // 找出共有的
    res := []int{}
    for k, _ := range m1 {
        if m2[k] > 0 {
            res = append(res, k)
        }
    }
    return res
}

func intersection1(nums1 []int, nums2 []int) []int {
    res := make([]int, 0)
    m := make(map[int]bool, len(nums1))

    // 遍历 nums1
    for _, v := range nums1 {
        m[v] = false
    }
    // 遍历 nums2
    for _, v := range nums2 {
        // 如果存在  num2 的 map 中 且没有添加到 res 中时 加入到 res
        used, exists := m[v]
        if exists && !used {
            res = append(res, v)
            m[v] = true
        }
    }
    return res
}

func main() {
    fmt.Println(intersection([]int{1,2,2,1},[]int{2,2})) // [2]
    fmt.Println(intersection([]int{4,9,5},[]int{9,4,9,8,4})) // [9,4]

    fmt.Println(intersection1([]int{1,2,2,1},[]int{2,2})) // [2]
    fmt.Println(intersection1([]int{4,9,5},[]int{9,4,9,8,4})) // [9,4]
}