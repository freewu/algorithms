package main

// 917. Reverse Only Letters
// Given a string s, reverse the string according to the following rules:
//     All the characters that are not English letters remain in the same position.
//     All the English letters (lowercase or uppercase) should be reversed.

// Return s after reversing it.

// Example 1:
// Input: s = "ab-cd"
// Output: "dc-ba"

// Example 2:
// Input: s = "a-bC-dEf-ghIj"
// Output: "j-Ih-gfE-dCba"

// Example 3:
// Input: s = "Test1ng-Leet=code-Q!"
// Output: "Qedo1ct-eeLg=ntse-T!"

// Constraints:
//     1 <= s.length <= 100
//     s consists of characters with ASCII values in the range [33, 122].
//     s does not contain '\"' or '\\'.

import "fmt"

func reverseOnlyLetters(s string) string {
    isLetter := func (c byte) bool {
        return (c >= 'a' && c <= 'z') || c >= 'A' && c <= 'Z'
    }
    arr := []byte(s)
    i, j := 0, len(arr) - 1
    for i < j {
        if !isLetter(arr[i]) {
            i++
        } else if !isLetter(arr[j]) {
            j--
        } else {
            arr[i], arr[j] = arr[j], arr[i]
            i++
            j--
        }
    }
    return string(arr)
}

func main() {
    // Example 1:
    // Input: s = "ab-cd"
    // Output: "dc-ba"
    fmt.Println(reverseOnlyLetters("ab-cd")) // "dc-ba"
    // Example 2:
    // Input: s = "a-bC-dEf-ghIj"
    // Output: "j-Ih-gfE-dCba"
    fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj")) // "j-Ih-gfE-dCba"
    // Example 3:
    // Input: s = "Test1ng-Leet=code-Q!"
    // Output: "Qedo1ct-eeLg=ntse-T!"
    fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) // "Qedo1ct-eeLg=ntse-T!"
}