package main

// 3335. Total Characters in String After Transformations I
// You are given a string s and an integer t, representing the number of transformations to perform. 
// In one transformation, every character in s is replaced according to the following rules:
//     1. If the character is 'z', replace it with the string "ab".
//     2. Otherwise, replace it with the next character in the alphabet. 
//        For example, 'a' is replaced with 'b', 'b' is replaced with 'c', and so on.

// Return the length of the resulting string after exactly t transformations.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "abcyy", t = 2
// Output: 7
// Explanation:
// 1. First Transformation (t = 1):
//     'a' becomes 'b'
//     'b' becomes 'c'
//     'c' becomes 'd'
//     'y' becomes 'z'
//     'y' becomes 'z'
//     String after the first transformation: "bcdzz"
// 2. Second Transformation (t = 2):
//     'b' becomes 'c'
//     'c' becomes 'd'
//     'd' becomes 'e'
//     'z' becomes "ab"
//     'z' becomes "ab"
//     String after the second transformation: "cdeabab"
// 3. Final Length of the string: The string is "cdeabab", which has 7 characters.

// Example 2:
// Input: s = "azbk", t = 1
// Output: 5
// Explanation:
// 1. First Transformation (t = 1):
//     'a' becomes 'b'
//     'z' becomes "ab"
//     'b' becomes 'c'
//     'k' becomes 'l'
//     String after the first transformation: "babcl"
// 2. Final Length of the string: The string is "babcl", which has 5 characters.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters.
//     1 <= t <= 10^5

import "fmt"

func lengthAfterTransformations(s string, t int) int {
    res, mod := 0, 1_000_000_007
    count := make([]int, 26)
    for _, v := range s {
        count[v-'a']++
    }
    for i := 0; i < t; i++ {
        newCount := make([]int, 26)
        for j := 0; j < 25; j++ {
            newCount[j + 1] = (newCount[j + 1] + count[j]) % mod
        }
        newCount[0] = (newCount[0] + count[25]) % mod
        newCount[1] = (newCount[1] + count[25]) % mod 
        count = newCount
    }
    for i := 0; i < 26; i++ {
        res = (res + count[i]) % mod
    }
    return res
}

func lengthAfterTransformations1(s string, t int) int {
    res, mod, arr := 0, 1_000_000_007, [26]int{}
    for _, v := range s {
        arr[v-'a']++
    }
    for i := 0; i < t; i++ {
        j := (i / 26 + 1) * 26 + 25 - i
        arr[(j + 1) % 26] = (arr[(j + 1) % 26] + arr[j % 26]) % mod
    }
    for i := 0; i < 26; i++ {
        res = (res + arr[i]) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcyy", t = 2
    // Output: 7
    // Explanation:
    // 1. First Transformation (t = 1):
    //     'a' becomes 'b'
    //     'b' becomes 'c'
    //     'c' becomes 'd'
    //     'y' becomes 'z'
    //     'y' becomes 'z'
    //     String after the first transformation: "bcdzz"
    // 2. Second Transformation (t = 2):
    //     'b' becomes 'c'
    //     'c' becomes 'd'
    //     'd' becomes 'e'
    //     'z' becomes "ab"
    //     'z' becomes "ab"
    //     String after the second transformation: "cdeabab"
    // 3. Final Length of the string: The string is "cdeabab", which has 7 characters.
    fmt.Println(lengthAfterTransformations("abcyy", 2)) // 7
    // Example 2:
    // Input: s = "azbk", t = 1
    // Output: 5
    // Explanation:
    // 1. First Transformation (t = 1):
    //     'a' becomes 'b'
    //     'z' becomes "ab"
    //     'b' becomes 'c'
    //     'k' becomes 'l'
    //     String after the first transformation: "babcl"
    // 2. Final Length of the string: The string is "babcl", which has 5 characters.
    fmt.Println(lengthAfterTransformations("azbk", 1)) // 5

    fmt.Println(lengthAfterTransformations("bluefrog", 1)) // 8
    fmt.Println(lengthAfterTransformations("leetcode", 1)) // 8

    fmt.Println(lengthAfterTransformations1("abcyy", 2)) // 7
    fmt.Println(lengthAfterTransformations1("azbk", 1)) // 5
    fmt.Println(lengthAfterTransformations1("bluefrog", 1)) // 8
    fmt.Println(lengthAfterTransformations1("leetcode", 1)) // 8
}