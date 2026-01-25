package main

// 1984. Minimum Difference Between Highest and Lowest of K Scores
// You are given a 0-indexed integer array nums, where nums[i] represents the score of the ith student. 
// You are also given an integer k.

// Pick the scores of any k students from the array so that the difference between the highest and the lowest of the k scores is minimized.

// Return the minimum possible difference.

// Example 1:
// Input: nums = [90], k = 1
// Output: 0
// Explanation: There is one way to pick score(s) of one student:
// - [90]. The difference between the highest and lowest score is 90 - 90 = 0.
// The minimum possible difference is 0.

// Example 2:
// Input: nums = [9,4,1,7], k = 2
// Output: 2
// Explanation: There are six ways to pick score(s) of two students:
// - [9,4,1,7]. The difference between the highest and lowest score is 9 - 4 = 5.
// - [9,4,1,7]. The difference between the highest and lowest score is 9 - 1 = 8.
// - [9,4,1,7]. The difference between the highest and lowest score is 9 - 7 = 2.
// - [9,4,1,7]. The difference between the highest and lowest score is 4 - 1 = 3.
// - [9,4,1,7]. The difference between the highest and lowest score is 7 - 4 = 3.
// - [9,4,1,7]. The difference between the highest and lowest score is 7 - 1 = 6.
// The minimum possible difference is 2.

// Constraints:
//     1 <= k <= nums.length <= 1000
//     0 <= nums[i] <= 10^5

import "fmt"
import "sort"

func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
    l, r, n := 0, k - 1, len(nums)
    mn, mx, res := nums[l], nums[r], nums[n - 1]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for r < n {
        mn, mx =  nums[l], nums[r]
        res = min(res, mx - mn)
        l++
        r++
    }
    return res
}

func minimumDifference1(nums []int, k int) int {
    if k == 1 { return 0 }
    sort.Ints(nums)
    res := int(^uint(0) >> 1)
    for i := 0; i <= len(nums)-k; i++ {
        diff := nums[i + k -1] - nums[i]
        if diff < res {
            res = diff
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [90], k = 1
    // Output: 0
    // Explanation: There is one way to pick score(s) of one student:
    // - [90]. The difference between the highest and lowest score is 90 - 90 = 0.
    // The minimum possible difference is 0.
    fmt.Println(minimumDifference([]int{90}, 1)) // 0
    // Example 2:
    // Input: nums = [9,4,1,7], k = 2
    // Output: 2
    // Explanation: There are six ways to pick score(s) of two students:
    // - [9,4,1,7]. The difference between the highest and lowest score is 9 - 4 = 5.
    // - [9,4,1,7]. The difference between the highest and lowest score is 9 - 1 = 8.
    // - [9,4,1,7]. The difference between the highest and lowest score is 9 - 7 = 2.
    // - [9,4,1,7]. The difference between the highest and lowest score is 4 - 1 = 3.
    // - [9,4,1,7]. The difference between the highest and lowest score is 7 - 4 = 3.
    // - [9,4,1,7]. The difference between the highest and lowest score is 7 - 1 = 6.
    // The minimum possible difference is 2.
    fmt.Println(minimumDifference([]int{9,4,1,7}, 2)) // 2

    fmt.Println(minimumDifference([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(minimumDifference([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1

    fmt.Println(minimumDifference1([]int{90}, 1)) // 0
    fmt.Println(minimumDifference1([]int{9,4,1,7}, 2)) // 2
    fmt.Println(minimumDifference1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(minimumDifference1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1
}