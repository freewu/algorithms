package main

// 3839. Number of Prefix Connected Groups
// You are given an array of strings words and an integer k.

// Two words a and b at distinct indices are prefix-connected if a[0..k-1] == b[0..k-1].

// A connected group is a set of words such that each pair of words is prefix-connected.

// Return the number of connected groups that contain at least two words, formed from the given words.

// Note:
//     1. Words with length less than k cannot join any group and are ignored.
//     2. Duplicate strings are treated as separate words.
//     3. A prefix of a string is a substring that starts from the beginning of the string and extends to any point within it.

// Example 1:
// Input: words = ["apple","apply","banana","bandit"], k = 2
// Output: 2
// Explanation:
// Words sharing the same first k = 2 letters are grouped together:
// words[0] = "apple" and words[1] = "apply" share prefix "ap".
// words[2] = "banana" and words[3] = "bandit" share prefix "ba".
// Thus, there are 2 connected groups, each containing at least two words.

// Example 2:
// Input: words = ["car","cat","cartoon"], k = 3
// Output: 1
// Explanation:
// Words are evaluated for a prefix of length k = 3:
// words[0] = "car" and words[2] = "cartoon" share prefix "car".
// words[1] = "cat" does not share a 3-length prefix with any other word.
// Thus, there is 1 connected group.

// Example 3:
// Input: words = ["bat","dog","dog","doggy","bat"], k = 3
// Output: 2
// Explanation:
// Words are evaluated for a prefix of length k = 3:
// words[0] = "bat" and words[4] = "bat" form a group.
// words[1] = "dog", words[2] = "dog" and words[3] = "doggy" share prefix "dog".
// Thus, there are 2 connected groups, each containing at least two words.

// Constraints:
//     1 <= words.length <= 5000
//     1 <= words[i].length <= 100
//     1 <= k <= 100
//     All strings in words consist of lowercase English letters.

import "fmt"

func prefixConnected(words []string, k int) int {
    res, mp := 0, make(map[string]int)
    for _, word := range words {
        if len(word) < k { continue }
        mp[word[:k]]++
    }
    for _, v := range mp {    
        if v > 1 {
            res++
        }
    }
    return res
}

func prefixConnected1(words []string, k int) int {
    res, mp := 0, make(map[string]int)
    for _, w := range words {
        if len(w) >= k {
            mp[w[:k]]++
        }
    }
    for _, c := range mp {
        if c > 1 {
            res++
        }
    }
    return res
} 

func main() {
    // Example 1:
    // Input: words = ["apple","apply","banana","bandit"], k = 2
    // Output: 2
    // Explanation:
    // Words sharing the same first k = 2 letters are grouped together:
    // words[0] = "apple" and words[1] = "apply" share prefix "ap".
    // words[2] = "banana" and words[3] = "bandit" share prefix "ba".
    // Thus, there are 2 connected groups, each containing at least two words.
    fmt.Println(prefixConnected([]string{"apple","apply","banana","bandit"}, 2)) // 2
    // Example 2:
    // Input: words = ["car","cat","cartoon"], k = 3
    // Output: 1
    // Explanation:
    // Words are evaluated for a prefix of length k = 3:
    // words[0] = "car" and words[2] = "cartoon" share prefix "car".
    // words[1] = "cat" does not share a 3-length prefix with any other word.
    // Thus, there is 1 connected group.
    fmt.Println(prefixConnected([]string{"car","cat","cartoon"}, 3)) // 1
    // Example 3:
    // Input: words = ["bat","dog","dog","doggy","bat"], k = 3
    // Output: 2
    // Explanation:
    // Words are evaluated for a prefix of length k = 3:
    // words[0] = "bat" and words[4] = "bat" form a group.
    // words[1] = "dog", words[2] = "dog" and words[3] = "doggy" share prefix "dog".
    // Thus, there are 2 connected groups, each containing at least two words.
    fmt.Println(prefixConnected([]string{"bat","dog","dog","doggy","bat"}, 3)) // 2

    fmt.Println(prefixConnected([]string{"a","bc","bcd","ef"}, 1)) // 1
    fmt.Println(prefixConnected([]string{"bluefrog","leetcode","freewu"}, 3)) // 0

    fmt.Println(prefixConnected1([]string{"apple","apply","banana","bandit"}, 2)) // 2
    fmt.Println(prefixConnected1([]string{"car","cat","cartoon"}, 3)) // 1
    fmt.Println(prefixConnected1([]string{"bat","dog","dog","doggy","bat"}, 3)) // 2
    fmt.Println(prefixConnected1([]string{"a","bc","bcd","ef"}, 1)) // 1
    fmt.Println(prefixConnected1([]string{"bluefrog","leetcode","freewu"}, 3)) // 0
}