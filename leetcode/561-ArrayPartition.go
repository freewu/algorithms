package main

// 561. Array Partition
// Given an integer array nums of 2n integers, 
// group these integers into n pairs (a1, b1), (a2, b2), ..., (an, bn) such that the sum of min(ai, bi) 
// for all i is maximized. Return the maximized sum.

// Example 1:
// Input: nums = [1,4,3,2]
// Output: 4
// Explanation: All possible pairings (ignoring the ordering of elements) are:
// 1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
// 2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
// 3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
// So the maximum possible sum is 4.

// Example 2:
// Input: nums = [6,2,6,5,1,2]
// Output: 9
// Explanation: The optimal pairing is (2, 1), (2, 5), (6, 6). min(2, 1) + min(2, 5) + min(6, 6) = 1 + 2 + 6 = 9.
 
// Constraints:
//     1 <= n <= 10^4
//     nums.length == 2 * n
//     -10^4 <= nums[i] <= 10^4

import "fmt"
import "sort"

func arrayPairSum(nums []int) int {
    res := 0
    sort.Ints(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(nums) - 1; i += 2 {
        res += min(nums[i],nums[i + 1])
    }
    return res
}

func arrayPairSum1(nums []int) int {
    sort.Ints(nums)
    res := 0
    for i := 0; i < len(nums); i += 2 {
        res += nums[i] // 因为排序过，所的 i 一定是 <= i + 1的
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,4,3,2]
    // Output: 4
    // Explanation: All possible pairings (ignoring the ordering of elements) are:
    // 1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
    // 2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
    // 3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
    // So the maximum possible sum is 4.
    fmt.Println(arrayPairSum([]int{1,4,3,2})) // 4
    // Example 2:
    // Input: nums = [6,2,6,5,1,2]
    // Output: 9
    // Explanation: The optimal pairing is (2, 1), (2, 5), (6, 6). min(2, 1) + min(2, 5) + min(6, 6) = 1 + 2 + 6 = 9.
    fmt.Println(arrayPairSum([]int{6,2,6,5,1,2})) // 9

    fmt.Println(arrayPairSum1([]int{1,4,3,2})) // 4
    fmt.Println(arrayPairSum1([]int{6,2,6,5,1,2})) // 9
}