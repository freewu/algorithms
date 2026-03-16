package main

// 3868. Minimum Cost to Equalize Arrays Using Swaps
// You are given two integer arrays nums1 and nums2 of size n.

// You can perform the following two operations any number of times on these two arrays:
//     1. Swap within the same array: Choose two indices i and j. 
//        Then, choose either to swap nums1[i] and nums1[j], or nums2[i] and nums2[j]. 
//        This operation is free of charge.
//     2. Swap between two arrays: Choose an index i. 
//        Then, swap nums1[i] and nums2[i]. This operation incurs a cost of 1.

// Return an integer denoting the minimum cost to make nums1 and nums2 identical. 
// If this is not possible, return -1.

// Example 1:
// Input: nums1 = [10,20], nums2 = [20,10]
// Output: 0
// Explanation:
// Swap nums2[0] = 20 and nums2[1] = 10.
// nums2 becomes [10, 20].
// This operation is free of charge.
// nums1 and nums2 are now identical. The cost is 0.

// Example 2:
// Input: nums1 = [10,10], nums2 = [20,20]
// Output: 1
// Explanation:
// Swap nums1[0] = 10 and nums2[0] = 20.
// nums1 becomes [20, 10].
// nums2 becomes [10, 20].
// This operation costs 1.
// Swap nums2[0] = 10 and nums2[1] = 20.
// nums2 becomes [20, 10].
// This operation is free of charge.
// nums1 and nums2 are now identical. The cost is 1.

// Example 3:
// Input: nums1 = [10,20], nums2 = [30,40]
// Output: -1
// Explanation:
// It is impossible to make the two arrays identical. Therefore, the answer is -1.

// Constraints:
//     2 <= n == nums1.length == nums2.length <= 8 * 10^4
//     1 <= nums1[i], nums2[i] <= 8 * 10^4

import "fmt"

func minCost(nums1 []int, nums2 []int) int {
    freq := map[int]int{}
    for _, v := range nums1 { // 统计 nums1 中每个元素的出现次数        
        freq[v]++
    }
    for _, v := range nums2 { // 减去 nums2 中每个元素的出现次数        
        freq[v]--
    }
    diff := 0
    for _, v := range freq {
        if v % 2 != 0 {
            return -1
        }
        if v < 0 {
            v = -v
        }
        diff += v
    }
    return diff / 4
}

func minCost1(nums1 []int, nums2 []int) int {
    mp := [80009]int{0}
    for _, v := range nums1 {
        mp[v]++
    }
    for _, v := range nums2 {
        mp[v]++
    }
    for _, v := range nums1 {
        if mp[v] % 2 == 1 {
            return -1
        }
    }
    for _, v := range nums2 {
        if mp[v] % 2 == 1 {
            return -1
        }
    }
    count := 0
    for i, v := range nums2 {
        if mp[v] == 0 {
            nums1[i], nums2[i] = nums2[i], nums1[i]
            count++
        } else {
            mp[v] -= 2
        }
    }
    return count
}

func main() {
    // Example 1:
    // Input: nums1 = [10,20], nums2 = [20,10]
    // Output: 0
    // Explanation:
    // Swap nums2[0] = 20 and nums2[1] = 10.
    // nums2 becomes [10, 20].
    // This operation is free of charge.
    // nums1 and nums2 are now identical. The cost is 0.
    fmt.Println(minCost([]int{10,20}, []int{20,10})) // 0
    // Example 2:
    // Input: nums1 = [10,10], nums2 = [20,20]
    // Output: 1
    // Explanation:
    // Swap nums1[0] = 10 and nums2[0] = 20.
    // nums1 becomes [20, 10].
    // nums2 becomes [10, 20].
    // This operation costs 1.
    // Swap nums2[0] = 10 and nums2[1] = 20.
    // nums2 becomes [20, 10].
    // This operation is free of charge.
    // nums1 and nums2 are now identical. The cost is 1.
    fmt.Println(minCost([]int{10,10}, []int{20,20})) // 1
    // Example 3:
    // Input: nums1 = [10,20], nums2 = [30,40]
    // Output: -1
    // Explanation:
    // It is impossible to make the two arrays identical. Therefore, the answer is -1.
    fmt.Println(minCost([]int{10,20}, []int{30,40})) // -1

    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(minCost1([]int{10,20}, []int{20,10})) // 0
    fmt.Println(minCost1([]int{10,10}, []int{20,20})) // 1
    fmt.Println(minCost1([]int{10,20}, []int{30,40})) // -1
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
}