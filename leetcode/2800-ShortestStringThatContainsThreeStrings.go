package main

// 2800. Shortest String That Contains Three Strings
// Given three strings a, b, and c, your task is to find a string that has the minimum length 
// and contains all three strings as substrings.

// If there are multiple such strings, return the lexicographically smallest one.

// Return a string denoting the answer to the problem.

// Notes
//     A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ, string a has a letter that appears earlier in the alphabet than the corresponding letter in b.
//     A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: a = "abc", b = "bca", c = "aaa"
// Output: "aaabca"
// Explanation:  We show that "aaabca" contains all the given strings: a = ans[2...4], b = ans[3..5], c = ans[0..2]. It can be shown that the length of the resulting string would be at least 6 and "aaabca" is the lexicographically smallest one.

// Example 2:
// Input: a = "ab", b = "ba", c = "aba"
// Output: "aba"
// Explanation: We show that the string "aba" contains all the given strings: a = ans[0..1], b = ans[1..2], c = ans[0..2]. Since the length of c is 3, the length of the resulting string would be at least 3. It can be shown that "aba" is the lexicographically smallest one.

// Constraints:
//     1 <= a.length, b.length, c.length <= 100
//     a, b, c consist only of lowercase English letters.

import "fmt"
import "strings"
import "sort"

func minimumString(a, b, c string) string {
    merge := func(s, t string) string {
        // 先特判完全包含的情况
        if strings.Contains(s, t) { return s }
        if strings.Contains(t, s) { return t }
        for i := min(len(s), len(t)); ; i-- {
            if s[len(s)-i:] == t[:i] { // 枚举：s 的后 i 个字母和 t 的前 i 个字母是一样的
                return s + t[i:]
            }
        }
    }
    res, arr := "", []string{ a, b, c }
    // 枚举 arr 的全排列
    for _, p := range [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}} {
        str := merge(merge(arr[p[0]], arr[p[1]]), arr[p[2]])
        if res == "" || len(str) < len(res) || len(str) == len(res) && str < res {
            res = str
        }
    }
    return res
}

func minimumString1(a string, b string, c string) string {
    res, mn := []string{}, 1 << 31
    arr := [][]string{{a, b, c}, {a, c, b}, {b, a, c}, {b, c, a}, {c, a, b}, {c, b, a}}
    merge := func(a string, b string) string {
        if strings.Contains(a, b) { return a }
        for i:= len(b); i > 0; i-- {
            if strings.HasSuffix(a, b[:i]) { return a + b[i:] }
        }
        return a + b
    }
    for _, v := range arr {
        str := merge(merge(v[0], v[1]), v[2])
        if len(str) < mn {
            mn, res = len(str), []string{ str }
        } else if len(str) == mn {
            res = append(res, str)
        }
    }
    sort.Strings(res) // sort the string based on lexographical order
    return res[0]
}

func minimumString2(a string, b string, c string) string {
    res, arr := "", []string{ a, b, c }
    sort.Strings(arr)
    containsRight := func(a, b string) bool {
        if len(a) - len(b) < 0 { return false }
        return a[len(a) - len(b):] == b
    }
    merge := func(a, b, c string) string {
        last := 0
        for i := range b {
            if containsRight(a, b[:i + 1]) { last = i + 1 }
        }
        if !strings.Contains(a, b) { a += b[last:] }
        last = 0
        for i := range c {
            if containsRight(a, c[:i + 1]) { last = i + 1 }
        }
        if !strings.Contains(a, c) { a += c[last:] }
        return a
    }
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr); j++ {
            for k := 0; k < len(arr); k++ {
                if i != j && j != k && i != k {
                    merged := merge(arr[i], arr[j], arr[k])
                    if len(merged) < len(res) || len(merged) == len(res) && strings.Compare(res, merged) == 1 || len(res) == 0 {
                        res = merged
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: a = "abc", b = "bca", c = "aaa"
    // Output: "aaabca"
    // Explanation:  We show that "aaabca" contains all the given strings: a = ans[2...4], b = ans[3..5], c = ans[0..2]. It can be shown that the length of the resulting string would be at least 6 and "aaabca" is the lexicographically smallest one.
    fmt.Println(minimumString("abc", "bca", "aaa")) // aaabca
    // Example 2:
    // Input: a = "ab", b = "ba", c = "aba"
    // Output: "aba"
    // Explanation: We show that the string "aba" contains all the given strings: a = ans[0..1], b = ans[1..2], c = ans[0..2]. Since the length of c is 3, the length of the resulting string would be at least 3. It can be shown that "aba" is the lexicographically smallest one.
    fmt.Println(minimumString("ab", "ba", "aba")) // aba

    fmt.Println(minimumString("ba", "b", "a")) // ba
    fmt.Println(minimumString("bluefrog", "leetcode", "abcdefghijk")) // abcdefghijkbluefrogleetcode

    fmt.Println(minimumString1("abc", "bca", "aaa")) // aaabca
    fmt.Println(minimumString1("ab", "ba", "aba")) // aba
    fmt.Println(minimumString1("ba", "b", "a")) // ba
    fmt.Println(minimumString1("bluefrog", "leetcode", "abcdefghijk")) // abcdefghijkbluefrogleetcode

    fmt.Println(minimumString2("abc", "bca", "aaa")) // aaabca
    fmt.Println(minimumString2("ab", "ba", "aba")) // aba
    fmt.Println(minimumString2("ba", "b", "a")) // ba
    fmt.Println(minimumString2("bluefrog", "leetcode", "abcdefghijk")) // abcdefghijkbluefrogleetcode
}