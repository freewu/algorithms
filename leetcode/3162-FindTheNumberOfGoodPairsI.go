package main

// 3162. Find the Number of Good Pairs I
// You are given 2 integer arrays nums1 and nums2 of lengths n and m respectively. 
// You are also given a positive integer k.

// A pair (i, j) is called good if nums1[i] is divisible by nums2[j] * k (0 <= i <= n - 1, 0 <= j <= m - 1).
// Return the total number of good pairs.

// Example 1:
// Input: nums1 = [1,3,4], nums2 = [1,3,4], k = 1
// Output: 5
// Explanation:
// The 5 good pairs are (0, 0), (1, 0), (1, 1), (2, 0), and (2, 2).

// Example 2:
// Input: nums1 = [1,2,4,12], nums2 = [2,4], k = 3
// Output: 2
// Explanation:
// The 2 good pairs are (3, 0) and (3, 1).

// Constraints:
//     1 <= n, m <= 50
//     1 <= nums1[i], nums2[j] <= 50
//     1 <= k <= 50

import "fmt"

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    res := 0
    for _, v1 := range nums1 {
        for _, v2 := range nums2 {
            if (v1 % (v2 * k)) == 0 { // 如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [1,3,4], nums2 = [1,3,4], k = 1
    // Output: 5
    // Explanation:
    // The 5 good pairs are (0, 0), (1, 0), (1, 1), (2, 0), and (2, 2).
    fmt.Println(numberOfPairs([]int{1,3,4},[]int{1,3,4},1)) // 5
    // Example 2:
    // Input: nums1 = [1,2,4,12], nums2 = [2,4], k = 3
    // Output: 2
    // Explanation:
    // The 2 good pairs are (3, 0) and (3, 1).
    fmt.Println(numberOfPairs([]int{1,2,4,12},[]int{2,4},3)) // 2
}