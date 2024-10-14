package main

// 1433. Check If a String Can Break Another String
// Given two strings: s1 and s2 with the same size, check if some permutation of string s1 can break some permutation of string s2 or vice-versa. In other words s2 can break s1 or vice-versa.

// A string x can break string y (both of size n) if x[i] >= y[i] (in alphabetical order) for all i between 0 and n-1.

// Example 1:
// Input: s1 = "abc", s2 = "xya"
// Output: true
// Explanation: "ayx" is a permutation of s2="xya" which can break to string "abc" which is a permutation of s1="abc".

// Example 2:
// Input: s1 = "abe", s2 = "acd"
// Output: false 
// Explanation: All permutations for s1="abe" are: "abe", "aeb", "bae", "bea", "eab" and "eba" and all permutation for s2="acd" are: "acd", "adc", "cad", "cda", "dac" and "dca". However, there is not any permutation from s1 which can break some permutation from s2 and vice-versa.

// Example 3:
// Input: s1 = "leetcodee", s2 = "interview"
// Output: true

// Constraints:
//     s1.length == n
//     s2.length == n
//     1 <= n <= 10^5
//     All strings consist of lowercase English letters.

import "fmt"
import "sort"

func checkIfCanBreak(s1 string, s2 string) bool {
    b1, b2 := []byte(s1), []byte(s2)
    a, b := 0, 0
    sort.Slice(b1, func(a, b int) bool { return b1[a] < b1[b] })
    sort.Slice(b2, func(a, b int) bool { return b2[a] < b2[b] })
    for i := 0; i < len(b1); i++ {
        if b1[i] <= b2[i] { a++ }
        if b1[i] >= b2[i] { b++ }
    }
    if a == len(b1) || b == len(b2) { return true }
    return false
}

func checkIfCanBreak1(s1 string, s2 string) bool {
    count1, count2 := make([]int, 26), make([]int, 26)
    for _, v := range s1 { count1[v - 'a']++ }
    for _, v := range s2 { count2[v - 'a']++ }
    cmp := func(count1, count2 []int) bool {
        n := 0
        for i := 0; i < 26; i++ {
            n += count1[i] - count2[i]
            if n < 0 { return false }
        }
        return true
    }
    return cmp(count1, count2) || cmp(count2, count1)
}

func main() {
    // Example 1:
    // Input: s1 = "abc", s2 = "xya"
    // Output: true
    // Explanation: "ayx" is a permutation of s2="xya" which can break to string "abc" which is a permutation of s1="abc".
    fmt.Println(checkIfCanBreak("abc", "xya")) // true
    // Example 2:
    // Input: s1 = "abe", s2 = "acd"
    // Output: false 
    // Explanation: All permutations for s1="abe" are: "abe", "aeb", "bae", "bea", "eab" and "eba" and all permutation for s2="acd" are: "acd", "adc", "cad", "cda", "dac" and "dca". However, there is not any permutation from s1 which can break some permutation from s2 and vice-versa.
    fmt.Println(checkIfCanBreak("abe", "acd")) // false
    // Example 3:
    // Input: s1 = "leetcodee", s2 = "interview"
    // Output: true
    fmt.Println(checkIfCanBreak("leetcodee", "interview")) // true

    fmt.Println(checkIfCanBreak1("abc", "xya")) // true
    fmt.Println(checkIfCanBreak1("abe", "acd")) // false
    fmt.Println(checkIfCanBreak1("leetcodee", "interview")) // true
}