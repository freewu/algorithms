package main

// 3853. Merge Close Characters
// You are given a string s consisting of lowercase English letters and an integer k.

// Two equal characters in the current string s are considered close if the distance between their indices is at most k.

// When two characters are close, the right one merges into the left. 
// Merges happen one at a time, and after each merge, the string updates until no more merges are possible.

// Return the resulting string after performing all possible merges.

// Note: 
//     If multiple merges are possible, always merge the pair with the smallest left index. 
//     If multiple pairs share the smallest left index, choose the pair with the smallest right index.

// Example 1:
// Input: s = "abca", k = 3
// Output: "abc"
// Explanation:
// ​​​​​​​Characters 'a' at indices i = 0 and i = 3 are close as 3 - 0 = 3 <= k.
// Merge them into the left 'a' and s = "abc".
// No other equal characters are close, so no further merges occur.

// Example 2:
// Input: s = "aabca", k = 2
// Output: "abca"
// Explanation:
// Characters 'a' at indices i = 0 and i = 1 are close as 1 - 0 = 1 <= k.
// Merge them into the left 'a' and s = "abca".
// Now the remaining 'a' characters at indices i = 0 and i = 3 are not close as k < 3, so no further merges occur.

// Example 3:
// Input: s = "yybyzybz", k = 2
// Output: "ybzybz"
// Explanation:
// Characters 'y' at indices i = 0 and i = 1 are close as 1 - 0 = 1 <= k.
// Merge them into the left 'y' and s = "ybyzybz".
// Now the characters 'y' at indices i = 0 and i = 2 are close as 2 - 0 = 2 <= k.
// Merge them into the left 'y' and s = "ybzybz".
// No other equal characters are close, so no further merges occur.

// Constraints:
//     1 <= s.length <= 100
//     1 <= k <= s.length
//     s consists of lowercase English letters.

import "fmt"

func mergeCharacters(s string, k int) string {
    res, i := []rune(s), 0
    for i < len(res) {
        curr := res[i]
        for j := i + 1; j <= i + k && j < len(res); j++ {
            if res[j] == curr {
                res = append(res[:j], res[j+1:]...)
                i = -1
                break
            }
        }
        i++
    }
    return string(res)
}

func mergeCharacters1(s string, k int) string {
    res, last := []byte{},[26]int{}
    for i := range last {
        last[i] = -k - 1 // 或者其他等价于 -inf 的值 
    }
    for _, v := range s {
        // v 在 res 中的下标是 len(res)
        if len(res) - last[v-'a'] > k {
            last[v-'a'] = len(res)
            res = append(res, byte(v))
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "abca", k = 3
    // Output: "abc"
    // Explanation:
    // ​​​​​​​Characters 'a' at indices i = 0 and i = 3 are close as 3 - 0 = 3 <= k.
    // Merge them into the left 'a' and s = "abc".
    // No other equal characters are close, so no further merges occur.
    fmt.Println(mergeCharacters("abca", 3)) // "abc"
    // Example 2:
    // Input: s = "aabca", k = 2
    // Output: "abca"
    // Explanation:
    // Characters 'a' at indices i = 0 and i = 1 are close as 1 - 0 = 1 <= k.
    // Merge them into the left 'a' and s = "abca".
    // Now the remaining 'a' characters at indices i = 0 and i = 3 are not close as k < 3, so no further merges occur.
    fmt.Println(mergeCharacters("aabca", 2)) // "abca"
    // Example 3:
    // Input: s = "yybyzybz", k = 2
    // Output: "ybzybz"
    // Explanation:
    // Characters 'y' at indices i = 0 and i = 1 are close as 1 - 0 = 1 <= k.
    // Merge them into the left 'y' and s = "ybyzybz".
    // Now the characters 'y' at indices i = 0 and i = 2 are close as 2 - 0 = 2 <= k.
    // Merge them into the left 'y' and s = "ybzybz".
    // No other equal characters are close, so no further merges occur.
    fmt.Println(mergeCharacters("yybyzybz", 2)) // "ybzybz"

    fmt.Println(mergeCharacters("bluefrog", 2)) // "bluefrog"
    fmt.Println(mergeCharacters("leetcode", 2)) // "letcode"
    fmt.Println(mergeCharacters("freewu", 2)) // "frewu"

    fmt.Println(mergeCharacters1("abca", 3)) // "abc"
    fmt.Println(mergeCharacters1("aabca", 2)) // "abca"
    fmt.Println(mergeCharacters1("yybyzybz", 2)) // "ybzybz"
    fmt.Println(mergeCharacters1("bluefrog", 2)) // "bluefrog"
    fmt.Println(mergeCharacters1("leetcode", 2)) // "letcode"
    fmt.Println(mergeCharacters1("freewu", 2)) // "frewu"
}