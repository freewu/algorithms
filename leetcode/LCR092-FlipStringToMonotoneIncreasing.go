package main

// LCR 092. 将字符串翻转到单调递增
// 如果一个由 '0' 和 '1' 组成的字符串，是以一些 '0'（可能没有 '0'）后面跟着一些 '1'（也可能没有 '1'）的形式组成的，那么该字符串是 单调递增 的。
// 我们给出一个由字符 '0' 和 '1' 组成的字符串 s，我们可以将任何 '0' 翻转为 '1' 或者将 '1' 翻转为 '0'。
// 返回使 s 单调递增 的最小翻转次数。

// 示例 1：
// 输入：s = "00110"
// 输出：1
// 解释：我们翻转最后一位得到 00111.

// 示例 2：
// 输入：s = "010110"
// 输出：2
// 解释：我们翻转得到 011111，或者是 000111。

// 示例 3：
// 输入：s = "00011000"
// 输出：2
// 解释：我们翻转得到 00000000。

// 提示：
//     1 <= s.length <= 20000
//     s 中只包含字符 '0' 和 '1'

import "fmt"

func minFlipsMonoIncr(s string) int {
    n, i, count0, count1, res := len(s), 0, 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i < n {
        if s[i] == '1' {
            count1++
        } else {
            count0++
        }
        i++
        if count1 < count0 {
            res += count1
            count1, count0 = 0, 0
        }
    }
    return res + min(count0, count1)
}


func minFlipsMonoIncr1(s string) int {
    cnt1, cnt2 := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, c := range s {
        t1, t2 := 0, 0
        if c == '0' {
            t1 = cnt1
            t2 = min(cnt1, cnt2) + 1
        } else {
            t1 = cnt1 + 1
            t2 = min(cnt1, cnt2)
        }
        cnt1, cnt2 = t1, t2
    }
    return min(cnt1, cnt2)
}

func main() {
    // Example 1:
    // Input: s = "00110"
    // Output: 1
    // Explanation: We flip the last digit to get 00111.
    fmt.Println(minFlipsMonoIncr("00110")) // 1
    // Example 2:
    // Input: s = "010110"
    // Output: 2
    // Explanation: We flip to get 011111, or alternatively 000111.
    fmt.Println(minFlipsMonoIncr("010110")) // 2
    // Example 3:
    // Input: s = "00011000"
    // Output: 2
    // Explanation: We flip to get 00000000.
    fmt.Println(minFlipsMonoIncr("00011000")) // 2

    fmt.Println(minFlipsMonoIncr1("00110")) // 1
    fmt.Println(minFlipsMonoIncr1("010110")) // 2
    fmt.Println(minFlipsMonoIncr1("00011000")) // 2
}