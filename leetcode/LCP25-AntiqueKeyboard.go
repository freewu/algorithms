package main

// LCP 25. 古董键盘
// 小扣在秋日市集购买了一个古董键盘。
// 由于古董键盘年久失修，键盘上只有 26 个字母 a~z 可以按下，且每个字母最多仅能被按 k 次。

// 小扣随机按了 n 次按键，请返回小扣总共有可能按出多少种内容。
// 由于数字较大，最终答案需要对 1000000007 (1e9 + 7) 取模。

// 示例 1：
// 输入：k = 1, n = 1
// 输出：26
// 解释：由于只能按一次按键，所有可能的字符串为 "a", "b", ... "z"

// 示例 2：
// 输入：k = 1, n = 2
// 输出：650
// 解释：由于只能按两次按键，且每个键最多只能按一次，所有可能的字符串（按字典序排序）为 "ab", "ac", ... "zy"

// 提示：
//     1 <= k <= 5
//     1 <= n <= 26*k

import "fmt"

func keyboard(k int, n int) int {
    comb := make([][]int, n + 1)
    for i := 0; i < len(comb); i++ {
        comb[i] = make([]int, n + 1)
    }
    comb[0][0] = 1
    for i := 1; i <= n; i++ {
        comb[i][0] = 1
        for j := 1; j <= i; j++ {
            comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j]
        }
    }
    dp := make([]int, n + 1)
    dp[0] = 1
    for i := 0; i < 26; i++ {
        for j := n; j >= 1; j-- {
            for l := 1; l <= k; l++ {
                if l > j { break }
                dp[j] += dp[j -l] * comb[j][l]
                dp[j] %= 1_000_000_007
            }
        }
    }
    return dp[n]
}

func main() {
    // 示例 1：
    // 输入：k = 1, n = 1
    // 输出：26
    // 解释：由于只能按一次按键，所有可能的字符串为 "a", "b", ... "z"
    fmt.Println(keyboard(1, 1)) // 26
    // 示例 2：
    // 输入：k = 1, n = 2
    // 输出：650
    // 解释：由于只能按两次按键，且每个键最多只能按一次，所有可能的字符串（按字典序排序）为 "ab", "ac", ... "zy"
    fmt.Println(keyboard(1, 2)) // 650

    fmt.Println(keyboard(5, 1)) // 26
    fmt.Println(keyboard(5, 130)) // 735365374
    fmt.Println(keyboard(1, 26)) // 459042011
    fmt.Println(keyboard(2, 52)) // 157365993
    fmt.Println(keyboard(3, 72)) // 383055294
    fmt.Println(keyboard(4, 104)) // 233347976
}