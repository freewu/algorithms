package main

// LCP 19. 秋叶收藏集
// 小扣出去秋游，途中收集了一些红叶和黄叶，他利用这些叶子初步整理了一份秋叶收藏集 leaves， 
// 字符串 leaves 仅包含小写字符 r 和 y， 其中字符 r 表示一片红叶，字符 y 表示一片黄叶。 
// 出于美观整齐的考虑，小扣想要将收藏集中树叶的排列调整成「红、黄、红」三部分。
// 每部分树叶数量可以不相等，但均需大于等于 1。
// 每次调整操作，小扣可以将一片红叶替换成黄叶或者将一片黄叶替换成红叶。
// 请问小扣最少需要多少次调整操作才能将秋叶收藏集调整完毕。

// 示例 1：
// 输入：leaves = "rrryyyrryyyrr"
// 输出：2
// 解释：调整两次，将中间的两片红叶替换成黄叶，得到 "rrryyyyyyyyrr"

// 示例 2：
// 输入：leaves = "ryr"
// 输出：0
// 解释：已符合要求，不需要额外操作

// 提示：
//     3 <= leaves.length <= 10^5
//     leaves 中只包含字符 'r' 和字符 'y'

import "fmt"

// dp
func minimumOperations(leaves string) int {
    n, inf := len(leaves), 1 << 31
    dp := make([][3]int, n)
    min := func(x, y int) int { if x < y { return x; }; return y; }
    boolToInt := func(x bool) int { if x { return 1; }; return 0; }
    dp[0] = [3]int{ boolToInt(leaves[0] == 'y'), inf, inf }
    dp[1] = [3]int{ 0, 0, inf }
    for i := 1; i < n; i++ {
        isRed, isYellow := boolToInt(leaves[i] == 'r'), boolToInt(leaves[i] == 'y')
        dp[i][0] = dp[i - 1][0] + isYellow
        dp[i][1] = min(dp[i - 1][0], dp[i - 1][1]) + isRed
        if i >= 2 {
            dp[i][2] = min(dp[i - 1][1], dp[i - 1][2]) + isYellow
        }
    }
    return dp[n - 1][2]
}

func minimumOperations1(leaves string) int {
    g, n, inf := -1, len(leaves), 1 << 31
    if leaves[0] == 'y' {
        g = 1
    }
    min := func(x, y int) int { if x < y { return x; }; return y; }
    boolToInt := func(x bool) int { if x { return 1; }; return 0; }
    mn, res := g, inf
    for i := 1; i < n; i++ {
        isYellow := boolToInt(leaves[i] == 'y')
        g += 2 * isYellow - 1
        if i != n - 1 {
            res = min(res, mn - g)
        }
        mn = min(mn, g)
    }
    return res + (g + n) / 2
}

func minimumOperations2(leaves string) int {
    n, inf := len(leaves), 1 << 31
    dp := make([][3]int, n)
    for i := range dp {
        dp[i] = [3]int{ inf, inf, inf }
    }
    if leaves[0] != 'r' {
        dp[0][0] = 1
    } else {
        dp[0][0] = 0
    }
    min := func(x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n; i++ {
        if leaves[i] == 'r' {
            dp[i][0] = dp[i - 1][0]
            dp[i][1] = min(dp[i - 1][0], dp[i - 1][1]) + 1
            dp[i][2] = min(dp[i - 1][1], dp[i - 1][2])
        }
        if leaves[i] == 'y' {
            dp[i][0] = dp[i - 1][0] + 1
            dp[i][1] = min(dp[i - 1][0], dp[i - 1][1])
            dp[i][2] = min(dp[i - 1][1], dp[i - 1][2]) + 1
        }
    }
    return dp[n - 1][2]
}

func main() {
    // 示例 1：
    // 输入：leaves = "rrryyyrryyyrr"
    // 输出：2
    // 解释：调整两次，将中间的两片红叶替换成黄叶，得到 "rrryyyyyyyyrr"
    fmt.Println(minimumOperations("rrryyyrryyyrr")) // 2
    // 示例 2：
    // 输入：leaves = "ryr"
    // 输出：0
    // 解释：已符合要求，不需要额外操作
    fmt.Println(minimumOperations("ryr")) // 0

    fmt.Println(minimumOperations("rrrrrrrrrr")) // 1
    fmt.Println(minimumOperations("yyyyyyyyyy")) // 2

    fmt.Println(minimumOperations1("rrryyyrryyyrr")) // 2
    fmt.Println(minimumOperations1("ryr")) // 0
    fmt.Println(minimumOperations1("rrrrrrrrrr")) // 1
    fmt.Println(minimumOperations1("yyyyyyyyyy")) // 2

    fmt.Println(minimumOperations2("rrryyyrryyyrr")) // 2
    fmt.Println(minimumOperations2("ryr")) // 0
    fmt.Println(minimumOperations2("rrrrrrrrrr")) // 1
    fmt.Println(minimumOperations2("yyyyyyyyyy")) // 2
}