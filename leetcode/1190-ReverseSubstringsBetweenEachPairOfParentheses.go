package main

// 1190. Reverse Substrings Between Each Pair of Parentheses
// You are given a string s that consists of lower case English letters and brackets.
// Reverse the strings in each pair of matching parentheses, starting from the innermost one.
// Your result should not contain any brackets.

// Example 1:
// Input: s = "(abcd)"
// Output: "dcba"

// Example 2:
// Input: s = "(u(love)i)"
// Output: "iloveu"
// Explanation: The substring "love" is reversed first, then the whole string is reversed.

// Example 3:
// Input: s = "(ed(et(oc))el)"
// Output: "leetcode"
// Explanation: First, we reverse the substring "oc", then "etco", and finally, the whole string.

// Constraints:
//     1 <= s.length <= 2000
//     s only contains lower case English characters and parentheses.
//     It is guaranteed that all parentheses are balanced.

import "fmt"


func reverseParentheses(s string) string {
    reverseString := func (s string) string {
        res := ""
        for _, char := range s[1:] { res = string(char) + res; }
        return res[1:]
    }
    recentL, pointer := 0, 0
    for pointer < len(s) {
        if s[pointer] == ')' {
            s = s[:recentL] + reverseString(s[recentL:pointer + 1]) + s[pointer + 1:]
            recentL, pointer = 0, 0
        }
        if s[pointer] == '(' { // 记录要翻转的开始位置
            recentL = pointer
        }
        pointer++
    } 
    return s
}

// stack
func reverseParentheses1(s string) string {
    res, n := []byte{}, len(s)
    pair, stack := make([]int, n), []int{}
    for i, c := range s {
        if c == '(' {
            stack = append(stack, i)
        } else if c == ')' {
            j := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            pair[i], pair[j] = j, i
        }
    }
    for i, step := 0, 1; i < n; i += step {
        if s[i] == '(' || s[i] == ')' {
            i = pair[i]
            step = -step
        } else {
            res = append(res, s[i])
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "(abcd)"
    // Output: "dcba"
    fmt.Println(reverseParentheses("(abcd)")) // "dcba"
    // Example 2:
    // Input: s = "(u(love)i)"
    // Output: "iloveu"
    // Explanation: The substring "love" is reversed first, then the whole string is reversed.
    fmt.Println(reverseParentheses("(u(love)i)")) // "iloveu"
    // Example 3:
    // Input: s = "(ed(et(oc))el)"
    // Output: "leetcode"
    // Explanation: First, we reverse the substring "oc", then "etco", and finally, the whole string.
    fmt.Println(reverseParentheses("(ed(et(oc))el)")) // "leetcode"

    fmt.Println(reverseParentheses1("(abcd)")) // "dcba"
    fmt.Println(reverseParentheses1("(u(love)i)")) // "iloveu"
    fmt.Println(reverseParentheses1("(ed(et(oc))el)")) // "leetcode"
}