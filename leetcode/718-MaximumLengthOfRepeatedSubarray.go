package main

// 718. Maximum Length of Repeated Subarray
// Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.

// Example 1:
// Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
// Output: 3
// Explanation: The repeated subarray with maximum length is [3,2,1].

// Example 2:
// Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
// Output: 5
// Explanation: The repeated subarray with maximum length is [0,0,0,0,0].
 
// Constraints:
//     1 <= nums1.length, nums2.length <= 1000
//     0 <= nums1[i], nums2[i] <= 100

import "fmt"

// dp
func findLength(nums1 []int, nums2 []int) int {
    arr := make([][]int, len(nums1)+1)
    for i := 0; i < len(arr); i++ {
        arr[i] = make([]int, len(nums2)+1)
    }
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := len(nums1) - 1; i > -1; i-- {
        for j := len(nums2) - 1; j > -1; j-- {
            if nums1[i] == nums2[j] {
                if arr[i+1][j+1] != 0 {
                    arr[i][j] = arr[i+1][j+1] + 1
                } else {
                    arr[i][j] = 1
                }
                res = max(res, arr[i][j])
            }
        }
    }
    return res
}

func findLength1(nums1 []int, nums2 []int) int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return 0
    }
    dp := make([]int, len(nums2)+1)
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= len(nums1); i++ {
        for j := len(nums2); j > 0; j-- {
            if nums1[i-1] == nums2[j-1] {
                dp[j] = dp[j-1] + 1
            } else {
                dp[j] = 0
            }
            res = max(res, dp[j])
        }
    }
    return res
}

func main() {
    // Explanation: The repeated subarray with maximum length is [3,2,1].
    fmt.Println(findLength([]int{1,2,3,2,1},[]int{3,2,1,4,7})) // 3
    // Explanation: The repeated subarray with maximum length is [0,0,0,0,0].
    fmt.Println(findLength([]int{0,0,0,0,0},[]int{0,0,0,0,0})) // 5

    fmt.Println(findLength1([]int{1,2,3,2,1},[]int{3,2,1,4,7})) // 3
    fmt.Println(findLength1([]int{0,0,0,0,0},[]int{0,0,0,0,0})) // 5
}