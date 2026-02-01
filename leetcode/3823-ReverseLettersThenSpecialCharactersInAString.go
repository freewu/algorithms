package main

// 3823. Reverse Letters Then Special Characters in a String
// You are given a string s consisting of lowercase English letters and special characters.

// Your task is to perform these in order:
//     1. Reverse the lowercase letters and place them back into the positions originally occupied by letters.
//     2. Reverse the special characters and place them back into the positions originally occupied by special characters.

// Return the resulting string after performing the reversals.

// Example 1:
// Input: s = ")ebc#da@f("
// Output: "(fad@cb#e)"
// Explanation:
// The letters in the string are ['e', 'b', 'c', 'd', 'a', 'f']:
// Reversing them gives ['f', 'a', 'd', 'c', 'b', 'e']
// s becomes ")fad#cb@e("
// ​​​​​​​The special characters in the string are [')', '#', '@', '(']:
// Reversing them gives ['(', '@', '#', ')']
// s becomes "(fad@cb#e)"

// Example 2:
// Input: s = "z"
// Output: "z"
// Explanation:
// The string contains only one letter, and reversing it does not change the string. There are no special characters.

// Example 3:
// Input: s = "!@#$%^&*()"
// Output: ")(*&^%$#@!"
// Explanation:
// The string contains no letters. The string contains all special characters, so reversing the special characters reverses the whole string.

// Constraints:
//     1 <= s.length <= 100
//     s consists only of lowercase English letters and the special characters in "!@#$%^&*()".

import "fmt"

// 双指针
func reverseByType(s string) string {
    res := []byte(s)
    reverse := func(t []byte, f func(byte) bool) {
        i, j := 0, len(t)-1
        for i < j {
            for i < j && f(t[i]) {
                i++
            }
            for i < j && f(t[j]) {
                j--
            }
            t[i], t[j] = t[j], t[i]
            i++
            j--
        }
    }
    reverse(res, func(b byte) bool { return 'a' <= b && b <= 'z' })
    reverse(res, func(b byte) bool { return !('a' <= b && b <= 'z') })
    return string(res)
}

func reverseByType1(s string) string {
    n := len(s)
    res := make([]byte, n)
    ai, si := n-1, n-1
    for ai >= 0 && s[ai] < 97 {
        ai--
    }
    for si >= 0 && s[si] >= 97 {
        si--
    }
    for i := range s {
        if s[i] >= 97 {
            res[ai] = s[i]
            ai--
            for ai >= 0 && s[ai] < 97 {
                ai--
            }
        } else {
            res[si] = s[i]
            si--
            for si >= 0 && s[si] >= 97 {
                si--
            }
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = ")ebc#da@f("
    // Output: "(fad@cb#e)"
    // Explanation:
    // The letters in the string are ['e', 'b', 'c', 'd', 'a', 'f']:
    // Reversing them gives ['f', 'a', 'd', 'c', 'b', 'e']
    // s becomes ")fad#cb@e("
    // ​​​​​​​The special characters in the string are [')', '#', '@', '(']:
    // Reversing them gives ['(', '@', '#', ')']
    // s becomes "(fad@cb#e)"
    fmt.Println(reverseByType(")ebc#da@f(")) // "(fad@cb#e)"
    // Example 2:
    // Input: s = "z"
    // Output: "z"
    // Explanation:
    // The string contains only one letter, and reversing it does not change the string. There are no special characters.
    fmt.Println(reverseByType("z")) // "z"
    // Example 3:
    // Input: s = "!@#$%^&*()"
    // Output: ")(*&^%$#@!"
    // Explanation:
    // The string contains no letters. The string contains all special characters, so reversing the special characters reverses the whole string.
    fmt.Println(reverseByType("!@#$%^&*()")) // ")(*&^%$#@!"

    fmt.Println(reverseByType("bluefrog")) // "gofreblu"
    fmt.Println(reverseByType("leetcode")) // "edocteel"

    fmt.Println(reverseByType1(")ebc#da@f(")) // "(fad@cb#e)"
    fmt.Println(reverseByType1("z")) // "z"
    fmt.Println(reverseByType1("!@#$%^&*()")) // ")(*&^%$#@!"
    fmt.Println(reverseByType1("bluefrog")) // "gofreblu"
    fmt.Println(reverseByType1("leetcode")) // "edocteel"
}