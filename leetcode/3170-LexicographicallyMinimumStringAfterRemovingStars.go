package main

// 3170. Lexicographically Minimum String After Removing Stars
// You are given a string s. It may contain any number of '*' characters. 
// Your task is to remove all '*' characters.

// While there is a '*', do the following operation:
//     1. Delete the leftmost '*' and the smallest non-'*' character to its left. 
//        If there are several smallest characters, you can delete any of them.

// Return the lexicographically smallest resulting string after removing all '*' characters.

// Example 1:
// Input: s = "aaba*"
// Output: "aab"
// Explanation:
// We should delete one of the 'a' characters with '*'. If we choose s[3], s becomes the lexicographically smallest.

// Example 2:
// Input: s = "abc"
// Output: "abc"
// Explanation:
// There is no '*' in the string.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters and '*'.
//     The input is generated such that it is possible to delete all '*' characters.

import "fmt"
import "sort"

func clearStars(s string) string {
    indexs, deleted := [26][]int{}, make(map[int]bool)
    for i := range s {
        if s[i] != '*' {
            indexs[int(s[i]-'a')] = append(indexs[int(s[i]-'a')], i)
            continue
        }
        for j := 0; j < 26 ; j++ {
            if len(indexs[j]) > 0 {
                pos :=  indexs[j][len(indexs[j]) - 1]
                indexs[j] = indexs[j][:len(indexs[j]) - 1]
                deleted[pos] = true
                break
            }
        }
    }
    res := []byte{}
    for i := range s {
        if !deleted[i] && s[i] != '*' {
            res = append(res, s[i])
        }
    }
    return string(res)
}

func clearStars1(s string) string {
    arr := make([][]int, 26)
    for i, v := range s {
        if v != '*' {
            index := v - 'a'
            arr[index] = append(arr[index], i)
            continue
        }
        for j, _ := range arr { // *
            if len(arr[j]) > 0 {
                arr[j] = arr[j][:len(arr[j]) - 1]
                break
            }
        }
    }
    order := []int{}
    for _, v := range arr {
        order = append(order, v...)
    }
    sort.Ints(order)
    res := make([]byte, len(order))
    for i, v := range order {
        res[i] = s[v]
    }
    return string(res)
}

func clearStars2(s string) string {
    arr, n, pos := []byte(s), len(s), [26][]int{}
    for i := 0; i < n; i++ {
        if arr[i] != '*' {
            c := arr[i] - 'a'
            pos[c] = append(pos[c], i)
            continue
        }
        j := 0
        for ; j < 26; j++ {
            if len(pos[j]) > 0 { break }
        }
        p := pos[j][len(pos[j]) - 1]
        pos[j] = pos[j][:len(pos[j])-1]
        arr[p] = '*'
    }
    res := make([]byte, 0, n)
    for i := 0; i < n; i++ {
        if arr[i] == '*' { continue }
        res = append(res, arr[i])
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "aaba*"
    // Output: "aab"
    // Explanation:
    // We should delete one of the 'a' characters with '*'. If we choose s[3], s becomes the lexicographically smallest.
    fmt.Println(clearStars("aaba*")) // aab
    // Example 2:
    // Input: s = "abc"
    // Output: "abc"
    // Explanation:
    // There is no '*' in the string.
    fmt.Println(clearStars("abc*")) // abc

    fmt.Println(clearStars("bleufrog*")) // leufrog
    fmt.Println(clearStars("leet*code")) // letcode

    fmt.Println(clearStars1("aaba*")) // aab
    fmt.Println(clearStars1("abc*")) // abc
    fmt.Println(clearStars1("bleufrog*")) // leufrog
    fmt.Println(clearStars1("leet*code")) // letcode

    fmt.Println(clearStars2("aaba*")) // aab
    fmt.Println(clearStars2("abc*")) // abc
    fmt.Println(clearStars2("bleufrog*")) // leufrog
    fmt.Println(clearStars2("leet*code")) // letcode
}