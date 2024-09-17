package main

// 880. Decoded String at Index
// You are given an encoded string s. 
// To decode the string to a tape, the encoded string is read one character at a time and the following steps are taken:
//     If the character read is a letter, that letter is written onto the tape.
//     If the character read is a digit d, the entire current tape is repeatedly written d - 1 more times in total.

// Given an integer k, return the kth letter (1-indexed) in the decoded string.

// Example 1:
// Input: s = "leet2code3", k = 10
// Output: "o"
// Explanation: The decoded string is "leetleetcodeleetleetcodeleetleetcode".
// The 10th letter in the string is "o".

// Example 2:
// Input: s = "ha22", k = 5
// Output: "h"
// Explanation: The decoded string is "hahahaha".
// The 5th letter is "h".

// Example 3:
// Input: s = "a2345678999999999999999", k = 1
// Output: "a"
// Explanation: The decoded string is "a" repeated 8301530446056247680 times.
// The 1st letter is "a".

// Constraints:
//     2 <= s.length <= 100
//     s consists of lowercase English letters and digits 2 through 9.
//     s starts with a letter.
//     1 <= k <= 10^9
//     It is guaranteed that k is less than or equal to the length of the decoded string.
//     The decoded string is guaranteed to have less than 2^63 letters.

import "fmt"

func decodeAtIndex(s string, k int) string {
    if k == 1 { return string(s[0]) }
    isLetter := func (char byte) bool { return char >= 'a' && char <= 'z' }
    n := 0
    for i := 0; i < len(s); i++ {
        if isLetter(s[i]) {
            n++
            if n == k { return string(s[i]) }
        } else { 
            if n * int(s[i]-'0') >= k { 
                if k % n != 0 {
                    return decodeAtIndex(s[:i], k % n) 
                } else {
                    return decodeAtIndex(s[:i], n)
                }
            }
            n *= int(s[i]-'0') 
        }
    }
    return ""
}

func main() {
    // Example 1:
    // Input: s = "leet2code3", k = 10
    // Output: "o"
    // Explanation: The decoded string is "leetleetcodeleetleetcodeleetleetcode".
    // The 10th letter in the string is "o".
    fmt.Println(decodeAtIndex("leet2code3", 10)) // "o" "leetleetc[o]deleetleetcodeleetleetcode".
    // Example 2:
    // Input: s = "ha22", k = 5
    // Output: "h"
    // Explanation: The decoded string is "hahahaha".
    // The 5th letter is "h".
    fmt.Println(decodeAtIndex("ha22", 5)) // "h"  "haha[h]aha".
    // Example 3:
    // Input: s = "a2345678999999999999999", k = 1
    // Output: "a"
    // Explanation: The decoded string is "a" repeated 8301530446056247680 times.
    // The 1st letter is "a".
    fmt.Println(decodeAtIndex("a2345678999999999999999", 5)) // "a"  "a2345678999999999999999".
}