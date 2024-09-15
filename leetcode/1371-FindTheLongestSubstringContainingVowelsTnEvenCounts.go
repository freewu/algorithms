package main

// 1371. Find the Longest Substring Containing Vowels in Even Counts
// Given the string s, return the size of the longest substring containing each vowel an even number of times. 
// That is, 'a', 'e', 'i', 'o', and 'u' must appear an even number of times.

// Example 1:
// Input: s = "eleetminicoworoep"
// Output: 13
// Explanation: The longest substring is "leetminicowor" which contains two each of the vowels: e, i and o and zero of the vowels: a and u.

// Example 2:
// Input: s = "leetcodeisgreat"
// Output: 5
// Explanation: The longest substring is "leetc" which contains two e's.

// Example 3:
// Input: s = "bcbcbc"
// Output: 6
// Explanation: In this case, the given string "bcbcbc" is the longest because all vowels: a, e, i, o and u appear zero times.

// Constraints:
//     1 <= s.length <= 5 x 10^5
//     s contains only lowercase English letters.

import "fmt"

func findTheLongestSubstring(s string) int {
    vowels := map[byte]int{'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}
    counts := map[int]int{0: -1}
    res, bitmask := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        char := s[i]
        if val, ok := vowels[char]; ok {
            bitmask ^= val
        }
        if val, ok := counts[bitmask]; ok {
            res = max(res, i - val)
        } else {
            counts[bitmask] = i
        }
    }
    return res
}

func findTheLongestSubstring1(s string) int {
    res, status := 0, 0
    pos := make([]int, 1 << 5)
    for i := 0; i < len(pos); i++ { pos[i] = -1 } // fill -1
    pos[0] = 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case 'a': status ^= 1 << 0
        case 'e': status ^= 1 << 1
        case 'i': status ^= 1 << 2
        case 'o': status ^= 1 << 3
        case 'u': status ^= 1 << 4
        }
        if pos[status] >= 0 {
            res = max(res, i + 1 - pos[status])
        } else {
            pos[status] = i + 1
        }
    }
    return res
}

func findTheLongestSubstring2(s string) int {
    res, cur, mp := 0, 0, make([]int,32)
    mp[0] = -1
    for i := 1; i < 32; i++ { 
        mp[i] = -2 // fill -2
    }
    for i, ch := range s {
        switch ch{
            case 97:  cur ^= 1 // a
            case 101: cur ^= 2 // e
            case 105: cur ^= 4 // i
            case 111: cur ^= 8 // o
            case 117: cur ^=16 // u
        }
        switch p := mp[cur]; p {
            case -2: mp[cur] = i // 第一次赋值
            default:
                if vs := i - p; res < vs {
                    res = vs
                }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "eleetminicoworoep"
    // Output: 13
    // Explanation: The longest substring is "leetminicowor" which contains two each of the vowels: e, i and o and zero of the vowels: a and u.
    fmt.Println(findTheLongestSubstring("eleetminicoworoep")) // 13
    // Example 2:
    // Input: s = "leetcodeisgreat"
    // Output: 5
    // Explanation: The longest substring is "leetc" which contains two e's.
    fmt.Println(findTheLongestSubstring("leetcodeisgreat")) // 5
    // Example 3:
    // Input: s = "bcbcbc"
    // Output: 6
    // Explanation: In this case, the given string "bcbcbc" is the longest because all vowels: a, e, i, o and u appear zero times.
    fmt.Println(findTheLongestSubstring("bcbcbc")) // 6

    fmt.Println(findTheLongestSubstring1("eleetminicoworoep")) // 13
    fmt.Println(findTheLongestSubstring1("leetcodeisgreat")) // 5
    fmt.Println(findTheLongestSubstring1("bcbcbc")) // 6

    fmt.Println(findTheLongestSubstring2("eleetminicoworoep")) // 13
    fmt.Println(findTheLongestSubstring2("leetcodeisgreat")) // 5
    fmt.Println(findTheLongestSubstring2("bcbcbc")) // 6
}