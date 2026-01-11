package main

// 3803. Count Residue Prefixes
// You are given a string s consisting only of lowercase English letters.

// A prefix of s is called a residue if the number of distinct characters in the prefix is equal to len(prefix) % 3.

// Return the count of residue prefixes in s.

// A prefix of a string is a non-empty substring that starts from the beginning of the string and extends to any point within it.

// Example 1:
// Input: s = "abc"
// Output: 2
// Explanation:​​​​​​​
// Prefix "a" has 1 distinct character and length modulo 3 is 1, so it is a residue.
// Prefix "ab" has 2 distinct characters and length modulo 3 is 2, so it is a residue.
// Prefix "abc" does not satisfy the condition. Thus, the answer is 2.

// Example 2:
// Input: s = "dd"
// Output: 1
// Explanation:
// Prefix "d" has 1 distinct character and length modulo 3 is 1, so it is a residue.
// Prefix "dd" has 1 distinct character but length modulo 3 is 2, so it is not a residue. Thus, the answer is 1.

// Example 3:
// Input: s = "bob"
// Output: 2
// Explanation:
// Prefix "b" has 1 distinct character and length modulo 3 is 1, so it is a residue.
// Prefix "bo" has 2 distinct characters and length mod 3 is 2, so it is a residue. Thus, the answer is 2.

// Constraints:
//     1 <= s.length <= 100
//     s contains only lowercase English letters.

import "fmt"

func residuePrefixes(s string) int {
    res, mask, n := 0, 0, len(s)
    ones := func(n int) int {
        res := 0
        for n > 0 {
            n &= n-1
            res++
        }
        return res
    }
    for i := range n {
        mask |= 1 << int(s[i]-'a')
        if ones(mask) == (i+1) % 3 { // 如果字符串 s 的某个 前缀 中 不同字符的数量 等于 len(prefix) % 3，则该前缀被称为残差前缀（residue）
            res++
        }
    }
    return res
}

func residuePrefixes1(s string) int {
    count := make([]bool,26)
    res,sum := 0, 0
    for i,ch := range s {
        if !count[ch - 'a'] {
            sum++
        }
        count[ch-'a'] = true
        if (i + 1) % 3 == sum {
            res++
        }
        if sum >= 3 {
            return res
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abc"
    // Output: 2
    // Explanation:​​​​​​​
    // Prefix "a" has 1 distinct character and length modulo 3 is 1, so it is a residue.
    // Prefix "ab" has 2 distinct characters and length modulo 3 is 2, so it is a residue.
    // Prefix "abc" does not satisfy the condition. Thus, the answer is 2.
    fmt.Println(residuePrefixes("abc")) // 2
    // Example 2:
    // Input: s = "dd"
    // Output: 1
    // Explanation:
    // Prefix "d" has 1 distinct character and length modulo 3 is 1, so it is a residue.
    // Prefix "dd" has 1 distinct character but length modulo 3 is 2, so it is not a residue. Thus, the answer is 1.
    fmt.Println(residuePrefixes("dd")) // 1
    // Example 3:
    // Input: s = "bob"
    // Output: 2
    // Explanation:
    // Prefix "b" has 1 distinct character and length modulo 3 is 1, so it is a residue.
    // Prefix "bo" has 2 distinct characters and length mod 3 is 2, so it is a residue. Thus, the answer is 2.
    fmt.Println(residuePrefixes("bob")) // 2

    fmt.Println(residuePrefixes("bluefrog")) // 2
    fmt.Println(residuePrefixes("leetcode")) // 2

    fmt.Println(residuePrefixes1("abc")) // 2
    fmt.Println(residuePrefixes1("dd")) // 1
    fmt.Println(residuePrefixes1("bob")) // 2
    fmt.Println(residuePrefixes1("bluefrog")) // 2
    fmt.Println(residuePrefixes1("leetcode")) // 2
}