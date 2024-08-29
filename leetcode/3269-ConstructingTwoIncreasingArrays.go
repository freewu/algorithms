package main

// 3269. Constructing Two Increasing Arrays
// Given 2 integer arrays nums1 and nums2 consisting only of 0 and 1,
// your task is to calculate the minimum possible largest number in arrays nums1 and nums2, after doing the following.

// Replace every 0 with an even positive integer and every 1 with an odd positive integer. 
// After replacement, both arrays should be increasing and each integer should be used at most once.

// Return the minimum possible largest number after applying the changes.

// Example 1:
// Input: nums1 = [], nums2 = [1,0,1,1]
// Output: 5
// Explanation:
// After replacing, nums1 = [], and nums2 = [1, 2, 3, 5].

// Example 2:
// Input: nums1 = [0,1,0,1], nums2 = [1,0,0,1]
// Output: 9
// Explanation:
// One way to replace, having 9 as the largest element is nums1 = [2, 3, 8, 9], and nums2 = [1, 4, 6, 7].

// Example 3:
// Input: nums1 = [0,1,0,0,1], nums2 = [0,0,0,1]
// Output: 13
// Explanation:
// One way to replace, having 13 as the largest element is nums1 = [2, 3, 4, 6, 7], and nums2 = [8, 10, 12, 13].

// Constraints:
//     0 <= nums1.length <= 1000
//     1 <= nums2.length <= 1000
//     nums1 and nums2 consist only of 0 and 1.

import "fmt"

func minLargest(nums1 []int, nums2 []int) int {
    nextNumber := func(value, odd int) int {
        value++
        if value % 2 != odd {
            value++
        }
        return value
    }
    n := len(nums1)
    dp, dp2 := make([]int, n + 1), make([]int, n + 1)
    for i := 0; i < len(nums1); i++ {
        dp[i + 1] = nextNumber(dp[i], nums1[i])
    }
    for i := 0; i < len(nums2); i++ {
        dp2[0] = nextNumber(dp[0], nums2[i])
        for j := 0; j < n; j++ {
            dp2[j + 1] = min(nextNumber(dp2[j], nums1[j]), nextNumber(dp[j + 1], nums2[i]))
        }
        for j := 0; j < len(dp); j++ {
            dp[j] = dp2[j]
        }
    }
    return dp[n]
}

// 
// object Solution {
//   def minLargest(nums1: Array[Int], nums2: Array[Int]): Int = {
//     def nextNumber(value: Int, odd: Int): Int = {
//       var newValue = value + 1
//       if ((newValue % 2) != odd) newValue += 1
//       newValue
//     }

//     val dp = Array.fill(nums1.length + 1)(0)
//     val dp2 = Array.fill(nums1.length + 1)(0)

//     nums1.zipWithIndex.foreach { case (v, i) => dp(i + 1) = nextNumber(dp(i), v) }

//     nums2.zipWithIndex.foreach { case (v2, _) => dp2(0) = nextNumber(dp.head, v2)
//       nums1.zipWithIndex.foreach { case (v, i) => dp2(i + 1) = nextNumber(dp2(i), v).min(nextNumber(dp(i + 1), v2)) }
//       dp.indices.foreach(i => dp(i) = dp2(i))
//     }

//     dp.last
//   }
// }

func main() {
    // Example 1:
    // Input: nums1 = [], nums2 = [1,0,1,1]
    // Output: 5
    // Explanation:
    // After replacing, nums1 = [], and nums2 = [1, 2, 3, 5].
    fmt.Println(minLargest([]int{}, []int{1,0,1,1})) // 5
    // Example 2:
    // Input: nums1 = [0,1,0,1], nums2 = [1,0,0,1]
    // Output: 9
    // Explanation:
    // One way to replace, having 9 as the largest element is nums1 = [2, 3, 8, 9], and nums2 = [1, 4, 6, 7].
    fmt.Println(minLargest([]int{0,1,0,1}, []int{1,0,0,1})) // 9
    // Example 3:
    // Input: nums1 = [0,1,0,0,1], nums2 = [0,0,0,1]
    // Output: 13
    // Explanation:
    // One way to replace, having 13 as the largest element is nums1 = [2, 3, 4, 6, 7], and nums2 = [8, 10, 12, 13].
    fmt.Println(minLargest([]int{0,1,0,0,1}, []int{0,0,0,1})) // 13
}