package main

// 5. Longest Palindromic Substring
// Given a string s, return the longest palindromic substring in s.

// Example 1:
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.

// Example 2:
// Input: s = "cbbd"
// Output: "bb"
 
// Constraints:
//     1 <= s.length <= 1000
//     s consist of only digits and English letters.

import "fmt"

// Manacher's algorithm，时间复杂度 O(n)，空间复杂度 O(n)
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    newS := make([]rune, 0)
    newS = append(newS, '#')
    for _, c := range s {
        newS = append(newS, c)
        newS = append(newS, '#')
    }
    // dp[i]:    以预处理字符串下标 i 为中心的回文半径(奇数长度时不包括中心)
    // maxRight: 通过中心扩散的方式能够扩散的最右边的下标
    // center:   与 maxRight 对应的中心字符的下标
    // maxLen:   记录最长回文串的半径
    // begin:    记录最长回文串在起始串 s 中的起始下标
    dp, maxRight, center, maxLen, begin := make([]int, len(newS)), 0, 0, 1, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(newS); i++ {
        if i < maxRight {
            // 这一行代码是 Manacher 算法的关键所在
            dp[i] = min(maxRight-i, dp[2*center-i])
        }
        // 中心扩散法更新 dp[i]
        left, right := i-(1+dp[i]), i+(1+dp[i])
        for left >= 0 && right < len(newS) && newS[left] == newS[right] {
            dp[i]++
            left--
            right++
        }
        // 更新 maxRight，它是遍历过的 i 的 i + dp[i] 的最大者
        if i+dp[i] > maxRight {
            maxRight = i + dp[i]
            center = i
        }
        // 记录最长回文子串的长度和相应它在原始字符串中的起点
        if dp[i] > maxLen {
            maxLen = dp[i]
            begin = (i - maxLen) / 2 // 这里要除以 2 因为有我们插入的辅助字符 #
        }
    }
    return s[begin : begin+maxLen]
}

// 滑动窗口，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome1(s string) string {
    if len(s) == 0 {
        return ""
    }
    left, right, pl, pr := 0, -1, 0, 0
    for left < len(s) {
        // 移动到相同字母的最右边（如果有相同字母）
        for right+1 < len(s) && s[left] == s[right+1] {
            right++
        }
        // 找到回文的边界
        for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
            left--
            right++
        }
        if right-left > pr-pl {
            pl, pr = left, right
        }
        // 重置到下一次寻找回文的中心
        left = (left + right)/2 + 1
        right = left
    }
    return s[pl : pr+1]
}

// 中心扩散法，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome2(s string) string {
    maxPalindrome := func (s string, i, j int, res string) string {
        sub := ""
        for i >= 0 && j < len(s) && s[i] == s[j] {
            sub = s[i : j+1]
            i--
            j++
        }
        if len(res) < len(sub) {
            return sub
        }
        return res
    }
    res := ""
    for i := 0; i < len(s); i++ {
        res = maxPalindrome(s, i, i, res)
        res = maxPalindrome(s, i, i+1, res)
    }
    return res
}

// DP，时间复杂度 O(n^2)，空间复杂度 O(n^2)
func longestPalindrome3(s string) string {
    res, dp := "", make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
    }
    for i := len(s) - 1; i >= 0; i-- {
        for j := i; j < len(s); j++ {
            dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
            if dp[i][j] && (res == "" || j-i+1 > len(res)) {
                res = s[i : j+1]
            }
        }
    }
    return res
}

func main() {
    fmt.Printf("longestPalindrome(\"babad\") = %v\n",longestPalindrome("babad")) // bab | aba
    fmt.Printf("longestPalindrome(\"cbbd\") = %v\n",longestPalindrome("cbbd")) // bb
    fmt.Printf("longestPalindrome(\"a\") = %v\n",longestPalindrome("a")) // a
    fmt.Printf("longestPalindrome(\"ac\") = %v\n",longestPalindrome("ac")) // a

    fmt.Printf("longestPalindrome1(\"babad\") = %v\n",longestPalindrome1("babad")) // bab | aba
    fmt.Printf("longestPalindrome1(\"cbbd\") = %v\n",longestPalindrome1("cbbd")) // bb
    fmt.Printf("longestPalindrome1(\"a\") = %v\n",longestPalindrome1("a")) // a
    fmt.Printf("longestPalindrome1(\"ac\") = %v\n",longestPalindrome1("ac")) // a

    fmt.Printf("longestPalindrome2(\"babad\") = %v\n",longestPalindrome2("babad")) // bab | aba
    fmt.Printf("longestPalindrome2(\"cbbd\") = %v\n",longestPalindrome2("cbbd")) // bb
    fmt.Printf("longestPalindrome2(\"a\") = %v\n",longestPalindrome2("a")) // a
    fmt.Printf("longestPalindrome2(\"ac\") = %v\n",longestPalindrome2("ac")) // a

    fmt.Printf("longestPalindrome3(\"babad\") = %v\n",longestPalindrome3("babad")) // bab | aba
    fmt.Printf("longestPalindrome3(\"cbbd\") = %v\n",longestPalindrome3("cbbd")) // bb
    fmt.Printf("longestPalindrome3(\"a\") = %v\n",longestPalindrome3("a")) // a
    fmt.Printf("longestPalindrome3(\"ac\") = %v\n",longestPalindrome3("ac")) // a
}
