package main

// 2981. Find Longest Special Substring That Occurs Thrice I
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
//     3 <= s.length <= 50
//     s consists of only lowercase English letters.

import "fmt"

func maximumLength(s string) int {
    type st struct {
        char byte
        length int
    }    
    res, stack, n, count := 0, make(map[st]int), len(s), 0
    for i := 0; i < n ; i++ {
        count = 1
        stack[st{char : s[i],length : count}]++
        for j := i; j < n-1; j++ {
            if s[j] == s[j+1] { 
                count++ 
                stack[st{char : s[i],length : count}]++
            } else {
                break
            }
        }
    }
    // fmt.Println(stack)
    for x,v := range stack {
        if v >= 3 {
            if x.length > res {
                res = x.length
            } 
        }
    }
    if res == 0 {
        return -1
    }
    return res
}

func maximumLength1(s string) int {
    cnt := make([]map[int]int, 26)
    for i := range cnt {
        cnt[i] = map[int]int{}
    }
    left, right, n, res := 0, 0, len(s), -1
    for right < n {
        for right < n && s[right] == s[left] {
            right++
        }
        cnt[s[left] - 'a'][right - left]++
        left = right
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    max_map_key := func (d map[int]int) (res int) {
        for k, _ := range d {
            res = max(res, k)
        }
        return
    }
    for _, d := range cnt {
        if len(d) == 0 {
            continue
        }
        for k := max_map_key(d); k > 0; k-- {
            if k < res { break }
            if d[k] >= 3 {
                res = k
                break
            }
            d[k-1] += 2 * d[k]
            d[k-2] += 3
        }
    }
    return res
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