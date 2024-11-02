package main

// 2083. Substrings That Begin and End With the Same Letter
// You are given a 0-indexed string s consisting of only lowercase English letters. 
// Return the number of substrings in s that begin and end with the same character.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "abcba"
// Output: 7
// Explanation:
// The substrings of length 1 that start and end with the same letter are: "a", "b", "c", "b", and "a".
// The substring of length 3 that starts and ends with the same letter is: "bcb".
// The substring of length 5 that starts and ends with the same letter is: "abcba".

// Example 2:
// Input: s = "abacad"
// Output: 9
// Explanation:
// The substrings of length 1 that start and end with the same letter are: "a", "b", "a", "c", "a", and "d".
// The substrings of length 3 that start and end with the same letter are: "aba" and "aca".
// The substring of length 5 that starts and ends with the same letter is: "abaca".

// Example 3:
// Input: s = "a"
// Output: 1
// Explanation:
// The substring of length 1 that starts and ends with the same letter is: "a".

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters.

import "fmt"

// 统计字符串中相同字母的数量。则相同字母中任意两个可作为某个子串的首尾字母,
// 另外，还需计算只含一个字母的子串。
func numberOfSubstrings(s string) int64 {
    count := make([]int, 26)
    for _, v := range s {
        count[v-'a']++
    }
    sum, n := 0, len(s)
    for _, v := range count {
        sum += (v * (v - 1) >> 1)
    }
    return int64(sum + n)
}

func main() {
    // Example 1:
    // Input: s = "abcba"
    // Output: 7
    // Explanation:
    // The substrings of length 1 that start and end with the same letter are: "a", "b", "c", "b", and "a".
    // The substring of length 3 that starts and ends with the same letter is: "bcb".
    // The substring of length 5 that starts and ends with the same letter is: "abcba".
    fmt.Println(numberOfSubstrings("abcba")) // 7
    // Example 2:
    // Input: s = "abacad"
    // Output: 9
    // Explanation:
    // The substrings of length 1 that start and end with the same letter are: "a", "b", "a", "c", "a", and "d".
    // The substrings of length 3 that start and end with the same letter are: "aba" and "aca".
    // The substring of length 5 that starts and ends with the same letter is: "abaca".
    fmt.Println(numberOfSubstrings("abacad")) // 9
    // Example 3:
    // Input: s = "a"
    // Output: 1
    // Explanation:
    // The substring of length 1 that starts and ends with the same letter is: "a".
    fmt.Println(numberOfSubstrings("a")) // 1
}