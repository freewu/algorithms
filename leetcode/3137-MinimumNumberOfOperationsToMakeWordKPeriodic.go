package main

// 3137. Minimum Number of Operations to Make Word K-Periodic
// You are given a string word of size n, and an integer k such that k divides n.

// In one operation, you can pick any two indices i and j, that are divisible by k, 
// then replace the substring of length k starting at i with the substring of length k starting at j. 
// That is, replace the substring word[i..i + k - 1] with the substring word[j..j + k - 1].

// Return the minimum number of operations required to make word k-periodic.

// We say that word is k-periodic if there is some string s of length k such that word can be obtained by concatenating s an arbitrary number of times. F、
// or example, if word == “ababab”, then word is 2-periodic for s = "ab".

// Example 1:
// Input: word = "leetcodeleet", k = 4
// Output: 1
// Explanation:
// We can obtain a 4-periodic string by picking i = 4 and j = 0. After this operation, word becomes equal to "leetleetleet".

// Example 2:
// Input: word = "leetcoleet", k = 2
// Output: 3
// Explanation:
// We can obtain a 2-periodic string by applying the operations in the table below.
// i	j	word
// 0	2	etetcoleet
// 4	0	etetetleet
// 6	0	etetetetet

// Constraints:
//     1 <= n == word.length <= 10^5
//     1 <= k <= word.length
//     k divides word.length.
//     word consists only of lowercase English letters.

import "fmt"

func minimumOperationsToMakeKPeriodic(word string, k int) int {
    mp, n := map[string]int{}, len(word)
    for i := 0; i < n; i+= k {
        mp[word[i:i+k]]++
    }
    res := n/k - 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range mp {
        res = min(res, n/k - v)
    }
    return res
}

func minimumOperationsToMakeKPeriodic1(word string, k int) int {
    n, mp := len(word), map[string]int{}
    for i := 0; i < n; i += k {
        mp[word[i:i+k]]++
    }
    mx := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, c := range mp {
        mx = max(mx, c)
    }
    return n/k - mx
}

func main() {
    // Example 1:
    // Input: word = "leetcodeleet", k = 4
    // Output: 1
    // Explanation:
    // We can obtain a 4-periodic string by picking i = 4 and j = 0. After this operation, word becomes equal to "leetleetleet".
    fmt.Println(minimumOperationsToMakeKPeriodic("leetcodeleet", 4)) // 1
    // Example 2:
    // Input: word = "leetcoleet", k = 2
    // Output: 3
    // Explanation:
    // We can obtain a 2-periodic string by applying the operations in the table below.
    // i	j	word
    // 0	2	etetcoleet
    // 4	0	etetetleet
    // 6	0	etetetetet
    fmt.Println(minimumOperationsToMakeKPeriodic("leetcoleet", 2)) // 3

    fmt.Println(minimumOperationsToMakeKPeriodic1("leetcodeleet", 4)) // 1
    fmt.Println(minimumOperationsToMakeKPeriodic1("leetcoleet", 2)) // 3
}