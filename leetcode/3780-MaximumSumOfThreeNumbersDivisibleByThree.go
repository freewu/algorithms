package main

// 3780. Maximum Sum of Three Numbers Divisible by Three
// You are given an integer array nums.

// Your task is to choose exactly three integers from nums such that their sum is divisible by three.

// Return the maximum possible sum of such a triplet. If no such triplet exists, return 0.

// Example 1:
// Input: nums = [4,2,3,1]
// Output: 9
// Explanation:
// The valid triplets whose sum is divisible by 3 are:
// (4, 2, 3) with a sum of 4 + 2 + 3 = 9.
// (2, 3, 1) with a sum of 2 + 3 + 1 = 6.
// Thus, the answer is 9.

// Example 2:
// Input: nums = [2,1,5]
// Output: 0
// Explanation:
// No triplet forms a sum divisible by 3, so the answer is 0.

// Constraints:
//     3 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func maximumSum(nums []int) int {
    var r [3][3]int
    for _, n := range nums {
        if n >= r[n%3][0]{
            r[n%3][0] = n
            if r[n%3][0] >= r[n%3][1] {
                r[n%3][0], r[n%3][1] = r[n%3][1], r[n%3][0]
                if r[n%3][1] >= r[n%3][2] {
                    r[n%3][1], r[n%3][2] = r[n%3][2], r[n%3][1]
                }
            }
        }
    }
    res := 0
    if r[0][2] > 0 && r[1][2] > 0 && r[2][2] > 0 {
        res = max(res, r[0][2]+r[1][2]+r[2][2])
    }
    if r[0][0] > 0 && r[0][1] > 0 && r[0][2] > 0 {
        res = max(res, r[0][0]+r[0][1]+r[0][2])
    }
    if r[1][0] > 0 && r[1][1] > 0 && r[1][2] > 0 {
        res = max(res, r[1][0]+r[1][1]+r[1][2])
    }
    if r[2][0] > 0 && r[2][1] > 0 && r[2][2] > 0 {
        res = max(res, r[2][0]+r[2][1]+r[2][2])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,2,3,1]
    // Output: 9
    // Explanation:
    // The valid triplets whose sum is divisible by 3 are:
    // (4, 2, 3) with a sum of 4 + 2 + 3 = 9.
    // (2, 3, 1) with a sum of 2 + 3 + 1 = 6.
    // Thus, the answer is 9.
    fmt.Println(maximumSum([]int{4,2,3,1})) // 9
    // Example 2:
    // Input: nums = [2,1,5]
    // Output: 0
    // Explanation:
    // No triplet forms a sum divisible by 3, so the answer is 0.
    fmt.Println(maximumSum([]int{2,1,5})) // 0

    fmt.Println(maximumSum([]int{1,2,3,4,5,6,7,8,9})) // 24
    fmt.Println(maximumSum([]int{9,8,7,6,5,4,3,2,1})) // 24
}