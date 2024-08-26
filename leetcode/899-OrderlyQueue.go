package main

// 899. Orderly Queue
// You are given a string s and an integer k. 
// You can choose one of the first k letters of s and append it at the end of the string.

// Return the lexicographically smallest string you could have after applying the mentioned step any number of moves.

// Example 1:
// Input: s = "cba", k = 1
// Output: "acb"
// Explanation: 
// In the first move, we move the 1st character 'c' to the end, obtaining the string "bac".
// In the second move, we move the 1st character 'b' to the end, obtaining the final result "acb".

// Example 2:
// Input: s = "baaca", k = 3
// Output: "aaabc"
// Explanation: 
// In the first move, we move the 1st character 'b' to the end, obtaining the string "aacab".
// In the second move, we move the 3rd character 'c' to the end, obtaining the final result "aaabc".

// Constraints:
//     1 <= k <= s.length <= 1000
//     s consist of lowercase English letters.

import "fmt"
import "strings"
import "sort"
import "bytes"

func orderlyQueue(s string, k int) string {
    min := func(w1, w2 string) string { if w1 < w2 { return w1; }; return w2; }
    if k == 1 {
        original := s
        for i := 0; i < len(s); i++ {
            s = min(s, original[i:] + original[:i])
        }
        return s
    }
    sw := strings.Split(s, "")
    sort.Strings(sw)
    return strings.Join(sw, "")
}

func orderlyQueue1(s string, k int) string {
    b := []byte(s)
    if k == 1 {
        mn := []byte(s)
        for i := 0; i < len(s)-1; i++ {
            b = append(b[1:], b[0])
            if bytes.Compare(b, mn) < 0 {
                copy(mn, b)
            }
        }
        return string(mn)
    } else {
        sort.Slice(b, func(i, j int) bool {
            return b[i] < b[j]
        })
        return string(b)
    }
}

func main() {
    // Example 1:
    // Input: s = "cba", k = 1
    // Output: "acb"
    // Explanation: 
    // In the first move, we move the 1st character 'c' to the end, obtaining the string "bac".
    // In the second move, we move the 1st character 'b' to the end, obtaining the final result "acb".
    fmt.Println(orderlyQueue("cba", 1)) // acb
    // Example 2:
    // Input: s = "baaca", k = 3
    // Output: "aaabc"
    // Explanation: 
    // In the first move, we move the 1st character 'b' to the end, obtaining the string "aacab".
    // In the second move, we move the 3rd character 'c' to the end, obtaining the final result "aaabc".
    fmt.Println(orderlyQueue("baaca", 3)) // aaabc

    fmt.Println(orderlyQueue1("cba", 1)) // acb
    fmt.Println(orderlyQueue1("baaca", 3)) // aaabc
}