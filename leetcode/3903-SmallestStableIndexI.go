package main

// 3903. Smallest Stable Index I
// You are given an integer array nums of length n and an integer k.

// For each index i, define its instability score as max(nums[0..i]) - min(nums[i..n - 1]).

// In other words:
//     1. max(nums[0..i]) is the largest value among the elements from index 0 to index i.
//     2. min(nums[i..n - 1]) is the smallest value among the elements from index i to index n - 1.

// An index i is called stable if its instability score is less than or equal to k.

// Return the smallest stable index. If no such index exists, return -1.

// Example 1:
// Input: nums = [5,0,1,4], k = 3
// Output: 3
// Explanation:
// At index 0: The maximum in [5] is 5, and the minimum in [5, 0, 1, 4] is 0, so the instability score is 5 - 0 = 5.
// At index 1: The maximum in [5, 0] is 5, and the minimum in [0, 1, 4] is 0, so the instability score is 5 - 0 = 5.
// At index 2: The maximum in [5, 0, 1] is 5, and the minimum in [1, 4] is 1, so the instability score is 5 - 1 = 4.
// At index 3: The maximum in [5, 0, 1, 4] is 5, and the minimum in [4] is 4, so the instability score is 5 - 4 = 1.
// This is the first index with an instability score less than or equal to k = 3. Thus, the answer is 3.

// Example 2:
// Input: nums = [3,2,1], k = 1
// Output: -1
// Explanation:
// At index 0, the instability score is 3 - 1 = 2.
// At index 1, the instability score is 3 - 1 = 2.
// At index 2, the instability score is 3 - 1 = 2.
// None of these values is less than or equal to k = 1, so the answer is -1.

// Example 3:
// Input: nums = [0], k = 0
// Output: 0
// Explanation:
// At index 0, the instability score is 0 - 0 = 0, which is less than or equal to k = 0. Therefore, the answer is 0.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 10^9
//     0 <= k <= 10^9

import "fmt"

func firstStableIndex(nums []int, k int) int {
    for i := 0; i < len(nums); i++ {
        left, right := nums[:i+1], nums[i:]
        h, l := 0, 0
        for j := 0; j < len(left); j++ {
            if j == 0 {
                h = left[j]
            }
            if left[j] > h {
                h = left[j]
            }
        }
        for j := 0; j < len(right); j++ {
            if j == 0 {
                l = right[j]
            }
            if right[j] < l {
                l = right[j]
            }
        }
        if h - l <= k {
            return i
        }
    }
    return -1
}

func firstStableIndex1(nums []int, k int) int {
    n := len(nums)
    suffix := make([]int, n)
    suffix[n-1] = nums[n-1]
    for i := n - 2; i >= 0; i-- {
        suffix[i] = min(nums[i], suffix[i+1])
    }
    mx := nums[0]
    for i, num := range nums {
        mx = max(mx, num)
        if mx-suffix[i] <= k {
            return i
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [5,0,1,4], k = 3
    // Output: 3
    // Explanation:
    // At index 0: The maximum in [5] is 5, and the minimum in [5, 0, 1, 4] is 0, so the instability score is 5 - 0 = 5.
    // At index 1: The maximum in [5, 0] is 5, and the minimum in [0, 1, 4] is 0, so the instability score is 5 - 0 = 5.
    // At index 2: The maximum in [5, 0, 1] is 5, and the minimum in [1, 4] is 1, so the instability score is 5 - 1 = 4.
    // At index 3: The maximum in [5, 0, 1, 4] is 5, and the minimum in [4] is 4, so the instability score is 5 - 4 = 1.
    // This is the first index with an instability score less than or equal to k = 3. Thus, the answer is 3.
    fmt.Println(firstStableIndex([]int{5,0,1,4}, 3)) // 3
    // Example 2:
    // Input: nums = [3,2,1], k = 1
    // Output: -1
    // Explanation:
    // At index 0, the instability score is 3 - 1 = 2.
    // At index 1, the instability score is 3 - 1 = 2.
    // At index 2, the instability score is 3 - 1 = 2.
    // None of these values is less than or equal to k = 1, so the answer is -1.
    fmt.Println(firstStableIndex([]int{3,2,1}, 1)) // -1
    // Example 3:
    // Input: nums = [0], k = 0
    // Output: 0
    // Explanation:
    // At index 0, the instability score is 0 - 0 = 0, which is less than or equal to k = 0. Therefore, the answer is 0.
    fmt.Println(firstStableIndex([]int{0}, 0)) // 0

    fmt.Println(firstStableIndex([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
    fmt.Println(firstStableIndex([]int{9,8,7,6,5,4,3,2,1}, 2)) // -1

    fmt.Println(firstStableIndex1([]int{5,0,1,4}, 3)) // 3
    fmt.Println(firstStableIndex1([]int{3,2,1}, 1)) // -1
    fmt.Println(firstStableIndex1([]int{0}, 0)) // 0
    fmt.Println(firstStableIndex1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
    fmt.Println(firstStableIndex1([]int{9,8,7,6,5,4,3,2,1}, 2)) // -1
}