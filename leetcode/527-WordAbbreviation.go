package main

// 527. Word Abbreviation
// Given an array of distinct strings words, return the minimal possible abbreviations for every word.
// The following are the rules for a string abbreviation:
//     1. The initial abbreviation for each word is: the first character, then the number of characters in between, followed by the last character.
//     2. If more than one word shares the same abbreviation, then perform the following operation:
//         2.1 Increase the prefix (characters in the first part) of each of their abbreviations by 1.
//             For example, say you start with the words ["abcdef","abndef"] both initially abbreviated as "a4f". 
//             Then, a sequence of operations would be ["a4f","a4f"] -> ["ab3f","ab3f"] -> ["abc2f","abn2f"].
//         2.2 This operation is repeated until every abbreviation is unique.
//     3. At the end, if an abbreviation did not make a word shorter, then keep it as the original word.
 
// Example 1:
// Input: words = ["like","god","internal","me","internet","interval","intension","face","intrusion"]
// Output: ["l2e","god","internal","me","i6t","interval","inte4n","f2e","intr4n"]

// Example 2:
// Input: words = ["aa","aaa"]
// Output: ["aa","aaa"]

// Constraints:
//     1 <= words.length <= 400
//     2 <= words[i].length <= 400
//     words[i] consists of lowercase English letters.
//     All the strings of words are unique.

import "fmt"
import "strconv"

// 如果有相同key的，用map分组，然后同时增加组内的所有词的前缀，每次加一，直到组内互不相同
// 期间检查是否短于原来的词长
func wordsAbbreviation(words []string) []string {
    res, prefix := make([]string, len(words)), make([]int, len(words))
    for i, w := range words {
        res[i] = w[:1] + strconv.Itoa(len(w) - 2) + w[len(w) - 1:]
        prefix[i] = 1
        if len(res[i]) >= len(words[i]) {
            res[i] = words[i]
        }
    }
    changed := true
    for changed {
        changed = false
        m := make(map[string][]int)
        for i, w := range res {
            if len(res[i]) >= len(words[i]) {
                res[i] = words[i]
            }
            if _, e := m[w]; e {
                m[w] = append(m[w], i)
                changed = true
            } else {
                m[w] = make([]int, 0)
                m[w] = append(m[w], i)
            }
        }
        for _, v := range m {
            if len(v) > 1 {
                for _, index := range v {
                    prefix[index] ++
                    res[index] = words[index][:prefix[index]] + strconv.Itoa(len(words[index]) - 1 - prefix[index]) + words[index][len(words[index]) - 1:]
                    if len(res[index]) >= len(words[index]) {
                        res[index] = words[index]
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["like","god","internal","me","internet","interval","intension","face","intrusion"]
    // Output: ["l2e","god","internal","me","i6t","interval","inte4n","f2e","intr4n"]
    fmt.Println(wordsAbbreviation([]string{"like","god","internal","me","internet","interval","intension","face","intrusion"})) // ["l2e","god","internal","me","i6t","interval","inte4n","f2e","intr4n"]
    // Example 2:
    // Input: words = ["aa","aaa"]
    // Output: ["aa","aaa"]
    fmt.Println(wordsAbbreviation([]string{"aa","aaa"})) // ["aa","aaa"]

    fmt.Println(wordsAbbreviation([]string{"bluefrog","leetcode"})) // [b6g l6e]
}