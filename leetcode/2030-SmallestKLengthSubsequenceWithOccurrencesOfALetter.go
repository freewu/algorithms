package main

// 2030. Smallest K-Length Subsequence With Occurrences of a Letter
// You are given a string s, an integer k, a letter letter, and an integer repetition.

// Return the lexicographically smallest subsequence of s of length k 
// that has the letter letter appear at least repetition times. 
// The test cases are generated so that the letter appears in s at least repetition times.

// A subsequence is a string that can be derived from another string by deleting some 
// or no characters without changing the order of the remaining characters.

// A string a is lexicographically smaller than a string b if in the first position where a and b differ, 
// string a has a letter that appears earlier in the alphabet than the corresponding letter in b.

// Example 1:
// Input: s = "leet", k = 3, letter = "e", repetition = 1
// Output: "eet"
// Explanation: There are four subsequences of length 3 that have the letter 'e' appear at least 1 time:
// - "lee" (from "leet")
// - "let" (from "leet")
// - "let" (from "leet")
// - "eet" (from "leet")
// The lexicographically smallest subsequence among them is "eet".

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/13/smallest-k-length-subsequence.png" />
// Input: s = "leetcode", k = 4, letter = "e", repetition = 2
// Output: "ecde"
// Explanation: "ecde" is the lexicographically smallest subsequence of length 4 that has the letter "e" appear at least 2 times.

// Example 3:
// Input: s = "bb", k = 2, letter = "b", repetition = 2
// Output: "bb"
// Explanation: "bb" is the only subsequence of length 2 that has the letter "b" appear at least 2 times.

// Constraints:
//     1 <= repetition <= k <= s.length <= 5 * 10^4
//     s consists of lowercase English letters.
//     letter is a lowercase English letter, and appears in s at least repetition times.

import "fmt"
import "strings"

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
    n := len(s)
    suffix := make([]int, n)
    if s[n - 1] == letter {
        suffix[n - 1] = 1
    }
    for i := n - 2; i >= 0; i-- {
        suffix[i] = suffix[i+1]
        if s[i] == letter {
            suffix[i]++
        }
    }
    stack, curr := []byte{}, 0
    for i := 0; i < n; i++ {
        for len(stack) + n - i - 1 >= k && len(stack) > 0 && stack[len(stack) - 1] > s[i] {
            if stack[len(stack) - 1] == letter {
                if curr - 1 + suffix[i] < repetition { break }
                stack = stack[:len(stack) - 1]
                curr--
            } else if curr + suffix[i] >= repetition {
                stack = stack[:len(stack)-1]
            } else {
                break
            }
        }
        if len(stack) < k {
            if s[i] == letter {
                curr++
                stack = append(stack, s[i])
            } else if repetition - curr <= 0 || repetition - curr <= k-len(stack) - 1 {
                stack = append(stack, s[i])
            }
        }
    }
    return string(stack)
}

func smallestSubsequence1(s string, k int, letter byte, repetition int) string {
    res, queue := []byte{}, []byte{}
    unread, in := strings.Count(s, string(letter)), 0
    for i, n := 0,len(s); len(res) < k;i++ {
        for i < n && len(queue) > 0 && s[i] < queue[len(queue) - 1] {
            if queue[len(queue) - 1] == letter {
                if in + unread <= repetition { break }
                in--
            }
            queue = queue[:len(queue) - 1]
        }
        if i < n {
            if s[i] == letter{
                unread--
                in++
            }
            queue = append(queue, s[i])
        }
        if i >= n - k {
            if queue[0] == letter {
                in--
                repetition--
                res = append(res, queue[0])
            }else if len(res) + repetition < k {
                res = append(res, queue[0])
            }
            queue = queue[1:]
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "leet", k = 3, letter = "e", repetition = 1
    // Output: "eet"
    // Explanation: There are four subsequences of length 3 that have the letter 'e' appear at least 1 time:
    // - "lee" (from "leet")
    // - "let" (from "leet")
    // - "let" (from "leet")
    // - "eet" (from "leet")
    // The lexicographically smallest subsequence among them is "eet".
    fmt.Println(smallestSubsequence("leet", 2, 'e', 1)) // "eet"
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/09/13/smallest-k-length-subsequence.png" />
    // Input: s = "leetcode", k = 4, letter = "e", repetition = 2
    // Output: "ecde"
    // Explanation: "ecde" is the lexicographically smallest subsequence of length 4 that has the letter "e" appear at least 2 times.
    fmt.Println(smallestSubsequence("leetcode", 4, 'e', 2)) // "ecde"
    // Example 3:
    // Input: s = "bb", k = 2, letter = "b", repetition = 2
    // Output: "bb"
    // Explanation: "bb" is the only subsequence of length 2 that has the letter "b" appear at least 2 times.
    fmt.Println(smallestSubsequence("bb", 2, 'b', 2)) // "bb"

    fmt.Println(smallestSubsequence1("leet", 2, 'e', 1)) // "eet"
    fmt.Println(smallestSubsequence1("leetcode", 4, 'e', 2)) // "ecde"
    fmt.Println(smallestSubsequence1("bb", 2, 'b', 2)) // "bb"
}