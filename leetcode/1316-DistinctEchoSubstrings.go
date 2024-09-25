package main

// 1316. Distinct Echo Substrings
// Return the number of distinct non-empty substrings of text 
// that can be written as the concatenation of some string with itself 
// (i.e. it can be written as a + a where a is some string).

// Example 1:
// Input: text = "abcabcabc"
// Output: 3
// Explanation: The 3 substrings are "abcabc", "bcabca" and "cabcab".

// Example 2:
// Input: text = "leetcodeleetcode"
// Output: 2
// Explanation: The 2 substrings are "ee" and "leetcodeleetcode".

// Constraints:
//     1 <= text.length <= 2000
//     text has only lowercase English letters.

import "fmt"
import "math/rand"

func distinctEchoSubstrings(text string) int {
    echo := make(map[string]struct{})
    for i := range text {
        for j := i + 1; j-i <= len(text)-j; j++ {
            if t := text[i:j]; t == text[j:2*j-i] {
                echo[t] = struct{}{}
            }
        }
    }
    return len(echo)
}

func distinctEchoSubstrings1(text string) int {
    // 字符串hash O(n^2)
    // 匹配任意相连相等长度的字符串是否相等
    res, n, mod := 0, len(text), 1_070_777_777
    base := 9e8 - rand.Intn(1e8)
    powBase, preHash := make([]int, n+1), make([]int, n+1)
    powBase[0] = 1
    for i, ch := range text {
        powBase[i+1] = (powBase[i] * base) % mod
        preHash[i+1] = (preHash[i]*base + int(ch)) % mod
    }
    subHash := func(l, r int) (h int) { // s[l:r)开区间的hash值
        h = preHash[r] - preHash[l]*powBase[r-l]
        return (h % mod + mod) % mod
    }
    has := map[int]bool{} // k: string的hash v:是否已经存在
    for i := 0; i < n; i++ {
        for le := 1; le <= (n-i)/2; le++ { // i到end还有n-i的长度,
            if leftHash := subHash(i, i+le); leftHash == subHash(i+le, i+2*le) && !has[leftHash] {
                res++
                has[leftHash] = true
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: text = "abcabcabc"
    // Output: 3
    // Explanation: The 3 substrings are "abcabc", "bcabca" and "cabcab".
    fmt.Println(distinctEchoSubstrings("abcabcabc")) // 3
    // Example 2:
    // Input: text = "leetcodeleetcode"
    // Output: 2
    // Explanation: The 2 substrings are "ee" and "leetcodeleetcode".
    fmt.Println(distinctEchoSubstrings("leetcodeleetcode")) // 2

    fmt.Println(distinctEchoSubstrings1("abcabcabc")) // 3
    fmt.Println(distinctEchoSubstrings1("leetcodeleetcode")) // 2
}