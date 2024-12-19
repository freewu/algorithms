package main

// 3138. Minimum Length of Anagram Concatenation
// You are given a string s, which is known to be a concatenation of anagrams of some string t.

// Return the minimum possible length of the string t.

// An anagram is formed by rearranging the letters of a string. 
// For example, "aab", "aba", and, "baa" are anagrams of "aab".

// Example 1:
// Input: s = "abba"
// Output: 2
// Explanation:
// One possible string t could be "ba".

// Example 2:
// Input: s = "cdef"
// Output: 4
// Explanation:
// One possible string t could be "cdef", notice that t can be equal to s.

// Constraints:
//     1 <= s.length <= 10^5
//     s consist only of lowercase English letters.

import "fmt"
import "slices"

// 解答错误 550 / 559 
// func minAnagramLength(s string) int {
//     n := len(s)
//     mp := make([]int, 26)
//     for _, v := range s {
//         mp[v - 'a']++
//     }
//     gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
//     freq := mp[s[0] - 'a']
//     for i := 0; i < 26; i++ {
//         if mp[i] == 0 { continue }
//         freq = gcd(freq, mp[i])
//     }
//     return n / freq
// }

func minAnagramLength(s string) int {
    n := len(s)
    count := [26]int{}
    for _, v := range s {
        count[v - 'a']++
    }
    check := func(k int) bool {
        for i := 0; i < n; i += k {
            count1 := [26]int{}
            for j := i; j < i + k; j++ {
                count1[s[j]-'a']++
            }
            for j, v := range count {
                if count1[j]*(n / k) != v {
                    return false
                }
            }
        }
        return true
    }
    for i := 1; ; i++ {
        if n % i == 0 && check(i) {
            return i
        }
    }
    return -1
}

func minAnagramLength1(s string) int {
    count := make([]int, 26)
    n := len(s)
    check := func(j int) bool {
        for i := j; i < n; i += j {
            cur := make([]int, 26)
            for i2 := i; i2 < i + j && i2 < n; i2++ {
                cur[s[i2]-'a']++
            }
            if slices.Compare(cur, count) != 0 {
                return false
            }
        }
        return true
    }
    for i, v :=  range s {
        count[v-'a']++
        if n % (i + 1) == 0 && check(i + 1) {
            return i + 1
        }
    }
    return n
}

func main() {
    // Example 1:
    // Input: s = "abba"
    // Output: 2
    // Explanation:
    // One possible string t could be "ba".
    fmt.Println(minAnagramLength("abba")) // 2
    // Example 2:
    // Input: s = "cdef"
    // Output: 4
    // Explanation:
    // One possible string t could be "cdef", notice that t can be equal to s.
    fmt.Println(minAnagramLength("cdef")) // 4

    fmt.Println(minAnagramLength("aabb")) // 4

    fmt.Println(minAnagramLength1("abba")) // 2
    fmt.Println(minAnagramLength1("cdef")) // 4
    fmt.Println(minAnagramLength1("aabb")) // 4
}