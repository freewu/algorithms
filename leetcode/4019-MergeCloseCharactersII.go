package main

// 4019. Merge Close Characters II
// You are given a string s consisting of lowercase English letters and an integer k.

// Two equal characters s[i] and s[j], where 0 <= i < j < s.length, are considered close if j - i <= k. 
// All indices refer to the current string.

// Repeatedly perform the following operation until no close pair remains:
//     1. Among all close pairs (i, j), choose the pair with the smallest i. 
//        If multiple pairs have the same i, choose the one with the smallest j.
//     2. Merge the right character into the left character by removing s[j] from s. 
//        The character s[i] remains unchanged, and the remaining characters are reindexed.

// Return the resulting string after performing all possible merges.

// Example 1:
// Input: s = "abca", k = 3
// Output: "abc"
// Explanation:
// The characters 'a' at indices 0 and 3 are close because 3 - 0 = 3 <= k.
// Remove the right 'a', resulting in s = "abc".
// No close pair remains, so no further merges are performed.

// Example 2:
// Input: s = "aabca", k = 2
// Output: "abca"
// Explanation:
// The characters 'a' at indices 0 and 1 are close because 1 - 0 = 1 <= k.
// Remove the right 'a', resulting in s = "abca".
// The remaining 'a' characters are at indices 0 and 3. Since 3 - 0 = 3 > k, no further merges are performed.

// Example 3:
// Input: s = "yybyzybz", k = 2
// Output: "ybzybz"
// Explanation:
// The characters 'y' at indices 0 and 1 are close because 1 - 0 = 1 <= k. This pair has the smallest left index among all close pairs.
// Remove the right 'y', resulting in s = "ybyzybz".
// The characters 'y' at indices 0 and 2 are now close because 2 - 0 = 2 <= k.
// Remove the right 'y', resulting in s = "ybzybz".
// No close pair remains, so no further merges are performed.

// Constraints:
//     1 <= s.length <= 5 * 10^5
//     1 <= k <= s.length
//     s consists of lowercase English letters.

import "fmt"

// 超出时间限制 974 / 999
func mergeCharacters(s string, k int) string {
    res := []byte(s)
    for {
        n, foundI, foundJ := len(res), -1, -1
        // 找最小i，然后该i下最小j
        for i := 0; i < n; i++ {
            for j := i + 1; j <= i+k && j < n; j++ {
                if res[i] == res[j] {
                    foundI = i
                    foundJ = j
                    goto endSearch
                }
            }
        }
    endSearch:
        if foundI == -1 {
            break
        }
        // 删除j位置
        res = append(res[:foundJ], res[foundJ + 1:]...)           
    }
    return string(res)
}

func mergeCharacters1(s string, k int) string {
    dawn := make([]int, 26)
    for i := range dawn { 
        dawn[i] = -1 
    }
    res := make([]byte, 0, len(s))
    for breath := 0; breath < len(s); breath++ {
        note := s[breath]
        val := note - 'a'
        if dawn[val] != -1 && len(res) - dawn[val] <= k { 
            continue 
        }
        dawn[val] = len(res)
        res = append(res, note)
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "abca", k = 3
    // Output: "abc"
    // Explanation:
    // The characters 'a' at indices 0 and 3 are close because 3 - 0 = 3 <= k.
    // Remove the right 'a', resulting in s = "abc".
    // No close pair remains, so no further merges are performed.
    fmt.Println(mergeCharacters("abca", 3)) // "abc"
    // Example 2:
    // Input: s = "aabca", k = 2
    // Output: "abca"
    // Explanation:
    // The characters 'a' at indices 0 and 1 are close because 1 - 0 = 1 <= k.
    // Remove the right 'a', resulting in s = "abca".
    // The remaining 'a' characters are at indices 0 and 3. Since 3 - 0 = 3 > k, no further merges are performed.
    fmt.Println(mergeCharacters("aabca", 2)) // "abca"
    // Example 3:
    // Input: s = "yybyzybz", k = 2
    // Output: "ybzybz"
    // Explanation:
    // The characters 'y' at indices 0 and 1 are close because 1 - 0 = 1 <= k. This pair has the smallest left index among all close pairs.
    // Remove the right 'y', resulting in s = "ybyzybz".
    // The characters 'y' at indices 0 and 2 are now close because 2 - 0 = 2 <= k.
    // Remove the right 'y', resulting in s = "ybzybz".
    // No close pair remains, so no further merges are performed.
    fmt.Println(mergeCharacters("yybyzybz", 2)) // "ybzybz"

    fmt.Println(mergeCharacters("bluefrog", 3)) // "bluefrog"
    fmt.Println(mergeCharacters("leetcode", 3)) // "leetcode"
    fmt.Println(mergeCharacters("freewu", 3)) // "frewu"

    fmt.Println(mergeCharacters1("abca", 3)) // "abc"
    fmt.Println(mergeCharacters1("yybyzybz", 2)) // "ybzybz"
    fmt.Println(mergeCharacters1("bluefrog", 3)) // "bluefrog"
    fmt.Println(mergeCharacters1("leetcode", 3)) // "leetcode"
    fmt.Println(mergeCharacters1("freewu", 3)) // "frewu"
}