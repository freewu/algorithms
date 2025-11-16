package main

// 3746. Minimum String Length After Balanced Removals
// You are given a string s consisting only of the characters 'a' and 'b'.

// You are allowed to repeatedly remove any substring where the number of 'a' characters is equal to the number of 'b' characters. 
// After each removal, the remaining parts of the string are concatenated together without gaps.

// Return an integer denoting the minimum possible length of the string after performing any number of such operations.

// Example 1:
// Input: s = "aabbab"
// Output: 0
// Explanation:
// The substring "aabbab" has three 'a' and three 'b'. Since their counts are equal, we can remove the entire string directly. The minimum length is 0.

// Example 2:
// Input: s = "aaaa"
// Output: 4
// Explanation:
// Every substring of "aaaa" contains only 'a' characters. No substring can be removed as a result, so the minimum length remains 4.

// Example 3:
// Input: s = "aaabb"
// Output: 1
// Explanation:
// First, remove the substring "ab", leaving "aab". Next, remove the new substring "ab", leaving "a". No further removals are possible, so the minimum length is 1.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either 'a' or 'b'.

import "fmt"
import "strings"

// time: O(n), space: O(1)
func minLengthAfterRemovals(s string) int {
    ac, bc := 0, 0
    for i := range s {
        if s[i] == 'a' {
            ac++
        } else {
            bc++
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    return abs(ac - bc) 
}

func minLengthAfterRemovals1(s string) int {
    a := strings.Count(s, "a")
    b := len(s) - a
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    return abs(a - b)
}

func main() {
    // Example 1:
    // Input: s = "aabbab"
    // Output: 0
    // Explanation:
    // The substring "aabbab" has three 'a' and three 'b'. Since their counts are equal, we can remove the entire string directly. The minimum length is 0.
    fmt.Println(minLengthAfterRemovals("aabbab")) // 0
    // Example 2:
    // Input: s = "aaaa"
    // Output: 4
    // Explanation:
    // Every substring of "aaaa" contains only 'a' characters. No substring can be removed as a result, so the minimum length remains 4.
    fmt.Println(minLengthAfterRemovals("aaaa")) // 4
    // Example 3:
    // Input: s = "aaabb"
    // Output: 1
    // Explanation:
    // First, remove the substring "ab", leaving "aab". Next, remove the new substring "ab", leaving "a". No further removals are possible, so the minimum length is 1.
    fmt.Println(minLengthAfterRemovals("aaabb")) // 1

    fmt.Println(minLengthAfterRemovals("aaaaaaaaaa")) // 10
    fmt.Println(minLengthAfterRemovals("bbbbbbbbbb")) // 10
    fmt.Println(minLengthAfterRemovals("aaaaabbbbb")) // 0
    fmt.Println(minLengthAfterRemovals("bbbbbaaaaa")) // 0
    fmt.Println(minLengthAfterRemovals("ababababab")) // 0
    fmt.Println(minLengthAfterRemovals("bababababa")) // 0

    fmt.Println(minLengthAfterRemovals1("aabbab")) // 0
    fmt.Println(minLengthAfterRemovals1("aaaa")) // 4
    fmt.Println(minLengthAfterRemovals1("aaabb")) // 1
    fmt.Println(minLengthAfterRemovals1("aaaaaaaaaa")) // 10
    fmt.Println(minLengthAfterRemovals1("bbbbbbbbbb")) // 10
    fmt.Println(minLengthAfterRemovals1("aaaaabbbbb")) // 0
    fmt.Println(minLengthAfterRemovals1("bbbbbaaaaa")) // 0
    fmt.Println(minLengthAfterRemovals1("ababababab")) // 0
    fmt.Println(minLengthAfterRemovals1("bababababa")) // 0
}