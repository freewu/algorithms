package main

// 1933. Check if String Is Decomposable Into Value-Equal Substrings
// A value-equal string is a string where all characters are the same.
//     For example, "1111" and "33" are value-equal strings.
//     In contrast, "123" is not a value-equal string.

// Given a digit string s, decompose the string into some number of consecutive value-equal substrings 
// where exactly one substring has a length of 2 and the remaining substrings have a length of 3.

// Return true if you can decompose s according to the above rules. Otherwise, return false.

// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: s = "000111000"
// Output: false
// Explanation: s cannot be decomposed according to the rules because ["000", "111", "000"] does not have a substring of length 2.

// Example 2:
// Input: s = "00011111222"
// Output: true
// Explanation: s can be decomposed into ["000", "111", "11", "222"].

// Example 3:
// Input: s = "011100022233"
// Output: false
// Explanation: s cannot be decomposed according to the rules because of the first '0'.

// Constraints:
//     1 <= s.length <= 1000
//     s consists of only digits '0' through '9'.

import "fmt"

func isDecomposable(s string) bool {
    n, count := len(s), 0
    for i := 0; i< n; {
        j := i + 1
        for ; j < n && s[j] == s[i]; j++ { } // 字符一样就一直累计
        k := j - i
        if k % 3 == 1 { return false } // 如果出现 余1 的 4  比需 每个一个的字符串为3个
        if k % 3 == 2 { count++ } // 可以允许有人 2个一样的字符串 存在
        i = j
    }
    return count == 1
}

func main() {
    // Example 1:
    // Input: s = "000111000"
    // Output: false
    // Explanation: s cannot be decomposed according to the rules because ["000", "111", "000"] does not have a substring of length 2.
    fmt.Println(isDecomposable("000111000")) // false
    // Example 2:
    // Input: s = "00011111222"
    // Output: true
    // Explanation: s can be decomposed into ["000", "111", "11", "222"].
    fmt.Println(isDecomposable("00011111222")) // true
    // Example 3:
    // Input: s = "011100022233"
    // Output: false
    // Explanation: s cannot be decomposed according to the rules because of the first '0'.
    fmt.Println(isDecomposable("011100022233")) // false
}