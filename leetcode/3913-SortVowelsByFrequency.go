package main

// 3913. Sort Vowels by Frequency
// You are given a string s consisting of lowercase English characters.

// Rearrange only the vowels in the string so that they appear in non-increasing order of their frequency.

// If multiple vowels have the same frequency, order them by the position of their first occurrence in s.

// Return the modified string.

// Vowels are 'a', 'e', 'i', 'o', and 'u'.

// The frequency of a letter is the number of times it occurs in the string.

// Example 1:
// Input: s = "leetcode"
// Output: "leetcedo"
// Explanation:​​​​​​​
// Vowels in the string are ['e', 'e', 'o', 'e'] with frequencies: e = 3, o = 1.
// Sorting in non-increasing order of frequency and placing them back into the vowel positions results in "leetcedo".

// Example 2:
// Input: s = "aeiaaioooa"
// Output: "aaaaoooiie"
// Explanation:​​​​​​​
// Vowels in the string are ['a', 'e', 'i', 'a', 'a', 'i', 'o', 'o', 'o', 'a'] with frequencies: a = 4, o = 3, i = 2, e = 1.
// Sorting them in non-increasing order of frequency and placing them back into the vowel positions results in "aaaaoooiie".

// Example 3:
// Input: s = "baeiou"
// Output: "baeiou"
// Explanation:
// Each vowel appears exactly once, so all have the same frequency.
// Thus, they retain their relative order based on first occurrence, and the string remains unchanged.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters

import "fmt"
import "sort"

func sortVowels(s string) string {
    res := make([]byte, len(s))
    mp, first := map[byte]int{}, map[byte]int{}
    isVowel := func(r byte) bool { return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' }
    x, n := []byte{}, 0
    for i := 0; i < len(s); i++ {
        if isVowel(s[i]) {
            x = append(x, s[i])
        }
    }
    for i := 0; i < len(s); i++ {
        if isVowel(s[i]) {
            if _, ok := first[s[i]]; !ok {
                first[s[i]] = i
            }
        }
    }
    for i := 0; i < len(s); i++ {
        if isVowel(s[i]) {
            mp[s[i]]++
        }
    }
    sort.Slice(x, func(i, j int) bool {
        if mp[x[i]] == mp[x[j]] {
            return first[x[i]] < first[x[j]]
        }
        return mp[x[i]] > mp[x[j]]
    })
    for i := 0; i < len(s); i++ {
        if isVowel(s[i]) {
            res[i] = x[n]
            n++
            continue
        }
        res[i] = s[i]
    }
    return string(res)
}

func sortVowels1(s string) string {
    freq, first := make(map[rune]int), make(map[rune]int)
    isVowel := func(r rune) bool { return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' }
    for i, ch := range s {
        if !isVowel(ch) { continue }
        freq[ch]++
        if _, ok := first[ch]; !ok {
            first[ch] = i
        }
    }
    type Item struct {
        freq, index int
        ch rune
    }
    // Store freq and ch
    items := make([]Item, 0, len(freq))
    for k, v := range freq {
        items = append(items, Item{
            freq: v,
            ch: k,
            index: first[k],
        })
    }
    sort.Slice(items, func(i, j int) bool {
        if items[i].freq == items[j].freq {
            return items[i].index < items[j].index
        }
        return items[i].freq > items[j].freq
    })
    chars := []rune(s)
    curr := 0
    for i, ch := range chars {
        if isVowel(ch) {
            if items[curr].freq == 0 {
                curr++
            }
            chars[i] = items[curr].ch
            items[curr].freq--
        }
    }
    return string(chars)
}

func main() {
    // Example 1:
    // Input: s = "leetcode"
    // Output: "leetcedo"
    // Explanation:​​​​​​​
    // Vowels in the string are ['e', 'e', 'o', 'e'] with frequencies: e = 3, o = 1.
    // Sorting in non-increasing order of frequency and placing them back into the vowel positions results in "leetcedo".
    fmt.Println(sortVowels("leetcode")) // "leetcedo"
    // Example 2:
    // Input: s = "aeiaaioooa"
    // Output: "aaaaoooiie"
    // Explanation:​​​​​​​
    // Vowels in the string are ['a', 'e', 'i', 'a', 'a', 'i', 'o', 'o', 'o', 'a'] with frequencies: a = 4, o = 3, i = 2, e = 1.
    // Sorting them in non-increasing order of frequency and placing them back into the vowel positions results in "aaaaoooiie".
    fmt.Println(sortVowels("aeiaaioooa")) // "aaaaoooiie"
    // Example 3:
    // Input: s = "baeiou"
    // Output: "baeiou"
    // Explanation:
    // Each vowel appears exactly once, so all have the same frequency.
    // Thus, they retain their relative order based on first occurrence, and the string remains unchanged.
    fmt.Println(sortVowels("baeiou")) // "baeiou"

    fmt.Println(sortVowels("blueffrog")) // "blueffrog"
    fmt.Println(sortVowels("freewu")) // "freewu"

    fmt.Println(sortVowels1("leetcode")) // "leetcedo"
    fmt.Println(sortVowels1("aeiaaioooa")) // "aaaaoooiie"
    fmt.Println(sortVowels1("baeiou")) // "baeiou"
    fmt.Println(sortVowels1("blueffrog")) // "blueffrog"
    fmt.Println(sortVowels1("freewu")) // "freewu"
}