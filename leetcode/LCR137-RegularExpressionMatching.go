package main

// LCR 137. 模糊搜索验证
// 请设计一个程序来支持用户在文本编辑器中的模糊搜索功能。用户输入内容中可能使用到如下两种通配符：
//     '.' 匹配任意单个字符。
//     '*' 匹配零个或多个前面的那一个元素。

// 请返回用户输入内容 input 所有字符是否可以匹配原文字符串 article。

// 示例 1:
// 输入: article = "aa", input = "a"
// 输出: false
// 解释: "a" 无法匹配 "aa" 整个字符串。

// 示例 2:
// 输入: article = "aa", input = "a*"
// 输出: true
// 解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

// 示例 3:
// 输入: article = "ab", input = ".*"
// 输出: true
// 解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

// 提示：
//     1 <= article.length <= 20
//     1 <= input.length <= 20
//     article 只包含从 a-z 的小写字母。
//     input 只包含从 a-z 的小写字母，以及字符 . 和 * 。
//     保证每次出现字符 * 时，前面都匹配到有效的字符

import "fmt"

// 递归
func articleMatch(s string, p string) bool {
    if len(p) == 0 {
        return len(s) == 0
    }
    if len(p) == 1 {
        return (len(s) == 1) && (s[0] == p[0] || p[0] == '.')
    }
    if p[1] != '*' {
        if len(s) == 0 {
            return false
        } 
        return (s[0] == p[0] || p[0] == '.') && articleMatch(s[1:], p[1:])
    }
    for len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
        if (articleMatch(s, p[2:])) {
            return true
        }
        s = s[1:]
    }
    return articleMatch(s, p[2:])
}

// dp
func articleMatch1(s string, p string) bool {
    n, m := len(s), len(p)
    dp := make([][]bool, n + 1)
    for i := 0; i <= n; i ++ {
        dp[i] = make([]bool, m + 1)
    }
    dp[0][0] = true
    match := func(i, j int) bool {
        if i == 0 || j == 0 {
            return false
        }
        return s[i - 1] == p[j - 1] || p[j - 1] == '.'
    }
    for i := 0; i <= n; i ++ {
        for j := 0; j <= m; j ++ {
            if match(i, j) {
                dp[i][j] = dp[i - 1][j - 1]
                continue
            }
            if j == 0 || p[j - 1] != '*' {
                continue
            }
            if match(i, j - 1) {
                dp[i][j] = dp[i][j - 1] || dp[i][j - 2] || dp[i - 1][j]
            } else {
                dp[i][j] = dp[i][j - 2]
            }
        }
    }
    return dp[n][m]
}

func articleMatch2(s string, p string) bool {
    n, m := len(s), len(p)
    dp := make([][]bool, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]bool, m+1)
    }
    dp[0][0] = true
    for i := 0; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if p[j-1] == '*' {
                dp[i][j] = dp[i][j] || dp[i][j-2]
                if i >= 1 &&(s[i-1] == p[j-2] || p[j-2] == '.') {
                    dp[i][j] = dp[i][j] || dp[i-1][j]
                }
            } else {
                if i >= 1 && (s[i-1] == p[j-1] || p[j-1] == '.') {
                    dp[i][j] = dp[i][j] || dp[i-1][j-1]
                }
            }
        }
    }
    return dp[len(s)][len(p)]
}

func main() {
    fmt.Println(articleMatch("aa","a")) // false
    fmt.Println(articleMatch("aa","a*")) // true
    fmt.Println(articleMatch("ab",".*")) // true

    fmt.Println(articleMatch1("aa","a")) // false
    fmt.Println(articleMatch1("aa","a*")) // true
    fmt.Println(articleMatch1("ab",".*")) // true

    fmt.Println(articleMatch2("aa","a")) // false
    fmt.Println(articleMatch2("aa","a*")) // true
    fmt.Println(articleMatch2("ab",".*")) // true
}