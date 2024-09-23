package main

// 2207. Maximize Number of Subsequences in a String
// You are given a 0-indexed string text and another 0-indexed string pattern of length 2, 
// both of which consist of only lowercase English letters.

// You can add either pattern[0] or pattern[1] anywhere in text exactly once. 
// Note that the character can be added even at the beginning or at the end of text.

// Return the maximum number of times pattern can occur as a subsequence of the modified text.

// A subsequence is a string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.

// Example 1:
// Input: text = "abdcdbc", pattern = "ac"
// Output: 4
// Explanation:
// If we add pattern[0] = 'a' in between text[1] and text[2], we get "abadcdbc". Now, the number of times "ac" occurs as a subsequence is 4.
// Some other strings which have 4 subsequences "ac" after adding a character to text are "aabdcdbc" and "abdacdbc".
// However, strings such as "abdcadbc", "abdccdbc", and "abdcdbcc", although obtainable, have only 3 subsequences "ac" and are thus suboptimal.
// It can be shown that it is not possible to get more than 4 subsequences "ac" by adding only one character.

// Example 2:
// Input: text = "aabb", pattern = "ab"
// Output: 6
// Explanation:
// Some of the strings which can be obtained from text and have 6 subsequences "ab" are "aaabb", "aaabb", and "aabbb".

// Constraints:
//     1 <= text.length <= 10^5
//     pattern.length == 2
//     text and pattern consist only of lowercase English letters.

import "fmt"

func maximumSubsequenceCount(text string, pattern string) int64 {
    cnt0, cnt1, res0, res1 := 1, 1, 0, 0
    for i, j := 0, len(text) - 1; j >= 0;  {
        if text[i] == pattern[1] { res0 += cnt0 }
        if text[j] == pattern[0] { res1 += cnt1 }
        if text[i] == pattern[0] { cnt0++ }
        if text[j] == pattern[1] { cnt1++ }
        i++
        j--
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return int64(max(res0, res1))
}

func maximumSubsequenceCount1(text string, pattern string) int64 {
    a, b := pattern[0], pattern[1]
    n, cnta, cntb := len(text), 0, 0
    for i := 0; i < n; i++ {
        if text[i] == a { cnta++ }
        if text[i] == b { cntb++ }
    }
    if a == b {
        return int64(cnta * (cnta + 1) / 2)
    }
    res, cntc := 0, cntb
    for i := 0; i <= n && cntb > 0; i++ {
        if text[i] == a { res += cntb }
        if text[i] == b { cntb-- }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res += max(cnta, cntc)
    return int64(res)
}

// func maximumSubsequenceCount(text string, pattern string) int64 {
//     if pattern[0] == pattern[1] { // // when pattern[0] == pattern[1]
//         freq := 1
//         for i := 0; i < len(text); i++ {
//             if text[i] == pattern[0] {
//                 freq++
//             }
//         }
//         return int64((freq * (freq - 1)) / 2); // number of subsequences : choose any two characters from freq nC2
//     }
//     // choice 1
//     //text1 := pattern.charAt(0) + text;
//     text1 := []byte{ pattern[0] }
//     text1 = append(text1, []byte(text)...)
//     fmt.Println(text, " ", string(text1))
//     freq1, count1 := 0, 0
//     for i := 0; i < len(text1); i++ {
//         if text1[i] == pattern[0] {
//             freq1++
//         } else if text1[i] == pattern[1] {
//             count1 += freq1
//         }
//     }
//     // choice 2
//     text2 := []byte(text)
//     text2 = append(text2, pattern[1])
//     fmt.Println(string(text2))
//     freq2, count2 := 0, 0
//     for  i := len(text2) - 1; i>= 0; i-- {
//         if text2[i] == pattern[1] {
//             freq2++
//         } else if text2[i] == pattern[0] {
//             count2 += freq2
//         }
//     }
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     return int64(max(freq1, freq2))
// }

func main() {
    // Example 1:
    // Input: text = "abdcdbc", pattern = "ac"
    // Output: 4
    // Explanation:
    // If we add pattern[0] = 'a' in between text[1] and text[2], we get "abadcdbc". Now, the number of times "ac" occurs as a subsequence is 4.
    // Some other strings which have 4 subsequences "ac" after adding a character to text are "aabdcdbc" and "abdacdbc".
    // However, strings such as "abdcadbc", "abdccdbc", and "abdcdbcc", although obtainable, have only 3 subsequences "ac" and are thus suboptimal.
    // It can be shown that it is not possible to get more than 4 subsequences "ac" by adding only one character.
    fmt.Println(maximumSubsequenceCount("abdcdbc", "ac")) // 4
    // Example 2:
    // Input: text = "aabb", pattern = "ab"
    // Output: 6
    // Explanation:
    // Some of the strings which can be obtained from text and have 6 subsequences "ab" are "aaabb", "aaabb", and "aabbb".
    fmt.Println(maximumSubsequenceCount("aabb", "ab")) // 6
    
    fmt.Println(maximumSubsequenceCount1("abdcdbc", "ac")) // 4
    fmt.Println(maximumSubsequenceCount1("aabb", "ab")) // 6
}