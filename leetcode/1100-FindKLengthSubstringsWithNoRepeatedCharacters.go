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
    i, res, left := 0, 0, 0
    mp := make(map[byte]int)
    for left <= len(s)-k {
        if index, ok := mp[s[i]]; !ok {
            mp[s[i]] = i
            if i - left + 1 == k {
                delete(mp, s[left])
                left++; i++; res++
            } else {
                i++
            }
        } else {
            left, i = index + 1,index + 1
            mp = make(map[byte]int)
        }
    }
    return res
}

func numKLenSubstrNoRepeats2(s string, k int) int {
    res, n := 0, len(s)
    if n < k { return 0 }
    count := make([]int, 26)
    for i := 0; i < k; i++ {
        count[s[i] - 'a']++
    }
    isValid := func(arr []int) bool {
        for i:= 0; i < 26; i++ {
            if arr[i] > 1 {
                return false
            }
        }
        return true
    }
    if isValid(count) {
        res++
    }
    for i := k; i < n; i++ {
        count[s[i] - 'a']++
        count[s[i - k] - 'a']--
        if isValid(count) {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "havefunonleetcode", k = 5
    // Output: 6
    // Explanation: There are 6 substrings they are: 'havef','avefu','vefun','efuno','etcod','tcode'.
    fmt.Println(numKLenSubstrNoRepeats("havefunonleetcode",5)) // 6
    // Example 2:
    // Input: s = "home", k = 5
    // Output: 0
    // Explanation: Notice k can be larger than the length of s. In this case, it is not possible to find any substring.
    fmt.Println(numKLenSubstrNoRepeats("home", 5)) // 0

    fmt.Println(numKLenSubstrNoRepeats("bluefrog", 5)) // 4
    fmt.Println(numKLenSubstrNoRepeats("leetcode", 5)) // 2

    fmt.Println(numKLenSubstrNoRepeats1("havefunonleetcode", 5)) // 6
    fmt.Println(numKLenSubstrNoRepeats1("home", 5)) // 0
    fmt.Println(numKLenSubstrNoRepeats1("bluefrog", 5)) // 4
    fmt.Println(numKLenSubstrNoRepeats1("leetcode", 5)) // 2

    fmt.Println(numKLenSubstrNoRepeats2("havefunonleetcode", 5)) // 6
    fmt.Println(numKLenSubstrNoRepeats2("home", 5)) // 0
    fmt.Println(numKLenSubstrNoRepeats2("bluefrog", 5)) // 4
    fmt.Println(numKLenSubstrNoRepeats2("leetcode", 5)) // 2
}