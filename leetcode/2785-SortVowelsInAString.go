package main

// 2785. Sort Vowels in a String
// Given a 0-indexed string s, permute s to get a new string t such that:
//     1. All consonants remain in their original places. 
//        More formally, if there is an index i with 0 <= i < s.length such that s[i] is a consonant, 
//        then t[i] = s[i].
//     2. The vowels must be sorted in the nondecreasing order of their ASCII values. 
//        More formally, for pairs of indices i, j with 0 <= i < j < s.length such that s[i] and s[j] are vowels, 
//        then t[i] must not have a higher ASCII value than t[j].

// Return the resulting string.

// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in lowercase or uppercase. 
// Consonants comprise all letters that are not vowels.

// Example 1:
// Input: s = "lEetcOde"
// Output: "lEOtcede"
// Explanation: 'E', 'O', and 'e' are the vowels in s; 'l', 't', 'c', and 'd' are all consonants. The vowels are sorted according to their ASCII values, and the consonants remain in the same places.

// Example 2:
// Input: s = "lYmpH"
// Output: "lYmpH"
// Explanation: There are no vowels in s (all characters in s are consonants), so we return "lYmpH".

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of letters of the English alphabet in uppercase and lowercase.

import "fmt"
import "strings"
import "sort"

func sortVowels(s string) string {
    res, vowels := []byte(s), []byte{}
    isVowel := func(b byte) bool { return strings.ContainsRune("aeiouAEIOU", rune(b)) }
    for i := 0; i < len(res); i++ {
        if isVowel(res[i]) {
            vowels = append(vowels, res[i])
        }
    }
    sort.Slice(vowels, func(i, j int) bool {
        return vowels[i] < vowels[j]
    })
    ptr := 0
    for i := 0; i < len(res); i++ {
        if isVowel(res[i]) {
            res[i] = vowels[ptr]
            ptr++
        }
    }
    return string(res)
}

func sortVowels1(s string) string {
    var dict []bool
    init := func() {
        dict = make([]bool, 58)
        for _, v := range "aeiouAEIOU" {
            dict[v - 'A'] = true
        }
    }
    init()
    index, count := make([]int, 0, len(s)), make([]int, 58)
    for i, v := range s {
        k := v - 'A'
        if !dict[k] { continue }
        index = append(index, i)
        count[k] += 1
    }
    if len(index) == 0 { return s } // 没有元音
    res, i := []byte(s), 0
    for k, v := range count {
        if v == 0 { continue }
        for j := 0; j < v; j++ {
            res[index[i]] = byte(k) + byte('A')
            i++
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "lEetcOde"
    // Output: "lEOtcede"
    // Explanation: 'E', 'O', and 'e' are the vowels in s; 'l', 't', 'c', and 'd' are all consonants. The vowels are sorted according to their ASCII values, and the consonants remain in the same places.
    fmt.Println(sortVowels("lEetcOde")) // lEOtcede
    // Example 2:
    // Input: s = "lYmpH"
    // Output: "lYmpH"
    // Explanation: There are no vowels in s (all characters in s are consonants), so we return "lYmpH".
    fmt.Println(sortVowels("lYmpH")) // lYmpH

    fmt.Println(sortVowels("leetcode")) // leetcedo
    fmt.Println(sortVowels("bluefrog")) // bleofrug

    fmt.Println(sortVowels1("lEetcOde")) // lEOtcede
    fmt.Println(sortVowels1("lYmpH")) // lYmpH
    fmt.Println(sortVowels1("leetcode")) // leetcedo
    fmt.Println(sortVowels1("bluefrog")) // bleofrug
}