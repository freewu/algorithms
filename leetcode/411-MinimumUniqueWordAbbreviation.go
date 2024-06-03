package main

// 411. Minimum Unique Word Abbreviation
// A string can be abbreviated by replacing any number of non-adjacent substrings with their lengths. 
// For example, a string such as "substitution" could be abbreviated as (but not limited to):
//     "s10n" ("s ubstitutio n")
//     "sub4u4" ("sub stit u tion")
//     "12" ("substitution")
//     "su3i1u2on" ("su bst i t u ti on")
//     "substitution" (no substrings replaced)

// Note that "s55n" ("s ubsti tutio n") is not a valid abbreviation of "substitution" because the replaced substrings are adjacent.
// The length of an abbreviation is the number of letters that were not replaced plus the number of substrings that were replaced. 
// For example, the abbreviation "s10n" has a length of 3 (2 letters + 1 substring) and "su3i1u2on" has a length of 9 (6 letters + 3 substrings).

// Given a target string target and an array of strings dictionary, 
// return an abbreviation of target with the shortest possible length such that it is not an abbreviation of any string in dictionary. 
// If there are multiple shortest abbreviations, return any of them.

// Example 1:
// Input: target = "apple", dictionary = ["blade"]
// Output: "a4"
// Explanation: The shortest abbreviation of "apple" is "5", but this is also an abbreviation of "blade".
// The next shortest abbreviations are "a4" and "4e". "4e" is an abbreviation of blade while "a4" is not.
// Hence, return "a4".

// Example 2:
// Input: target = "apple", dictionary = ["blade","plain","amber"]
// Output: "1p3"
// Explanation: "5" is an abbreviation of both "apple" but also every word in the dictionary.
// "a4" is an abbreviation of "apple" but also "amber".
// "4e" is an abbreviation of "apple" but also "blade".
// "1p3", "2p2", and "3l1" are the next shortest abbreviations of "apple".
// Since none of them are abbreviations of words in the dictionary, returning any of them is correct.

// Constraints:
//     m == target.length
//     n == dictionary.length
//     1 <= m <= 21
//     0 <= n <= 1000
//     1 <= dictionary[i].length <= 100
//     log2(n) + m <= 21 if n > 0
//     target and dictionary[i] consist of lowercase English letters.
//     dictionary does not contain target.

import "fmt"
import "strconv"

func minAbbreviation(target string, dictionary []string) string {
    if 0 == len(dictionary) {
        return strconv.Itoa(len(target))
    }
    m := make([][]string, len(target))
    for i := 0; i < len(target)-1; i++ {
        m[i] = []string{}
    }
    var cal func(string, int, int)
    cal = func(s string, i, c int) {
        if i != len(target) {
            cal(s, i+1, c+1)
            if c > 0 {
                s += strconv.Itoa(c)
            }
            s += target[i : i+1]
            cal(s, i+1, 0)
        } else {
            if c > 0 {
                s += strconv.Itoa(c)
            }
            m[len(s)-1] = append(m[len(s)-1], s)
        }
    }
    cal("", 0, 0)
    isAbbr := func (word, abbr string) bool {
        j, c, lw := 0, 0, len(word)
        for _, b := range []byte(abbr) {
            if b >= '0' && b <= '9' {
                if b == '0' && c == 0 {
                    return false
                }
                c = c * 10 + int(b-'0')
            } else {
                j += c
                if j >= lw || b != word[j] {
                    return false
                }
                j, c = j + 1, 0
            }
        }
        return j + c == lw
    }
    for _, abbrs := range m {
        for _, abbr := range abbrs {
            i := 0
            for ; i < len(dictionary); i++ {
                if isAbbr(dictionary[i], abbr) {
                    break
                }
            }
            if i == len(dictionary) {
                return abbr
            }
        }
    }
    return ""
}

func main() {
    // Example 1:
    // Input: target = "apple", dictionary = ["blade"]
    // Output: "a4"
    // Explanation: The shortest abbreviation of "apple" is "5", but this is also an abbreviation of "blade".
    // The next shortest abbreviations are "a4" and "4e". "4e" is an abbreviation of blade while "a4" is not.
    // Hence, return "a4".
    fmt.Println(minAbbreviation("apple",[]string{"blade"})) // a4
    // Example 2:
    // Input: target = "apple", dictionary = ["blade","plain","amber"]
    // Output: "1p3"
    // Explanation: "5" is an abbreviation of both "apple" but also every word in the dictionary.
    // "a4" is an abbreviation of "apple" but also "amber".
    // "4e" is an abbreviation of "apple" but also "blade".
    // "1p3", "2p2", and "3l1" are the next shortest abbreviations of "apple".
    // Since none of them are abbreviations of words in the dictionary, returning any of them is correct.
    fmt.Println(minAbbreviation("apple",[]string{"blade","plain","amber"})) // 1p3
}