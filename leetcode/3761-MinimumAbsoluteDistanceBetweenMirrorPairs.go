package main

// 3761. Minimum Absolute Distance Between Mirror Pairs
// You are given an integer array nums.

// A mirror pair is a pair of indices (i, j) such that:
//     1. 0 <= i < j < nums.length, and
//     2. reverse(nums[i]) == nums[j], where reverse(x) denotes the integer formed by reversing the digits of x. 
//        Leading zeros are omitted after reversing, for example reverse(120) = 21.

// Return the minimum absolute distance between the indices of any mirror pair. 
// The absolute distance between indices i and j is abs(i - j).

// If no mirror pair exists, return -1.

// Example 1:
// Input: nums = [12,21,45,33,54]
// Output: 1
// Explanation:
// The mirror pairs are:
// (0, 1) since reverse(nums[0]) = reverse(12) = 21 = nums[1], giving an absolute distance abs(0 - 1) = 1.
// (2, 4) since reverse(nums[2]) = reverse(45) = 54 = nums[4], giving an absolute distance abs(2 - 4) = 2.
// The minimum absolute distance among all pairs is 1.

// Example 2:
// Input: nums = [120,21]
// Output: 1
// Explanation:
// There is only one mirror pair (0, 1) since reverse(nums[0]) = reverse(120) = 21 = nums[1].
// The minimum absolute distance is 1.

// Example 3:
// Input: nums = [21,120]
// Output: -1
// Explanation:
// There are no mirror pairs in the array.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9​​​​​​​

import "fmt"

func minMirrorPairDistance(nums []int) int {
    res, mp := 1 << 31, make(map[int]int)
    reverse := func(n int) int {
        res := 0
        for ; 0 < n; n /= 10 {
            res = 10 * res + n % 10
        }
        return res
    }
    for i, v := range nums {
        if index, ok := mp[v]; ok {
            res = min(res, i - index)
        }
        mp[(reverse(v))] = i
    }
    if res == 1 << 31{
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [12,21,45,33,54]
    // Output: 1
    // Explanation:
    // The mirror pairs are:
    // (0, 1) since reverse(nums[0]) = reverse(12) = 21 = nums[1], giving an absolute distance abs(0 - 1) = 1.
    // (2, 4) since reverse(nums[2]) = reverse(45) = 54 = nums[4], giving an absolute distance abs(2 - 4) = 2.
    // The minimum absolute distance among all pairs is 1.
    fmt.Println(minMirrorPairDistance([]int{12,21,45,33,54})) // 1
    // Example 2:
    // Input: nums = [120,21]
    // Output: 1
    // Explanation:
    // There is only one mirror pair (0, 1) since reverse(nums[0]) = reverse(120) = 21 = nums[1].
    // The minimum absolute distance is 1.
    fmt.Println(minMirrorPairDistance([]int{120,21})) // 1
    // Example 3:
    // Input: nums = [21,120]
    // Output: -1
    // Explanation:
    // There are no mirror pairs in the array.
    fmt.Println(minMirrorPairDistance([]int{21,120})) // -1

    fmt.Println(minMirrorPairDistance([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(minMirrorPairDistance([]int{9,8,7,6,5,4,3,2,1})) // -1
}