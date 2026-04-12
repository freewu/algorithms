package main

// 3897. Maximum Value of Concatenated Binary Segments
// You are given two integer arrays nums1 and nums0, each of size n.
//     1. nums1[i] represents the number of '1's in the ith segment.
//     2. nums0[i] represents the number of '0's in the ith segment.

// For each index i, construct a binary segment consisting of:
//     1. nums1[i] occurrences of '1' followed by
//     2. nums0[i] occurrences of '0'.

// You may rearrange the order of these segments in any way. 
// After rearranging, concatenate all segments to form a single binary string.

// Return the maximum possible integer value of the concatenated binary string.

// Since the result can be very large, return the answer modulo 10^9 + 7.

// Example 1:
// Input: nums1 = [1,2], nums0 = [1,0]
// Output: 14
// Explanation:
// At index 0, nums1[0] = 1 and nums0[0] = 1, so the segment formed is "10".
// At index 1, nums1[1] = 2 and nums0[1] = 0, so the segment formed is "11".
// Reordering the segments as "11" followed by "10" produces the binary string "1110".
// The binary number "1110" has value 14 which is the maximum possible value.

// Example 2:
// Input: nums1 = [3,1], nums0 = [0,3]
// Output: 120
// Explanation:
// At index 0, nums1[0] = 3 and nums0[0] = 0, so the segment formed is "111".
// At index 1, nums1[1] = 1 and nums0[1] = 3, so the segment formed is "1000".
// Reordering the segments as "111" followed by "1000" produces the binary string "1111000".
// The binary number "1111000" has value 120 which is the maximum possible value.
 
// Constraints:
//     1 <= n == nums1.length == nums0.length <= 10^5
//     0 <= nums1[i], nums0[i] <= 10^4
//     nums1[i] + nums0[i] > 0
//     The total sum of all elements in nums1 and nums0 does not exceed 2 * 10^5.

import "fmt"
import "slices"
import "cmp"

const MOD = 1_000_000_007
const MX = 10001

var pow2 = [MX]int{1}

func init() {
    // 预处理 2 的幂
    for i := 1; i < MX; i++ {
        pow2[i] = pow2[i-1] * 2 % MOD   
    }
}

func maxValue(nums1, nums0 []int) int {
    res, index := 0, make([]int, len(nums1))
    for i := range index {
        index[i] = i
    }
    slices.SortFunc(index, func(i, j int) int {
        if nums0[i] == 0 {
            return -1
        }
        if nums0[j] == 0 {
            return 1
        }
        return cmp.Or(nums1[j]-nums1[i], nums0[i]-nums0[j])
    })
    for _, i := range index {
        res = ((res + 1) * pow2[nums1[i]] - 1) % MOD * pow2[nums0[i]] % MOD
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [1,2], nums0 = [1,0]
    // Output: 14
    // Explanation:
    // At index 0, nums1[0] = 1 and nums0[0] = 1, so the segment formed is "10".
    // At index 1, nums1[1] = 2 and nums0[1] = 0, so the segment formed is "11".
    // Reordering the segments as "11" followed by "10" produces the binary string "1110".
    // The binary number "1110" has value 14 which is the maximum possible value.
    fmt.Println(maxValue([]int{1,2}, []int{1,0})) // 14
    // Example 2:
    // Input: nums1 = [3,1], nums0 = [0,3]
    // Output: 120
    // Explanation:
    // At index 0, nums1[0] = 3 and nums0[0] = 0, so the segment formed is "111".
    // At index 1, nums1[1] = 1 and nums0[1] = 3, so the segment formed is "1000".
    // Reordering the segments as "111" followed by "1000" produces the binary string "1111000".
    // The binary number "1111000" has value 120 which is the maximum possible value.
    fmt.Println(maxValue([]int{3,1}, []int{0,3})) // 120

    fmt.Println(maxValue([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 17184768
    fmt.Println(maxValue([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 375250857
    fmt.Println(maxValue([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 375250857
    fmt.Println(maxValue([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 17184768
}