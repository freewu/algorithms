package main

// 3561. Resulting String After Adjacent Removals
// You are given a string s consisting of lowercase English letters.

// You must repeatedly perform the following operation while the string s has at least two consecutive characters:
//     1. Remove the leftmost pair of adjacent characters in the string that are consecutive in the alphabet, in either order (e.g., 'a' and 'b', or 'b' and 'a').
//     2. Shift the remaining characters to the left to fill the gap.

// Return the resulting string after no more operations can be performed.

// Note: Consider the alphabet as circular, thus 'a' and 'z' are consecutive.

// Example 1:
// Input: s = "abc"
// Output: "c"
// Explanation:
// Remove "ab" from the string, leaving "c" as the remaining string.
// No further operations are possible. Thus, the resulting string after all possible removals is "c".

// Example 2:
// Input: s = "adcb"
// Output: ""
// Explanation:
// Remove "dc" from the string, leaving "ab" as the remaining string.
// Remove "ab" from the string, leaving "" as the remaining string.
// No further operations are possible. Thus, the resulting string after all possible removals is "".

// Example 3:
// Input: s = "zadb"
// Output: "db"
// Explanation:
// Remove "za" from the string, leaving "db" as the remaining string.
// No further operations are possible. Thus, the resulting string after all possible removals is "db".

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters.

import "fmt"

func resultingString(s string) string {
    stack := []byte{}
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    isConsecutive := func(a byte, b byte) bool{
        return ( a == 'z' && b == 'a' ) || ( a == 'a' && b == 'z' ) || abs(int(a) - int(b)) == 1
    }
    for i := 0; i < len(s); i++ {
        if len(stack) > 0 && isConsecutive(stack[len(stack) - 1],s[i]) {
            stack = stack[:len(stack) - 1]
        }else{
            stack = append(stack, s[i])
        }
    }
    return string(stack)
}

func resultingString1(s string) string {
    arr := make([]byte, 0, len(s))
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(s); i++ {
        if len(arr) > 0 {
            if (abs(int(s[i]) - int(arr[len(arr) - 1])) == 1) || 
               (s[i] == 'a' && arr[len(arr) - 1] == 'z') || 
               (s[i] == 'z' && arr[len(arr) - 1] == 'a') {
                arr = arr[:len(arr) - 1]
            } else {
                arr = append(arr, s[i])
            }
        } else {
            arr = append(arr, s[i])
        }
    }
    return string(arr)
}

func main() {
    // Example 1:
    // Input: s = "abc"
    // Output: "c"
    // Explanation:
    // Remove "ab" from the string, leaving "c" as the remaining string.
    // No further operations are possible. Thus, the resulting string after all possible removals is "c".
    fmt.Println(resultingString("abc")) // "c"
    // Example 2:
    // Input: s = "adcb"
    // Output: ""
    // Explanation:
    // Remove "dc" from the string, leaving "ab" as the remaining string.
    // Remove "ab" from the string, leaving "" as the remaining string.
    // No further operations are possible. Thus, the resulting string after all possible removals is "".
    fmt.Println(resultingString("adcb")) // ""
    // Example 3:
    // Input: s = "zadb"
    // Output: "db"
    // Explanation:
    // Remove "za" from the string, leaving "db" as the remaining string.
    // No further operations are possible. Thus, the resulting string after all possible removals is "db".
    fmt.Println(resultingString("zadb")) // "db"

    fmt.Println(resultingString("bluefrog")) // "blurog"
    fmt.Println(resultingString("leetcode")) // "leetco"

    fmt.Println(resultingString1("abc")) // "c"
    fmt.Println(resultingString1("adcb")) // ""
    fmt.Println(resultingString1("zadb")) // "db"
    fmt.Println(resultingString1("bluefrog")) // "blurog"
    fmt.Println(resultingString1("leetcode")) // "leetco"
}