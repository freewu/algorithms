package main

// 1625. Lexicographically Smallest String After Applying Operations 
// You are given a string s of even length consisting of digits from 0 to 9, and two integers a and b.

// You can apply either of the following two operations any number of times and in any order on s:
//     1. Add a to all odd indices of s (0-indexed). Digits post 9 are cycled back to 0. 
//        For example, if s = "3456" and a = 5, s becomes "3951".
//     2. Rotate s to the right by b positions. For example, if s = "3456" and b = 1, s becomes "6345".

// Return the lexicographically smallest string you can obtain by applying the above operations any number of times on s.

// A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ, 
// string a has a letter that appears earlier in the alphabet than the corresponding letter in b. 
// For example, "0158" is lexicographically smaller than "0190" because the first position they differ is at the third letter, 
// and '5' comes before '9'.

// Example 1:
// Input: s = "5525", a = 9, b = 2
// Output: "2050"
// Explanation: We can apply the following operations:
// Start:  "5525"
// Rotate: "2555"
// Add:    "2454"
// Add:    "2353"
// Rotate: "5323"
// Add:    "5222"
// Add:    "5121"
// Rotate: "2151"
// Add:    "2050"​​​​​
// There is no way to obtain a string that is lexicographically smaller than "2050".

// Example 2:
// Input: s = "74", a = 5, b = 1
// Output: "24"
// Explanation: We can apply the following operations:
// Start:  "74"
// Rotate: "47"
// ​​​​​​​Add:    "42"
// ​​​​​​​Rotate: "24"​​​​​​​​​​​​
// There is no way to obtain a string that is lexicographically smaller than "24".

// Example 3:
// Input: s = "0011", a = 4, b = 2
// Output: "0011"
// Explanation: There are no sequence of operations that will give us a lexicographically smaller string than "0011".

// Constraints:
//     2 <= s.length <= 100
//     s.length is even.
//     s consists of digits from 0 to 9 only.
//     1 <= a <= 9
//     1 <= b <= s.length - 1

import "fmt"

func findLexSmallestString(s string, a int, b int) string {
    res, n := "", len(s)
    for i := 0; i < n; i++ {
        res += "9"
    }
    add := func(s string, a int, start int) string {
        res := []byte(s)
        for i := start; i < len(s); i += 2 {
            res[i] = byte('0' + (int(res[i] - '0') + a) % 10)
        }
        return string(res)
    }
    min := func (x, y string) string { if x < y { return x; }; return y; }
    for count1 := 0; count1 < 10; count1++ {
        s1 := add(s, a * count1, 1)  
        i := 0
        for {
            s2 := s1[i:] + s1[:i]
            res = min(res, s2)
            i = (i + b) % n
            if i == 0 { break }
        }
        if b % 2 != 0 {
            for count2 := 0; count2 < 10; count2++ {
                s2 := add(s1, a * count2, 0)
                i = 0
                for {
                    s3 := s2[i:] + s2[:i]
                    res = min(res, s3)
                    i = (i + b) % n
                    if i == 0 { break }
                }
            }
        }
    }
    return res
}

func findLexSmallestString1(s string, a int, b int) string {
    res, n := s, len(s)
    s = s + s
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    add := func(t []byte, start int) {
        mn, times := 10, 0
        original := int(t[start] - '0')
        for i := 0; i < 10; i++ {
            added := (original + i*a) % 10
            if added < mn {
                mn, times = added, i
            }
        }
        for i := start; i < n; i += 2 {
            t[i] = byte('0' + (int(t[i] - '0') + times * a) % 10)
        }
    }
    g := gcd(b, n)
    for i := 0; i < n; i += g {
        t := []byte(s[i : i+n])
        add(t, 1)
        if b % 2 != 0 {
            add(t, 0)
        }
        tmp := string(t)
        if tmp < res {
            res = tmp
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "5525", a = 9, b = 2
    // Output: "2050"
    // Explanation: We can apply the following operations:
    // Start:  "5525"
    // Rotate: "2555"
    // Add:    "2454"
    // Add:    "2353"
    // Rotate: "5323"
    // Add:    "5222"
    // Add:    "5121"
    // Rotate: "2151"
    // Add:    "2050"​​​​​
    // There is no way to obtain a string that is lexicographically smaller than "2050".
    fmt.Println(findLexSmallestString("5525", 9, 2)) // 2050
    // Example 2:
    // Input: s = "74", a = 5, b = 1
    // Output: "24"
    // Explanation: We can apply the following operations:
    // Start:  "74"
    // Rotate: "47"
    // ​​​​​​​Add:    "42"
    // ​​​​​​​Rotate: "24"​​​​​​​​​​​​
    // There is no way to obtain a string that is lexicographically smaller than "24".
    fmt.Println(findLexSmallestString("74", 5, 1)) // 24
    // Example 3:
    // Input: s = "0011", a = 4, b = 2
    // Output: "0011"
    // Explanation: There are no sequence of operations that will give us a lexicographically smaller string than "0011".
    fmt.Println(findLexSmallestString("0011", 4, 2)) // 0011

    fmt.Println(findLexSmallestString("1024", 4, 2)) // 1024
    fmt.Println(findLexSmallestString("2048", 4, 2)) // 2048
    fmt.Println(findLexSmallestString("000000000", 4, 2)) // 000000000
    fmt.Println(findLexSmallestString("111111111", 4, 2)) // 111111111
    fmt.Println(findLexSmallestString("123456789", 4, 2)) // 103254769
    fmt.Println(findLexSmallestString("987654321", 4, 2)) // 118967452
    fmt.Println(findLexSmallestString("999999999", 4, 2)) // 919191919
    fmt.Println(findLexSmallestString("1000000007", 4, 2)) // 0000000710

    fmt.Println(findLexSmallestString1("5525", 9, 2)) // 2050
    fmt.Println(findLexSmallestString1("74", 5, 1)) // 24
    fmt.Println(findLexSmallestString1("0011", 4, 2)) // 0011
    fmt.Println(findLexSmallestString1("1024", 4, 2)) // 1024
    fmt.Println(findLexSmallestString1("2048", 4, 2)) // 2048
    fmt.Println(findLexSmallestString1("000000000", 4, 2)) // 000000000
    fmt.Println(findLexSmallestString1("111111111", 4, 2)) // 111111111
    fmt.Println(findLexSmallestString1("123456789", 4, 2)) // 103254769
    fmt.Println(findLexSmallestString1("987654321", 4, 2)) // 118967452
    fmt.Println(findLexSmallestString1("999999999", 4, 2)) // 919191919
    fmt.Println(findLexSmallestString1("1000000007", 4, 2)) // 0000000710
}