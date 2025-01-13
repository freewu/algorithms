package main

// 3395. Subsequences with a Unique Middle Mode I
// Given an integer array nums, find the number of subsequences of size 5 of nums with a unique middle mode.

// Since the answer may be very large, return it modulo 10^9 + 7.

// A mode of a sequence of numbers is defined as the element that appears the maximum number of times in the sequence.

// A sequence of numbers contains a unique mode if it has only one mode.

// A sequence of numbers seq of size 5 contains a unique middle mode if the middle element (seq[2]) is a unique mode.

// Example 1:
// Input: nums = [1,1,1,1,1,1]
// Output: 6
// Explanation:
// [1, 1, 1, 1, 1] is the only subsequence of size 5 that can be formed, and it has a unique middle mode of 1. This subsequence can be formed in 6 different ways, so the output is 6. 

// Example 2:
// Input: nums = [1,2,2,3,3,4]
// Output: 4
// Explanation:
// [1, 2, 2, 3, 4] and [1, 2, 3, 3, 4] each have a unique middle mode because the number at index 2 has the greatest frequency in the subsequence. [1, 2, 2, 3, 3] does not have a unique middle mode because 2 and 3 appear twice.

// Example 3:
// Input: nums = [0,1,2,3,4,5,6,7,8]
// Output: 0
// Explanation:
// There is no subsequence of length 5 with a unique middle mode.

// Constraints:
//     5 <= nums.length <= 1000
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func subsequencesWithMiddleMode(nums []int) int {
    n := len(nums)
    res := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120 // 所有方案数
    suf := map[int]int{}
    for _, num := range nums {
        suf[num]++
    }
    comb2 := func(num int) int { return num * (num - 1) / 2 }
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
        res -= comb2(left-px) * comb2(right-sx)
        res -= (cp - comb2(px)) * sx * (right - sx)
        res -= (cs - comb2(sx)) * px * (left - px)
        res -= ((ps-px*sx)*(right-sx) - (ps2 - px*sx*sx)) * px
        res -= ((ps-px*sx)*(left-px)  - (p2s - px*px*sx)) * sx

        cp += px
        ps += sx
        ps2 += sx * sx
        p2s += (px*2 + 1) * sx

        pre[x]++
    }
    return res % 1_000_000_007
}

func subsequencesWithMiddleMode1(nums []int) int {
    n := len(nums)
    res := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120 // 所有方案数
    suf := map[int]int{}
    for _, x := range nums {
        suf[x]++
    }
    comb2 := func(num int) int { return num * (num - 1) / 2 }
    pre := make(map[int]int, len(suf)) // 预分配空间
    for left, x := range nums[:n-2] { // 枚举 x，作为子序列正中间的数
        suf[x]--
        if left > 1 {
            right := n - 1 - left
            preX, sufX := pre[x], suf[x]
            // 不合法：只有一个 x
            res -= comb2(left-preX) * comb2(right-sufX)
            // 不合法：只有两个 x，且至少有两个 y（y != x）
            for y, sufY := range suf { // 注意 sufY 可能是 0
                if y == x { continue }
                preY := pre[y]
                // 左边有两个 y，右边有一个 x，即 yy x xz（z 可以等于 y）
                res -= comb2(preY) * sufX * (right - sufX)
                // 右边有两个 y，左边有一个 x，即 zx x yy（z 可以等于 y）
                res -= comb2(sufY) * preX * (left - preX)
                // 左右各有一个 y，另一个 x 在左边，即 xy x yz（z != y）
                res -= preY * sufY * preX * (right - sufX - sufY)
                // 左右各有一个 y，另一个 x 在右边，即 zy x xy（z != y）
                res -= preY * sufY * sufX * (left - preX - preY)
            }
        }
        pre[x]++
    }
    return res % 1_000_000_007
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

    fmt.Println(subsequencesWithMiddleMode1([]int{1,1,1,1,1,1})) // 6
    fmt.Println(subsequencesWithMiddleMode1([]int{1,2,2,3,3,4})) // 4
    fmt.Println(subsequencesWithMiddleMode1([]int{0,1,2,3,4,5,6,7,8})) // 0
    fmt.Println(subsequencesWithMiddleMode1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(subsequencesWithMiddleMode1([]int{9,8,7,6,5,4,3,2,1})) // 0
}