package main

// 1898. Maximum Number of Removable Characters
// You are given two strings s and p where p is a subsequence of s. 
// You are also given a distinct 0-indexed integer array removable containing a subset of indices of s (s is also 0-indexed).

// You want to choose an integer k (0 <= k <= removable.length) such that, 
// after removing k characters from s using the first k indices in removable, p is still a subsequence of s. 
// More formally, you will mark the character at s[removable[i]] for each 0 <= i < k, then remove all marked characters and check if p is still a subsequence.

// Return the maximum k you can choose such that p is still a subsequence of s after the removals.

// A subsequence of a string is a new string generated from the original string 
// with some characters (can be none) deleted without changing the relative order of the remaining characters.

// Example 1:
// Input: s = "abcacb", p = "ab", removable = [3,1,0]
// Output: 2
// Explanation: After removing the characters at indices 3 and 1, "abcacb" becomes "accb".
// "ab" is a subsequence of "accb".
// If we remove the characters at indices 3, 1, and 0, "abcacb" becomes "ccb", and "ab" is no longer a subsequence.
// Hence, the maximum k is 2.

// Example 2:
// Input: s = "abcbddddd", p = "abcd", removable = [3,2,1,4,5,6]
// Output: 1
// Explanation: After removing the character at index 3, "abcbddddd" becomes "abcddddd".
// "abcd" is a subsequence of "abcddddd".

// Example 3:
// Input: s = "abcab", p = "abc", removable = [0,1,2,3,4]
// Output: 0
// Explanation: If you remove the first index in the array removable, "abc" is no longer a subsequence.

// Constraints:
//     1 <= p.length <= s.length <= 105
//     0 <= removable.length < s.length
//     0 <= removable[i] < s.length
//     p is a subsequence of s.
//     s and p both consist of lowercase English letters.
//     The elements in removable are distinct.

import "fmt"

func maximumRemovals(s string, p string, removable []int) int {
    isSubSeq := func(remove map[int]bool) bool {
        i, j := 0, 0
        for i < len(s) && j < len(p) {
            if remove[i] == true || s[i] != p[j] {
                i++
                continue
            }
            i++
            j++
        }
        return j == len(p)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, l, r := 0, 0, len(removable)-1
    for l <= r {
        mid := l + (r - l) / 2
        remove := make(map[int]bool)
        for i := 0; i <= mid; i++ {
            remove[removable[i]] = true
        }
        if isSubSeq(remove) {
            l = mid + 1
            res = max(res, mid + 1)
        } else {
            r = mid - 1
        }
    }
    return res
}

func maximumRemovals1(s string, p string, removable []int) int {
    check := func (k int) bool {
        arr := []byte(s)
        for _, v := range removable[:k] {
            arr[v] = 0
        }
        i, j := 0, 0
        for i < len(arr) && j < len(p) {
            if arr[i] == p[j] { j++ }
            i++
        }
        return j == len(p)
    }
    l, r := 0, len(removable)
    for l <= r {
        mid := l + (r - l) >> 1
        // 求最大, check == true 时更新 l = mid + 1, 当check == false时 r = l - 1 刚好是答案
        // 求最小, check == true 时更新 r = mid - 1, 当check == false时 l = r + 1 刚好是答案
        if check(mid) {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return r
}

func maximumRemovals2(s string, p string, removable []int) int {
    n, bufs, bufp := len(removable), []byte(s), []byte(p)
    res, l, r := 0, 0, n
    isSubseq := func(s, p []byte) bool {
        index, n := 0, len(s)
        for i := range p {
            for index < n && s[index] != p[i] { index++ }
            if index >= n { return false }
            index++
        }
        return true
    }
    match := func(s, p []byte, nums []int, k int) bool {
        index := 0
        for i := 0; i < k; i++ {
            index = nums[i]
            s[index] = s[index] - 'a' + 'A'
        }
        res := isSubseq(s, p)
        for i := 0; i < k; i++ {
            index = nums[i]
            s[index] = s[index] - 'A' + 'a'
        }
        return res
    }
    for l <= r {
        mid := (l + r) / 2
        if match(bufs, bufp, removable, mid) {
            res = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcacb", p = "ab", removable = [3,1,0]
    // Output: 2
    // Explanation: After removing the characters at indices 3 and 1, "abcacb" becomes "accb".
    // "ab" is a subsequence of "accb".
    // If we remove the characters at indices 3, 1, and 0, "abcacb" becomes "ccb", and "ab" is no longer a subsequence.
    // Hence, the maximum k is 2.
    fmt.Println(maximumRemovals("abcacb", "ab", []int{3,1,0})) // 2
    // Example 2:
    // Input: s = "abcbddddd", p = "abcd", removable = [3,2,1,4,5,6]
    // Output: 1
    // Explanation: After removing the character at index 3, "abcbddddd" becomes "abcddddd".
    // "abcd" is a subsequence of "abcddddd".
    fmt.Println(maximumRemovals("abcbddddd", "abcd", []int{3,2,1,4,5,6})) // 1
    // Example 3:
    // Input: s = "abcab", p = "abc", removable = [0,1,2,3,4]
    // Output: 0
    // Explanation: If you remove the first index in the array removable, "abc" is no longer a subsequence.
    fmt.Println(maximumRemovals("abcab", "abc", []int{0,1,2,3,4})) // 1

    fmt.Println(maximumRemovals1("abcacb", "ab", []int{3,1,0})) // 2
    fmt.Println(maximumRemovals1("abcbddddd", "abcd", []int{3,2,1,4,5,6})) // 1
    fmt.Println(maximumRemovals1("abcab", "abc", []int{0,1,2,3,4})) // 1

    fmt.Println(maximumRemovals2("abcacb", "ab", []int{3,1,0})) // 2
    fmt.Println(maximumRemovals2("abcbddddd", "abcd", []int{3,2,1,4,5,6})) // 1
    fmt.Println(maximumRemovals2("abcab", "abc", []int{0,1,2,3,4})) // 1
}