package main

// 1537. Get the Maximum Score
// You are given two sorted arrays of distinct integers nums1 and nums2.

// A valid path is defined as follows:
//     1. Choose array nums1 or nums2 to traverse (from index-0).
//     2. Traverse the current array from left to right.
//     3. If you are reading any value that is present in nums1 and nums2 you are allowed to change your path to the other array. 
//        (Only one repeated value is considered in the valid path).

// The score is defined as the sum of unique values in a valid path.

// Return the maximum score you can obtain of all possible valid paths. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/16/sample_1_1893.png" />
// Input: nums1 = [2,4,5,8,10], nums2 = [4,6,8,9]
// Output: 30
// Explanation: Valid paths:
// [2,4,5,8,10], [2,4,5,8,9], [2,4,6,8,9], [2,4,6,8,10],  (starting from nums1)
// [4,6,8,9], [4,5,8,10], [4,5,8,9], [4,6,8,10]    (starting from nums2)
// The maximum is obtained with the path in green [2,4,6,8,10].

// Example 2:
// Input: nums1 = [1,3,5,7,9], nums2 = [3,5,100]
// Output: 109
// Explanation: Maximum sum is obtained with the path [1,3,5,100].

// Example 3:
// Input: nums1 = [1,2,3,4,5], nums2 = [6,7,8,9,10]
// Output: 40
// Explanation: There are no common elements between nums1 and nums2.
// Maximum sum is obtained with the path [6,7,8,9,10].

// Constraints:
//     1 <= nums1.length, nums2.length <= 10^5
//     1 <= nums1[i], nums2[i] <= 10^7
//     nums1 and nums2 are strictly increasing.

import "fmt"

// 双指针
func maxSum(nums1 []int, nums2 []int) int {
    i, j, sum1, sum2 := 0, 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            sum1 += nums1[i]
            i++
        } else if nums1[i] > nums2[j] {
            sum2 += nums2[j]
            j++
        }  else { // nums1[i] == nums2[j] 交点
            mx := max(sum1 + nums1[i], sum2 + nums2[j])
            sum1, sum2 = mx, mx
            i++
            j++
        }
    }
    for ; i < len(nums1); i++ { // 没走完则, 累加剩余的
        sum1 += nums1[i] 
    }
    for ; j < len(nums2); j++ { // 没走完则, 累加剩余的
        sum2 += nums2[j] 
    }
    return max(sum1, sum2) % 1_000_000_007
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/16/sample_1_1893.png" />
    // Input: nums1 = [2,4,5,8,10], nums2 = [4,6,8,9]
    // Output: 30
    // Explanation: Valid paths:
    // [2,4,5,8,10], [2,4,5,8,9], [2,4,6,8,9], [2,4,6,8,10],  (starting from nums1)
    // [4,6,8,9], [4,5,8,10], [4,5,8,9], [4,6,8,10]    (starting from nums2)
    // The maximum is obtained with the path in green [2,4,6,8,10].
    fmt.Println(maxSum([]int{2,4,5,8,10}, []int{4,6,8,9})) // 30
    // Example 2:
    // Input: nums1 = [1,3,5,7,9], nums2 = [3,5,100]
    // Output: 109
    // Explanation: Maximum sum is obtained with the path [1,3,5,100].
    fmt.Println(maxSum([]int{1,3,5,7,9}, []int{3,5,100})) // 109
    // Example 3:
    // Input: nums1 = [1,2,3,4,5], nums2 = [6,7,8,9,10]
    // Output: 40
    // Explanation: There are no common elements between nums1 and nums2.
    // Maximum sum is obtained with the path [6,7,8,9,10].
    fmt.Println(maxSum([]int{1,2,3,4,5}, []int{6,7,8,9,10})) // 40
}