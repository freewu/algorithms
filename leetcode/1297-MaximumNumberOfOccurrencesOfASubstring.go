package main

// 1297. Maximum Number of Occurrences of a Substring
// Given a string s, return the maximum number of occurrences of any substring under the following rules:
//     The number of unique characters in the substring must be less than or equal to maxLetters.
//     The substring size must be between minSize and maxSize inclusive.

// Example 1:
// Input: s = "aababcaab", maxLetters = 2, minSize = 3, maxSize = 4
// Output: 2
// Explanation: Substring "aab" has 2 occurrences in the original string.
// It satisfies the conditions, 2 unique letters and size 3 (between minSize and maxSize).

// Example 2:
// Input: s = "aaaa", maxLetters = 1, minSize = 3, maxSize = 3
// Output: 2
// Explanation: Substring "aaa" occur 2 times in the string. It can overlap.

// Constraints:
//     1 <= s.length <= 10^5
//     1 <= maxLetters <= 26
//     1 <= minSize <= maxSize <= min(26, s.length)
//     s consists of only lowercase English letters.

import "fmt"

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
    magic, counter := map[string]int{}, [123]int{}
    res, kind := 0, 0
    for i := 0; i < minSize; i++ {
        counter[s[i]]++
        if counter[s[i]] == 1 {
            kind++
        }
    }
    if kind <= maxLetters {
        magic[s[0:minSize]]++
    }
    for i := minSize; i < len(s); i++ {
        counter[s[i]]++
        if counter[s[i]] == 1 {
            kind++
        }
        counter[s[i-minSize]]--
        if counter[s[i-minSize]] == 0 {
            kind--
        }
        if kind <= maxLetters {
            magic[s[i-minSize+1:i+1]]++
        }
    }
    for _, v := range magic {
        if v > res {
            res = v
        }
    }
    return res
}

// sliding window
func maxFreq1(s string, maxLetters int, minSize int, maxSize int) int {
    if len(s) < minSize {
        return 0
    }
    count, letter := make([]int, 26), 0
    for i := 0; i < minSize - 1; i++ {
        if count[s[i] - 'a'] == 0 {
            letter++
        }
        count[s[i] - 'a'] ++
    }
    res, mp := 0, make(map[string]int) // use map to save all the substrings. 
    for i := 0; i < len(s) - minSize + 1; i++ {
        if count[s[i + minSize - 1] - 'a'] == 0 {
            letter++
        }
        count[s[i + minSize - 1] - 'a']++
        if letter <= maxLetters {
            mp[s[i : i + minSize]]++ 
            if mp[s[i : i + minSize]] > res {
                res = mp[s[i : i + minSize]]  
            }
        }
        //slide the left pointer        
        if count[s[i] - 'a'] == 1 {
            letter--
        }
        count[s[i] - 'a']--
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aababcaab", maxLetters = 2, minSize = 3, maxSize = 4
    // Output: 2
    // Explanation: Substring "aab" has 2 occurrences in the original string.
    // It satisfies the conditions, 2 unique letters and size 3 (between minSize and maxSize).
    fmt.Println(maxFreq("aababcaab", 2, 3, 4)) // 2
    // Example 2:
    // Input: s = "aaaa", maxLetters = 1, minSize = 3, maxSize = 3
    // Output: 2
    // Explanation: Substring "aaa" occur 2 times in the string. It can overlap.
    fmt.Println(maxFreq("aaaa", 1, 3, 3)) // 2

    fmt.Println(maxFreq1("aababcaab", 2, 3, 4)) // 2
    fmt.Println(maxFreq1("aaaa", 1, 3, 3)) // 2
}