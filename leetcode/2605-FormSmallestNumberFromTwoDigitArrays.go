package main

// 2605. Form Smallest Number From Two Digit Arrays
// Given two arrays of unique digits nums1 and nums2, return the smallest number that contains at least one digit from each array.

// Example 1:
// Input: nums1 = [4,1,3], nums2 = [5,7]
// Output: 15
// Explanation: The number 15 contains the digit 1 from nums1 and the digit 5 from nums2. It can be proven that 15 is the smallest number we can have.

// Example 2:
// Input: nums1 = [3,5,2,6], nums2 = [3,1,7]
// Output: 3
// Explanation: The number 3 contains the digit 3 which exists in both arrays.

// Constraints:
//     1 <= nums1.length, nums2.length <= 9
//     1 <= nums1[i], nums2[i] <= 9
//     All digits in each array are unique.

import "fmt"

func minNumber(nums1 []int, nums2 []int) int {
    mn1, mn2 := 9, 9
    mp := make(map[int]bool)
    for _, v := range nums1 {
        if v < mn1 {
            mn1 = v
        }
        mp[v] = true
    }
    dupeCheck := 10
    for _, v := range nums2 {
        if v < mn2 {
            mn2 = v
        }
        if mp[v] {
            if v < dupeCheck {
                dupeCheck = v
            }
        }
    }
    if dupeCheck != 10 {
        return dupeCheck
    }
    if mn1 < mn2 {
        return mn1 * 10 + mn2
    }
    return mn2 * 10 + mn1
}

func main() {
    // Example 1:
    // Input: nums1 = [4,1,3], nums2 = [5,7]
    // Output: 15
    // Explanation: The number 15 contains the digit 1 from nums1 and the digit 5 from nums2. It can be proven that 15 is the smallest number we can have.
    fmt.Println(minNumber([]int{4,1,3}, []int{5,7})) // 15
    // Example 2:
    // Input: nums1 = [3,5,2,6], nums2 = [3,1,7]
    // Output: 3
    // Explanation: The number 3 contains the digit 3 which exists in both arrays.
    fmt.Println(minNumber([]int{3,5,2,6}, []int{3,1,7})) // 3
}