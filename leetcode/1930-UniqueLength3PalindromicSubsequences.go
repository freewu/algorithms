package main

// 1930. Unique Length-3 Palindromic Subsequences
// Given a string s, return the number of unique palindromes of length three that are a subsequence of s.

// Note that even if there are multiple ways to obtain the same subsequence, it is still only counted once.

// A palindrome is a string that reads the same forwards and backwards.

// A subsequence of a string is a new string generated from the original string 
// with some characters (can be none) deleted without changing the relative order of the remaining characters.
//     For example, "ace" is a subsequence of "abcde".

// Example 1:
// Input: s = "aabca"
// Output: 3
// Explanation: The 3 palindromic subsequences of length 3 are:
// - "aba" (subsequence of "aabca")
// - "aaa" (subsequence of "aabca")
// - "aca" (subsequence of "aabca")

// Example 2:
// Input: s = "adc"
// Output: 0
// Explanation: There are no palindromic subsequences of length 3 in "adc".

// Example 3:
// Input: s = "bbcbaba"
// Output: 4
// Explanation: The 4 palindromic subsequences of length 3 are:
// - "bbb" (subsequence of "bbcbaba")
// - "bcb" (subsequence of "bbcbaba")
// - "bab" (subsequence of "bbcbaba")
// - "aba" (subsequence of "bbcbaba")

// Constraints:
//     3 <= s.length <= 10^5
//     s consists of only lowercase English letters.

import "fmt"
import "math/bits"

func countPalindromicSubsequence(s string) int {
    res, inf := 0, 1 << 31
    first, last := make([]int, 26),  make([]int, 26) // 记录每个字符第一次出现的位置和最后一次出现的位置
    for i := range first{
        first[i] = inf
    }
    for i := range s {
        index := int(s[i]) - int('a')
        first[index], last[index] = min(first[index], i), i
    }
    for i := 0; i < 26;i++ {
        if first[i] < last[i] {
            substring := s[first[i] + 1 : last[i]]
            set := make(map[rune]bool)
            for _, v := range substring {
                set[v] = true
            }
            res += len(set)
        }
    }
    return res
}

func countPalindromicSubsequence1(s string) int {
    // 思路：因为题目只要求长度为3，那么就意味着，该回文肯定会有一个中心字符，核心思想是只要在该字符两侧找到一个相同字符即可，
    // 维护两个数组pre[i]和suff[i],表示前缀的字符比特位(不包括当前位)，如果对应字符的比特位为1，表示前缀拥有该字符；后缀也如此
    // 那么以该字符为中心的回文子序列为：pre[i]&suff[i] 中为1的个数
    // 可能会出现重复，因此还需要维护一个[26]int arr的数组 ，arr[i]表示以第i个字符为中心的回文子序列个数(1的个数就是回文子序列的个数)
    pre, suff := make([]int,len(s)), make([]int,len(s))
    for i := 1; i < len(s); i++ {
        pre[i] = pre[i-1]|(1 << (s[i - 1] - 'a'))
    }
    for i := len(s) - 2; i >= 0; i--{
        suff[i] = suff[i+1]|(1 << (s[i + 1] - 'a'))
    }
    arr := [26]int{} // arr表示以第i个字符为中心的回文子序列的情况
    for i := 1;i < len(s) - 1; i++ {
        arr[s[i] - 'a'] |= (pre[i] & suff[i])
    }
    bitCount := func(bit int) int {
        res := 0
        for i := 0; i < 32; i++{
            if (bit >> i & 1) == 1 { res++ }
        }
        return res
    }
    res := 0
    for i := 0; i < 26; i++ {
        res += bitCount(arr[i])
    }
    return res
}

func countPalindromicSubsequence2(s string) int {
    res, suffix, prefix, n := 0, 0, 0, len(s)
    count,has := [26]int{}, [26]int{} // 统计后缀每个字母的个数
    for _, v := range s[1:] {
        v -= 'a'
        count[v]++
        suffix |= 1 << v
    }
    for i := 1; i < n-1; i++ {
        prefix |= 1 << (s[i-1] - 'a')
        v := s[i] - 'a'
        count[v]--
        if count[v] == 0 { // 现在，后缀 [i+1,n-1] 不包含字母 v
            suffix ^= 1 << v // 从 suffix 中去掉 v
        }
        has[v] |= prefix & suffix
    }
    for _, mask := range has {
        res += bits.OnesCount(uint(mask))
    }
    return res  
}

func main() {
    // Example 1:
    // Input: s = "aabca"
    // Output: 3
    // Explanation: The 3 palindromic subsequences of length 3 are:
    // - "aba" (subsequence of "aabca")
    // - "aaa" (subsequence of "aabca")
    // - "aca" (subsequence of "aabca")
    fmt.Println(countPalindromicSubsequence("aabca")) // 3
    // Example 2:
    // Input: s = "adc"
    // Output: 0
    // Explanation: There are no palindromic subsequences of length 3 in "adc".
    fmt.Println(countPalindromicSubsequence("adc")) // 0
    // Example 3:
    // Input: s = "bbcbaba"
    // Output: 4
    // Explanation: The 4 palindromic subsequences of length 3 are:
    // - "bbb" (subsequence of "bbcbaba")
    // - "bcb" (subsequence of "bbcbaba")
    // - "bab" (subsequence of "bbcbaba")
    // - "aba" (subsequence of "bbcbaba")
    fmt.Println(countPalindromicSubsequence("bbcbaba")) // 4

    fmt.Println(countPalindromicSubsequence("leetcode")) // 5
    fmt.Println(countPalindromicSubsequence("bluefrog")) // 0

    fmt.Println(countPalindromicSubsequence1("aabca")) // 3
    fmt.Println(countPalindromicSubsequence1("adc")) // 0
    fmt.Println(countPalindromicSubsequence1("bbcbaba")) // 4
    fmt.Println(countPalindromicSubsequence1("leetcode")) // 5
    fmt.Println(countPalindromicSubsequence1("bluefrog")) // 0

    fmt.Println(countPalindromicSubsequence2("aabca")) // 3
    fmt.Println(countPalindromicSubsequence2("adc")) // 0
    fmt.Println(countPalindromicSubsequence2("bbcbaba")) // 4
    fmt.Println(countPalindromicSubsequence2("leetcode")) // 5
    fmt.Println(countPalindromicSubsequence2("bluefrog")) // 0
}