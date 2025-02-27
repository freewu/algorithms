package main

// 面试题 17.13. Re-Space LCCI
// Oh, no! You have accidentally removed all spaces, punctuation, and capitalization in a lengthy document. 
// A sentence like "I reset the computer. It still didn't boot!" became "iresetthecomputeritstilldidntboot". 
// You'll deal with the punctuation and capi­talization later; right now you need to re-insert the spaces. 
// Most of the words are in a dictionary but a few are not. 
// Given a dictionary (a list of strings) and the document (a string), design an algorithm to unconcatenate the document in a way that minimizes the number of unrecognized characters. 
// Return the number of unrecognized characters.

// Note: This problem is slightly different from the original one in the book.

// Example:
// Input: 
// dictionary = ["looked","just","like","her","brother"]
// sentence = "jesslookedjustliketimherbrother"
// Output:  7
// Explanation:  After unconcatenating, we got "jess looked just like tim her brother", which containing 7 unrecognized characters.

// Note:
//     0 <= len(sentence) <= 1000
//     The total number of characters in dictionary is less than or equal to 150000.
//     There are only lowercase letters in dictionary and sentence.

import "fmt"

func respace(dictionary []string, sentence string) int {
    n := len(sentence)
    dp := make([]int, n + 1)
    for i := 1; i < n + 1; i++ {
        dp[i] = 1 << 31
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for _, v := range dictionary {
            if i - len(v) + 1 >= 0 && sentence[i - len(v) + 1:i + 1] == v {
                dp[i + 1] = min(dp[i + 1], dp[i - len(v) + 1])
            }
        }
        dp[i + 1] = min(dp[i + 1], dp[i] + 1)
    }
    return dp[n]
}

func respace1(dictionary []string, sentence string) int {
    n := len(sentence)
    dp := make([]int, n+1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        dp[i] = dp[i - 1] +1
        for _, v := range dictionary {
            if i < len(v) || sentence[i - len(v):i] != v { continue }
            dp[i] = min(dp[i], dp[i - len(v)])
        }
    }
    return dp[n]
}

func main() {
    // Example:
    // Input: 
    // dictionary = ["looked","just","like","her","brother"]
    // sentence = "jesslookedjustliketimherbrother"
    // Output:  7
    // Explanation:  After unconcatenating, we got "jess looked just like tim her brother", which containing 7 unrecognized characters.
    fmt.Println(respace([]string{"looked","just","like","her","brother"}, "jesslookedjustliketimherbrother")) // 7

    fmt.Println(respace1([]string{"looked","just","like","her","brother"}, "jesslookedjustliketimherbrother")) // 7
}