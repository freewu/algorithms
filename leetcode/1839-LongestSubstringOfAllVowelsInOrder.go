package main

// 1839. Longest Substring Of All Vowels in Order
// A string is considered beautiful if it satisfies the following conditions:
//     1. Each of the 5 English vowels ('a', 'e', 'i', 'o', 'u') must appear at least once in it.
//     2. The letters must be sorted in alphabetical order (i.e. all 'a's before 'e's, all 'e's before 'i's, etc.).

// For example, strings "aeiou" and "aaaaaaeiiiioou" are considered beautiful, but "uaeio", "aeoiu", and "aaaeeeooo" are not beautiful.

// Given a string word consisting of English vowels, return the length of the longest beautiful substring of word. 
// If no such substring exists, return 0.

// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
// Output: 13
// Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of length 13.

// Example 2:
// Input: word = "aeeeiiiioooauuuaeiou"
// Output: 5
// Explanation: The longest beautiful substring in word is "aeiou" of length 5.

// Example 3:
// Input: word = "a"
// Output: 0
// Explanation: There is no beautiful substring, so return 0.

// Constraints:
//     1 <= word.length <= 5 * 10^5
//     word consists of characters 'a', 'e', 'i', 'o', and 'u'.

import "fmt"

func longestBeautifulSubstring(word string) int {
    vowels := []byte{'a','e','i','o','u'}
    res, n, needed := 0, len(word), 0 // current needed vowel inside window
    left, right := 0, 0 // left and right pointers
    for right = 0; right < n; right++ {
        right = left
        for right < n && word[right] != vowels[needed] { // move right pointer until we found 'a' vowel
            right++
        }
        left = right // set left pointer to the first appearance of 'a'
        for right < n - 1 && word[right] == vowels[needed] && needed < len(vowels) { // now we moving right pointer until vowels order is correct
            if word[right] != word[right + 1] && needed + 1 < len(vowels) && word[right + 1] == vowels[needed + 1] { // if next vowel is different so we move 'needed' pointer to the next searching value
                needed++
            }
            right++
        }
        if needed >= len(vowels) - 1 { // we calculate longest substring only if we had met all vowels
            if right == len(word)-1 && word[right] == 'u' { // don't forget about edge case when last letter could be the last vowel
                res = max(res, right - left + 1)
            } else {
                res = max(res, right - left)
            }
        }
        needed, left = 0, right // reset needed && move left pointer
    }
    return res
}

func longestBeautifulSubstring1(word string) int {
    vowel := "aeiou"
    res, cur, sum := 0, 0, 0
    for i, n := 0, len(word); i < n; {
        start, ch := i, word[i]
        for i < n && word[i] == ch { i++ }
        if ch != vowel[cur] {
            cur, sum = 0, 0
            if ch != vowel[0] { continue } // 
        }
        sum += (i - start)
        cur++
        if cur == 5 {
            if sum > res {
                res = sum
            }
            cur, sum = 0, 0
        }
    }
    return res
}

func longestBeautifulSubstring2(word string) int {
    res, unique, curr, n := 0, 1, 1, len(word)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        if word[i] == word[i-1] {
            curr++
        } else if word[i] > word[i - 1] {
            curr++
            unique++
        } else {
            curr = 1
            unique = 1
        }
        if unique == 5 {
            res = max(res, curr)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
    // Output: 13
    // Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of length 13.
    fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu")) // 13
    // Example 2:
    // Input: word = "aeeeiiiioooauuuaeiou"
    // Output: 5
    // Explanation: The longest beautiful substring in word is "aeiou" of length 5.
    fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou")) // 5
    // Example 3:
    // Input: word = "a"
    // Output: 0
    // Explanation: There is no beautiful substring, so return 0.
    fmt.Println(longestBeautifulSubstring("a")) // 0

    fmt.Println(longestBeautifulSubstring1("aeiaaioaaaaeiiiiouuuooaauuaeiu")) // 13
    fmt.Println(longestBeautifulSubstring1("aeeeiiiioooauuuaeiou")) // 5
    fmt.Println(longestBeautifulSubstring1("a")) // 0

    fmt.Println(longestBeautifulSubstring2("aeiaaioaaaaeiiiiouuuooaauuaeiu")) // 13
    fmt.Println(longestBeautifulSubstring2("aeeeiiiioooauuuaeiou")) // 5
    fmt.Println(longestBeautifulSubstring2("a")) // 0
}