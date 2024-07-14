package main

// 1061. Lexicographically Smallest Equivalent String
// You are given two strings of the same length s1 and s2 and a string baseStr.
// We say s1[i] and s2[i] are equivalent characters.
//     For example, if s1 = "abc" and s2 = "cde", then we have 'a' == 'c', 'b' == 'd', and 'c' == 'e'.

// Equivalent characters follow the usual rules of any equivalence relation:
//     Reflexivity: 'a' == 'a'.
//     Symmetry: 'a' == 'b' implies 'b' == 'a'.
//     Transitivity: 'a' == 'b' and 'b' == 'c' implies 'a' == 'c'.

// For example, given the equivalency information from s1 = "abc" and s2 = "cde", "acd" and "aab" are equivalent strings of baseStr = "eed", and "aab" is the lexicographically smallest equivalent string of baseStr.
// Return the lexicographically smallest equivalent string of baseStr by using the equivalency information from s1 and s2.

// Example 1:
// Input: s1 = "parker", s2 = "morris", baseStr = "parser"
// Output: "makkek"
// Explanation: Based on the equivalency information in s1 and s2, we can group their characters as [m,p], [a,o], [k,r,s], [e,i].
// The characters in each group are equivalent and sorted in lexicographical order.
// So the answer is "makkek".

// Example 2:
// Input: s1 = "hello", s2 = "world", baseStr = "hold"
// Output: "hdld"
// Explanation: Based on the equivalency information in s1 and s2, we can group their characters as [h,w], [d,e,o], [l,r].
// So only the second letter 'o' in baseStr is changed to 'd', the answer is "hdld".

// Example 3:
// Input: s1 = "leetcode", s2 = "programs", baseStr = "sourcecode"
// Output: "aauaaaaada"
// Explanation: We group the equivalent characters in s1 and s2 as [a,o,e,r,s,c], [l,p], [g,t] and [d,m], thus all letters in baseStr except 'u' and 'd' are transformed to 'a', the answer is "aauaaaaada".

// Constraints:
//     1 <= s1.length, s2.length, baseStr <= 1000
//     s1.length == s2.length
//     s1, s2, and baseStr consist of lowercase English letters.

import "fmt"
import "strings"

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    var sb strings.Builder
    parent := make([]int, 26)
    for i := 0; i < 26; i++ {
        parent[i] = -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var find func(x int, parent *[]int) int
    find = func(x int, parent *[]int) int {
        if (*parent)[x] == -1 {
            return x
        }
        (*parent)[x] = find((*parent)[x], parent)
        return (*parent)[x]
    }
    union := func (x, y int, parent *[]int) {
        x, y = find(x, parent), find(y, parent)
        if x != y {
            (*parent)[max(x, y)] = min(x, y)
        }
    }
    for i := 0; i < len(s1); i++ {
        union(int(s1[i]-'a'), int(s2[i]-'a'), &parent)
    }
    for i := 0; i < len(baseStr); i++ {
        sb.WriteByte(byte(find(int(baseStr[i]-'a'), &parent) + 'a'))
    }
    return sb.String()
}

func main() {
    // Example 1:
    // Input: s1 = "parker", s2 = "morris", baseStr = "parser"
    // Output: "makkek"
    // Explanation: Based on the equivalency information in s1 and s2, we can group their characters as [m,p], [a,o], [k,r,s], [e,i].
    // The characters in each group are equivalent and sorted in lexicographical order.
    // So the answer is "makkek".
    fmt.Println(smallestEquivalentString("parker","morris","parser")) // "makkek"
    // Example 2:
    // Input: s1 = "hello", s2 = "world", baseStr = "hold"
    // Output: "hdld"
    // Explanation: Based on the equivalency information in s1 and s2, we can group their characters as [h,w], [d,e,o], [l,r].
    // So only the second letter 'o' in baseStr is changed to 'd', the answer is "hdld".
    fmt.Println(smallestEquivalentString("hello","world","hold")) // "hdld"
    // Example 3:
    // Input: s1 = "leetcode", s2 = "programs", baseStr = "sourcecode"
    // Output: "aauaaaaada"
    // Explanation: We group the equivalent characters in s1 and s2 as [a,o,e,r,s,c], [l,p], [g,t] and [d,m], thus all letters in baseStr except 'u' and 'd' are transformed to 'a', the answer is "aauaaaaada".
    fmt.Println(smallestEquivalentString("leetcode","programs","sourcecode")) // "aauaaaaada"
}