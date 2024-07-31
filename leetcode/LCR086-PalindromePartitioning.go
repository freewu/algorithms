package main

// LCR 086. 分割回文串
// 给定一个字符串 s ，请将 s 分割成一些子串，使每个子串都是 回文串 ，返回 s 所有可能的分割方案。
// 回文串 是正着读和反着读都一样的字符串。

// 示例 1：
// 输入：s = "google"
// 输出：[["g","o","o","g","l","e"],["g","oo","g","l","e"],["goog","l","e"]]

// 示例 2：
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]

// 示例 3：
// 输入：s = "a"
// 输出：[["a"]]

// 提示：
//     1 <= s.length <= 16
//     s 仅由小写英文字母组成

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

func main() {
    start := time.Now() // 获取当前时间
    fmt.Println(partition("aab")) // [[a a b] [aa b]]
    fmt.Println(partition("a")) // [[a]]
    fmt.Printf("ladderLength use : %v \r\n",time.Since(start))

    start = time.Now() // 获取当前时间
    fmt.Println(partition1("aab")) // [[a a b] [aa b]]
    fmt.Println(partition1("a")) // [[a]]
    fmt.Printf("ladderLength use : %v \r\n",time.Since(start))
}