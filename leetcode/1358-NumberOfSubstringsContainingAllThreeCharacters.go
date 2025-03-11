package main

// 1358. Number of Substrings Containing All Three Characters
// Given a string s consisting only of characters a, b and c.

// Return the number of substrings containing at least one occurrence of all these characters a, b and c.

// Example 1:
// Input: s = "abcabc"
// Output: 10
// Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again). 

// Example 2:
// Input: s = "aaacb"
// Output: 3
// Explanation: The substrings containing at least one occurrence of the characters a, b and c are "aaacb", "aacb" and "acb". 

// Example 3:
// Input: s = "abc"
// Output: 1

// Constraints:
//     3 <= s.length <= 5 x 10^4
//     s only consists of a, b or c characters.

import "fmt"

func numberOfSubstrings(s string) int {
    count := make([]int, 3)
    res, h, t, n := 0, 0, -1, len(s)
    for t < n {
        for count[0] * count[1] * count[2] == 0 {
            if t == n - 1 {
                return res
            }
            t++
            count[s[t] - 'a'] ++ 
        }
        res += n - t
        count[s[h] - 'a'] -- 
        h++ 
    }
    return res
}

func numberOfSubstrings1(s string) int {
    count := ['d']int{}
    res, l := 0, 0
    for r, b := range s {
        count[b] += 1
        if count['a'] > 0 && count['b'] > 0 && count['c'] > 0 {
            for l < r && count[s[l]] > 1 {
                count[s[l]]--
                l++
            }
            res += l + 1
        }
    }
    return res
}

func numberOfSubstrings2(s string) int {
    res, mp := 0, make([]int, 3)
    for i := 0; i < 3; i++ {
        mp[i] = -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, v := range s {
        mp[v - 'a'] = i 
        if mp[0] != -1 && mp[1] != -1 && mp[2] != -1 {
            // substrings  = ending at max of these lastSeen and starting from 0 to minIndex
            // so total minIndex+1 variations
            res += 1 + min(mp[0], min(mp[1], mp[2])) 
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcabc"
    // Output: 10
    // Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again). 
    fmt.Println(numberOfSubstrings("abcabc")) // 10
    // Example 2:
    // Input: s = "aaacb"
    // Output: 3
    // Explanation: The substrings containing at least one occurrence of the characters a, b and c are "aaacb", "aacb" and "acb". 
    fmt.Println(numberOfSubstrings("aaacb")) // 3
    // Example 3:
    // Input: s = "abc"
    // Output: 1
    fmt.Println(numberOfSubstrings("abc")) // 1

    fmt.Println(numberOfSubstrings1("abcabc")) // 10
    fmt.Println(numberOfSubstrings1("aaacb")) // 3
    fmt.Println(numberOfSubstrings1("abc")) // 1

    fmt.Println(numberOfSubstrings2("abcabc")) // 10
    fmt.Println(numberOfSubstrings2("aaacb")) // 3
    fmt.Println(numberOfSubstrings2("abc")) // 1
}