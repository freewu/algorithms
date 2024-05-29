package main

// 2982. Find Longest Special Substring That Occurs Thrice II
// You are given a string s that consists of lowercase English letters.
// A string is called special if it is made up of only a single character. 
// For example, the string "abc" is not special, whereas the strings "ddd", "zz", and "f" are special.

// Return the length of the longest special substring of s which occurs at least thrice, 
// or -1 if no special substring occurs at least thrice.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "aaaa"
// Output: 2
// Explanation: The longest special substring which occurs thrice is "aa": substrings "aaaa", "aaaa", and "aaaa".
// It can be shown that the maximum length achievable is 2.

// Example 2:
// Input: s = "abcdef"
// Output: -1
// Explanation: There exists no special substring which occurs at least thrice. Hence return -1.

// Example 3:
// Input: s = "abcaba"
// Output: 1
// Explanation: The longest special substring which occurs thrice is "a": substrings "abcaba", "abcaba", and "abcaba".
// It can be shown that the maximum length achievable is 1.

// Constraints:
//     3 <= s.length <= 5 * 10^5
//     s consists of only lowercase English letters.

import "fmt"

func maximumLength(s string) int {
    res, mem := -1, make([]map[int]int, 26) // Initialize the memory
    for i := range mem {
        mem[i] = make(map[int]int)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    s = s + "-" // Add '-' so that the last character gets processed.
    for l, r :=0, 0; r < len(s); r++ { // Classic sliding window from here
        if s[r] != s[l] {
            mem[s[l]-'a'][r-l]++  // Consider substring of length n
            if mem[s[l]-'a'][r-l] >= 3 {
                res = max(res, r - l)
            }
            t := r-l-1
            if t > 0 { // Consider substring of length n-1
                mem[s[l]-'a'][t]+=2
                if mem[s[l]-'a'][t] >= 3{
                    res = max(res, t)
                }
            }
            if t - 1  > 0 { // Consider substring of length n-2
                res = max(res, t - 1)
            }
            l = r
        }
    }
    return res
}

func maximumLength1(s string) int {
    n, pre, cnt := len(s), 1, make([][]int, 26)
    for i := range cnt {
        cnt[i] = make([]int, 3)
    }
    for i := range s[:n-1] {
        if s[i] == s[i+1] {
            pre++
        } else {
            c := s[i] - 97
            if pre >= cnt[c][0] {
                cnt[c][0], cnt[c][1], cnt[c][2] = pre, cnt[c][0], cnt[c][1]
            } else if pre >= cnt[c][1] {
                cnt[c][1], cnt[c][2] = pre, cnt[c][1]
            } else if pre >= cnt[c][2] {
                cnt[c][2] = pre
            }
            pre = 1
        }
    }
    last := s[n-1] - 'a'
    if pre >= cnt[last][0] {
        cnt[last][0], cnt[last][1], cnt[last][2] = pre, cnt[last][0], cnt[last][1]
    } else if pre >= cnt[last][1] {
        cnt[last][1], cnt[last][2] = pre, cnt[last][1]
    } else if pre >= cnt[last][2] {
        cnt[last][2] = pre
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    maxCnt := -1
    for k := range cnt {
        switch {
        case cnt[k][2] > 0:
            if cnt[k][0] == cnt[k][1] && cnt[k][1] == cnt[k][2] {
                maxCnt = max(maxCnt, cnt[k][0])
            } else if cnt[k][0] <= cnt[k][1]+1 {
                maxCnt = max(maxCnt, cnt[k][0]-1)
            } else {
                maxCnt = max(maxCnt, cnt[k][0]-2)
            }
        case cnt[k][1] > 0:
            if cnt[k][0] <= cnt[k][1]+1 {
                maxCnt = max(maxCnt, cnt[k][0]-1)
            } else {
                maxCnt = max(maxCnt, cnt[k][0]-2)
            }
        case cnt[k][0] > 0:
            maxCnt = max(maxCnt, cnt[k][0]-2)
        }
    }
    if maxCnt == 0 {
        return -1
    } else {
        return maxCnt
    }
}

func main() {
    // Example 1:
    // Input: s = "aaaa"
    // Output: 2
    // Explanation: The longest special substring which occurs thrice is "aa": substrings "aaaa", "aaaa", and "aaaa".
    // It can be shown that the maximum length achievable is 2.
    fmt.Println(maximumLength("aaaa")) // 2
    // Example 2:
    // Input: s = "abcdef"
    // Output: -1
    // Explanation: There exists no special substring which occurs at least thrice. Hence return -1.
    fmt.Println(maximumLength("abcdef")) // -1
    // Example 3:
    // Input: s = "abcaba"
    // Output: 1
    // Explanation: The longest special substring which occurs thrice is "a": substrings "abcaba", "abcaba", and "abcaba".
    // It can be shown that the maximum length achievable is 1.
    fmt.Println(maximumLength("abcaba")) // 1

    fmt.Println(maximumLength1("aaaa")) // 2
    fmt.Println(maximumLength1("abcdef")) // -1
    fmt.Println(maximumLength1("abcaba")) // 1
}