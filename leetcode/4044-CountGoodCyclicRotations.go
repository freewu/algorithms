package main

// 4044. Count Good Cyclic Rotations
// You are given an integer array nums of even length n.

// A cyclic rotation of nums is obtained by choosing a prefix of nums whose length is between 0 and n - 1 (inclusive), 
// and moving it to the end of the array while preserving the order of all elements.

// A cyclic rotation is good if the sum of its first n / 2 elements is strictly greater than the sum of its last n / 2 elements.

// Return the number of cyclic rotations of nums that are good.

// Example 1:
// Input: nums = [1,2,3,4,5,6]
// Output: 3
// Explanation:
// The cyclic rotations of nums are:
// Cyclic rotation     | Sum of first n / 2 elements   | Sum of last n / 2 elements
// [1, 2, 3, 4, 5, 6]	| 1 + 2 + 3 = 6                 | 4 + 5 + 6 = 15
// [2, 3, 4, 5, 6, 1]	| 2 + 3 + 4 = 9                 | 5 + 6 + 1 = 12
// [3, 4, 5, 6, 1, 2]	| 3 + 4 + 5 = 12                | 6 + 1 + 2 = 9
// [4, 5, 6, 1, 2, 3]	| 4 + 5 + 6 = 15                | 1 + 2 + 3 = 6
// [5, 6, 1, 2, 3, 4]	| 5 + 6 + 1 = 12                | 2 + 3 + 4 = 9
// [6, 1, 2, 3, 4, 5]	| 6 + 1 + 2 = 9                 | 3 + 4 + 5 = 12
// The first half has a greater sum than the second half for 3 rotations. Thus, the answer is 3.

// Example 2:
// Input: nums = [1,2,1,2]
// Output: 0
// Explanation:
// The cyclic rotations of nums are:
// Cyclic rotation | Sum of first n / 2 elements   | Sum of last n / 2 elements
// [1, 2, 1, 2]    | 1 + 2 = 3                     | 1 + 2 = 3
// [2, 1, 2, 1]    | 2 + 1 = 3                     | 2 + 1 = 3
// [1, 2, 1, 2]	| 1 + 2 = 3                     | 1 + 2 = 3
// [2, 1, 2, 1]	| 2 + 1 = 3                     | 2 + 1 = 3
// No cyclic rotation is good because the two sums are equal for every rotation. Thus, the answer is 0.

// Constraints:
//     2 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     n is even.

import "fmt"

func countGoodRotations(nums []int) int {
    res, n := 0, len(nums)
    h := n / 2
    sumA, sumB := 0, 0
    for i := 0; i < h; i++ {
        sumA += nums[i]
    }
    for i := h; i < n; i++ {
        sumB += nums[i]
    }
    i, x := 0, h
    for s := 0; s < n; s++ {
        if sumA > sumB {
            res++
        }
        sumA = sumA - nums[i] + nums[x]
        sumB = sumB - nums[x] + nums[i]
        i = (i + 1) % n
        x = (x + 1) % n
    }
    return res
}

func countGoodRotations1(nums []int) int {
    res, sum, fir, n := 0, int64(0), int64(0), len(nums)
    h := n / 2
    for _, v := range nums {
        sum += int64(v)
    }
    for i := 0; i < h; i++ {
        fir += int64(nums[i])
    }
    for i := 0; i < n; i++ {
        sec := sum - fir
        if fir > sec {
            res++
        }
        fir += int64(nums[(i+h)%n]) - int64(nums[i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5,6]
    // Output: 3
    // Explanation:
    // The cyclic rotations of nums are:
    // Cyclic rotation     | Sum of first n / 2 elements   | Sum of last n / 2 elements
    // [1, 2, 3, 4, 5, 6]	| 1 + 2 + 3 = 6                 | 4 + 5 + 6 = 15
    // [2, 3, 4, 5, 6, 1]	| 2 + 3 + 4 = 9                 | 5 + 6 + 1 = 12
    // [3, 4, 5, 6, 1, 2]	| 3 + 4 + 5 = 12                | 6 + 1 + 2 = 9
    // [4, 5, 6, 1, 2, 3]	| 4 + 5 + 6 = 15                | 1 + 2 + 3 = 6
    // [5, 6, 1, 2, 3, 4]	| 5 + 6 + 1 = 12                | 2 + 3 + 4 = 9
    // [6, 1, 2, 3, 4, 5]	| 6 + 1 + 2 = 9                 | 3 + 4 + 5 = 12
    // The first half has a greater sum than the second half for 3 rotations. Thus, the answer is 3.
    fmt.Println(countGoodRotations([]int{1,2,3,4,5,6})) // 3
    // Example 2:
    // Input: nums = [1,2,1,2]
    // Output: 0
    // Explanation:
    // The cyclic rotations of nums are:
    // Cyclic rotation | Sum of first n / 2 elements   | Sum of last n / 2 elements
    // [1, 2, 1, 2]    | 1 + 2 = 3                     | 1 + 2 = 3
    // [2, 1, 2, 1]    | 2 + 1 = 3                     | 2 + 1 = 3
    // [1, 2, 1, 2]	| 1 + 2 = 3                     | 1 + 2 = 3
    // [2, 1, 2, 1]	| 2 + 1 = 3                     | 2 + 1 = 3
    // No cyclic rotation is good because the two sums are equal for every rotation. Thus, the answer is 0.
    fmt.Println(countGoodRotations([]int{1,2,1,2})) // 0

    fmt.Println(countGoodRotations([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(countGoodRotations([]int{9,8,7,6,5,4,3,2,1})) // 3

    fmt.Println(countGoodRotations1([]int{1,2,3,4,5,6})) // 3
    fmt.Println(countGoodRotations1([]int{1,2,1,2})) // 0
    fmt.Println(countGoodRotations1([]int{1,2,3,4,5,6,7,8,9})) // 3
    fmt.Println(countGoodRotations1([]int{9,8,7,6,5,4,3,2,1})) // 3
}