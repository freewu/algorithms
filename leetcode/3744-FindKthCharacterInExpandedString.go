package main

// 3744. Find Kth Character in Expanded String
// You are given a string s consisting of one or more words separated by single spaces. 
// Each word in s consists of lowercase English letters.

// We obtain the expanded string t from s as follows:
//     For each word in s, repeat its first character once, then its second character twice, and so on.

// For example, if s = "hello world", then t = "heelllllllooooo woorrrllllddddd".

// You are also given an integer k, representing a valid index of the string t.

// Return the kth character of the string t.

// Example 1:
// Input: s = "hello world", k = 0
// Output: "h"
// Explanation:
// t = "heelllllllooooo woorrrllllddddd". Therefore, the answer is t[0] = "h".

// Example 2:
// Input: s = "hello world", k = 15
// Output: " "
// Explanation:
// t = "heelllllllooooo woorrrllllddddd". Therefore, the answer is t[15] = " ".

// Constraints:
//     1 <= s.length <= 10^5
//     s contains only lowercase English letters and spaces ' '.
//     s does not contain any leading or trailing spaces.
//     All the words in s are separated by a single space.
//     0 <= k < t.length. That is, k is a valid index of t.

import "fmt"
import "strings"

// brute force 超出时间限制 922 / 1000 个通过的测试用例
func kthCharacter(s string, k int64) byte {
    str, i := "", int64(1)
    for _, c := range s {
        if c == ' ' { // 如果是 ' ' 从头开始
            i = 1
            str += string(c)
            continue
        }
        str += strings.Repeat(string(c), int(i))
        if len(str) > int(k) { // 满足长度了
            return str[k]
        }
        i++
    }
    return s[k]
}

func kthCharacter1(s string, k int64) byte {
    curr, index := 0, 0
    for _, c := range s {
        expand := 0
        if c == ' ' { // 如果是 ' ' 从头开始
            expand, index = 1, 0
        } else {
            expand = index + 1
            index++
        }
        // 检查k是否在当前字符的展开范围内 [curr, curr + expand - 1]
        if int64(curr + expand) > k {
            return byte(c)
        }
        curr += expand; // 累加长度，继续下一个字符
    }
    return s[k]
}

func main() {
    // Example 1:
    // Input: s = "hello world", k = 0
    // Output: "h"
    // Explanation:
    // t = "heelllllllooooo woorrrllllddddd". Therefore, the answer is t[0] = "h".
    fmt.Printf("kthCharacter(\"hello world\", 0) = %c\n", kthCharacter("hello world", 0)) // "h"
    // Example 2:
    // Input: s = "hello world", k = 15
    // Output: " "
    // Explanation:
    // t = "heelllllllooooo woorrrllllddddd". Therefore, the answer is t[15] = " ".
    fmt.Printf("kthCharacter(\"hello world\", 15) = %c\n", kthCharacter("hello world", 15)) // " "

    fmt.Printf("kthCharacter(\"leetcode\", 2) = %c\n", kthCharacter("leetcode", 2)) // "e"
    fmt.Printf("kthCharacter(\"bluefrog\", 2) = %c\n", kthCharacter("bluefrog", 2)) // "l"
    fmt.Printf("kthCharacter(\"a c\", 2) = %c\n", kthCharacter("a c", 2)) // "c"
    fmt.Printf("kthCharacter(\"a\", 0) = %c\n", kthCharacter("a", 0)) // "a"
    fmt.Printf("kthCharacter(\"ab\", 1) = %c\n", kthCharacter("ab", 1)) // "b"
    fmt.Printf("kthCharacter(\"ab\", 2) = %c\n", kthCharacter("ab", 2)) // "a"


    fmt.Printf("kthCharacter1(\"hello world\", 0) = %c\n", kthCharacter1("hello world", 0)) // "h"
    fmt.Printf("kthCharacter1(\"hello world\", 15) = %c\n", kthCharacter1("hello world", 15)) // " "
    fmt.Printf("kthCharacter1(\"leetcode\", 2) = %c\n", kthCharacter1("leetcode", 2)) // "e"
    fmt.Printf("kthCharacter1(\"bluefrog\", 2) = %c\n", kthCharacter1("bluefrog", 2)) // "l"
    fmt.Printf("kthCharacter1(\"a c\", 2) = %c\n", kthCharacter1("a c", 2)) // "c"
    fmt.Printf("kthCharacter1(\"a\", 0) = %c\n", kthCharacter1("a", 0)) // "a"
    fmt.Printf("kthCharacter1(\"ab\", 1) = %c\n", kthCharacter1("ab", 1)) // "b"
    fmt.Printf("kthCharacter1(\"ab\", 2) = %c\n", kthCharacter1("ab", 2)) // "a"
}