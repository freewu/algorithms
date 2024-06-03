package main

// 418. Sentence Screen Fitting
// Given a rows x cols screen and a sentence represented as a list of strings, 
// return the number of times the given sentence can be fitted on the screen.

// The order of words in the sentence must remain unchanged, 
// and a word cannot be split into two lines.
// A single space must separate two consecutive words in a line.

// Example 1:
// Input: sentence = ["hello","world"], rows = 2, cols = 8
// Output: 1
// Explanation:
// hello---
// world---
// The character '-' signifies an empty space on the screen.

// Example 2:
// Input: sentence = ["a", "bcd", "e"], rows = 3, cols = 6
// Output: 2
// Explanation:
// a-bcd- 
// e-a---
// bcd-e-
// The character '-' signifies an empty space on the screen.

// Example 3:
// Input: sentence = ["i","had","apple","pie"], rows = 4, cols = 5
// Output: 1
// Explanation:
// i-had
// apple
// pie-i
// had--
// The character '-' signifies an empty space on the screen.

// Constraints:
//     1 <= sentence.length <= 100
//     1 <= sentence[i].length <= 10
//     sentence[i] consists of lowercase English letters.
//     1 <= rows, cols <= 2 * 10^4

import "fmt"

// 模拟
func wordsTyping(sentence []string, rows int, cols int) int {
    res := 0
    for i, j := 0, cols; rows > 0; {
        if len(sentence[i]) <= j {
            j -= len(sentence[i]) + 1
            i++
        } else {
            rows--
            j = cols
        }
        if i == len(sentence) {
            res++
            i = 0
        }
    }
    return res
}

// dp
func wordsTyping1(sentence []string, rows int, cols int) int {
    res, dp := 0, make([][]int, len(sentence))
    for i := range sentence {
        wordIdx, sumLen, cnt := i, len(sentence[i]), 0
        for sumLen <= cols {
            wordIdx++
            if wordIdx == len(sentence) {
                cnt++
                wordIdx = 0
            }
            sumLen += len(sentence[wordIdx]) + 1
        }
        dp[i] = []int{cnt, wordIdx}
    }
    for i, idx := 0, 0; i < rows; i++ {
        res, idx = res + dp[idx][0], dp[idx][1]
    }
    return res
}

func main() {
    // Example 1:
    // Input: sentence = ["hello","world"], rows = 2, cols = 8
    // Output: 1
    // Explanation:
    // hello---
    // world---
    // The character '-' signifies an empty space on the screen.
    fmt.Println(wordsTyping([]string{"hello","world"},2,8)) // 1
    // Example 2: 
    // Input: sentence = ["a", "bcd", "e"], rows = 3, cols = 6
    // Output: 2
    // Explanation:
    // a-bcd- 
    // e-a---
    // bcd-e-
    // The character '-' signifies an empty space on the screen.
    fmt.Println(wordsTyping([]string{"a", "bcd", "e"},3,6)) // 2
    // Example 3:
    // Input: sentence = ["i","had","apple","pie"], rows = 4, cols = 5
    // Output: 1
    // Explanation:
    // i-had
    // apple
    // pie-i
    // had--
    // The character '-' signifies an empty space on the screen.
    fmt.Println(wordsTyping([]string{"i","had","apple","pie"},4,5)) // 1

    fmt.Println(wordsTyping1([]string{"hello","world"},2,8)) // 1
    fmt.Println(wordsTyping1([]string{"a", "bcd", "e"},3,6)) // 2
    fmt.Println(wordsTyping1([]string{"i","had","apple","pie"},4,5)) // 1
}