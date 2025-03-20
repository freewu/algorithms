package main

// 3104. Find Longest Self-Contained Substring
// Given a string s, your task is to find the length of the longest self-contained substring of s.
// A substring t of a string s is called self-contained if t != s and for every character in t, it doesn't exist in the rest of s.
// Return the length of the longest self-contained substring of s if it exists, otherwise, return -1.

// Example 1:
// Input: s = "abba"
// Output: 2
// Explanation:
// Let's check the substring "bb". You can see that no other "b" is outside of this substring. Hence the answer is 2.

// Example 2:
// Input: s = "abab"
// Output: -1
// Explanation:
// Every substring we choose does not satisfy the described property (there is some character which is inside and outside of that substring). So the answer would be -1.

// Example 3:
// Input: s = "abacd"
// Output: 4
// Explanation:
// Let's check the substring "abac". There is only one character outside of this substring and that is "d". There is no "d" inside the chosen substring, so it satisfies the condition and the answer is 4.

// Constraints:
//     2 <= s.length <= 5 * 10^4
//     s consists only of lowercase English letters.

import "fmt"

func maxSubstringLength(s string) int {
    res, n := -1, len(s)
    pre, mx := [26][]int{}, [26][]int{}
    for i := 0; i < 26; i++ {
        pre[i] = make([]int, n + 1)
        mx[i] = []int{ -1, -1, 0 } 
    }
    for i := 0; i < n; i++ {
        v := int(s[i] - 'a')
        if mx[v][0] == -1 {
            mx[v][0], mx[v][1] = i, i
        } else {
            mx[v][1] = i
        }
        mx[v][2]++
        for j := 0; j < 26; j++ {
            if v == j {
                pre[j][i + 1] = pre[j][i] + 1
            } else {
                pre[j][i + 1] = pre[j][i]
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < 26; i++ {
        for j := 0; j < 26; j++ {
            l, r := mx[i][0], mx[j][1]
            if l == -1 || r == -1 || l > r { continue }
            flag := true
            for k := 0; k < 26; k++ {
                count := pre[k][r + 1] - pre[k][l]
                if count > 0 && count < mx[k][2] {
                    flag = false
                    break
                }
            }
            if flag && r - l < n - 1 {
                res = max(res, r - l + 1)
            }
        }
    }
    return res
}

func maxSubstringLength1(s string) int {
    locs := [26][]int{}
    for i, c := range s {
        locs[c-'a'] = append(locs[c-'a'], i)
    }
    check := func(locs [26][]int, l int, r int) bool {
        for _, poses := range locs {
            flag := 0
            for _, pos := range poses {
                if pos >= l && pos <= r {
                    if flag == 0 {
                        flag = 1
                    } else if flag == 2 {
                        return false
                    }
                } else {
                    if flag == 0 {
                        flag = 2
                    } else if flag == 1 {
                        return false
                    }
                }
            }
        }
        return true
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := -1
    for _, pos1 := range locs {
        if len(pos1) == 0 { continue }
        left := pos1[0]
        for _, row := range locs {
            if len(row) == 0 { continue }
            right := row[len(row) - 1]
            if left > right || right - left + 1 == len(s) { continue }
            if check(locs, left, right) {
                res = max(res, right - left + 1)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abba"
    // Output: 2
    // Explanation:
    // Let's check the substring "bb". You can see that no other "b" is outside of this substring. Hence the answer is 2.
    fmt.Println(maxSubstringLength("abba")) // 2
    // Example 2:
    // Input: s = "abab"
    // Output: -1
    // Explanation:
    // Every substring we choose does not satisfy the described property (there is some character which is inside and outside of that substring). So the answer would be -1.
    fmt.Println(maxSubstringLength("abab")) // -1
    // Example 3:
    // Input: s = "abacd"
    // Output: 4
    // Explanation:
    // Let's check the substring "abac". There is only one character outside of this substring and that is "d". There is no "d" inside the chosen substring, so it satisfies the condition and the answer is 4.
    fmt.Println(maxSubstringLength("abacd")) // 4

    fmt.Println(maxSubstringLength("bluefrog")) // 7
    fmt.Println(maxSubstringLength("leetcode")) // 7

    fmt.Println(maxSubstringLength1("abba")) // 2
    fmt.Println(maxSubstringLength1("abab")) // -1
    fmt.Println(maxSubstringLength1("abacd")) // 4
    fmt.Println(maxSubstringLength1("bluefrog")) // 7
    fmt.Println(maxSubstringLength1("leetcode")) // 7
}