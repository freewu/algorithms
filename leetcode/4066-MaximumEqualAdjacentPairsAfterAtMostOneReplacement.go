package main

// 4066. Maximum Equal Adjacent Pairs After at Most One Replacement
// You are given a 1-indexed integer array nums.

// You can choose two distinct values x and y and perform the following operation at most once:
//     1. Replace every occurrence of x in nums with y.

// Return the maximum possible number of pairs of adjacent elements that are equal after performing the operation.

// Example 1:
// Input: nums = [1,2,3,2]
// Output: 2
// Explanation:
// One optimal solution is to choose x = 3 and y = 2.
// The resulting array is [1, 2, 2, 2].
// There are 2 pairs of adjacent elements that are equal: (nums[2], nums[3]) and (nums[3], nums[4]).
// Therefore, the answer is 2.

// Example 2:
// Input: nums = [1,2,1,2,1]
// Output: 4
// Explanation:
// One optimal solution is to choose x = 1 and y = 2.
// The resulting array is [2, 2, 2, 2, 2].
// There are 4 pairs of adjacent elements that are equal: (nums[1], nums[2]), (nums[2], nums[3]), (nums[3], nums[4]), and (nums[4], nums[5]).
// Therefore, the answer is 4.

// Example 3:
// Input: nums = [1,1,1]
// Output: 2
// Explanation:
// One optimal solution is to perform no operation.
// Thus, the resulting array is [1, 1, 1].
// There are 2 pairs of adjacent elements that are equal: (nums[1], nums[2]) and (nums[2], nums[3]).
// Therefore, the answer is 2.

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func maxEqualAdjacentPairs(nums []int) int {
    base, mx := 0, 0
    type Pair struct{ x, y int }
    count := map[Pair]int{}
    for i := 1; i < len(nums); i++ {
        x, y := nums[i-1], nums[i]
        if x == y {
            base++
        } else {
            // 把 (x,y) 和 (y,x) 都统一为 (x,y)
            if x > y {
                x, y = y, x
            }
            // 统计相邻且不相等的数对个数
            count[Pair{x, y}]++
        }
    }
    for _, c := range count {
        mx = max(mx, c)
    }
    return base + mx
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,2]
    // Output: 2
    // Explanation:
    // One optimal solution is to choose x = 3 and y = 2.
    // The resulting array is [1, 2, 2, 2].
    // There are 2 pairs of adjacent elements that are equal: (nums[2], nums[3]) and (nums[3], nums[4]).
    // Therefore, the answer is 2.
    fmt.Println(maxEqualAdjacentPairs([]int{1,2,3,2})) // 2
    // Example 2:
    // Input: nums = [1,2,1,2,1]
    // Output: 4
    // Explanation:
    // One optimal solution is to choose x = 1 and y = 2.
    // The resulting array is [2, 2, 2, 2, 2].
    // There are 4 pairs of adjacent elements that are equal: (nums[1], nums[2]), (nums[2], nums[3]), (nums[3], nums[4]), and (nums[4], nums[5]).
    // Therefore, the answer is 4.
    fmt.Println(maxEqualAdjacentPairs([]int{1,2,1,2,1})) // 4
    // Example 3:
    // Input: nums = [1,1,1]
    // Output: 2
    // Explanation:
    // One optimal solution is to perform no operation.
    // Thus, the resulting array is [1, 1, 1].
    // There are 2 pairs of adjacent elements that are equal: (nums[1], nums[2]) and (nums[2], nums[3]).
    // Therefore, the answer is 2.
    fmt.Println(maxEqualAdjacentPairs([]int{1,1,1})) // 2

    fmt.Println(maxEqualAdjacentPairs([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(maxEqualAdjacentPairs([]int{9,8,7,6,5,4,3,2,1})) // 1
}