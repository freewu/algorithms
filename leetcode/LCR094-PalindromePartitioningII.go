package main

// LCR 094. 分割回文串 II
// 给定一个字符串 s，请将 s 分割成一些子串，使每个子串都是回文串。
// 返回符合要求的 最少分割次数 。

// 示例 1：
// 输入：s = "aab"
// 输出：1
// 解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

// 示例 2：
// 输入：s = "a"
// 输出：0

// 示例 3：
// 输入：s = "ab"
// 输出：1

// 提示：
//     1 <= s.length <= 2000
//     s 仅由小写英文字母组成

import "fmt"

// dfs
func minCut(str string) int {
    mem := map[int]int{}
    isPalindrome := func(str string) bool {
        for i := 0; i < len(str)/2; i++ {
            if str[i] != str[len(str)-1-i] {
                return false
            }
        }
        return true
    }
    var dfs func(string, int) int
    dfs = func(str string, cutsDone int) int {
        if cutsToDo, ok := mem[len(str)]; ok {
            return cutsDone + cutsToDo
        }
        cutsToDo := 0
        for i := 1; i <= len(str); i++ {
            if isPalindrome(str[:i]) {
                ret := dfs(str[i:], cutsDone + 1)
                if cutsToDo == 0 || ret - cutsDone < cutsToDo {
                    cutsToDo = ret - cutsDone
                }
            }
        }
        mem[len(str)] = cutsToDo
        return cutsDone + cutsToDo
    }
    return dfs(str, 0) - 1
}

// dp
func minCut1(s string) int {
    dp := make([]int, len(s)) 
    for i := range dp { // 最差的情况 每个字符切一下
        dp[i] = i
    }
    for m := 1; m < len(s); m++ {
        {
            i, j := m, m
            for i >= 0 && j < len(s) && s[i] == s[j] {
                n := 0
                if i != 0 {
                    n = dp[i - 1] + 1
                }
                if n < dp[j] {
                    dp[j] = n
                }
                i -= 1
                j += 1
            }
        }
        {
            i, j := m - 1, m
            for i >= 0 && j < len(s) && s[i] == s[j] {
                n := 0
                if i != 0 {
                    n = dp[i - 1] + 1
                }
                if n < dp[j] {
                    dp[j] = n
                }
                i -= 1
                j += 1
            }
        }
    }
    return dp[len(s) - 1]
}

func main() {
    // Example 1:
    // Input: s = "aab"
    // Output: 1
    // Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
    fmt.Println(minCut("aab")) // 1
    // Example 2:
    // Input: s = "a"
    // Output: 0
    fmt.Println(minCut("a")) // 0
    // Example 3:
    // Input: s = "ab"
    // Output: 1
    fmt.Println(minCut("ab")) // 1

    fmt.Println(minCut1("aab")) // 1
    fmt.Println(minCut1("a")) // 0
    fmt.Println(minCut1("ab")) // 1
}