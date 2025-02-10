package main

// LCP 69. Hello LeetCode!
// 力扣嘉年华同样准备了纪念品展位，参观者只需要集齐 helloleetcode 的 13 张字母卡片即可获得力扣纪念章。

// 在展位上有一些由字母卡片拼成的单词，words[i][j] 表示第 i 个单词的第 j 个字母。

// 你可以从这些单词中取出一些卡片，但每次拿取卡片都需要消耗游戏代币，规则如下：
//     1. 从一个单词中取一个字母所需要的代币数量，为该字母左边和右边字母数量之积
//     2. 可以从一个单词中多次取字母，每个字母仅可被取一次

// 例如：从 example 中取出字母 a，需要消耗代币 2*4=8，字母取出后单词变为 exmple； 再从中取出字母 m，需要消耗代币 2*3=6，字母取出后单词变为 exple；

// 请返回取得 helloleetcode 这些字母需要消耗代币的 最少 数量。如果无法取得，返回 -1。

// 注意：
//     1. 取出字母的顺序没有要求
//     2. 取出的所有字母恰好可以拼成 helloleetcode

// 示例 1：
// 输入：words = ["hold","engineer","cost","level"]
// 输出：5
// 解释：最优方法为： 从 hold 依次取出 h、o、l、d， 代价均为 0 从 engineer 依次取出第 1 个 e 与最后一个 e， 代价为 0 和 5*1=5 从 cost 取出 c、o、t， 代价均为 0 从 level 依次取出 l、l、e、e， 代价均为 0 所有字母恰好可以拼成 helloleetcode，因此最小的代价为 5

// 示例 2：
// 输入：words = ["hello","leetcode"]
// 输出：0

// 提示：
//     n == words.length
//     m == words[i].length
//     1 <= n <= 24
//     1 <= m <= 8
//     words[i][j] 仅为小写字母

import "fmt"

func Leetcode(words []string) int {
    const keys = "elohtcd"
    const full = 2012 // 0b11111011100，每个字母都选到了对应的上限
    // pos：字母在二进制上的起始位置
    // limit：这个字母能选择的上限
    // mask：位掩码
    rules := ['z' + 1]struct{ pos, limit, mask int } {
        'e': {0, 4, 7},
        'l': {3, 3, 3},
        'o': {5, 2, 3},
        'h': {7, 1, 1},
        't': {8, 1, 1},
        'c': {9, 1, 1},
        'd': {10, 1, 1},
    }
    merge := func(cur, add int) int { // 合并两种选择字母的方案
        for _, c := range keys {
            r := rules[c]
            c1 := cur >> r.pos & r.mask
            c2 := add >> r.pos & r.mask
            if c1+c2 > r.limit {
                return -1
            }
            cur += c2 << r.pos
        }
        return cur
    }
    n, inf := len(words), 1 << 31
    // 预处理每个单词的每种选择字母的方案所消耗的代币的最小值
    costs := make([][1 << 11]int, n)
    for i, word := range words {
        for j := range costs[i] {
            costs[i][j] = inf
        }
        var f func(string, int, int)
        f = func(s string, mask, tot int) {
            costs[i][mask] = min(costs[i][mask], tot)
            for j, c := range s { // 枚举选择字母的位置
                r := rules[c]
                if mask>>r.pos&r.mask < r.limit { // 可以选字母 c
                    f(s[:j]+s[j+1:], mask+1<<r.pos, tot+j*(len(s)-1-j))
                }
            }
        }
        f(word, 0, 0)
    }
    dp := make([][1 << 11]int, n)
    for i := range dp {
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    var f func(int, int) int
    f = func(i, mask int) int {
        if i == n {
            if mask == full {
                return 0
            }
            return inf // inf 表示不合法，没有选完要求的字母
        }
        ptr := &dp[i][mask]
        if *ptr != -1 {
            return *ptr
        }
        res := inf
        for add, tot := range costs[i][:] {
            if tot >= res { // 剪枝
                continue
            }
            m2 := merge(mask, add)
            if m2 >= 0 {
                res = min(res, f(i+1, m2)+tot)
            }
        }
        *ptr = res
        return res
    }
    res := f(0, 0)
    if res == inf {
        return -1
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：words = ["hold","engineer","cost","level"]
    // 输出：5
    // 解释：最优方法为： 从 hold 依次取出 h、o、l、d， 代价均为 0 从 engineer 依次取出第 1 个 e 与最后一个 e， 代价为 0 和 5*1=5 从 cost 取出 c、o、t， 代价均为 0 从 level 依次取出 l、l、e、e， 代价均为 0 所有字母恰好可以拼成 helloleetcode，因此最小的代价为 5
    fmt.Println(Leetcode([]string{"hold","engineer","cost","level"})) // 5
    // 示例 2：
    // 输入：words = ["hello","leetcode"]
    // 输出：0
    fmt.Println(Leetcode([]string{"hello","leetcode"})) // 0

    fmt.Println(Leetcode([]string{"bluefrog","leetcode"})) // -1
}