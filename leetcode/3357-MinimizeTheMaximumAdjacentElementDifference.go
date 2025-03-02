package main

// 3357. Minimize the Maximum Adjacent Element Difference
// You are given an array of integers nums. Some values in nums are missing and are denoted by -1.

// You can choose a pair of positive integers (x, y) exactly once and replace each missing element with either x or y.

// You need to minimize the maximum absolute difference between adjacent elements of nums after replacements.

// Return the minimum possible difference.

// Example 1:
// Input: nums = [1,2,-1,10,8]
// Output: 4
// Explanation:
// By choosing the pair as (6, 7), nums can be changed to [1, 2, 6, 10, 8].
// The absolute differences between adjacent elements are:
// |1 - 2| == 1
// |2 - 6| == 4
// |6 - 10| == 4
// |10 - 8| == 2

// Example 2:
// Input: nums = [-1,-1,-1]
// Output: 0
// Explanation:
// By choosing the pair as (4, 4), nums can be changed to [4, 4, 4].

// Example 3:
// Input: nums = [-1,10,-1,8]
// Output: 1
// Explanation:
// By choosing the pair as (11, 9), nums can be changed to [11, 10, 9, 8].

// Constraints:
//     2 <= nums.length <= 10^5
//     nums[i] is either -1 or in the range [1, 10^9].

import "fmt"

func minDifference(nums []int) int {
    res, pre, n := 0, -1, len(nums)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 和空位相邻的最小数字 mn 和最大数字 mx
    mn, mx := 1 << 31, 0
    for i, v := range nums {
        if v != -1 && (i > 0 && nums[i - 1] == -1 || i < n - 1 && nums[i + 1] == -1) {
            mn = min(mn, v)
            mx = max(mx, v)
        }
    }
    update := func(l, r int, big bool) {
        d := (min(r - mn, mx - l) + 1) / 2
        if big {
            d = min(d, (mx - mn + 2) / 3) // d 不能超过上界
        }
        res = max(res, d)
    }
    for i, v := range nums {
        if v == -1 { continue }
        if pre >= 0 {
            if i - pre == 1 {
                res = max(res, abs(v - nums[pre]))
            } else {
                update(min(nums[pre], v), max(nums[pre], v), i - pre > 2)
            }
        } else if i > 0 {
            update(v, v, false)
        }
        pre = i
    }
    if 0 <= pre && pre < n - 1 {
        update(nums[pre], nums[pre], false)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,-1,10,8]
    // Output: 4
    // Explanation:
    // By choosing the pair as (6, 7), nums can be changed to [1, 2, 6, 10, 8].
    // The absolute differences between adjacent elements are:
    // |1 - 2| == 1
    // |2 - 6| == 4
    // |6 - 10| == 4
    // |10 - 8| == 2
    fmt.Println(minDifference([]int{1,2,-1,10,8})) // 4
    // Example 2:
    // Input: nums = [-1,-1,-1]
    // Output: 0
    // Explanation:
    // By choosing the pair as (4, 4), nums can be changed to [4, 4, 4].
    fmt.Println(minDifference([]int{-1,-1,-1})) // 0
    // Example 3:
    // Input: nums = [-1,10,-1,8]
    // Output: 1
    // Explanation:
    // By choosing the pair as (11, 9), nums can be changed to [11, 10, 9, 8].
    fmt.Println(minDifference([]int{-1,10,-1,8})) // 1

    fmt.Println(minDifference([]int{2,-1,4,-1,-1,6})) // 1
}