package main

// 1960. Maximum Product of the Length of Two Palindromic Substrings
// You are given a 0-indexed string s and are tasked with finding two non-intersecting palindromic substrings of odd length such that the product of their lengths is maximized.

// More formally, you want to choose four integers i, j, k, l such that 0 <= i <= j < k <= l < s.length 
// and both the substrings s[i...j] and s[k...l] are palindromes and have odd lengths. s[i...j] denotes a substring from index i to index j inclusive.

// Return the maximum possible product of the lengths of the two non-intersecting palindromic substrings.

// A palindrome is a string that is the same forward and backward. 
// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: s = "ababbb"
// Output: 9
// Explanation: Substrings "aba" and "bbb" are palindromes with odd length. product = 3 * 3 = 9.

// Example 2:
// Input: s = "zaaaxbbby"
// Output: 9
// Explanation: Substrings "aaa" and "bbb" are palindromes with odd length. product = 3 * 3 = 9.

// Constraints:
//     2 <= s.length <= 10^5
//     s consists of lowercase English letters.

import "fmt"

func maxProduct(s string) int64 {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // Manacher's Algorithm (confined to odd-only results)
    // The goal of this algorithm is to re-use sub-palindromes confined within
    // larger palindromes.
    //
    // For example, when visiting i=3 for:
    // 'aaabaaa'
    //  0123456
    // A palindrome is found with the size 7.
    // Once i=5, then there will be another palindrome of size 3 ('aaa')
    manachers := func(s string, n int) []int {
        m, maxLeft := make([]int, n), make([]int, n)
        for i := range maxLeft {
            maxLeft[i] = 1
        }
        for i, l, r := 0, 0, -1; i < n; i++ {
            k := 1
            if i <= r {
                k = min(m[l + r - i], r - i + 1)
            }
            for i - k >= 0 && i + k < n && s[i - k] == s[i + k] {
                // Not part of Manacher's Algo: keep track of the maximum palindrome
                // found to the left of where the current palindrome ends:
                maxLeft[i + k] = 2 * k + 1
                k++
            }
            m[i] = k
            if i + k > r {
                l = i - k + 1
                r = i + k - 1
            }
        }
        for i := 1; i < n; i++ { // Fill graps in the prefix max with longest palindrome.
            maxLeft[i] = max(maxLeft[i], maxLeft[i-1])
        }
        return maxLeft
    }
    n := len(s)
    maxLeft := manachers(s, n)
    // To get the maximum palindrome to the right of the string, reverse the input
    // and the result to get maxRight.
    rev := []byte(s)
    for l, r := 0, len(rev)-1; l < r; l, r = l+1, r-1 {
        rev[l], rev[r] = rev[r], rev[l]
    }
    maxRight := manachers(string(rev), n)
    for l, r := 0, len(maxRight)-1; l < r; l, r = l+1, r-1 {
        maxRight[l], maxRight[r] = maxRight[r], maxRight[l]
    }
    // Find maximum product
    res := 1
    for i := 0; i < len(maxLeft)-1; i++ {
        res = max(res, maxLeft[i]*maxRight[i+1])
    }
    return int64(res)
}

func maxProduct1(s string) int64 {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    manachar := func(s string) ([]int, []int) {
        n := len(s)
        d1, d2, curStr := make([]int, n), make([]int, n), []byte{}
        for i := 0; i < n; i++ {
            curStr = append(curStr, '#')
            curStr = append(curStr, s[i])
        }
        curStr = append(curStr, '#')
        m := len(curStr)
        d := make([]int, m)	// palindrome radius
        l, r := 0, -1
        for i := range curStr {
            // curStr[i + gap] and curStr[i - gap] will be the first to compare
            gap := 1
            if i <= r {
                gap = min(r - i + 1, d[l + r - i])
            }
            for i - gap >= 0 && i + gap < m && curStr[i - gap] == curStr[i + gap] {
                gap++
            }
            // d[(i - gap + 1) ... (i + gap - 1)] is palindrome
            d[i] = gap
            if i + gap - 1 > r {
                r, l = i + gap - 1, i - gap + 1
            }	
        }
        for i := range s {
            d1[i], d2[i] = d[2 * i + 1] / 2, (d[2 * i] - 1) / 2
        }
        return d1, d2
    }
    d1, _ := manachar(s)
    res, n := 0, len(s)
    prefix, suffix := make([]int, n), make([]int, n)
    for i := range prefix {
        prefix[i], suffix[i] = 1, 1
    }
    for i, v := range d1 {
        prefix[i + v - 1], suffix[i - v + 1] = max(prefix[i + v - 1], 2 * v - 1), max(suffix[i - v + 1], 2 * v - 1)
    }
    for i := range prefix {
        if i == 0 { continue }
        prefix[i] = max(prefix[i], prefix[i - 1])
    }
    for i := n - 2; i >= 0; i-- {
        prefix[i] = max(prefix[i], prefix[i + 1] - 2)
    }
    for i := n - 2; i >= 0; i-- {
        suffix[i] = max(suffix[i], suffix[i + 1])
    }
    for i := 1; i < n; i++ {
        suffix[i] = max(suffix[i], suffix[i - 1] - 2)
    }
    for i := 0; i < n - 1; i++ {
        res = max(res, prefix[i] * suffix[i + 1])
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: s = "ababbb"
    // Output: 9
    // Explanation: Substrings "aba" and "bbb" are palindromes with odd length. product = 3 * 3 = 9.
    fmt.Println(maxProduct("ababbb")) // 9
    // Example 2:
    // Input: s = "zaaaxbbby"
    // Output: 9
    // Explanation: Substrings "aaa" and "bbb" are palindromes with odd length. product = 3 * 3 = 9.
    fmt.Println(maxProduct("zaaaxbbby")) // 9

    fmt.Println(maxProduct1("ababbb")) // 9
    fmt.Println(maxProduct1("zaaaxbbby")) // 9
}