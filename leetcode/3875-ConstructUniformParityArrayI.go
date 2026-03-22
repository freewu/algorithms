package main

// 3875. Construct Uniform Parity Array I
// You are given an array nums1 of n distinct integers.

// You want to construct another array nums2 of length n such that the elements in nums2 are either all odd or all even.

// For each index i, you must choose exactly one of the following (in any order):
//     1. nums2[i] = nums1[i]
//     2. nums2[i] = nums1[i] - nums1[j], for an index j != i

// Return true if it is possible to construct such an array, otherwise, return false.

// Example 1:
// Input: nums1 = [2,3]
// Output: true
// Explanation:
// Choose nums2[0] = nums1[0] - nums1[1] = 2 - 3 = -1.
// Choose nums2[1] = nums1[1] = 3.
// nums2 = [-1, 3], and both elements are odd. Thus, the answer is true​​​​​​​.

// Example 2:
// Input: nums1 = [4,6]
// Output: true
// Explanation:​​​​​​​
// Choose nums2[0] = nums1[0] = 4.
// Choose nums2[1] = nums1[1] = 6.
// nums2 = [4, 6], and all elements are even. Thus, the answer is true.

// Constraints:
//     1 <= n == nums1.length <= 100
//     1 <= nums1[i] <= 100
//     nums1 consists of distinct integers.

import "fmt"

func uniformArray(nums1 []int) bool {
    // 如果 nums1 全是奇数或者全是偶数，那么只用第一种操作 nums2[i] = nums1[i] 即可满足要求。 否则，nums1 奇数偶数都有。
    // 由于偶数减去奇数等于奇数，那么随便选一个奇数 x，把每个偶数都减去 x（第二种操作），即可让所有偶数都变成奇数。其余每个奇数用第一种操作。
    // 所以，一定可以满足题目要求，返回 true 即可
    return true
}

func main() {
    // Example 1:
    // Input: nums1 = [2,3]
    // Output: true
    // Explanation:
    // Choose nums2[0] = nums1[0] - nums1[1] = 2 - 3 = -1.
    // Choose nums2[1] = nums1[1] = 3.
    // nums2 = [-1, 3], and both elements are odd. Thus, the answer is true​​​​​​​.
    fmt.Println(uniformArray([]int{2,3})) // true
    // Example 2:
    // Input: nums1 = [4,6]
    // Output: true
    // Explanation:​​​​​​​
    // Choose nums2[0] = nums1[0] = 4.
    // Choose nums2[1] = nums1[1] = 6.
    // nums2 = [4, 6], and all elements are even. Thus, the answer is true.
    fmt.Println(uniformArray([]int{4,6})) // true

    fmt.Println(uniformArray([]int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(uniformArray([]int{9,8,7,6,5,4,3,2,1})) // true
}