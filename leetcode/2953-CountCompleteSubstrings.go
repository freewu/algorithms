package main

// 2953. Count Complete Substrings
// You are given a string word and an integer k.

// A substring s of word is complete if:
//     1. Each character in s occurs exactly k times.
//     2. The difference between two adjacent characters is at most 2. 
//        That is, for any two adjacent characters c1 and c2 in s, the absolute difference in their positions in the alphabet is at most 2.

// Return the number of complete substrings of word.

// A substring is a non-empty contiguous sequence of characters in a string.

// Example 1:
// Input: word = "igigee", k = 2
// Output: 3
// Explanation: The complete substrings where each character appears exactly twice and the difference between adjacent characters is at most 2 are: igigee, igigee, igigee.

// Example 2:
// Input: word = "aaabbbccc", k = 3
// Output: 6
// Explanation: The complete substrings where each character appears exactly three times and the difference between adjacent characters is at most 2 are: aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc.

// Constraints:
//     1 <= word.length <= 10^5
//     word consists only of lowercase English letters.
//     1 <= k <= word.length

import "fmt"

func countCompleteSubstrings(word string, k int) int {
    res, n := 0, len(word)
    helper := func(s string) int {
        res, m := 0, len(s)
        for i := 1; i <= 26 && i * k <= m; i++ {
            l := i * k
            count := [26]int{}
            for j := 0; j < l; j++ {
                count[int(s[j] - 'a')]++
            }
            freq := make(map[int]int)
            for _, v := range count {
                if v > 0 {
                    freq[v]++
                }
            }
            if freq[k] == i {
                res++
            }
            for j := l; j < m; j++ {
                a, b := int(s[j] - 'a'), int(s[j-l] - 'a')
                freq[count[a]]--
                count[a]++
                freq[count[a]]++
                freq[count[b]]--
                count[b]--
                freq[count[b]]++
                if freq[k] == i {
                    res++
                }
            }
        }
        return res
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n; {
        j := i + 1
        for j < n && abs(int(word[j]) - int(word[j - 1])) <= 2 {
            j++
        }
        res += helper(word[i:j])
        i = j
    }
    return res
}

func countCompleteSubstrings1(word string, k int) int {
    res := 0
    helper := func(s string, k int) int {
        res := 0
        for m := 1; m <= 26 && k*m <= len(s); m++ {
            count := [26]int{}
            check := func() {
                for i := range count {
                    if count[i] > 0 && count[i] != k { return }
                }
                res++
            }
            for right, in := range s {
                count[in - 'a']++
                if left := right + 1 - k * m; left >= 0 {
                    check()
                    count[s[left] - 'a']--
                }
            }
        }
        return res
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, n := 0, len(word); i < n; {
        st := i
        for i++; i < n && abs(int(word[i]) - int(word[i-1])) <= 2; i++ {
        }
        res += helper(word[st:i], k)
    }
    return res 
}

func main() {
    // Example 1:
    // Input: word = "igigee", k = 2
    // Output: 3
    // Explanation: The complete substrings where each character appears exactly twice and the difference between adjacent characters is at most 2 are: igigee, igigee, igigee.
    fmt.Println(countCompleteSubstrings("igigee", 2)) // 3
    // Example 2:
    // Input: word = "aaabbbccc", k = 3
    // Output: 6
    // Explanation: The complete substrings where each character appears exactly three times and the difference between adjacent characters is at most 2 are: aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc, aaabbbccc.
    fmt.Println(countCompleteSubstrings("aaabbbccc", 3)) // 6

    fmt.Println(countCompleteSubstrings("bluefrog", 2)) // 0
    fmt.Println(countCompleteSubstrings("leetcode", 2)) // 1

    fmt.Println(countCompleteSubstrings1("igigee", 2)) // 3
    fmt.Println(countCompleteSubstrings1("aaabbbccc", 3)) // 6
    fmt.Println(countCompleteSubstrings1("bluefrog", 2)) // 0
    fmt.Println(countCompleteSubstrings1("leetcode", 2)) // 1
}