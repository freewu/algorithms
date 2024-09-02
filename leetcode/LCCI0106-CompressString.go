package main

// 面试题 01.06. Compress String LCCI
// Implement a method to perform basic string compression using the counts of repeated characters. 
// For example, the string aabcccccaaa would become a2blc5a3. 
// If the "compressed" string would not become smaller than the original string, your method should return the original string. 
// You can assume the string has only uppercase and lowercase letters (a - z).

// Example 1:
// Input: "aabcccccaaa"
// Output: "a2b1c5a3"

// Example 2:
// Input: "abbccd"
// Output: "abbccd"
// Explanation: 
// The compressed string is "a1b2c2d1", which is longer than the original string.

// Note:
//     0 <= S.length <= 50000

import "fmt"
import "strconv"
import "strings"

func compressString(S string) string {
    n := len(S)
    if n == 0 || n == 1 {
        return S
    }
    res, count := "", 1
    for i := 1; i < n; i++ {
        if i > 0 && S[i-1] == S[i] { // 这符相同则累加
            count++
        } else {
            res += (string(S[i-1]) + strconv.Itoa(count)) // 遇见不同字符了 压缩字符
            count = 1
        }
    }
    res = res + string(S[n-1]) + strconv.Itoa(count)
    if len(res) >= n { // 压缩后的长度一样，使用原始字符返回
        return S
    } else {
        return res
    }
}

func compressString1(S string) string {
    if len(S) == 0 {
        return ""
    }
    str := S + " "
    count, char := 0, S[0]
    sb := &strings.Builder{}
    for i := 0; i != len(str); i++ {
        if char != str[i] { // 遇见不一样的，压缩
            sb.WriteByte(char)
            sb.WriteString(strconv.Itoa(count))
            count, char = 1, str[i]
        } else { // 相同则累加
            count++
        }
    }
    res := sb.String()
    if len(res) >= len(S) {
        return S
    }
    return res
}

func main() {
    // Example 1:
    // Input: "aabcccccaaa"
    // Output: "a2b1c5a3"
    fmt.Println(compressString("aabcccccaaa")) // a2b1c5a3
    // Example 2:
    // Input: "abbccd"
    // Output: "abbccd"
    // Explanation: 
    // The compressed string is "a1b2c2d1", which is longer than the original string.
    fmt.Println(compressString("abbccd")) // abbccd

    fmt.Println(compressString("abcdefggg")) // abcdefggg
    fmt.Println(compressString("aa")) // aa

    fmt.Println(compressString1("aabcccccaaa")) // a2b1c5a3
    fmt.Println(compressString1("abbccd")) // abbccd
    fmt.Println(compressString1("abcdefggg")) // abcdefggg
    fmt.Println(compressString1("aa")) // aa
}