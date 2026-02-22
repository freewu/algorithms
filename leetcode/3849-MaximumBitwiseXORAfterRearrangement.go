package main

// 3849. Maximum Bitwise XOR After Rearrangement
// You are given two binary strings s and t​​​​​​​, each of length n.

// You may rearrange the characters of t in any order, but s must remain unchanged.

// Return a binary string of length n representing the maximum integer value obtainable by taking the bitwise XOR of s and rearranged t.

// Example 1:
// Input: s = "101", t = "011"
// Output: "110"
// Explanation:
// One optimal rearrangement of t is "011".
// The bitwise XOR of s and rearranged t is "101" XOR "011" = "110", which is the maximum possible.

// Example 2:
// Input: s = "0110", t = "1110"
// Output: "1101"
// Explanation:
// One optimal rearrangement of t is "1011".
// The bitwise XOR of s and rearranged t is "0110" XOR "1011" = "1101", which is the maximum possible.

// Example 3:
// Input: s = "0101", t = "1001"
// Output: "1111"
// Explanation:
// One optimal rearrangement of t is "1010".
// The bitwise XOR of s and rearranged t is "0101" XOR "1010" = "1111", which is the maximum possible.
 
// Constraints:
//     1 <= n == s.length == t.length <= 2 * 10^5
//     s[i] and t[i] are either '0' or '1'.

import "fmt"
import "strings"

func maximumXor(s string, t string) string {
    zero, one := 0, 0
    for _, v := range t { // 统计数量
        if v == '1' {
            one++
        } else {
            zero++
        }
    }
    var sb strings.Builder
    for _, v := range s {
        if v == '1' {
            if zero > 0 {
                zero--
                sb.WriteByte('1')
            } else {
                one--
                sb.WriteByte('0')
            }
        } else {
            if one > 0 {
                one--
                sb.WriteByte('1')
            } else {
                zero--
                sb.WriteByte('0')
            }
        }
    }
    return sb.String()
}

func maximumXor1(s string, t string) string {
    res, n, one := []byte(s), len(s), strings.Count(t, "1")
    visited := make([]bool, n)
    for i := range res {
        if res[i] == '0' && one > 0 {
            one--
            res[i] = '1'
            visited[i] = true
        }
    }
    for i := n - 1; i >= 0; i-- {
        if res[i] == '1' && one > 0 && !visited[i] {
            one--
            res[i] = '0'
        }
    }
    return string(res)
}

func maximumXor2(s string, t string) string {
    zero, one, n := 0, 0, len(s)
    res := make([]byte, n)
    for i := 0; i < len(t); i++ {
        if t[i] == '0' {
            zero++
        } else {
            one++
        }
    }
    for i := 0; i < n; i++ {
        if s[i] == '0' {
            if one > 0 {
                res[i] = '1'
                one--
            } else {
                res[i] = '0'
                zero--
            }
        } else {
            if zero > 0 {
                res[i] = '1'
                zero--
            } else {
                res[i] = '0'
                one--
            }
        }
    }
    return string(res)  
}

func main() {
    // Example 1:
    // Input: s = "101", t = "011"
    // Output: "110"
    // Explanation:
    // One optimal rearrangement of t is "011".
    // The bitwise XOR of s and rearranged t is "101" XOR "011" = "110", which is the maximum possible.
    fmt.Println(maximumXor("101","011")) // "110"
    // Example 2:
    // Input: s = "0110", t = "1110"
    // Output: "1101"
    // Explanation:
    // One optimal rearrangement of t is "1011".
    // The bitwise XOR of s and rearranged t is "0110" XOR "1011" = "1101", which is the maximum possible.
    fmt.Println(maximumXor("0110","1110")) // "1101"
    // Example 3:
    // Input: s = "0101", t = "1001"
    // Output: "1111"
    // Explanation:
    // One optimal rearrangement of t is "1010".
    // The bitwise XOR of s and rearranged t is "0101" XOR "1010" = "1111", which is the maximum possible.
    fmt.Println(maximumXor("0101","1001")) // "1111"

    fmt.Println(maximumXor("0000000000","0000000000")) // "0000000000"
    fmt.Println(maximumXor("0000000000","1111111111")) // "1111111111"
    fmt.Println(maximumXor("1111111111","0000000000")) // "1111111111"
    fmt.Println(maximumXor("1111111111","1111111111")) // "0000000000"
    fmt.Println(maximumXor("0101010101","0101010101")) // "1111111111"
    fmt.Println(maximumXor("0101010101","1010101010")) // "1111111111"
    fmt.Println(maximumXor("1010101010","0101010101")) // "1111111111"
    fmt.Println(maximumXor("1010101010","1010101010")) // "1111111111"
    fmt.Println(maximumXor("0000011111","1111100000")) // "1111111111"
    fmt.Println(maximumXor("0000011111","1111100000")) // "1111111111"
    fmt.Println(maximumXor("1111100000","0000011111")) // "1111111111"
    fmt.Println(maximumXor("1111100000","1111100000")) // "1111111111"

    fmt.Println(maximumXor1("101","011")) // "110"
    fmt.Println(maximumXor1("0110","1110")) // "1101"
    fmt.Println(maximumXor1("0101","1001")) // "1111"
    fmt.Println(maximumXor1("0000000000","0000000000")) // "0000000000"
    fmt.Println(maximumXor1("0000000000","1111111111")) // "1111111111"
    fmt.Println(maximumXor1("1111111111","0000000000")) // "1111111111"
    fmt.Println(maximumXor1("1111111111","1111111111")) // "0000000000"
    fmt.Println(maximumXor1("0101010101","0101010101")) // "1111111111"
    fmt.Println(maximumXor1("0101010101","1010101010")) // "1111111111"
    fmt.Println(maximumXor1("1010101010","0101010101")) // "1111111111"
    fmt.Println(maximumXor1("1010101010","1010101010")) // "1111111111"
    fmt.Println(maximumXor1("0000011111","1111100000")) // "1111111111"
    fmt.Println(maximumXor1("0000011111","1111100000")) // "1111111111"
    fmt.Println(maximumXor1("1111100000","0000011111")) // "1111111111"
    fmt.Println(maximumXor1("1111100000","1111100000")) // "1111111111"

    fmt.Println(maximumXor2("101","011")) // "110"
    fmt.Println(maximumXor2("0110","1110")) // "1101"
    fmt.Println(maximumXor2("0101","1001")) // "1111"
    fmt.Println(maximumXor2("0000000000","0000000000")) // "0000000000"
    fmt.Println(maximumXor2("0000000000","1111111111")) // "1111111111"
    fmt.Println(maximumXor2("1111111111","0000000000")) // "1111111111"
    fmt.Println(maximumXor2("1111111111","1111111111")) // "0000000000"
    fmt.Println(maximumXor2("0101010101","0101010101")) // "1111111111"
    fmt.Println(maximumXor2("0101010101","1010101010")) // "1111111111"
    fmt.Println(maximumXor2("1010101010","0101010101")) // "1111111111"
    fmt.Println(maximumXor2("1010101010","1010101010")) // "1111111111"
    fmt.Println(maximumXor2("0000011111","1111100000")) // "1111111111"
    fmt.Println(maximumXor2("0000011111","1111100000")) // "1111111111"
    fmt.Println(maximumXor2("1111100000","0000011111")) // "1111111111"
    fmt.Println(maximumXor2("1111100000","1111100000")) // "1111111111"
}