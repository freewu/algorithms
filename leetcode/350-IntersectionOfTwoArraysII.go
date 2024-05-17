package main

// 350. Intersection of Two Arrays II
// Given two integer arrays nums1 and nums2, return an array of their intersection. 
// Each element in the result must appear as many times as it shows in both arrays 
// and you may return the result in any order.

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.
 
// Constraints:
//     1 <= nums1.length, nums2.length <= 1000
//     0 <= nums1[i], nums2[i] <= 1000
    
// Follow up:
//     What if the given array is already sorted? How would you optimize your algorithm?
//     What if nums1's size is small compared to nums2's size? Which algorithm is better?
//     What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
    unique, res := make(map[int]int), []int{}
    for _, v := range nums1 {
        unique[v]++
    }
    for _, v := range nums2 {
        if _, ok := unique[v]; ok && unique[v] > 0 {
            unique[v]--
            res = append(res, v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [1,2,2,1], nums2 = [2,2]
    // Output: [2,2]
    fmt.Println(intersect([]int{1,2,2,1}, []int{2,2})) // [2,2]
    // Example 2:
    // Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
    // Output: [4,9]
    // Explanation: [9,4] is also accepted.
    fmt.Println(intersect([]int{4,9,5}, []int{9,4,9,8,4})) // [4,9]
}