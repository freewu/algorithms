package main

// 1525. Number of Good Ways to Split a String
// You are given a string s.

// A split is called good if you can split s into two non-empty strings sleft 
// and sright where their concatenation is equal to s (i.e., sleft + sright = s) 
// and the number of distinct letters in sleft and sright is the same.

// Return the number of good splits you can make in s.

// Example 1:
// Input: s = "aacaba"
// Output: 2
// Explanation: There are 5 ways to split "aacaba" and 2 of them are good. 
// ("a", "acaba") Left string and right string contains 1 and 3 different letters respectively.
// ("aa", "caba") Left string and right string contains 1 and 3 different letters respectively.
// ("aac", "aba") Left string and right string contains 2 and 2 different letters respectively (good split).
// ("aaca", "ba") Left string and right string contains 2 and 2 different letters respectively (good split).
// ("aacab", "a") Left string and right string contains 3 and 1 different letters respectively.

// Example 2:
// Input: s = "abcd"
// Output: 1
// Explanation: Split the string as follows ("ab", "cd").

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of only lowercase English letters.

import "fmt"

func numSplits(s string) int {
    mp, lmp, rmp := make(map[rune]int), make(map[rune]int), make(map[rune]int)
    for _, v := range s {
        mp[v]++
        rmp[v]++
    }
    res := 0
    for _, v := range s {
        lmp[v]++
        rmp[v]--
        if rmp[v] == 0 {
            delete(rmp, v)
        }
        if len(lmp) == len(rmp) {
            res++
        }
    }
    return res
}

func numSplits1(s string) int {
    left, right := make([]int, 26), make([]int, 26)
    res, countLeft, countRight := 0, 0, 0
    for _, c := range s {
        i := c - 'a'
        if left[i] == 0 { countLeft++ }
        left[i]++
    }
    for _, c := range s {
        i := c - 'a'
        left[i]--
        if left[i] == 0 { countLeft-- }
        if right[i] == 0 {
            countRight++
            right[i]++
        }
        if countLeft == countRight { res++ }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aacaba"
    // Output: 2
    // Explanation: There are 5 ways to split "aacaba" and 2 of them are good. 
    // ("a", "acaba") Left string and right string contains 1 and 3 different letters respectively.
    // ("aa", "caba") Left string and right string contains 1 and 3 different letters respectively.
    // ("aac", "aba") Left string and right string contains 2 and 2 different letters respectively (good split).
    // ("aaca", "ba") Left string and right string contains 2 and 2 different letters respectively (good split).
    // ("aacab", "a") Left string and right string contains 3 and 1 different letters respectively.
    fmt.Println(numSplits("aacaba")) // 2
    // Example 2:
    // Input: s = "abcd"
    // Output: 1
    // Explanation: Split the string as follows ("ab", "cd").
    fmt.Println(numSplits("abcd")) // 1

    fmt.Println(numSplits1("aacaba")) // 2
    fmt.Println(numSplits1("abcd")) // 1
}