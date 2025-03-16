package main

// 1220. Count Vowels Permutation
// Given an integer n, your task is to count how many strings of length n can be formed under the following rules:
//     Each character is a lower case vowel ('a', 'e', 'i', 'o', 'u')
//     Each vowel 'a' may only be followed by an 'e'.
//     Each vowel 'e' may only be followed by an 'a' or an 'i'.
//     Each vowel 'i' may not be followed by another 'i'.
//     Each vowel 'o' may only be followed by an 'i' or a 'u'.
//     Each vowel 'u' may only be followed by an 'a'.

// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 1
// Output: 5
// Explanation: All possible strings are: "a", "e", "i" , "o" and "u".

// Example 2:
// Input: n = 2
// Output: 10
// Explanation: All possible strings are: "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" and "ua".

// Example 3: 
// Input: n = 5
// Output: 68
 
// Constraints:
//     1 <= n <= 2 * 10^4

import "fmt"

func countVowelPermutation(n int) int { 
    a, e, i, o, u, mod := 1, 1, 1, 1, 1, 1_000_000_007
    for j := 1; j < n; j++ {
        a_next := e // 每个元音 'a' 后面都只能跟着 'e'
        e_next := (a + i) % mod // 每个元音 'e' 后面只能跟着 'a' 或者是 'i'
        i_next := (a + e + o + u) % mod // 每个元音 'i' 后面 不能 再跟着另一个 'i'
        o_next := (i + u) % mod // 每个元音 'o' 后面只能跟着 'i' 或者是 'u'
        u_next := a // 每个元音 'u' 后面只能跟着 'a'
        a, e, i, o, u = a_next, e_next, i_next, o_next, u_next
    }
    return (a + e + i + o + u) % mod
}

func countVowelPermutation1(n int) int {
    a, e, i, o, u, mod := 1, 1, 1, 1, 1, 1_000_000_007
    for k := 1; k < n; k++ {
        a, e, i, o, u = e, (a + i) % mod, (a + e + o + u) % mod, (i + u) % mod, a
    }
    return (a + e + i + o + u) % mod
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 5
    // Explanation: All possible strings are: "a", "e", "i" , "o" and "u".
    fmt.Println(countVowelPermutation(1)) // 5
    // Example 2:
    // Input: n = 2
    // Output: 10
    // Explanation: All possible strings are: "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" and "ua".
    fmt.Println(countVowelPermutation(2)) // 10
    // Example 3: 
    // Input: n = 5
    // Output: 68
    fmt.Println(countVowelPermutation(5)) // 68

    fmt.Println(countVowelPermutation(3)) // 19
    fmt.Println(countVowelPermutation(1024)) // 245633651
    fmt.Println(countVowelPermutation(9999)) // 68
    fmt.Println(countVowelPermutation(20000)) // 68

    fmt.Println(countVowelPermutation1(1)) // 5
    fmt.Println(countVowelPermutation1(2)) // 10
    fmt.Println(countVowelPermutation1(3)) // 19
    fmt.Println(countVowelPermutation1(5)) // 68
    fmt.Println(countVowelPermutation1(1024)) // 245633651
    fmt.Println(countVowelPermutation1(9999)) // 876173704
    fmt.Println(countVowelPermutation1(20000)) // 759959057
}