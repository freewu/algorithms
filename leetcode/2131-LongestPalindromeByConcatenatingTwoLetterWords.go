package main

// 2131. Longest Palindrome by Concatenating Two Letter Words
// You are given an array of strings words. 
// Each element of words consists of two lowercase English letters.

// Create the longest possible palindrome by selecting some elements from words and concatenating them in any order. 
// Each element can be selected at most once.

// Return the length of the longest palindrome that you can create. 
// If it is impossible to create any palindrome, return 0.

// A palindrome is a string that reads the same forward and backward.

// Example 1:
// Input: words = ["lc","cl","gg"]
// Output: 6
// Explanation: One longest palindrome is "lc" + "gg" + "cl" = "lcggcl", of length 6.
// Note that "clgglc" is another longest palindrome that can be created.

// Example 2:
// Input: words = ["ab","ty","yt","lc","cl","ab"]
// Output: 8
// Explanation: One longest palindrome is "ty" + "lc" + "cl" + "yt" = "tylcclyt", of length 8.
// Note that "lcyttycl" is another longest palindrome that can be created.

// Example 3:
// Input: words = ["cc","ll","xx"]
// Output: 2
// Explanation: One longest palindrome is "cc", of length 2.
// Note that "ll" is another longest palindrome that can be created, and so is "xx".

// Constraints:
//     1 <= words.length <= 10^5
//     words[i].length == 2
//     words[i] consists of lowercase English letters.

import "fmt"

func longestPalindrome(words []string) int {
    mp := make(map[string]int)
    for _, word := range words {
        mp[word]++
    }
    res, central := 0, false
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for word, count := range mp {
        if word[0] == word[1] { // If this is a palindrome like 'nn'
            // If the count of the palindrome word is even, then
            // just add the count into result as this will not form a central
            if count % 2 == 0 {
                res += count
            } else {
                // If this is odd number, then this forms a central
                // add the count after subtracting 1 from it and make central falg as true
                res += count - 1
                central = true
            }
        } else if word[0] < word[1] { // If this is not the palindrome word
            reverseStr := string(word[1]) + string(word[0])
            if rcount, ok := mp[reverseStr]; ok { // If the reverse word is available into the map
                // Update the result with the minimum of word and reverse word's count
                res += 2 * min(count, rcount)
            }
        }
    }
    if central { // If this is central sentence, add one more to result
        res++
    }
    return 2 * res // Return the result by multiplying to 2 as each word contains two chars
}

func longestPalindrome1(words []string) int {
    count := [26][26]int{}
    for _, s := range words {
        count[s[0] - 'a'][s[1] - 'a']++
    }
    res, odd := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < 26; i++ {
        res += count[i][i] & ^1
        odd |= count[i][i] & 1
        for j := i + 1; j < 26; j++ {
            res += min(count[i][j], count[j][i]) * 2
        }
    }
    return (res + odd) * 2
}

func main() {
    // Example 1:
    // Input: words = ["lc","cl","gg"]
    // Output: 6
    // Explanation: One longest palindrome is "lc" + "gg" + "cl" = "lcggcl", of length 6.
    // Note that "clgglc" is another longest palindrome that can be created.
    fmt.Println(longestPalindrome([]string{"lc","cl","gg"})) // 6
    // Example 2:
    // Input: words = ["ab","ty","yt","lc","cl","ab"]
    // Output: 8
    // Explanation: One longest palindrome is "ty" + "lc" + "cl" + "yt" = "tylcclyt", of length 8.
    // Note that "lcyttycl" is another longest palindrome that can be created.
    fmt.Println(longestPalindrome([]string{"ab","ty","yt","lc","cl","ab"})) // 8
    // Example 3:
    // Input: words = ["cc","ll","xx"]
    // Output: 2
    // Explanation: One longest palindrome is "cc", of length 2.
    // Note that "ll" is another longest palindrome that can be created, and so is "xx".
    fmt.Println(longestPalindrome([]string{"cc","ll","xx"})) // 2

    fmt.Println(longestPalindrome1([]string{"lc","cl","gg"})) // 6
    fmt.Println(longestPalindrome1([]string{"ab","ty","yt","lc","cl","ab"})) // 8
    fmt.Println(longestPalindrome1([]string{"cc","ll","xx"})) // 2
}