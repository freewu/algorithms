package main

// 1100. Find K-Length Substrings With No Repeated Characters
// Given a string s and an integer k, return the number of substrings in s of length k with no repeated characters.

// Example 1:
// Input: s = "havefunonleetcode", k = 5
// Output: 6
// Explanation: There are 6 substrings they are: 'havef','avefu','vefun','efuno','etcod','tcode'.

// Example 2:
// Input: s = "home", k = 5
// Output: 0
// Explanation: Notice k can be larger than the length of s. In this case, it is not possible to find any substring.

// Constraints:
//     1 <= s.length <= 10^4
//     s consists of lowercase English letters.
//     1 <= k <= 10^4

import "fmt"

// 滑动窗口
func numKLenSubstrNoRepeats(s string, k int) int {
    if k > len(s) {
        return 0
    }
    m := make(map[byte]int) // 改字符所在下标
    left, right := 0, 0
    res := 0
    for ; left < len(s) && right < len(s); right++ {
        c := s[right]
        if v, ok := m[c]; ok && v >= left {
            left = v + 1
        }
        m[c] = right
        // 长度满足,向 -> 滑动
        if right - left + 1 == k {
            res++
            left++
        }
    }
    return res
}

// best solution
func numKLenSubstrNoRepeats1(s string, k int) int {
    i, res, m, left := 0, 0, make(map[byte]int), 0
    for left <= len(s)-k {
        if index, ok := m[s[i]]; !ok {
            m[s[i]] = i
            if i - left + 1 == k {
                delete(m, s[left])
                left++; i++; res++
            } else {
                i++
            }
        } else {
            left, i = index + 1,index + 1
            m = make(map[byte]int)
        }
    }
    return res
}

func main() {
    // There are 6 substrings they are: 'havef','avefu','vefun','efuno','etcod','tcode'.
    fmt.Println(numKLenSubstrNoRepeats("havefunonleetcode",5))
    // Notice k can be larger than the length of s. In this case, it is not possible to find any substring.
    fmt.Println(numKLenSubstrNoRepeats("home",5)) // 0

    // There are 6 substrings they are: 'havef','avefu','vefun','efuno','etcod','tcode'.
    fmt.Println(numKLenSubstrNoRepeats1("havefunonleetcode",5))
    // Notice k can be larger than the length of s. In this case, it is not possible to find any substring.
    fmt.Println(numKLenSubstrNoRepeats1("home",5)) // 0
}