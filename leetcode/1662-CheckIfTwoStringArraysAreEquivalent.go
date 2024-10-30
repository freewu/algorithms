package main

// 1662. Check If Two String Arrays are Equivalent
// Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.

// A string is represented by an array if the array elements concatenated in order forms the string.

// Example 1:
// Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
// Output: true
// Explanation:
// word1 represents string "ab" + "c" -> "abc"
// word2 represents string "a" + "bc" -> "abc"
// The strings are the same, so return true.

// Example 2:
// Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
// Output: false

// Example 3:
// Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
// Output: true

// Constraints:
//     1 <= word1.length, word2.length <= 10^3
//     1 <= word1[i].length, word2[i].length <= 10^3
//     1 <= sum(word1[i].length), sum(word2[i].length) <= 10^3
//     word1[i] and word2[i] consist of lowercase letters.

import "fmt"
// import "sort"

// func arrayStringsAreEqual(word1 []string, word2 []string) bool {
//     sort.Strings(word1)
//     sort.Strings(word2)
//     str1, str2 := "", ""
//     for _, v := range word1 {
//         str1 += v
//     }
//     for _, v := range word2 {
//         str2 += v
//     }
//     return str1 == str2
// }

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    i,j, k,z := 0,0,0,0
    for i < len(word1) && j < len(word2) {
        if word1[i][k] != word2[j][z] { return false }
        k++
        z++
        if k >= len(word1[i]) {
            k = 0
            i++
        }
        if z >= len(word2[j]) {
            z = 0
            j++
        }
    }
    return i == len(word1) && j == len(word2)
}

func main() {
    // Example 1:
    // Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
    // Output: true
    // Explanation:
    // word1 represents string "ab" + "c" -> "abc"
    // word2 represents string "a" + "bc" -> "abc"
    // The strings are the same, so return true.
    fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"})) // true
    // Example 2:
    // Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
    // Output: false
    fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"})) // false
    // Example 3:
    // Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
    // Output: true
    fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"})) // true

    fmt.Println(arrayStringsAreEqual([]string{"jbboxe","yshcrtanrtlzyyp","vudsssnzuef","lde"}, []string{"jbboxeyshcrtanrt","lzyypvudsssnzueflde"})) // true
}