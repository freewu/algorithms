package main

// 2423. Remove Letter To Equalize Frequency
// You are given a 0-indexed string word, consisting of lowercase English letters. 
// You need to select one index and remove the letter at that index from word so that the frequency of every letter present in word is equal.

// Return true if it is possible to remove one letter so that the frequency of all letters in word are equal, and false otherwise.

// Note:
//     The frequency of a letter x is the number of times it occurs in the string.
//     You must remove exactly one letter and cannot choose to do nothing.

// Example 1:
// Input: word = "abcc"
// Output: true
// Explanation:
// Select index 3 and delete it: word becomes "abc" and each character has a frequency of 1.

// Example 2:
// Input: word = "aazz"
// Output: false
// Explanation:
// We must delete a character, so either the frequency of "a" is 1 and the frequency of "z" is 2, or vice versa. 
// It is impossible to make all present letters have equal frequency.

// Constraints:
//     2 <= word.length <= 100
//     word consists of lowercase English letters only.

import "fmt"
import "sort"

func equalFrequency(word string) bool {
    freq := [26]int{}
    for _, v := range word {
        freq[v - 'a']++
    }
    for i := range freq {
        if freq[i] > 0 {
            freq[i]--
            x, flag := 0, true
            for _, v := range freq {
                if v == 0 { continue }
                if x > 0 && v != x {
                    flag = false
                    break
                }
                x = v
            }
            if flag { return true }
            freq[i]++
        }
    }
    return false
}

func equalFrequency1(word string) bool {
    freq := make([]int, 26)
    for _, v := range word {
        freq[v - 'a']++
    }
    sort.Ints(freq)
    for len(freq) > 0 && freq[0] == 0 {
        freq = freq[1:]
    }
    n := len(freq)
    if n == 1 { return true }
    return freq[n - 1] - 1 == freq[n - 2] && freq[0] == freq[n - 2] || freq[0] == 1 && freq[1] == freq[n - 1]
}

func main() {
    // Example 1:
    // Input: word = "abcc"
    // Output: true
    // Explanation:
    // Select index 3 and delete it: word becomes "abc" and each character has a frequency of 1.
    fmt.Println(equalFrequency("abcc")) // true
    // Example 2:
    // Input: word = "aazz"
    // Output: false
    // Explanation:
    // We must delete a character, so either the frequency of "a" is 1 and the frequency of "z" is 2, or vice versa. 
    // It is impossible to make all present letters have equal frequency.
    fmt.Println(equalFrequency("aazz")) // false

    fmt.Println(equalFrequency("bluefrog")) // true
    fmt.Println(equalFrequency("leetcode")) // false

    fmt.Println(equalFrequency1("abcc")) // true
    fmt.Println(equalFrequency1("aazz")) // false
    fmt.Println(equalFrequency1("bluefrog")) // true
    fmt.Println(equalFrequency1("leetcode")) // false
}