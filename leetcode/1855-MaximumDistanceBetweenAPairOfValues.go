package main

// 1855. Maximum Distance Between a Pair of Values
// You are given two non-increasing 0-indexed integer arrays nums1​​​​​​ and nums2​​​​​​.

// A pair of indices (i, j), where 0 <= i < nums1.length and 0 <= j < nums2.length, 
// is valid if both i <= j and nums1[i] <= nums2[j]. 
// The distance of the pair is j - i​​​​.

// Return the maximum distance of any valid pair (i, j). 
// If there are no valid pairs, return 0.

// An array arr is non-increasing if arr[i-1] >= arr[i] for every 1 <= i < arr.length.

// Example 1:
// Input: nums1 = [55,30,5,4,2], nums2 = [100,20,10,10,5]
// Output: 2
// Explanation: The valid pairs are (0,0), (2,2), (2,3), (2,4), (3,3), (3,4), and (4,4).
// The maximum distance is 2 with pair (2,4).

// Example 2:
// Input: nums1 = [2,2,2], nums2 = [10,10,1]
// Output: 1
// Explanation: The valid pairs are (0,0), (0,1), and (1,1).
// The maximum distance is 1 with pair (0,1).

// Example 3:
// Input: nums1 = [30,29,19,5], nums2 = [25,25,25,25,25]
// Output: 2
// Explanation: The valid pairs are (2,2), (2,3), (2,4), (3,3), and (3,4).
// The maximum distance is 2 with pair (2,4).

// Constraints:
//     1 <= nums1.length, nums2.length <= 10^5
//     1 <= nums1[i], nums2[j] <= 10^5
//     Both nums1 and nums2 are non-increasing.

import "fmt"

func maxDistance(nums1 []int, nums2 []int) int {
    res := 0
    for i := 0 ; i < len(nums1); i++ {
        low, high := i, len(nums2) - 1
        for low <= high {
            mid := (low + high) / 2
            if nums2[mid] >= nums1[i] {
                if mid - i > res { res = mid - i }
                low = mid + 1
            } else { 
                high = mid - 1
            }
        }
    }
    return res 
}

func maxDistance1(nums1 []int, nums2 []int) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, j := 0, 0; j < len(nums2); j++ {
        for i < len(nums1) && nums1[i] > nums2[j] { i++ }
        if i == len(nums1) { break }
        res = max(res, j - i)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [55,30,5,4,2], nums2 = [100,20,10,10,5]
    // Output: 2
    // Explanation: The valid pairs are (0,0), (2,2), (2,3), (2,4), (3,3), (3,4), and (4,4).
    // The maximum distance is 2 with pair (2,4).
    fmt.Println(maxDistance([]int{55,30,5,4,2}, []int{100,20,10,10,5})) // 2
    // Example 2:
    // Input: nums1 = [2,2,2], nums2 = [10,10,1]
    // Output: 1
    // Explanation: The valid pairs are (0,0), (0,1), and (1,1).
    // The maximum distance is 1 with pair (0,1).
    fmt.Println(maxDistance([]int{2,2,2}, []int{10,10,1})) // 1
    // Example 3:
    // Input: nums1 = [30,29,19,5], nums2 = [25,25,25,25,25]
    // Output: 2
    // Explanation: The valid pairs are (2,2), (2,3), (2,4), (3,3), and (3,4).
    // The maximum distance is 2 with pair (2,4).
    fmt.Println(maxDistance([]int{30,29,19,5}, []int{25,25,25,25,25})) // 2

    fmt.Println(maxDistance1([]int{55,30,5,4,2}, []int{100,20,10,10,5})) // 2
    fmt.Println(maxDistance1([]int{2,2,2}, []int{10,10,1})) // 1
    fmt.Println(maxDistance1([]int{30,29,19,5}, []int{25,25,25,25,25})) // 2
}