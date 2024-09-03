package main

// 925. Long Pressed Name
// Your friend is typing his name into a keyboard. 
// Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

// You examine the typed characters of the keyboard. 
// Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.

// Example 1:
// Input: name = "alex", typed = "aaleex"
// Output: true
// Explanation: 'a' and 'e' in 'alex' were long pressed.

// Example 2:
// Input: name = "saeed", typed = "ssaaedd"
// Output: false
// Explanation: 'e' must have been pressed twice, but it was not in the typed output.

// Constraints:
//     1 <= name.length, typed.length <= 1000
//     name and typed consist of only lowercase English letters.

import "fmt"

// 双指针
func isLongPressedName(name string, typed string) bool {
    i, j, n, m  := 0, 0, len(name), len(typed)
    for i < n && j < m {
        ch := name[i]
        if ch != typed[j] {
            return false
        }
        c1 := 0
        for i < n && name[i] == ch {
            i++
            c1++
        }
        c2 := 0
        for j < m && typed[j] == ch {
            j++
            c2++
        }
        if c1 > c2 {
            return false
        }
    }
    return j == m && i == n
}

func main() {
    // Example 1:
    // Input: name = "alex", typed = "aaleex"
    // Output: true
    // Explanation: 'a' and 'e' in 'alex' were long pressed.
    fmt.Println(isLongPressedName("alex", "aaleex")) // true
    // Example 2:
    // Input: name = "saeed", typed = "ssaaedd"
    // Output: false
    // Explanation: 'e' must have been pressed twice, but it was not in the typed output.
    fmt.Println(isLongPressedName("saeed", "ssaaedd")) // true
}