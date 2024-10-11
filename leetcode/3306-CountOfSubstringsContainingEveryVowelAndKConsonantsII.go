package main

// 3306. Count of Substrings Containing Every Vowel and K Consonants II
// You are given a string word and a non-negative integer k.

// Return the total number of substrings of word that contain every vowel ('a', 'e', 'i', 'o', and 'u') at least once and exactly k consonants.

// Example 1:
// Input: word = "aeioqq", k = 1
// Output: 0
// Explanation:
// There is no substring with every vowel.

// Example 2:
// Input: word = "aeiou", k = 0
// Output: 1
// Explanation:
// The only substring with every vowel and zero consonants is word[0..4], which is "aeiou".

// Example 3:
// Input: word = "ieaouqqieaouqq", k = 1
// Output: 3
// Explanation:
// The substrings with every vowel and one consonant are:
// word[0..5], which is "ieaouq".
// word[6..11], which is "qieaou".
// word[7..12], which is "ieaouq".

// Constraints:
//     5 <= word.length <= 2 * 10^5
//     word consists only of lowercase English letters.
//     0 <= k <= word.length - 5

import "fmt"
import "strings"

func countOfSubstrings(word string, k int) int64 {
    calc := func(w string, k int) int64 {
        res, consonants, n := 0, 0, len(w)
        vowels := make(map[byte]int)
        for l, r := 0, 0; r < n; r++ { // 枚举子串右端点
            if strings.ContainsRune("aeiou", rune(w[r])) {
                vowels[w[r]]++
            } else {
                consonants++
            }
            for len(vowels) == 5 && consonants >= k { // 每种元音字母都出现，才可能进入循环
                if strings.ContainsRune("aeiou", rune(w[l])) {
                    vowels[w[l]]--
                    if vowels[w[l]] == 0 {
                        delete(vowels, w[l])
                    }
                } else {
                    consonants--
                }
                l++
            }
            res += l
        }
        return int64(res)
    }
    return calc(word, k) - calc(word, k+1) // { >= k 的答案 } - { >= k+1 的答案 }
}

func countOfSubstrings1(s string, k int) int64 {
    const vowelMask = 1065233
    var cntVowel1, cntVowel2 ['u' - 'a' + 1]int
    sizeVowel1, sizeVowel2 := 0, 0 // 元音种类数
    cntConsonant1, cntConsonant2 := 0, 0
    res, left1, left2 := 0, 0, 0
    for _, b := range s {
        b -= 'a'
        if vowelMask>>b&1 > 0 {
            if cntVowel1[b] == 0 {
                sizeVowel1++
            }
            cntVowel1[b]++
            if cntVowel2[b] == 0 {
                sizeVowel2++
            }
            cntVowel2[b]++
        } else {
            cntConsonant1++
            cntConsonant2++
        }
        for sizeVowel1 == 5 && cntConsonant1 >= k {
            out := s[left1] - 'a'
            if vowelMask>>out&1 > 0 {
                cntVowel1[out]--
                if cntVowel1[out] == 0 {
                    sizeVowel1--
                }
            } else {
                cntConsonant1--
            }
            left1++
        }
        for sizeVowel2 == 5 && cntConsonant2 > k {
            out := s[left2] - 'a'
            if vowelMask>>out&1 > 0 {
                cntVowel2[out]--
                if cntVowel2[out] == 0 {
                    sizeVowel2--
                }
            } else {
                cntConsonant2--
            }
            left2++
        }
        res += left1 - left2
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: word = "aeioqq", k = 1
    // Output: 0
    // Explanation:
    // There is no substring with every vowel.
    fmt.Println(countOfSubstrings("aeioqq", 1)) // 0
    // Example 2:
    // Input: word = "aeiou", k = 0
    // Output: 1
    // Explanation:
    // The only substring with every vowel and zero consonants is word[0..4], which is "aeiou".
    fmt.Println(countOfSubstrings("aeiou", 0)) // 1
    // Example 3:
    // Input: word = "ieaouqqieaouqq", k = 1
    // Output: 3
    // Explanation:
    // The substrings with every vowel and one consonant are:
    // word[0..5], which is "ieaouq".
    // word[6..11], which is "qieaou".
    // word[7..12], which is "ieaouq".
    fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1)) // 3

    fmt.Println(countOfSubstrings1("aeioqq", 1)) // 0
    fmt.Println(countOfSubstrings1("aeiou", 0)) // 1
    fmt.Println(countOfSubstrings1("ieaouqqieaouqq", 1)) // 3
}