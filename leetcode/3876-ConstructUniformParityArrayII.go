package main

// 3876. Construct Uniform Parity Array II
// You are given an array nums1 of n distinct integers.

// You want to construct another array nums2 of length n such that the elements in nums2 are either all odd or all even.

// For each index i, you must choose exactly one of the following (in any order):
//     1. nums2[i] = nums1[i]​​​​​​​
//     2. nums2[i] = nums1[i] - nums1[j], for an index j != i, such that nums1[i] - nums1[j] >= 1

// Return true if it is possible to construct such an array, otherwise return false.

// Example 1:
// Input: nums1 = [1,4,7]
// Output: true
// Explanation:​​​​​​​​​​​​​​
// Set nums2[0] = nums1[0] = 1.
// Set nums2[1] = nums1[1] - nums1[0] = 4 - 1 = 3.
// Set nums2[2] = nums1[2] = 7.
// nums2 = [1, 3, 7], and all elements are odd. Thus, the answer is true.

// Example 2:
// Input: nums1 = [2,3]
// Output: false
// Explanation:
// It is not possible to construct nums2 such that all elements have the same parity. Thus, the answer is false.

// Example 3:
// Input: nums1 = [4,6]
// Output: true
// Explanation:
// Set nums2[0] = nums1[0] = 4.
// Set nums2[1] = nums1[1] = 6.
// nums2 = [4, 6], and all elements are even. Thus, the answer is true.
 
// Constraints:
//     1 <= n == nums1.length <= 10^5
//     1 <= nums1[i] <= 10^9
//     nums1 consists of distinct integers.

import "fmt"

func uniformArray(nums1 []int) bool {
    const inf = 1 << 61
    odd, even := inf, inf
    for _, v := range nums1 {
        if v % 2 != 0 {
            odd = min(odd, v)
        } else {
            even = min(even, v)
        }
    }
    return odd == inf || // all even
        even == inf || // all odd
        even-odd >= 1 // can construct all odd
    // all even cannot be constructed
}

func uniformArray1(nums1 []int) bool {
    count, mn := 0, 1 << 61
    for i := 0; i < len(nums1); i++ {
        mn = min(mn, nums1[i])
        if nums1[i] % 2 == 1 {
            count++
        }
    }
    if mn % 2 == 1 {
        return true
    }
    if count == 0 {
        return true
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums1 = [1,4,7]
    // Output: true
    // Explanation:​​​​​​​​​​​​​​
    // Set nums2[0] = nums1[0] = 1.
    // Set nums2[1] = nums1[1] - nums1[0] = 4 - 1 = 3.
    // Set nums2[2] = nums1[2] = 7.
    // nums2 = [1, 3, 7], and all elements are odd. Thus, the answer is true.
    fmt.Println(uniformArray([]int{1,4,7})) // true 
    // Example 2:
    // Input: nums1 = [2,3]
    // Output: false
    // Explanation:
    // It is not possible to construct nums2 such that all elements have the same parity. Thus, the answer is false.
    fmt.Println(uniformArray([]int{2,3})) // false
    // Example 3:
    // Input: nums1 = [4,6]
    // Output: true
    // Explanation:
    // Set nums2[0] = nums1[0] = 4.
    // Set nums2[1] = nums1[1] = 6.
    // nums2 = [4, 6], and all elements are even. Thus, the answer is true.
    fmt.Println(uniformArray([]int{4,6})) // true

    fmt.Println(uniformArray([]int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(uniformArray([]int{9,8,7,6,5,4,3,2,1})) // true

    fmt.Println(uniformArray1([]int{1,4,7})) // true 
    fmt.Println(uniformArray1([]int{2,3})) // false
    fmt.Println(uniformArray1([]int{4,6})) // true
    fmt.Println(uniformArray1([]int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(uniformArray1([]int{9,8,7,6,5,4,3,2,1})) // true
}
