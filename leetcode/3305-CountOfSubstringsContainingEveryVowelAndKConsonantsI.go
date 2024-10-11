package main

// 3305. Count of Substrings Containing Every Vowel and K Consonants I
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
//     5 <= word.length <= 250
//     word consists only of lowercase English letters.
//     0 <= k <= word.length - 5

import "fmt"
import "strings"

// brute force
func countOfSubstrings(word string, k int) int {
    res, n := 0, len(word)
    for l := 0; l + 5 + k <= n; l++ { // 暴力枚举所有潜在合法子串的左端点
        vowels := map[byte]bool{} // 重置
        k1 := k
        for r := l; r < n && k1>=0; r++ { // k1 >= 0 剪枝优化
            if strings.ContainsRune("aeiou", rune(word[r])) {
                vowels[word[r]] = true
            } else {
                k1--
            }
            if len(vowels) == 5 && k1 == 0 {
                res++
            }
        }
    }
    return res
}

func countOfSubstrings1(word string, k int) int {
    wordMap, nextConsonant, lastConsonant := make(map[byte]int), make([]int, len(word)), len(word)
    for i := len(word)-1; i>=0; i-- {
        nextConsonant[i] = lastConsonant
        switch word[i] {
            case 'a', 'e', 'i','o', 'u':
                continue
            default:
                lastConsonant = i
        }
    }
    vowel, consonants := 0, 0
    res, left, right := 0, 0, 0
    for right < len(word) {
        switch word[right] {
        case 'a', 'e', 'i', 'o', 'u':
            wordMap[word[right]] = wordMap[word[right]]+1
            if wordMap[word[right]] == 1 {
                vowel++
            }
        default:
            consonants++
        }
        for left <= right && consonants > k {
            switch word[left] {
            case 'a', 'e', 'i', 'o', 'u':
                wordMap[word[left]] = wordMap[word[left]] - 1
                if wordMap[word[left]] == 0 {
                    vowel--
                }
            default:
                consonants--
            }
            left++
        }
        for left < right && consonants == k && vowel == 5 {
            res += (nextConsonant[right] - right)
            switch word[left] {
            case 'a', 'e', 'i', 'o', 'u':
                wordMap[word[left]] = wordMap[word[left]] - 1
                if wordMap[word[left]] == 0 {
                    vowel--
                }
            default:
                consonants--
            }
            left++
        }
        right++
    }
    return res
}

func countOfSubstrings2(word string, k int) int {
    res := 0
    pLeft, pRight := 0, 0
    m := make(map[byte]int)
    n := 0
    isVowel := func(ch byte) bool {
        return 'a' == ch || 'e' == ch || 'i' == ch || 'o' == ch || 'u' == ch ||
                'A' == ch || 'E' == ch || 'I' == ch || 'O' == ch || 'U' == ch
    }
    for pRight = 0; pRight < len(word); pRight++ {
        if isVowel(word[pRight]) {
            m[word[pRight]]++
            n++
        }
        returnOri := false
        oriL := 0
        oriM := make(map[byte]int)
        oriN := 0
        for m['a'] > 0 && m['e'] > 0 && m['i'] > 0 &&
            m['o'] > 0 && m['u'] > 0 {
            tmpN := pRight - pLeft + 1 - n
            if tmpN < k {
                break
            } else {
                if tmpN == k {
                    //fmt.Println(word[pLeft : pRight+1])
                    res++
                    
                    if returnOri == false {
                        returnOri = true
                        oriL = pLeft
                        oriM['a'] = m['a']
                        oriM['e'] = m['e']
                        oriM['i'] = m['i']
                        oriM['o'] = m['o']
                        oriM['u'] = m['u']
                        oriN = n
                    }
                }
                ch := word[pLeft]
                if ch == 'a' || ch == 'e' ||
                    ch == 'i' || ch == 'o' ||
                    ch == 'u' {
                    if m[ch] > 1 {
                        m[ch]--
                        n--
                        pLeft++
                    } else {
                        break
                    }
                } else {
                    if tmpN > k || k == 0 {
                        pLeft++
                    } else {
                        break
                    }
                }
            }
        }
        if returnOri == true {
            pLeft = oriL
            m = oriM
            n = oriN
        }
    }
    return res
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

    fmt.Println(countOfSubstrings2("aeioqq", 1)) // 0
    fmt.Println(countOfSubstrings2("aeiou", 0)) // 1
    fmt.Println(countOfSubstrings2("ieaouqqieaouqq", 1)) // 3
}