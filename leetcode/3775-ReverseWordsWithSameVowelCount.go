package main

// 3775. Reverse Words With Same Vowel Count
// You are given a string s consisting of lowercase English words, each separated by a single space.

// Determine how many vowels appear in the first word. 
// Then, reverse each following word that has the same vowel count. 
// Leave all remaining words unchanged.

// Return the resulting string.

// Vowels are 'a', 'e', 'i', 'o', and 'u'.

// Example 1:
// Input: s = "cat and mice"
// Output: "cat dna mice"
// Explanation:​​​​​​​
// The first word "cat" has 1 vowel.
// "and" has 1 vowel, so it is reversed to form "dna".
// "mice" has 2 vowels, so it remains unchanged.
// Thus, the resulting string is "cat dna mice".

// Example 2:
// Input: s = "book is nice"
// Output: "book is ecin"
// Explanation:
// The first word "book" has 2 vowels.
// "is" has 1 vowel, so it remains unchanged.
// "nice" has 2 vowels, so it is reversed to form "ecin".
// Thus, the resulting string is "book is ecin".

// Example 3:
// Input: s = "banana healthy"
// Output: "banana healthy"
// Explanation:
// The first word "banana" has 3 vowels.
// "healthy" has 2 vowels, so it remains unchanged.
// Thus, the resulting string is "banana healthy".

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters and spaces.
//     Words in s are separated by a single space.
//     s does not contain leading or trailing spaces.

import "fmt"
import "strings"
import "slices"

func reverseWords(s string) string {
    arr := strings.Split(s, " ")
    countVowel := func(s string) int { // 统计元音个数
        res := 0
        for _, c := range s {
            if strings.IndexRune("aeiou", c) >= 0 {
                res++
            }
        }
        return res
    }
    count := countVowel(arr[0]) // 第一个单词的元音数量
    for i := 1; i < len(arr); i++ {
        if countVowel(arr[i]) == count { // 如果它们的元音字母数与第一个单词相同，则将它们 反转
            t := []byte(arr[i])
            slices.Reverse(t)
            arr[i] = string(t)
        }
    }
    return strings.Join(arr, " ")
}

func reverseWords1(s string) string {
    s += " "
    count, index, n := 0, 0, len(s)
    res := make([]byte, 0, n)    
    isVowel := func(c byte) bool { return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' }
    for ;; index++ {
        ch := s[index]
        if isVowel(ch) {
            count++
        }
        if ch == ' ' {
            break
        }
        res = append(res, ch)
    }
    count2, old := 0, index
    for ; index < n; index++ {
        ch := s[index]
        if ch == ' ' {
            if count2 == count {
                for j := index-1; j >= old; j-- {
                    res = append(res, s[j])
                }
            } else {
                res = append(res, []byte(s[old:index])...)
            }
            res = append(res, ' ')
            count2 = 0
            old = index+1
        } else {
            if isVowel(ch) {
                count2++
            }
        }
    }
    return string(res[:n-1])
}

func main() {
    // Example 1:
    // Input: s = "cat and mice"
    // Output: "cat dna mice"
    // Explanation:​​​​​​​
    // The first word "cat" has 1 vowel.
    // "and" has 1 vowel, so it is reversed to form "dna".
    // "mice" has 2 vowels, so it remains unchanged.
    // Thus, the resulting string is "cat dna mice".
    fmt.Println(reverseWords("cat and mice")) // "cat dna mice"
    // Example 2:
    // Input: s = "book is nice"
    // Output: "book is ecin"
    // Explanation:
    // The first word "book" has 2 vowels.
    // "is" has 1 vowel, so it remains unchanged.
    // "nice" has 2 vowels, so it is reversed to form "ecin".
    // Thus, the resulting string is "book is ecin".
    fmt.Println(reverseWords("book is nice")) // "book is ecin"
    // Example 3:
    // Input: s = "banana healthy"
    // Output: "banana healthy"
    // Explanation:
    // The first word "banana" has 3 vowels.
    // "healthy" has 2 vowels, so it remains unchanged.
    // Thus, the resulting string is "banana healthy".
    fmt.Println(reverseWords("banana healthy")) // "banana healthy"

    fmt.Println(reverseWords("leetcode bluefrog")) // "leetcode bluefrog"

    fmt.Println(reverseWords1("cat and mice")) // "cat dna mice"
    fmt.Println(reverseWords1("book is nice")) // "book is ecin"
    fmt.Println(reverseWords1("banana healthy")) // "banana healthy"
    fmt.Println(reverseWords1("leetcode bluefrog")) // "leetcode bluefrog"
}