package main

// 1209. Remove All Adjacent Duplicates in String II
// You are given a string s and an integer k, 
// a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them, 
// causing the left and the right side of the deleted substring to concatenate together.

// We repeatedly make k duplicate removals on s until we no longer can.

// Return the final string after all such duplicate removals have been made. 
// It is guaranteed that the answer is unique.

// Example 1:
// Input: s = "abcd", k = 2
// Output: "abcd"
// Explanation: There's nothing to delete.

// Example 2:
// Input: s = "deeedbbcccbdaa", k = 3
// Output: "aa"
// Explanation: 
// First delete "eee" and "ccc", get "ddbbbdaa"
// Then delete "bbb", get "dddaa"
// Finally delete "ddd", get "aa"

// Example 3:
// Input: s = "pbbcggttciiippooaais", k = 2
// Output: "ps"

// Constraints:
//     1 <= s.length <= 10^5
//     2 <= k <= 10^4
//     s only contains lowercase English letters.

import "fmt"

func removeDuplicates(s string, k int) string {
    arr := []byte{}
    for index, _ := range s {
        arr = append(arr, s[index])
        if len(arr) >= k {
            allEqualsCharacter := true
            for i := 0; i < k; i++ {
                if arr[len(arr) - (i+1)] != s[index] {
                    allEqualsCharacter = false
                    break
                }
            }
            if allEqualsCharacter {
                arr = arr[:len(arr)-k]
            }
        }
    }
    return string(arr)
}

func removeDuplicates1(s string, k int) string {
    n, p := len(s), 1
    stack := make([]int, n + 2)
    ch := make([]rune, n + 2)
    ch[0] = ' '
    for _, v := range s {
        if v == ch[p - 1] {
            stack[p - 1]++
            if stack[p - 1] == k {
                p--
            }
        } else {
            stack[p] =  1
            ch[p] = v
            p++
        }
    }
    res := []rune{}
    for i := 1; i < p; i++ {
        for j := 0; j < stack[i]; j++ {
            res = append(res, ch[i])
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "abcd", k = 2
    // Output: "abcd"
    // Explanation: There's nothing to delete.
    fmt.Println(removeDuplicates("abcd", 2)) // "abcd"
    // Example 2:
    // Input: s = "deeedbbcccbdaa", k = 3
    // Output: "aa"
    // Explanation: 
    // First delete "eee" and "ccc", get "ddbbbdaa"
    // Then delete "bbb", get "dddaa"
    // Finally delete "ddd", get "aa"
    fmt.Println(removeDuplicates("deeedbbcccbdaa", 3)) // "aa"
    // Example 3:
    // Input: s = "pbbcggttciiippooaais", k = 2
    // Output: "ps"
    fmt.Println(removeDuplicates("pbbcggttciiippooaais", 2)) // "ps"

    fmt.Println(removeDuplicates1("abcd", 2)) // "abcd"
    fmt.Println(removeDuplicates1("deeedbbcccbdaa", 3)) // "aa"
    fmt.Println(removeDuplicates1("pbbcggttciiippooaais", 2)) // "ps"
}