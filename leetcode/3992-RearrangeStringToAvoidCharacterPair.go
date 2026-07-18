package main

// 3992. Rearrange String to Avoid Character Pair
// You are given a string s and two distinct lowercase English letters x and y.

// Rearrange the characters of s to construct a new string t such that:
//     1. t is a permutation of s.
//     2. Every occurrence of y appears before every occurrence of x in t.

// Return any valid string t.

// Example 1:
// Input: s = "aabc", x = "a", y = "c"
// Output: "cbaa"
// Explanation:
// The string "cbaa" is a permutation of "aabc", and every occurrence of 'c' appears before every occurrence of 'a'.

// Example 2:
// Input: s = "dcab", x = "d", y = "b"
// Output: "cabd"
// Explanation:
// The string "cabd" is a permutation of "dcab", and every occurrence of 'b' appears before every occurrence of 'd'.

// Example 3:
// Input: s = "axe", x = "o", y = "x"
// Output: "axe"
// Explanation:
// The string "axe" is already valid. Since 'o' does not occur in the string, the required condition is automatically satisfied.

// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters.
//     x and y are lowercase English letters.
//     x != y

import "fmt"
import "slices"

func rearrangeString(s string, x, y byte) string {
    t := []byte(s)
    if x < y {
        slices.SortFunc(t, func(a, b byte) int { 
            return int(b) - int(a) 
        })
    } else {
        slices.Sort(t)
    }
    return string(t)
}

// 相向双指针
func rearrangeString1(s string, x, y byte) string {
    t := []byte(s)
    l, r := 0, len(t)-1
    for l < r { // 循环直到不足两个字母
        if t[l] != x { // 寻找最左边的 x
            l++
        } else if t[r] != y { // 寻找最右边的 y
            r--
        } else {
            t[l], t[r] = t[r], t[l]
            // 交换后，[0,l] 都不含 x，[r,n-1] 都不含 y，问题变成 [l+1,r-1] 的子问题
            l++
            r--
        }
    }
    return string(t)
}

func main() {
    // Example 1:
    // Input: s = "aabc", x = "a", y = "c"
    // Output: "cbaa"
    // Explanation:
    // The string "cbaa" is a permutation of "aabc", and every occurrence of 'c' appears before every occurrence of 'a'.
    fmt.Println(rearrangeString("aabc", 'a', 'c')) // "cbaa"
    // Example 2:
    // Input: s = "dcab", x = "d", y = "b"
    // Output: "cabd"
    // Explanation:
    // The string "cabd" is a permutation of "dcab", and every occurrence of 'b' appears before every occurrence of 'd'.
    fmt.Println(rearrangeString("dcab", 'd', 'b')) // "cabd"    
    // Example 3:
    // Input: s = "axe", x = "o", y = "x"
    // Output: "axe"
    // Explanation:
    // The string "axe" is already valid. Since 'o' does not occur in the string, the required condition is automatically satisfied.
    fmt.Println(rearrangeString("axe", 'o', 'x')) // "axe"

    fmt.Println(rearrangeString("bluefrog", 'b', 'e')) // "urolgfeb"
    fmt.Println(rearrangeString("leetcode", 'c', 'e')) // "toleeedc"
    fmt.Println(rearrangeString("free", 'e', 'f')) // "rfee"

    fmt.Println(rearrangeString1("aabc", 'a', 'c')) // "cbaa"
    fmt.Println(rearrangeString1("dcab", 'd', 'b')) // "cabd"    
    fmt.Println(rearrangeString1("axe", 'o', 'x')) // "axe"
    fmt.Println(rearrangeString1("bluefrog", 'b', 'e')) // "urolgfeb"
    fmt.Println(rearrangeString1("leetcode", 'c', 'e')) // "toleeedc"
    fmt.Println(rearrangeString1("free", 'e', 'f')) // "rfee"
}