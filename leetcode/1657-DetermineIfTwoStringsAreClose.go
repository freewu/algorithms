package main

// 1657. Determine if Two Strings Are Close
// Two strings are considered close if you can attain one from the other using the following operations:
//     Operation 1: Swap any two existing characters.
//         For example, abcde -> aecdb
//     Operation 2: Transform every occurrence of one existing character into another existing character, 
//     and do the same with the other character.
//         For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)

// You can use the operations on either string as many times as necessary.
// Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

// Example 1:
// Input: word1 = "abc", word2 = "bca"
// Output: true
// Explanation: You can attain word2 from word1 in 2 operations.
// Apply Operation 1: "abc" -> "acb"
// Apply Operation 1: "acb" -> "bca"

// Example 2:
// Input: word1 = "a", word2 = "aa"
// Output: false
// Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.

// Example 3:
// Input: word1 = "cabbba", word2 = "abbccc"
// Output: true
// Explanation: You can attain word2 from word1 in 3 operations.
// Apply Operation 1: "cabbba" -> "caabbb"
// Apply Operation 2: "caabbb" -> "baaccc"
// Apply Operation 2: "baaccc" -> "abbccc"

// Constraints:
//     1 <= word1.length, word2.length <= 10^5
//     word1 and word2 contain only lowercase English letters.

import "fmt"
import "sort"
import "slices"

func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) { // Check if the lengths of both words are equal
        return false
    }
    freq1, freq2 := make([]int, 26), make([]int, 26) // Create slices to store the frequency of characters for both words
    for _, char := range word1 { freq1[char-'a']++; } // Populate the frequency slices for word1
    for _, char := range word2 { freq2[char-'a']++; } // Populate the frequency slices for word2
    for i := 0; i < 26; i++ { // Check if the sets of characters (keys) in both words are the same
        if (freq1[i] > 0 && freq2[i] == 0) || (freq1[i] == 0 && freq2[i] > 0) { // word1 存在字符 不存在 word2 就不能做置换处理
            return false
        }
    }
    // Sort the frequency slices in ascending order
    sort.Ints(freq1)
    sort.Ints(freq2)
    for i := 0; i < 26; i++ { // Compare the sorted frequency slices
        if freq1[i] != freq2[i] { // 置换数量不能匹配
            return false
        }
    }
    return true
}

func closeStrings1(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false 
    }
    l1, l2, s1, s2 := make([]int, 26), make([]int, 26), 0 , 0
    for i := 0 ; i < len(word1) ; i ++ {
        l1[word1[i] -'a'] ++ 
        s1 |= 1 << (word1[i] -'a')
        l2[word2[i] -'a'] ++ 
        s2 |= 1 << (word2[i]-'a')
    }
    if s1 != s2 {
        return false 
    }
    slices.Sort(l1[:])
    slices.Sort(l2[:])
    return slices.Equal(l1[:] , l2[:])
}

func main() {
    // Example 1:
    // Input: word1 = "abc", word2 = "bca"
    // Output: true
    // Explanation: You can attain word2 from word1 in 2 operations.
    // Apply Operation 1: "abc" -> "acb"
    // Apply Operation 1: "acb" -> "bca"
    fmt.Println(closeStrings("abc","bca")) // true
    // Example 2:
    // Input: word1 = "a", word2 = "aa"
    // Output: false
    // Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
    fmt.Println(closeStrings("a","aa")) // false
    // Example 3:
    // Input: word1 = "cabbba", word2 = "abbccc"
    // Output: true
    // Explanation: You can attain word2 from word1 in 3 operations.
    // Apply Operation 1: "cabbba" -> "caabbb"
    // Apply Operation 2: "caabbb" -> "baaccc"
    // Apply Operation 2: "baaccc" -> "abbccc"
    fmt.Println(closeStrings("cabbba","abbccc")) // true

    fmt.Println(closeStrings1("abc","bca")) // true
    fmt.Println(closeStrings1("a","aa")) // false
    fmt.Println(closeStrings1("cabbba","abbccc")) // true
}