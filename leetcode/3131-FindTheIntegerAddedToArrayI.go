package main

// 3131. Find the Integer Added to Array I
// You are given two arrays of equal length, nums1 and nums2.

// Each element in nums1 has been increased (or decreased in the case of negative) by an integer, represented by the variable x.

// As a result, nums1 becomes equal to nums2. 
// Two arrays are considered equal when they contain the same integers with the same frequencies.

// Return the integer x.

// Example 1:
// Input: nums1 = [2,6,4], nums2 = [9,7,5]
// Output: 3
// Explanation:
// The integer added to each element of nums1 is 3.

// Example 2:
// Input: nums1 = [10], nums2 = [5]
// Output: -5
// Explanation:
// The integer added to each element of nums1 is -5.

// Example 3:
// Input: nums1 = [1,1,1,1], nums2 = [1,1,1,1]
// Output: 0
// Explanation:
// The integer added to each element of nums1 is 0.

// Constraints:
//     1 <= nums1.length == nums2.length <= 100
//     0 <= nums1[i], nums2[i] <= 1000
//     The test cases are generated in a way that there is an integer x such that nums1 can become equal to nums2 by adding x to each element of nums1.

import "fmt"

func addedInteger(nums1 []int, nums2 []int) int {
    sum1, sum2, n := 0, 0, len(nums1)
    for _, v := range nums1 { sum1 += v }
    for _, v := range nums2 { sum2 += v } 
    return (sum2 - sum1) / n
}

func main() {
    // Example 1:
    // Input: nums1 = [2,6,4], nums2 = [9,7,5]
    // Output: 3
    // Explanation:
    // The integer added to each element of nums1 is 3.
    fmt.Println(addedInteger([]int{2,6,4},[]int{9,7,5})) // 3
    // Example 2:
    // Input: nums1 = [10], nums2 = [5]
    // Output: -5
    // Explanation:
    // The integer added to each element of nums1 is -5.
    fmt.Println(addedInteger([]int{10},[]int{5})) // -5
    // Example 3:
    // Input: nums1 = [1,1,1,1], nums2 = [1,1,1,1]
    // Output: 0
    // Explanation:
    // The integer added to each element of nums1 is 0.
    fmt.Println(addedInteger([]int{1,1,1,1},[]int{1,1,1,1})) // 0
}