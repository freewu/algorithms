package main

// 2516. Take K of Each Character From Left and Right
// You are given a string s consisting of the characters 'a', 'b', and 'c' and a non-negative integer k. 
// Each minute, you may take either the leftmost character of s, or the rightmost character of s.

// Return the minimum number of minutes needed for you to take at least k of each character, 
// or return -1 if it is not possible to take k of each character.

// Example 1:
// Input: s = "aabaaaacaabc", k = 2
// Output: 8
// Explanation: 
// Take three characters from the left of s. You now have two 'a' characters, and one 'b' character.
// Take five characters from the right of s. You now have four 'a' characters, two 'b' characters, and two 'c' characters.
// A total of 3 + 5 = 8 minutes is needed.
// It can be proven that 8 is the minimum number of minutes needed.

// Example 2:
// Input: s = "a", k = 1
// Output: -1
// Explanation: It is not possible to take one 'b' or 'c' so return -1.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of only the letters 'a', 'b', and 'c'.
//     0 <= k <= s.length

import "fmt"

// sliding window
func takeCharacters(s string, k int) int {
    target := make(map[byte]int)
    for i := range s {
        target[s[i]]++
    }
    for _, c := range []byte{'a', 'b', 'c'} {
        if target[c] - k < 0 {
            return -1
        }
        target[c] = target[c] - k
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res, l, window := len(s), 0, make(map[byte]int)
    for r := 0; r < len(s); r++ {
        window[s[r]]++
        for window['a'] > target['a'] || window['b'] > target['b'] || window['c'] > target['c'] {
            window[s[l]]--
            l++
        }
        res = min(res, len(s) - (r - l + 1))
    }
    return res
}

func takeCharacters1(s string, k int) int {
    n := len(s)
    count := [3]int{}
    j := n
    for count[0] < k || count[1] < k || count[2] < k {
        if j == 0 {
            return -1
        }
        j--
        count[s[j] - 'a'] ++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := n - j
    for i := 0; i < n && j < n; i++ {
        count[s[i] - 'a']++
        for j < n && count[s[j] - 'a'] > k {
            count[s[j] - 'a']--
            j++
        }
        res = min(res, i + 1 + n - j)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aabaaaacaabc", k = 2
    // Output: 8
    // Explanation: 
    // Take three characters from the left of s. You now have two 'a' characters, and one 'b' character.
    // Take five characters from the right of s. You now have four 'a' characters, two 'b' characters, and two 'c' characters.
    // A total of 3 + 5 = 8 minutes is needed.
    // It can be proven that 8 is the minimum number of minutes needed.
    fmt.Println(takeCharacters("aabaaaacaabc", 2)) // 8
    // Example 2:
    // Input: s = "a", k = 1
    // Output: -1
    // Explanation: It is not possible to take one 'b' or 'c' so return -1.
    fmt.Println(takeCharacters("a", 1)) // -1

    fmt.Println(takeCharacters1("aabaaaacaabc", 2)) // 8
    fmt.Println(takeCharacters1("a", 1)) // -1
}