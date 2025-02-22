package main

// 3031. Minimum Time to Revert Word to Initial State II
// You are given a 0-indexed string word and an integer k.

// At every second, you must perform the following operations:
//     Remove the first k characters of word.
//     Add any k characters to the end of word.

// Note that you do not necessarily need to add the same characters that you removed. 
// However, you must perform both operations at every second.

// Return the minimum time greater than zero required for word to revert to its initial state.

// Example 1:
// Input: word = "abacaba", k = 3
// Output: 2
// Explanation: At the 1st second, we remove characters "aba" from the prefix of word, and add characters "bac" to the end of word. Thus, word becomes equal to "cababac".
// At the 2nd second, we remove characters "cab" from the prefix of word, and add "aba" to the end of word. Thus, word becomes equal to "abacaba" and reverts to its initial state.
// It can be shown that 2 seconds is the minimum time greater than zero required for word to revert to its initial state.

// Example 2:
// Input: word = "abacaba", k = 4
// Output: 1
// Explanation: At the 1st second, we remove characters "abac" from the prefix of word, and add characters "caba" to the end of word. Thus, word becomes equal to "abacaba" and reverts to its initial state.
// It can be shown that 1 second is the minimum time greater than zero required for word to revert to its initial state.

// Example 3:
// Input: word = "abcbabcd", k = 2
// Output: 4
// Explanation: At every second, we will remove the first 2 characters of word, and add the same characters to the end of word.
// After 4 seconds, word becomes equal to "abcbabcd" and reverts to its initial state.
// It can be shown that 4 seconds is the minimum time greater than zero required for word to revert to its initial state.

// Constraints:
//     1 <= word.length <= 10^6
//     1 <= k <= word.length
//     word consists only of lowercase English letters.

import "fmt"
import "strings"

// // Time Limit Exceeded 898 / 911 
// func minimumTimeToInitialState(word string, k int) int {
//     res := 1
//     for i := k ; i < len(word); i +=k {
//         if strings.HasPrefix(word, word[i:]) {
//             break
//         } 
//         res++
//     }
//     return res
// }

type Hashing struct {
    p   []int64
    h   []int64
    mod int64
}

func NewHashing(word string, base int64, mod int64) *Hashing {
    n := len(word)
    p, h := make([]int64, n + 1), make([]int64, n + 1)
    p[0] = 1
    for i := 1; i <= n; i++ {
        p[i] = (p[i - 1] * base) % mod
        h[i] = (h[i - 1] * base + int64(word[i-1] - 'a')) % mod
    }
    return &Hashing{p, h, mod}
}

func (t *Hashing) Query(l, r int) int64 {
    return (t.h[r] - t.h[l-1] * t.p[r-l+1] % t.mod + t.mod) % t.mod
}

func minimumTimeToInitialState(word string, k int) int {
    hashing := NewHashing(word, 13331, 998244353)
    n := len(word)
    for i := k; i < n; i += k {
        if hashing.Query(1, n - i) == hashing.Query(i + 1, n) {
            return i / k
        }
    }
    return (n + k - 1) / k
}

func minimumTimeToInitialState1(word string, k int) int {
    n, m := len(word), 0
    pi := make([]int, n)
    for i := 1; i < n; i++ {
        for m > 0 && word[m] != word[i] {
            m = pi[m-1]
        }
        if word[m] == word[i] {
            m++
        }
        pi[i] = m
    }
    mn, mx := (n - pi[n-1] + k - 1) / k, (n + k - 1) / k
    for mn < mx {
        if strings.HasPrefix(word, word[mn * k:]) {
            return mn
        }
        mn++
    }
    return mx
}

func main() {
    // Example 1:
    // Input: word = "abacaba", k = 3
    // Output: 2
    // Explanation: At the 1st second, we remove characters "aba" from the prefix of word, and add characters "bac" to the end of word. Thus, word becomes equal to "cababac".
    // At the 2nd second, we remove characters "cab" from the prefix of word, and add "aba" to the end of word. Thus, word becomes equal to "abacaba" and reverts to its initial state.
    // It can be shown that 2 seconds is the minimum time greater than zero required for word to revert to its initial state.
    fmt.Println(minimumTimeToInitialState("abacaba", 3)) // 2
    // Example 2:
    // Input: word = "abacaba", k = 4
    // Output: 1
    // Explanation: At the 1st second, we remove characters "abac" from the prefix of word, and add characters "caba" to the end of word. Thus, word becomes equal to "abacaba" and reverts to its initial state.
    // It can be shown that 1 second is the minimum time greater than zero required for word to revert to its initial state.
    fmt.Println(minimumTimeToInitialState("abacaba", 4)) // 1
    // Example 3:
    // Input: word = "abcbabcd", k = 2
    // Output: 4
    // Explanation: At every second, we will remove the first 2 characters of word, and add the same characters to the end of word.
    // After 4 seconds, word becomes equal to "abcbabcd" and reverts to its initial state.
    // It can be shown that 4 seconds is the minimum time greater than zero required for word to revert to its initial state.
    fmt.Println(minimumTimeToInitialState("abcbabcd", 2)) // 4

    fmt.Println(minimumTimeToInitialState("bluefrog", 2)) // 4
    fmt.Println(minimumTimeToInitialState("leetcode", 2)) // 4

    fmt.Println(minimumTimeToInitialState1("abacaba", 3)) // 2
    fmt.Println(minimumTimeToInitialState1("abacaba", 4)) // 1
    fmt.Println(minimumTimeToInitialState1("abcbabcd", 2)) // 4
    fmt.Println(minimumTimeToInitialState1("bluefrog", 2)) // 4
    fmt.Println(minimumTimeToInitialState1("leetcode", 2)) // 4
}