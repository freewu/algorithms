package main

// LCR 018. 验证回文串
// 给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
// 本题中，将空字符串定义为有效的 回文串 。

// 示例 1:
// 输入: s = "A man, a plan, a canal: Panama"
// 输出: true
// 解释："amanaplanacanalpanama" 是回文串

// 示例 2:
// 输入: s = "race a car"
// 输出: false
// 解释："raceacar" 不是回文串

// 提示：
// 1 <= s.length <= 2 * 10^5
// 字符串 s 由 ASCII 字符组成

import "fmt"
import "strings"

func isPalindrome(s string) bool {
    s = strings.TrimSpace(s)
    s = strings.ToLower(s)
    // 新建一个只有 0-9 a-z的数组
    m := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 97 && s[i] <= 122) {
            m = append(m, s[i])
        }
    }
    if len(m) <= 1 {
        return true
    }
    i, l := 0, len(m) - 1
    // 判断是否是回文件
    for i < l {
        // 从 i -> <- l 判断是否相符
        if m[i] != m[l] {
            return false
        }
        i++
        l--
    }
    return true
}

func isPalindrome1(s string) bool {
    s = strings.ToLower(s)
    // 判断 c 是否是字符或者数字
    isChar := func (c byte) bool { return ('a' <= c && c <= 'z') || ('0' <= c && c <= '9'); }
    i, j := 0, len(s)-1
    // 如果 -> <- 相交了说明判断完，说明输入的是回文
    for i < j {
        // i -> 遇到特殊字符跳过
        for i < j && !isChar(s[i]) {
            i++
        }
        // <- j 遇到特殊字符跳过
        for i < j && !isChar(s[j]) {
            j--
        }
        // 判断收尾是否是相同字符，不是返回false 
        if s[i] != s[j] {
            return false
        }
        // 相符则继续向内收 i-> <-j
        i++
        j--
    }
    return true
}

func main() {
    fmt.Println(isPalindrome("ab"))   // false
    fmt.Println(isPalindrome(" Abc")) // false

    fmt.Println(isPalindrome(""))                               // true
    fmt.Println(isPalindrome("a"))                              // true
    fmt.Println(isPalindrome(" Aba"))                           // true
    fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
    fmt.Println(isPalindrome("12321"))                          // true
    fmt.Println(isPalindrome("123321"))                         // true
    fmt.Println(isPalindrome(".,"))                             // true


    fmt.Println(isPalindrome1("ab"))   // false
    fmt.Println(isPalindrome1(" Abc")) // false

    fmt.Println(isPalindrome1(""))                               // true
    fmt.Println(isPalindrome1("a"))                              // true
    fmt.Println(isPalindrome1(" Aba"))                           // true
    fmt.Println(isPalindrome1("A man, a plan, a canal: Panama")) // true
    fmt.Println(isPalindrome1("12321"))                          // true
    fmt.Println(isPalindrome1("123321"))                         // true
    fmt.Println(isPalindrome1(".,"))                             // true
}