package main

// 2067. Number of Equal Count Substrings
// You are given a 0-indexed string s consisting of only lowercase English letters, and an integer count. 
// A substring of s is said to be an equal count substring if, for each unique letter in the substring, it appears exactly count times in the substring.

// Return the number of equal count substrings in s.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "aaabcbbcc", count = 3
// Output: 3
// Explanation:
// The substring that starts at index 0 and ends at index 2 is "aaa".
// The letter 'a' in the substring appears exactly 3 times.
// The substring that starts at index 3 and ends at index 8 is "bcbbcc".
// The letters 'b' and 'c' in the substring appear exactly 3 times.
// The substring that starts at index 0 and ends at index 8 is "aaabcbbcc".
// The letters 'a', 'b', and 'c' in the substring appear exactly 3 times.

// Example 2:
// Input: s = "abcd", count = 2
// Output: 0
// Explanation:
// The number of times each letter appears in s is less than count.
// Therefore, no substrings in s are equal count substrings, so return 0.

// Example 3:
// Input: s = "a", count = 5
// Output: 0
// Explanation:
// The number of times each letter appears in s is less than count.
// Therefore, no substrings in s are equal count substrings, so return 0

// Constraints:
//     1 <= s.length <= 3 * 10^4
//     1 <= count <= 3 * 10^4
//     s consists only of lowercase English letters.

import "fmt"

func equalCountSubstrings(s string, count int) int {
    arr := make([][26]int, len(s))
    check := func(t [26]int) bool {
        for i := 0; i < 26; i++ {
            if t[i] == 0 { continue }
            if t[i] != count { return false }
        }
        return true
    }
    res, cur := 0, [26]int{}
    for i := 0; i < len(s); i++ {
        cur[s[i]-'a']++
        if check(cur) { res++ }
        arr[i] = cur
        for j := 1; j <= 26 && i - j * count >= 0; j++ {
            t := [26]int{}
            for k := 0; k < 26; k++ {
                t[k] = cur[k] - arr[i - j * count][k]
            }
            if check(t) { res++ }
        }
    }
    return res
}

func equalCountSubstrings1(s string, count int) int {
    res, n := 0, len(s)
    for k := count; k <= 26 * count; k += count {
        if k > n { break }
        t, mp := 0, make(map[byte]int)
        for i := 0; i < k; i++ {
            mp[s[i]-'a']++
        }
        for i := k; i <= n; i++ {
            if len(mp) == k / count {
                add := true
                for _, c := range mp {
                    if c != count {
                        add = false
                    }
                }
                if add {
                    t++
                }
            }
            if i == n { break }
            mp[s[i] - 'a']++
            mp[s[i - k] - 'a']--
            if mp[s[i - k] - 'a'] == 0 {
                delete(mp, s[i - k] - 'a')
            }
        }
        res += t
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aaabcbbcc", count = 3
    // Output: 3
    // Explanation:
    // The substring that starts at index 0 and ends at index 2 is "aaa".
    // The letter 'a' in the substring appears exactly 3 times.
    // The substring that starts at index 3 and ends at index 8 is "bcbbcc".
    // The letters 'b' and 'c' in the substring appear exactly 3 times.
    // The substring that starts at index 0 and ends at index 8 is "aaabcbbcc".
    // The letters 'a', 'b', and 'c' in the substring appear exactly 3 times.
    fmt.Println(equalCountSubstrings("aaabcbbcc", 3)) // 3
    // Example 2:
    // Input: s = "abcd", count = 2
    // Output: 0
    // Explanation:
    // The number of times each letter appears in s is less than count.
    // Therefore, no substrings in s are equal count substrings, so return 0.
    fmt.Println(equalCountSubstrings("abcd", 2)) // 0
    // Example 3:
    // Input: s = "a", count = 5
    // Output: 0
    // Explanation:
    // The number of times each letter appears in s is less than count.
    // Therefore, no substrings in s are equal count substrings, so return 0
    fmt.Println(equalCountSubstrings("a", 5)) // 0

    fmt.Println(equalCountSubstrings1("aaabcbbcc", 3)) // 3
    fmt.Println(equalCountSubstrings1("abcd", 2)) // 0
    fmt.Println(equalCountSubstrings1("a", 5)) // 0
}