package main

// 1031. Maximum Sum of Two Non-Overlapping Subarrays
// Given an integer array nums and two integers firstLen and secondLen, 
// return the maximum sum of elements in two non-overlapping subarrays with lengths firstLen and secondLen.

// The array with length firstLen could occur before or after the array with length secondLen, 
// but they have to be non-overlapping.

// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [0,6,5,2,2,5,1,9,4], firstLen = 1, secondLen = 2
// Output: 20
// Explanation: One choice of subarrays is [9] with length 1, and [6,5] with length 2.

// Example 2:
// Input: nums = [3,8,1,3,2,1,8,9,0], firstLen = 3, secondLen = 2
// Output: 29
// Explanation: One choice of subarrays is [3,8,1] with length 3, and [8,9] with length 2.

// Example 3:
// Input: nums = [2,1,5,6,0,9,5,0,3,8], firstLen = 4, secondLen = 3
// Output: 31
// Explanation: One choice of subarrays is [5,6,0,9] with length 4, and [0,3,8] with length 3.

// Constraints:
//     1 <= firstLen, secondLen <= 1000
//     2 <= firstLen + secondLen <= 1000
//     firstLen + secondLen <= nums.length <= 1000
//     0 <= nums[i] <= 1000

import "fmt"

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
    n := len(nums)
    prefix := make([]int, n + 1)
    for i := 1; i <= len(nums); i++ { // 生成前缀和数组
        prefix[i] = prefix[i-1] + nums[i-1]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    find := func (prefix []int, n1 int, n2 int) int {
        mxl, mx := 0, 0
        for i := n1 + n2; i < len(prefix); i++ {
            mxl = max(mxl, prefix[i - n2] - prefix[i - n2 - n1])
            mx = max(mx, mxl + prefix[i] - prefix[i - n2])
        }
        return mx
    }
    return max(find(prefix, firstLen, secondLen), find(prefix,secondLen, firstLen))
}

func maxSumTwoNoOverlap1(nums []int, firstLen int, secondLen int) int {
    res, n := 0, len(nums)
    prefix := make([]int, n + 1)
    for i := 1; i <= len(nums); i++ { // 生成前缀和数组
        prefix[i] = prefix[i-1] + nums[i-1]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    find := func(firstLen, secondLen int) {
        sum := 0
        for i := firstLen + secondLen; i <= n; i++ {
            sum = max(sum, prefix[i - secondLen] - prefix[i - secondLen - firstLen])
            res = max(res, sum + prefix[i] - prefix[i - secondLen])
        }
    }
    find(firstLen, secondLen)
    find(secondLen, firstLen)
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,6,5,2,2,5,1,9,4], firstLen = 1, secondLen = 2
    // Output: 20
    // Explanation: One choice of subarrays is [9] with length 1, and [6,5] with length 2.
    fmt.Println(maxSumTwoNoOverlap([]int{0,6,5,2,2,5,1,9,4}, 1, 2)) // 20
    // Example 2:
    // Input: nums = [3,8,1,3,2,1,8,9,0], firstLen = 3, secondLen = 2
    // Output: 29
    // Explanation: One choice of subarrays is [3,8,1] with length 3, and [8,9] with length 2.
    fmt.Println(maxSumTwoNoOverlap([]int{3,8,1,3,2,1,8,9,0}, 3, 2)) // 29
    // Example 3:
    // Input: nums = [2,1,5,6,0,9,5,0,3,8], firstLen = 4, secondLen = 3
    // Output: 31
    // Explanation: One choice of subarrays is [5,6,0,9] with length 4, and [0,3,8] with length 3.
    fmt.Println(maxSumTwoNoOverlap([]int{2,1,5,6,0,9,5,0,3,8}, 4, 3)) // 31

    fmt.Println(maxSumTwoNoOverlap1([]int{0,6,5,2,2,5,1,9,4}, 1, 2)) // 20
    fmt.Println(maxSumTwoNoOverlap1([]int{3,8,1,3,2,1,8,9,0}, 3, 2)) // 29
    fmt.Println(maxSumTwoNoOverlap1([]int{2,1,5,6,0,9,5,0,3,8}, 4, 3)) // 31
}