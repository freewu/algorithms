package main

// 1768. Merge Strings Alternately
// You are given two strings word1 and word2. 
// Merge the strings by adding letters in alternating order, starting with word1. 
// If a string is longer than the other, append the additional letters onto the end of the merged string.

// Return the merged string.

// Example 1:
// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r

// Example 2:
// Input: word1 = "ab", word2 = "pqrs"
// Output: "apbqrs"
// Explanation: Notice that as word2 is longer, "rs" is appended to the end.
// word1:  a   b 
// word2:    p   q   r   s
// merged: a p b q   r   s

// Example 3:
// Input: word1 = "abcd", word2 = "pq"
// Output: "apbqcd"
// Explanation: Notice that as word1 is longer, "cd" is appended to the end.
// word1:  a   b   c   d
// word2:    p   q 
// merged: a p b q c   d
 
// Constraints:
//     1 <= word1.length, word2.length <= 100
//     word1 and word2 consist of lowercase English letters.

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    l1, l2 := len(word1), len(word2)
    res := []byte{}
    ml := max(l1, l2)
    for i := 0; i < ml; i++ {
        if i < l1 {
            res = append(res, word1[i])
        }
        if i < l2 {
            res = append(res, word2[i])
        }
    }
    return string(res)
}

func mergeAlternately1(word1, word2 string) string {
    l1,l2 := len(word1),len(word2)
    i,j := 0,0
    res := make([]byte,0 , l1 + l2)
    for i < l1 || j < l2 {
        if i < l1 {
            res = append(res,word1[i])
        }
        if j < l2 {
            res = append(res,word2[j])
        }
        i++
        j++
    }
    return string(res)
}

func mergeAlternately2(word1 string, word2 string) string {
    res, i, j := "", 0, 0
    for i < len(word1) && j < len(word2) {
        res += string(word1[i]) + string(word2[j])
        i++
        j++
    }
    for i == len(word1) && j < len(word2) {
        res += string(word2[j])
        j++
    }
    for j == len(word2) && i < len(word1) {
        res += string(word1[i])
        i++
    }
    return res
}

func main() {
    // word1:  a   b   c
    // word2:    p   q   r
    // merged: a p b q c r
    fmt.Println(mergeAlternately("abc","pqr")) // apbqcr
    // word1:  a   b 
    // word2:    p   q   r   s
    // merged: a p b q   r   s
    fmt.Println(mergeAlternately("ab","pqrs")) // apbqrs
    // word1:  a   b   c   d
    // word2:    p   q 
    // merged: a p b q c   d
    fmt.Println(mergeAlternately("abcd","pq")) // apbqcd

    fmt.Println(mergeAlternately("bluefrog","leetcode")) // blleueetfcroodge

    fmt.Println(mergeAlternately1("abc","pqr")) // apbqcr
    fmt.Println(mergeAlternately1("ab","pqrs")) // apbqrs
    fmt.Println(mergeAlternately1("abcd","pq")) // apbqcd
    fmt.Println(mergeAlternately1("bluefrog","leetcode")) // blleueetfcroodge

    fmt.Println(mergeAlternately2("abc","pqr")) // apbqcr
    fmt.Println(mergeAlternately2("ab","pqrs")) // apbqrs
    fmt.Println(mergeAlternately2("abcd","pq")) // apbqcd
    fmt.Println(mergeAlternately2("bluefrog","leetcode")) // blleueetfcroodge
}