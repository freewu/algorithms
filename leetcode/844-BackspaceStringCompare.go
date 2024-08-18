package main

// 844. Backspace String Compare
// Given two strings s and t, return true if they are equal when both are typed into empty text editors. 
// '#' means a backspace character.

// Note that after backspacing an empty text, the text will continue empty.

// Example 1:
// Input: s = "ab#c", t = "ad#c"
// Output: true
// Explanation: Both s and t become "ac".

// Example 2:
// Input: s = "ab##", t = "c#d#"
// Output: true
// Explanation: Both s and t become "".

// Example 3:
// Input: s = "a#c", t = "b"
// Output: false
// Explanation: s becomes "c" while t becomes "b".

// Constraints:
//     1 <= s.length, t.length <= 200
//     s and t only contain lowercase letters and '#' characters.

// Follow up: Can you solve it in O(n) time and O(1) space?

import "fmt"

func backspaceCompare(s string, t string) bool {
    format := func(str string) string {
        res := []byte{}
        for i := 0; i < len(str); i++ {
            if str[i] == '#' { // 删除前一个字符
                if len(res) > 0 {
                    res = res[:len(res) - 1] // pop
                }
            } else {
                res = append(res, str[i])
            }
        }
        // fmt.Println(string(res))
        return string(res)
    } 
    return format(s) == format(t)
}

func backspaceCompare1(s string, t string) bool {
    skips, skipt,i, j := 0, 0, len(s) - 1, len(t) - 1
    for i >= 0 || j >= 0 {
        for i >= 0 {
            if s[i] == '#' {
                skips++
                i--
            } else if skips > 0 {
                skips--
                i--
            } else {
                break
            }
        }
        for j >= 0 {
            if t[j] == '#' {
                skipt++
                j--
            } else if skipt > 0 {
                skipt--
                j--
            } else {
                break
            }
        }
        if i < 0 && j < 0 {
            return true
        } else if i >= 0 && j >= 0 {
            if s[i] != t[j] {
                return false
            } else {
                i--
                j--
            }
        } else {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "ab#c", t = "ad#c"
    // Output: true
    // Explanation: Both s and t become "ac".
    fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
    // Example 2:
    // Input: s = "ab##", t = "c#d#"
    // Output: true
    // Explanation: Both s and t become "".
    fmt.Println(backspaceCompare("ab##", "c#d#")) // true
    // Example 3:
    // Input: s = "a#c", t = "b"
    // Output: false
    // Explanation: s becomes "c" while t becomes "b".
    fmt.Println(backspaceCompare("a#c", "b")) // false

    fmt.Println(backspaceCompare("y#fo##f", "y#f#o##f")) // true

    fmt.Println(backspaceCompare1("ab#c", "ad#c")) // true
    fmt.Println(backspaceCompare1("ab##", "c#d#")) // true
    fmt.Println(backspaceCompare1("a#c", "b")) // false
    fmt.Println(backspaceCompare1("y#fo##f", "y#f#o##f")) // true
}