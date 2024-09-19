package main

// 1234. Replace the Substring for Balanced String
// You are given a string s of length n containing only four kinds of characters: 'Q', 'W', 'E', and 'R'.

// A string is said to be balanced if each of its characters appears n / 4 times where n is the length of the string.

// Return the minimum length of the substring that can be replaced with any other string of the same length to make s balanced. 
// If s is already balanced, return 0.

// Example 1:
// Input: s = "QWER"
// Output: 0
// Explanation: s is already balanced.

// Example 2:
// Input: s = "QQWE"
// Output: 1
// Explanation: We need to replace a 'Q' to 'R', so that "RQWE" (or "QRWE") is balanced.

// Example 3:
// Input: s = "QQQW"
// Output: 2
// Explanation: We can replace the first "QQ" to "ER". 

// Constraints:
//     n == s.length
//     4 <= n <= 10^5
//     n is a multiple of 4.
//     s contains only 'Q', 'W', 'E', and 'R'.

import "fmt"

// func balancedString(s string) int {
//     mp, res, avg :=make(map[byte]int, 4), 0, len(s) / 4
//     for i := range s {
//         mp[s[i]]++
//     }
//     abs := func(x int) int { if x < 0 { return -x; }; return x; }
//     for _, v := range mp {
//         res += abs(v - avg)
//     }
//     return res
// }

func balancedString(s string) int {
    n := len(s)
    mp := make(map[byte]int,4)
    res, k, j := n, n / 4, 0
    for i := range s {
        mp[s[i]]++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        mp[s[i]]--
        for j < n && mp['Q'] <= k && mp['W'] <= k && mp['E'] <= k && mp['R'] <= k {
            res = min(res, i - j + 1)
            mp[s[j]]++
            j++
        }
    }
    return res
}

func balancedString1(s string) int {
    // 不定长滑窗-求最短 => 两次转换
    // 转换为最小的窗口储存所有多余的字符 => 转换为除窗口外剩余的每种字符数量都<=m
    avg := len(s) / 4
    count := ['Z' + 1]int{}
    for _, ch := range s {
        count[ch]++
    }
    if count['Q'] == avg && count['W'] == avg && count['E'] == avg && count['R'] == avg {
        return 0
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    l, res := 0, 1 << 31
    for r, ch := range s {
        count[ch]-- // 入窗口, cnt现在记录除窗口外的其余字符数量
        // trick!! 这样进入时都是满足的,但是会缩过头,但是r结尾的满足条件更短的不会错过 => 不会错过最佳答案
        // 因为会缩过头,r位置可以求出最佳答案,但是扩张到r+1位置后,r+1引入的字符对结果无帮助,便无法再r+1位置求出答案.
        // 但没关系,它需要l位置往后退一格,但l和之前的r已经求出了更佳答案了.r+1无法求出也不会错过最佳答案
        for count['Q'] <= avg && count['W'] <= avg && count['E'] <= avg && count['R'] <= avg {
            res = min(res, r-l+1)
            count[s[l]]++
            l++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "QWER"
    // Output: 0
    // Explanation: s is already balanced.
    fmt.Println(balancedString("QWER")) // 0
    // Example 2:
    // Input: s = "QQWE"
    // Output: 1
    // Explanation: We need to replace a 'Q' to 'R', so that "RQWE" (or "QRWE") is balanced.
    fmt.Println(balancedString("QQWE")) // 1
    // Example 3:
    // Input: s = "QQQW"
    // Output: 2
    // Explanation: We can replace the first "QQ" to "ER". 
    fmt.Println(balancedString("QQQW")) // 2
    fmt.Println(balancedString("WQWRQQQW")) // 3

    fmt.Println(balancedString1("QWER")) // 0
    fmt.Println(balancedString1("QQWE")) // 1
    fmt.Println(balancedString1("QQQW")) // 2
    fmt.Println(balancedString1("WQWRQQQW")) // 3
}