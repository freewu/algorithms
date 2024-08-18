package main

// 848. Shifting Letters
// You are given a string s of lowercase English letters and an integer array shifts of the same length.
// Call the shift() of a letter, the next letter in the alphabet, (wrapping around so that 'z' becomes 'a').
//     For example, shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'.

// Now for each shifts[i] = x, we want to shift the first i + 1 letters of s, x times.
// Return the final string after all such shifts to s are applied.

// Example 1:
// Input: s = "abc", shifts = [3,5,9]
// Output: "rpl"
// Explanation: We start with "abc".
// After shifting the first 1 letters of s by 3, we have "dbc".
// After shifting the first 2 letters of s by 5, we have "igc".
// After shifting the first 3 letters of s by 9, we have "rpl", the answer.

// Example 2:
// Input: s = "aaa", shifts = [1,2,3]
// Output: "gfd"

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.
//     shifts.length == s.length
//     0 <= shifts[i] <= 10^9

import "fmt"

func shiftingLetters(s string, shifts []int) string {
    bytes, acc := []byte(s), 0
    for i := len(s) - 1; i >= 0; i-- {
        acc += shifts[i]
        bytes[i] = byte((int(bytes[i]) - 'a' + acc ) % 26 + 'a')
    }
    return string(bytes)
}

func shiftingLetters1(s string, shifts []int) string {
    n, suffix, res := len(shifts), 0, []byte{}
    for i := n - 1; i >= 0; i-- {
        suffix = (suffix + shifts[i]) % 26
        move := (int(s[i]) - 'a' + suffix) % 26
        res = append(res, byte( 'a' + move))
    }
    for i := 0; i < n/2; i++ {
        res[i], res[n-i-1] = res[n-i-1], res[i]
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "abc", shifts = [3,5,9]
    // Output: "rpl"
    // Explanation: We start with "abc".
    // After shifting the first 1 letters of s by 3, we have "dbc".
    // After shifting the first 2 letters of s by 5, we have "igc".
    // After shifting the first 3 letters of s by 9, we have "rpl", the answer.
    fmt.Println(shiftingLetters("abc",[]int{3,5,9})) // "rpl"
    // Example 2:
    // Input: s = "aaa", shifts = [1,2,3]
    // Output: "gfd"
    fmt.Println(shiftingLetters("aaa",[]int{1,2,3})) // "gfd"

    fmt.Println(shiftingLetters1("abc",[]int{3,5,9})) // "rpl"
    fmt.Println(shiftingLetters1("aaa",[]int{1,2,3})) // "gfd"
}