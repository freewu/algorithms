package main 

// 1717. Maximum Score From Removing Substrings
// You are given a string s and two integers x and y. 
// You can perform two types of operations any number of times.
//     Remove substring "ab" and gain x points.
//         For example, when removing "ab" from "cabxbae" it becomes "cxbae".
//     Remove substring "ba" and gain y points.
//         For example, when removing "ba" from "cabxbae" it becomes "cabxe".

// Return the maximum points you can gain after applying the above operations on s.

// Example 1:
// Input: s = "cdbcbbaaabab", x = 4, y = 5
// Output: 19
// Explanation:
// - Remove the "ba" underlined in "cdbcbbaaa{ba}b". Now, s = "cdbcbbaaab" and 5 points are added to the score.
// - Remove the "ab" underlined in "cdbcbbaa{ab}". Now, s = "cdbcbbaa" and 4 points are added to the score.
// - Remove the "ba" underlined in "cdbcb{ba}a". Now, s = "cdbcba" and 5 points are added to the score.
// - Remove the "ba" underlined in "cdbc{ba}". Now, s = "cdbc" and 5 points are added to the score.
// Total score = 5 + 4 + 5 + 5 = 19.

// Example 2:
// Input: s = "aabbaaxybbaabb", x = 5, y = 4
// Output: 20

// Constraints:
//     1 <= s.length <= 10^5
//     1 <= x, y <= 10^4
//     s consists of lowercase English letters.

import "fmt"

func maximumGain(s string, x int, y int) int {
    b0, b1 := byte('a'), byte('b')
    if x < y {
        b0, b1 = b1, b0
        x, y = y, x
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res, cnt0, cnt1 := 0, 0, 0
    for i := range s {
        if s[i] == b0 { // 优先处理分高的
            cnt0++
            continue
        }
        if s[i] != b1 {
            res += y * min(cnt0, cnt1)
            cnt0, cnt1 = 0, 0
            continue
        }
        if cnt0 > 0 {
            cnt0--
            res += x
        } else {
            cnt1++
        }
    }
    res += y * min(cnt0, cnt1)
    return res
}

func main() {
    // Example 1:
    // Input: s = "cdbcbbaaabab", x = 4, y = 5
    // Output: 19
    // Explanation:
    // - Remove the "ba" underlined in "cdbcbbaaabab". Now, s = "cdbcbbaaab" and 5 points are added to the score.
    // - Remove the "ab" underlined in "cdbcbbaaab". Now, s = "cdbcbbaa" and 4 points are added to the score.
    // - Remove the "ba" underlined in "cdbcbbaa". Now, s = "cdbcba" and 5 points are added to the score.
    // - Remove the "ba" underlined in "cdbcba". Now, s = "cdbc" and 5 points are added to the score.
    // Total score = 5 + 4 + 5 + 5 = 19.
    fmt.Println(maximumGain("cdbcbbaaabab", 4, 5)) // 19
    // Example 2:
    // Input: s = "aabbaaxybbaabb", x = 5, y = 4
    // Output: 20
    fmt.Println(maximumGain("aabbaaxybbaabb", 5, 4)) // 20
}