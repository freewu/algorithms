package main

// 1554. Strings Differ by One Character
// Given a list of strings dict where all the strings are of the same length.
// Return true if there are 2 strings that only differ by 1 character in the same index, otherwise return false.

// Example 1:
// Input: dict = ["abcd","acbd", "aacd"]
// Output: true
// Explanation: Strings "abcd" and "aacd" differ only by one character in the index 1.

// Example 2:
// Input: dict = ["ab","cd","yz"]
// Output: false

// Example 3:
// Input: dict = ["abcd","cccc","abyd","abab"]
// Output: true

// Constraints:
//     The number of characters in dict <= 10^5
//     dict[i].length == dict[j].length
//     dict[i] should be unique.
//     dict[i] contains only lowercase English letters.

// Follow up: Could you solve this problem in O(n * m) where n is the length of dict and m is the length of each string.

import "fmt"

func differByOne(dict []string) bool {
    n := len(dict)
    diffOne := func(s1, s2 string) bool { // 两个字符串只有一个字符不相同
        if len(s1) != len(s2) { return false }
        diff := 0
        for i, size := 0, len(s1); i < size; i++ {
            if s1[i] != s2[i] { diff++ }
        }
        return diff == 1
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if diffOne(dict[i], dict[j]) {
                return true
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: dict = ["abcd","acbd", "aacd"]
    // Output: true
    // Explanation: Strings "abcd" and "aacd" differ only by one character in the index 1.
    fmt.Println(differByOne([]string{"abcd","acbd", "aacd"})) // true
    // Example 2:
    // Input: dict = ["ab","cd","yz"]
    // Output: false
    fmt.Println(differByOne([]string{"ab","cd","yz"})) // false
    // Example 3:
    // Input: dict = ["abcd","cccc","abyd","abab"]
    // Output: true
    fmt.Println(differByOne([]string{"abcd","cccc","abyd","abab"})) // true
}