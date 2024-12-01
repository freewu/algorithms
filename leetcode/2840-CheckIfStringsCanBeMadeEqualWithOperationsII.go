package main

// 2840. Check if Strings Can be Made Equal With Operations II
// You are given two strings s1 and s2, both of length n, consisting of lowercase English letters.

// You can apply the following operation on any of the two strings any number of times:
//     Choose any two indices i and j such that i < j and the difference j - i is even, 
//     then swap the two characters at those indices in the string.

// Return true if you can make the strings s1 and s2 equal, and false otherwise.

// Example 1:
// Input: s1 = "abcdba", s2 = "cabdab"
// Output: true
// Explanation: We can apply the following operations on s1:
// - Choose the indices i = 0, j = 2. The resulting string is s1 = "cbadba".
// - Choose the indices i = 2, j = 4. The resulting string is s1 = "cbbdaa".
// - Choose the indices i = 1, j = 5. The resulting string is s1 = "cabdab" = s2.

// Example 2:
// Input: s1 = "abe", s2 = "bea"
// Output: false
// Explanation: It is not possible to make the two strings equal.

// Constraints:
//     n == s1.length == s2.length
//     1 <= n <= 10^5
//     s1 and s2 consist only of lowercase English letters.

import "fmt"

func checkStrings(s1 string, s2 string) bool {
    mp := make(map[byte]int)
    update := func(h map[byte]int, b byte, v int) {
        h[b] += v
        if h[b] == 0 {
            delete(h, b)
        }
    }
    for i := 0; i < len(s1); i += 2 {
        update(mp, s1[i], 1)
        update(mp, s2[i], -1)
    }
    if len(mp) > 0 {
        return false
    }
    for i := 1; i < len(s1); i += 2 {
        update(mp, s1[i], 1)
        update(mp, s2[i], -1)
    }
    return len(mp) == 0
}

func checkStrings1(s1 string, s2 string) bool {
    count := [2][26]int{}
    for i := 0; i < len(s1); i++ {
        count[i&1][s1[i] - 'a']++
        count[i&1][s2[i] - 'a']--
    }
    for i := 0; i < 26; i++ {
        if count[0][i] != 0 || count[1][i] != 0 {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s1 = "abcdba", s2 = "cabdab"
    // Output: true
    // Explanation: We can apply the following operations on s1:
    // - Choose the indices i = 0, j = 2. The resulting string is s1 = "cbadba".
    // - Choose the indices i = 2, j = 4. The resulting string is s1 = "cbbdaa".
    // - Choose the indices i = 1, j = 5. The resulting string is s1 = "cabdab" = s2.
    fmt.Println(checkStrings("abcdba", "cabdab")) // true
    // Example 2:
    // Input: s1 = "abe", s2 = "bea"
    // Output: false
    // Explanation: It is not possible to make the two strings equal.
    fmt.Println(checkStrings("abe", "bea")) // false

    fmt.Println(checkStrings1("abcdba", "cabdab")) // true
    fmt.Println(checkStrings1("abe", "bea")) // false
}