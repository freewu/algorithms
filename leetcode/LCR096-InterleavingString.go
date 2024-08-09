package main

// LCR 096. 交错字符串
// 给定三个字符串 s1、s2、s3，请判断 s3 能不能由 s1 和 s2 交织（交错） 组成。
// 两个字符串 s 和 t 交织 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//     s = s1 + s2 + ... + sn
//     t = t1 + t2 + ... + tm
//     |n - m| <= 1
//     交织 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...

// 提示：a + b 意味着字符串 a 和 b 连接。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/09/02/interleave.jpg" />
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出：true

// 示例 2：
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 输出：false

// 示例 3：
// 输入：s1 = "", s2 = "", s3 = ""
// 输出：true

// 提示：
//     0 <= s1.length, s2.length <= 100
//     0 <= s3.length <= 200
//     s1、s2、和 s3 都由小写英文字母组成

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }
    visited := make(map[int]bool)
    var dfs func(s1, s2, s3 string, p1, p2 int, visited map[int]bool) bool
    dfs = func (s1, s2, s3 string, p1, p2 int, visited map[int]bool) bool {
        if p1+p2 == len(s3) {
            return true
        }
        if _, ok := visited[(p1*len(s3))+p2]; ok {
            return false
        }
        visited[(p1*len(s3))+p2] = true
        var match1, match2 bool
        if p1 < len(s1) && s3[p1+p2] == s1[p1] {
            match1 = true
        }
        if p2 < len(s2) && s3[p1+p2] == s2[p2] {
            match2 = true
        }
        if match1 && match2 {
            return dfs(s1, s2, s3, p1+1, p2, visited) || dfs(s1, s2, s3, p1, p2+1, visited)
        } else if match1 {
            return dfs(s1, s2, s3, p1+1, p2, visited)
        } else if match2 {
            return dfs(s1, s2, s3, p1, p2+1, visited)
        } else {
            return false
        }
    }
    return dfs(s1, s2, s3, 0, 0, visited)
}

// best solution
func isInterleave1(s1 string, s2 string, s3 string) bool {
    if len(s1)+len(s2) != len(s3) {
        return false
    }
    dp := make([][]bool, len(s1)+1)
    for i := 0; i < len(s1)+1; i++ {
        dp[i] = make([]bool, len(s2)+1)
        for j := 0; j < len(s2)+1; j++ {
            if i == 0 && j == 0 {
                dp[i][j] = true
                continue
            }
            if i == 0 {
                dp[i][j] = dp[i][j-1] && s3[j-1] == s2[j-1]
                continue
            }
            if j == 0 {
                dp[i][j] = dp[i-1][j] && s3[i-1] == s1[i-1]
                continue
            }
            dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1] || dp[i][j-1] && s2[j-1] == s3[i+j-1]
        }
    }
    return dp[len(s1)][len(s2)]
}

func main() {
    fmt.Printf("isInterleave(\"aabcc\",\"dbbca\",\"aadbbcbcac\") = %v\n",isInterleave("aabcc","dbbca","aadbbcbcac")) // true
    fmt.Printf("isInterleave(\"aabcc\",\"dbbca\",\"aadbbbaccc\") = %v\n",isInterleave("aabcc","dbbca","aadbbbaccc")) // false
    fmt.Printf("isInterleave(\"\",\"\",\"\") = %v\n",isInterleave("","","")) // true

    fmt.Printf("isInterleave1(\"aabcc\",\"dbbca\",\"aadbbcbcac\") = %v\n",isInterleave1("aabcc","dbbca","aadbbcbcac")) // true
    fmt.Printf("isInterleave1(\"aabcc\",\"dbbca\",\"aadbbbaccc\") = %v\n",isInterleave1("aabcc","dbbca","aadbbbaccc")) // false
    fmt.Printf("isInterleave1(\"\",\"\",\"\") = %v\n",isInterleave1("","","")) // true
}