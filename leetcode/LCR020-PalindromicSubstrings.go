package main

// LCR 020. 回文子串
// 给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

// 示例 1：
// 输入：s = "abc"
// 输出：3
// 解释：三个回文子串: "a", "b", "c"

// 示例 2：
// 输入：s = "aaa"
// 输出：6
// 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 
// 提示：
//     1 <= s.length <= 1000
//     s 由小写英文字母组成

import "fmt"

// 递归
func countSubstrings(s string) int {
    res := 0
    // 从左往右扫一遍字符串，以每个字符做轴，用中心扩散法，依次遍历计数回文子串
    countPalindrome := func (s string, left, right int) int {
        res := 0
        for left >= 0 && right < len(s) {
            if s[left] != s[right] {
                break
            }
            left--
            right++
            res++
        }
        return res
    }
    for i := 0; i < len(s); i++ {
        res += countPalindrome(s, i, i)
        res += countPalindrome(s, i, i+1)
    }
    return res
}

func main() {
    fmt.Println(countSubstrings("abc")) // 3
    fmt.Println(countSubstrings("aaa")) // 6
}