package main

// 3734. Lexicographically Smallest Palindromic Permutation Greater Than Target
// You are given two strings s and target, each of length n, consisting of lowercase English letters.

// Return the lexicographically smallest string that is both a palindromic permutation of s and strictly greater than target. 
// If no such permutation exists, return an empty string.

// Example 1:
// Input: s = "baba", target = "abba"
// Output: "baab"
// Explanation:
// The palindromic permutations of s (in lexicographical order) are "abba" and "baab".
// The lexicographically smallest permutation that is strictly greater than target is "baab".

// Example 2:
// Input: s = "baba", target = "bbaa"
// Output: ""
// Explanation:
// The palindromic permutations of s (in lexicographical order) are "abba" and "baab".
// None of them is lexicographically strictly greater than target. Therefore, the answer is "".

// Example 3:
// Input: s = "abc", target = "abb"
// Output: ""
// Explanation:
// s has no palindromic permutations. Therefore, the answer is "".

// Example 4:
// Input: s = "aac", target = "abb"
// Output: "aca"
// Explanation:
// The only palindromic permutation of s is "aca".
// "aca" is strictly greater than target. Therefore, the answer is "aca".
 
// Constraints:
//     1 <= n == s.length == target.length <= 300
//     s and target consist of only lowercase English letters.

import "fmt"
import "slices"
import "strings"

func lexPalindromicPermutation(s string, target string) string {
    left := make([]int, 26)
    for _, b := range s {
        left[b-'a']++
    }
    midCh := ""
    for i, c := range left {
        if c % 2 == 0 { continue }
        if midCh != "" { return "" } // s 不能有超过一个字母出现奇数次
        // 记录填在正中间的字母
        midCh = string('a' + byte(i))
        left[i]--
    }
    n := len(s)
    // 先假设答案左半与 t 的左半（不含正中间）相同
    for _, b := range target[:n/2] {
        left[b-'a'] -= 2
    }
    neg, leftMax := 0, byte(0)
    for i, cnt := range left {
        if cnt < 0 {
            neg++ // 统计 left 中的负数个数
        } else if cnt > 0 {
            leftMax = max(leftMax, byte(i)) // 剩余可用字母的最大值
        }
    }
    if neg == 0 {
        // 特殊情况：把 target 左半翻转到右半，能否比 target 大？
        leftS := target[:n/2]
        tmp := []byte(leftS)
        slices.Reverse(tmp)
        rightS := midCh + string(tmp)
        if rightS > target[n/2:] { // 由于左半是一样的，所以只需比右半
            return leftS + rightS
        }
    }
    for i := n/2 - 1; i >= 0; i-- {
        b := target[i] - 'a'
        left[b] += 2 // 撤销消耗
        if left[b] == 0 {
            neg--
        } else if left[b] == 2 {
            leftMax = max(leftMax, b)
        }
        // left 有负数 or 没有大于 target[i] 的字母
        if neg > 0 || leftMax <= b {
            continue
        }
        // 找到答案（下面的循环在整个算法中只会跑一次）
        j := b + 1
        for left[j] == 0 {
            j++
        }
        // 把 target[i] 增大到 j
        left[j] -= 2
        res := []byte(target[:i+1])
        res[i] = 'a' + j
        // 中间可以随便填
        for k, c := range left {
            ch := string('a' + byte(k))
            res = append(res, strings.Repeat(ch, c/2)...)
        }
        // 镜像翻转
        rightS := slices.Clone(res)
        slices.Reverse(rightS)
        res = append(res, midCh...)
        res = append(res, rightS...)
        return string(res)
    }
    return ""
}

func lexPalindromicPermutation1(s, target string) string {
    count := [26]int{}
    for _, v := range s {
        count[v-'a']++
    }
    check := func() bool {
        for _, v := range count {
            if v < 0 {
                return false
            }
        }
        return true
    }
    midCh := byte(0)
    for i, v := range count {
        if v%2 == 0 {
            continue
        }
        if midCh != byte(0) {
            return ""
        }

        midCh = byte(i + 'a')
        count[i]--  
    }
    res, n := []byte(target), len(s)
    for i, b := range target[:len(target)/2] {
        count[b-'a'] -= 2
        res[n-1-i] = byte(b)
    }
    if midCh != byte(0) {
        res[n/2] = midCh
    }
    if check() && string(res) > target {
        return string(res)
    }
    for i := len(s)/2 - 1; i >= 0; i-- {
        b := target[i] - 'a'
        count[b] += 2
        if !check() { continue }
        for j := b + 1; j < 26; j++ {
            if count[j] == 0 { continue }   
            count[j] -= 2
            res[i] = byte(j + 'a')
            res[n-1-i] = res[i]
            t := make([]byte, 0)
            for k, v := range count {       
                t = append(t, strings.Repeat(string(k+'a'), v/2)...)
            }
            a := append(res[:i+1], t...)
            if midCh != byte(0) {
                a = append(a, midCh)
            }
            slices.Reverse(t)
            a = append(a, t...)
            return string(res)
        }
    }
    return ""
}

func main() {
    // Example 1:
    // Input: s = "baba", target = "abba"
    // Output: "baab"
    // Explanation:
    // The palindromic permutations of s (in lexicographical order) are "abba" and "baab".
    // The lexicographically smallest permutation that is strictly greater than target is "baab".
    fmt.Println(lexPalindromicPermutation("baba", "abba")) // "baab"
    // Example 2:
    // Input: s = "baba", target = "bbaa"
    // Output: ""
    // Explanation:
    // The palindromic permutations of s (in lexicographical order) are "abba" and "baab".
    // None of them is lexicographically strictly greater than target. Therefore, the answer is "".
    fmt.Println(lexPalindromicPermutation("baba", "bbaa")) // ""
    // Example 3:
    // Input: s = "abc", target = "abb"
    // Output: ""
    // Explanation:
    // s has no palindromic permutations. Therefore, the answer is "".
    fmt.Println(lexPalindromicPermutation("abc", "abb")) // ""
    // Example 4:
    // Input: s = "aac", target = "abb"
    // Output: "aca"
    // Explanation:
    // The only palindromic permutation of s is "aca".
    // "aca" is strictly greater than target. Therefore, the answer is "aca".
    fmt.Println(lexPalindromicPermutation("aac", "abb")) // "aca"

    fmt.Println(lexPalindromicPermutation("bluefrog", "leetcode")) // ""
    fmt.Println(lexPalindromicPermutation("bluefrog", "bluefrog")) // ""
    fmt.Println(lexPalindromicPermutation("leetcode", "bluefrog")) // ""
    fmt.Println(lexPalindromicPermutation("leetcode", "leetcode")) // ""

    fmt.Println(lexPalindromicPermutation1("baba", "abba")) // "baab"
    fmt.Println(lexPalindromicPermutation1("baba", "bbaa")) // ""
    fmt.Println(lexPalindromicPermutation1("abc", "abb")) // ""
    fmt.Println(lexPalindromicPermutation1("aac", "abb")) // "aca"
    fmt.Println(lexPalindromicPermutation1("bluefrog", "leetcode")) // ""
    fmt.Println(lexPalindromicPermutation1("bluefrog", "bluefrog")) // ""
    fmt.Println(lexPalindromicPermutation1("leetcode", "bluefrog")) // ""
    fmt.Println(lexPalindromicPermutation1("leetcode", "leetcode")) // ""
}