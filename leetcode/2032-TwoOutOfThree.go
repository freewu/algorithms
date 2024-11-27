package main

// 2032. Two Out of Three
// Given three integer arrays nums1, nums2, and nums3, 
// return a distinct array containing all the values that are present in at least two out of the three arrays. 
// You may return the values in any order.
 
// Example 1:
// Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
// Output: [3,2]
// Explanation: The values that are present in at least two arrays are:
// - 3, in all three arrays.
// - 2, in nums1 and nums2.

// Example 2:
// Input: nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
// Output: [2,3,1]
// Explanation: The values that are present in at least two arrays are:
// - 2, in nums2 and nums3.
// - 3, in nums1 and nums2.
// - 1, in nums1 and nums3.

// Example 3:
// Input: nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
// Output: []
// Explanation: No value is present in at least two arrays.
 
// Constraints:
//     1 <= nums1.length, nums2.length, nums3.length <= 100
//     1 <= nums1[i], nums2[j], nums3[k] <= 100

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    count, visited := make(map[int]int), make(map[int]bool)
    for _, v := range nums1 {
        if !visited[v] {
            count[v]++
            visited[v] = true
        }
    }
    visited = make(map[int]bool) // 保证每次循环只被统计到一次即可
    for _, v := range nums2 {
        if !visited[v] {
            count[v]++
            visited[v] = true
        }
    }
    visited = make(map[int]bool) 
    for _, v := range nums3 {
        if !visited[v] {
            count[v]++
            visited[v] = true
        }
    }
    res := []int{}
    for k, v := range count {
        if v >= 2 { res = append(res, k)}
    }
    return res
}

func twoOutOfThree1(nums1 []int, nums2 []int, nums3 []int) []int {
    m1, m2 := 1_000, 1_000_000
    res, count := []int{}, make([]int, 101)
    for _, v := range nums1 { count[v] += 1  }
    for _, v := range nums2 { count[v] += m1 }
    for _, v := range nums3 { count[v] += m2 }
    for i := 1; i < 101; i++ {
        v := count[i]
        if v == 0 { continue }
        c1, c2, c3 := v % m1, (v / m1) % m1, (v / m2) % m1
        if (c1 > 0 && c2 > 0) || (c2 > 0 && c3 > 0) || (c1 > 0 && c3 > 0) {
            res = append(res, i)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
    // Output: [3,2]
    // Explanation: The values that are present in at least two arrays are:
    // - 3, in all three arrays.
    // - 2, in nums1 and nums2.
    fmt.Println(twoOutOfThree([]int{1,1,3,2}, []int{2,3}, []int{3})) // [3,2]
    // Example 2:
    // Input: nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
    // Output: [2,3,1]
    // Explanation: The values that are present in at least two arrays are:
    // - 2, in nums2 and nums3.
    // - 3, in nums1 and nums2.
    // - 1, in nums1 and nums3.
    fmt.Println(twoOutOfThree([]int{3,1}, []int{2,3}, []int{1,2})) // [2,3,1]
    // Example 3:
    // Input: nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
    // Output: []
    // Explanation: No value is present in at least two arrays.
    fmt.Println(twoOutOfThree([]int{1,2,2}, []int{4,3,3}, []int{5})) // []

    fmt.Println(twoOutOfThree1([]int{1,1,3,2}, []int{2,3}, []int{3})) // [3,2]
    fmt.Println(twoOutOfThree1([]int{3,1}, []int{2,3}, []int{1,2})) // [2,3,1]
    fmt.Println(twoOutOfThree1([]int{1,2,2}, []int{4,3,3}, []int{5})) // []
}