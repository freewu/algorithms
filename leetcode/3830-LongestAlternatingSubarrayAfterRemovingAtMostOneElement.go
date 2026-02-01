package main

// 3830. Longest Alternating Subarray After Removing At Most One Element
// You are given an integer array nums.

// A subarray nums[l..r] is alternating if one of the following holds:
//     1. nums[l] < nums[l + 1] > nums[l + 2] < nums[l + 3] > ...
//     2. nums[l] > nums[l + 1] < nums[l + 2] > nums[l + 3] < ...

// In other words, if we compare adjacent elements in the subarray, then the comparisons alternate between strictly greater and strictly smaller.

// You can remove at most one element from nums. Then, you select an alternating subarray from nums.

// Return an integer denoting the maximum length of the alternating subarray you can select.

// A subarray of length 1 is considered alternating.

// Example 1:
// Input: nums = [2,1,3,2]
// Output: 4
// Explanation:
// Choose not to remove elements.
// Select the entire array [2, 1, 3, 2], which is alternating because 2 > 1 < 3 > 2.

// Example 2:
// Input: nums = [3,2,1,2,3,2,1]
// Output: 4
// Explanation:
// Choose to remove nums[3] i.e., [3, 2, 1, 2, 3, 2, 1]. The array becomes [3, 2, 1, 3, 2, 1].
// Select the subarray [3, 2, 1, 3, 2, 1].

// Example 3:
// Input: nums = [100000,100000]
// Output: 1
// Explanation:
// Choose not to remove elements.
// Select the subarray [100000, 100000].

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"
import "slices"
import "cmp"

func longestAlternating(nums []int) int {
    calc := func(arr []int) []int {
        f := make([]int, len(arr)) // f[i] 表示以 i 结尾的最长交替子数组的长度
        for i, x := range arr {
            if i == 0 || arr[i-1] == x {
                f[i] = 1
            } else if i > 1 && arr[i-2] != arr[i-1] && (arr[i-2] < arr[i-1]) == (arr[i-1] > x) {
                f[i] = f[i-1] + 1
            } else {
                f[i] = 2
            }
        }
        return f
    }
    n := len(nums)
    pre := calc(nums) // pre[i] 表示以 i 结尾的最长交替子数组的长度
    slices.Reverse(nums)
    suf := calc(nums) // suf[i] 表示以 i 开头的最长交替子数组的长度
    slices.Reverse(suf)
    slices.Reverse(nums)
    // 不删除元素时的最长交替子数组的长度
    res := slices.Max(pre)
    // 枚举删除 nums[i]
    for i := 1; i < n-1; i++ {
        if nums[i-1] == nums[i+1] { continue } // 无法拼接
        // 计算 (i-2,i-1), (i-1,i+1), (i+1,i+2) 的大小关系
        x := 0
        if i > 1 {
            x = cmp.Compare(nums[i-2], nums[i-1])
        }
        y := cmp.Compare(nums[i-1], nums[i+1])
        z := 0
        if i < n-2 {
            z = cmp.Compare(nums[i+1], nums[i+2])
        }
        if x == -y && x == z { // 左右两边可以拼接
            res = max(res, pre[i-1]+suf[i+1])
        } else {
            if x == -y {
                res = max(res, pre[i-1]+1) // 只拼接 nums[i+1] 
            }
            if z == -y {
                res = max(res, suf[i+1]+1) // 只拼接 nums[i-1] 
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3,2]
    // Output: 4
    // Explanation:
    // Choose not to remove elements.
    // Select the entire array [2, 1, 3, 2], which is alternating because 2 > 1 < 3 > 2.
    fmt.Println(longestAlternating([]int{2,1,3,2})) // 4
    // Example 2:
    // Input: nums = [3,2,1,2,3,2,1]
    // Output: 4
    // Explanation:
    // Choose to remove nums[3] i.e., [3, 2, 1, 2, 3, 2, 1]. The array becomes [3, 2, 1, 3, 2, 1].
    // Select the subarray [3, 2, 1, 3, 2, 1].
    fmt.Println(longestAlternating([]int{3,2,1,2,3,2,1})) // 4
    // Example 3:
    // Input: nums = [100000,100000]
    // Output: 1
    // Explanation:
    // Choose not to remove elements.
    // Select the subarray [100000, 100000].
    fmt.Println(longestAlternating([]int{100000,100000})) // 1

    fmt.Println(longestAlternating([]int{1,2,3,4,5,6,7,8,9})) // 2
    fmt.Println(longestAlternating([]int{9,8,7,6,5,4,3,2,1})) // 2
}