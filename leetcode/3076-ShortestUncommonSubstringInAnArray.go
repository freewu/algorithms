package main

// 3076. Shortest Uncommon Substring in an Array
// You are given an array arr of size n consisting of non-empty strings.

// Find a string array answer of size n such that:
//     1. answer[i] is the shortest substring of arr[i] that does not occur as a substring in any other string in arr. 
//        If multiple such substrings exist, answer[i] should be the lexicographically smallest. 
//        And if no such substring exists, answer[i] should be an empty string.

// Return the array answer.

// Example 1:
// Input: arr = ["cab","ad","bad","c"]
// Output: ["ab","","ba",""]
// Explanation: We have the following:
// - For the string "cab", the shortest substring that does not occur in any other string is either "ca" or "ab", we choose the lexicographically smaller substring, which is "ab".
// - For the string "ad", there is no substring that does not occur in any other string.
// - For the string "bad", the shortest substring that does not occur in any other string is "ba".
// - For the string "c", there is no substring that does not occur in any other string.

// Example 2:
// Input: arr = ["abc","bcd","abcd"]
// Output: ["","","abcd"]
// Explanation: We have the following:
// - For the string "abc", there is no substring that does not occur in any other string.
// - For the string "bcd", there is no substring that does not occur in any other string.
// - For the string "abcd", the shortest substring that does not occur in any other string is "abcd".

// Constraints:
//     n == arr.length
//     2 <= n <= 100
//     1 <= arr[i].length <= 20
//     arr[i] consists only of lowercase English letters.

import "fmt"
import "strings"

func shortestSubstrings(arr []string) []string {
    n := len(arr)
    res := make([]string, n)
    for i := 0; i < n; i++ {
        for j := 1; j <= len(arr[i]); j++ {
            for start := 0; start + j <= len(arr[i]); start++ {
                s := arr[i][start : start + j]
                flag := false
                for k := 0; k < n; k++ {
                    if i == k { continue }
                    if strings.Contains(arr[k], s) {
                        flag = true
                        break
                    }
                }
                if !flag && (res[i] == "" || res[i] > s) {
                    res[i] = s
                }
            }
            if res[i] != "" { break }
        }
    }
    return res
}

func shortestSubstrings1(arr []string) []string {
    mp := make(map[string]map[int]bool)
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr[i]); j++ {
            for k := j + 1; k <= len(arr[i]); k++ {
                cur := arr[i][j:k]
                if len(mp[cur]) == 0 {
                    mp[cur] = make(map[int]bool, 2)
                }
                mp[cur][i] = true
            }
        }
    }
    res := make([]string, len(arr))
    for i := 0; i < len(arr); i++ {
        tmp := []string{}
        for j := 0; j < len(arr[i]); j++ {
            for k := j + 1; k <= len(arr[i]); k++ {
                cur := arr[i][j:k]
                if len(mp[cur]) == 1 {
                    tmp = append(tmp, cur)
                }
            }
        }
        if len(tmp) == 0 { continue }
        mn := tmp[0]
        for _, str := range tmp {
            if len(str) < len(mn) {
                mn = str
            } else if len(str) == len(mn) && strings.Compare(str, mn) == -1 {
                mn = str
            }
        }
        res[i] = mn
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = ["cab","ad","bad","c"]
    // Output: ["ab","","ba",""]
    // Explanation: We have the following:
    // - For the string "cab", the shortest substring that does not occur in any other string is either "ca" or "ab", we choose the lexicographically smaller substring, which is "ab".
    // - For the string "ad", there is no substring that does not occur in any other string.
    // - For the string "bad", the shortest substring that does not occur in any other string is "ba".
    // - For the string "c", there is no substring that does not occur in any other string.
    fmt.Println(shortestSubstrings([]string{"cab","ad","bad","c"})) // ["ab","","ba",""]
    // Example 2:
    // Input: arr = ["abc","bcd","abcd"]
    // Output: ["","","abcd"]
    // Explanation: We have the following:
    // - For the string "abc", there is no substring that does not occur in any other string.
    // - For the string "bcd", there is no substring that does not occur in any other string.
    // - For the string "abcd", the shortest substring that does not occur in any other string is "abcd".
    fmt.Println(shortestSubstrings([]string{"abc","bcd","abcd"})) // ["","","abcd"]

    fmt.Println(shortestSubstrings([]string{"bluefrog","leetcode","freewu"})) // ["b","c","w"]

    fmt.Println(shortestSubstrings1([]string{"cab","ad","bad","c"})) // ["ab","","ba",""]
    fmt.Println(shortestSubstrings1([]string{"abc","bcd","abcd"})) // ["","","abcd"]
    fmt.Println(shortestSubstrings1([]string{"bluefrog","leetcode","freewu"})) // ["b","c","w"]
}