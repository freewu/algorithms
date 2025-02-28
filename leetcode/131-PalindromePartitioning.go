package main

// 131. Palindrome Partitioning
// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.

// Example 1:
// Input: s = "aab"
// Output: [["a","a","b"],["aa","b"]]

// Example 2:
// Input: s = "a"
// Output: [["a"]]

// Constraints:
//     1 <= s.length <= 16
//     s contains only lowercase English letters.

import "fmt"
import "time"

// DFS 递归求解
func partition(s string) [][]string {
    res := [][]string{}
    size := len(s)
    if size == 0 {
        return res
    }
    // 判断是否是回文
    isPalindrome := func (str string, start, end int) bool {
        for start < end {
            if str[start] != str[end] {
                return false
            }
            start++
            end--
        }
        return true
    }
    var dfs func(s string, idx int, cur []string, result *[][]string)
    dfs = func(s string, idx int, cur []string, result *[][]string) {
        start, end := idx, len(s)
        if start == end {
            temp := make([]string, len(cur))
            copy(temp, cur)
            *result = append(*result, temp)
            return
        }
        for i := start; i < end; i++ {
            if isPalindrome(s, start, i) { // 只处理回文的情况
                dfs(s, i+1, append(cur, s[start:i+1]), result)
            }
        }
    }
    current := make([]string, 0, size)
    dfs(s, 0, current, &res)
    return res
}

// best solution
func partition1(s string) [][]string {
    path, res := []string{}, [][]string {}
    check := func(i, j int) bool {
        for i < j {
            if s[i] != s[j] {
                return false
            }
            i++
            j--
        }
        return true
    }
    var dfs func(start int) 
    dfs = func(start int) {
        if start >= len(s) {
            tmp := make([]string, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        for i := start; i < len(s); i++ {
            if check(start, i) {
                path = append(path, s[start:i+1])
                dfs(i + 1)
                path = path[:len(path) - 1]
            }
        }
    }
    dfs(0)
    return res
}

func partition2(s string) [][]string {
    n := len(s)
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }
    for i := n - 2; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                if j == i+1 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            }
        }
    }
    res, track := [][]string{}, []string{}
    var dfs func(int)
    dfs = func(index int) {
        if index == n {
            res = append(res, append([]string{}, track...))
            return
        }
        for i := index; i < n; i++ {
            if dp[index][i] {
                track = append(track, s[index:i + 1])
                dfs(i + 1)
                track = track[:len(track) - 1]
            }
        }
    }
    dfs(0)
    return res
}

func main() {
    start := time.Now() // 获取当前时间
    // Example 1:
    // Input: s = "aab"
    // Output: [["a","a","b"],["aa","b"]]
    fmt.Println(partition("aab")) // [[a a b] [aa b]]
    // Example 2:
    // Input: s = "a"
    // Output: [["a"]]
    fmt.Println(partition("a")) // [[a]]
    fmt.Println(partition("bluefrog")) // [[a]]
    fmt.Println(partition("leetcode")) // [[a]]
    fmt.Printf("partition use : %v \r\n",time.Since(start))

    start = time.Now() // 获取当前时间
    fmt.Println(partition1("aab")) // [[a a b] [aa b]]
    fmt.Println(partition1("a")) // [[a]]
    fmt.Println(partition1("bluefrog")) // [[b l u e f r o g]]
    fmt.Println(partition1("leetcode")) // [[l e e t c o d e] [l ee t c o d e]]
    fmt.Printf("partition1 use : %v \r\n",time.Since(start))

    start = time.Now() // 获取当前时间
    fmt.Println(partition2("aab")) // [[a a b] [aa b]]
    fmt.Println(partition2("a")) // [[a]]
    fmt.Println(partition2("bluefrog")) // [[b l u e f r o g]]
    fmt.Println(partition2("leetcode")) // [[l e e t c o d e] [l ee t c o d e]]
    fmt.Printf("partition2 use : %v \r\n",time.Since(start))
}