package main

// 3416. Subsequences with a Unique Middle Mode II
// Given an integer array nums, find the number of subsequences of size 5 of nums with a unique middle mode.

// Since the answer may be very large, return it modulo 10^9 + 7.

// A mode of a sequence of numbers is defined as the element that appears the maximum number of times in the sequence.

// A sequence of numbers contains a unique mode if it has only one mode.

// A sequence of numbers seq of size 5 contains a unique middle mode if the middle element (seq[2]) is a unique mode.

// Example 1:
// Input: nums = [1,1,1,1,1,1]
// Output: 6
// Explanation:
// [1, 1, 1, 1, 1] is the only subsequence of size 5 that can be formed from this list, and it has a unique middle mode of 1.

// Example 2:
// Input: nums = [1,2,2,3,3,4]
// Output: 4
// Explanation:
// [1, 2, 2, 3, 4] and [1, 2, 3, 3, 4] have unique middle modes because the number at index 2 has the greatest frequency in the subsequence. [1, 2, 2, 3, 3] does not have a unique middle mode because 2 and 3 both appear twice in the subsequence.

// Example 3:
// Input: nums = [0,1,2,3,4,5,6,7,8]
// Output: 0
// Explanation:
// There does not exist a subsequence of length 5 with a unique middle mode.

// Constraints:
//     5 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

// 解答错误 834 / 843 
func subsequencesWithMiddleMode(nums []int) int {
    n, mod := len(nums), 1_000_000_007
    res := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120 // 所有方案数
    suf := map[int]int{}
    for _, num := range nums {
        suf[num]++
    }
    comb2 := func (num int) int { return num * (num - 1) / 2 }
    pre := make(map[int]int, len(suf)) // 预分配空间
    var cp, cs, ps, p2s, ps2 int
    for _, c := range suf {
        cs += comb2(c)
    }
    // 枚举 x，作为子序列正中间的数
    for left, x := range nums[:n-2] {
        suf[x]--
        px, sx := pre[x], suf[x]
        cs -= sx
        ps -= px
        p2s -= px * px
        ps2 -= (sx*2 + 1) * px

        right := n - 1 - left
        res -= (comb2(left-px) * comb2(right-sx)) % mod
        res -= ((cp - comb2(px)) * sx * (right - sx)) % mod
        res -= ((cs - comb2(sx)) * px * (left - px)) % mod
        res -= (((ps-px*sx)*(right-sx) - (ps2 - px*sx*sx)) * px) % mod
        res -= (((ps-px*sx)*(left-px)  - (p2s - px*px*sx)) * sx) % mod

        cp += px
        ps += sx
        ps2 += sx * sx
        p2s += (px*2 + 1) * sx

        pre[x]++
    }
    return res % mod
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1,1,1,1]
    // Output: 6
    // Explanation:
    // [1, 1, 1, 1, 1] is the only subsequence of size 5 that can be formed from this list, and it has a unique middle mode of 1.
    fmt.Println(subsequencesWithMiddleMode([]int{1,1,1,1,1,1})) // 6
    // Example 2:
    // Input: nums = [1,2,2,3,3,4]
    // Output: 4
    // Explanation:
    // [1, 2, 2, 3, 4] and [1, 2, 3, 3, 4] have unique middle modes because the number at index 2 has the greatest frequency in the subsequence. [1, 2, 2, 3, 3] does not have a unique middle mode because 2 and 3 both appear twice in the subsequence.
    fmt.Println(subsequencesWithMiddleMode([]int{1,2,2,3,3,4})) // 4
    // Example 3:
    // Input: nums = [0,1,2,3,4,5,6,7,8]
    // Output: 0
    // Explanation:
    // There does not exist a subsequence of length 5 with a unique middle mode.
    fmt.Println(subsequencesWithMiddleMode([]int{0,1,2,3,4,5,6,7,8})) // 0

    fmt.Println(subsequencesWithMiddleMode([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(subsequencesWithMiddleMode([]int{9,8,7,6,5,4,3,2,1})) // 0
}