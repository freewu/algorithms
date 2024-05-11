package main

// 1055. Shortest Way to Form String
// A subsequence of a string is a new string that is formed from the original string 
// by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. 
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

// Given two strings source and target, return the minimum number of subsequences of source such that their concatenation equals target. 
// If the task is impossible, return -1.
 
// Example 1:
// Input: source = "abc", target = "abcbc"
// Output: 2
// Explanation: The target "abcbc" can be formed by "abc" and "bc", which are subsequences of source "abc".

// Example 2:
// Input: source = "abc", target = "acdbc"
// Output: -1
// Explanation: The target string cannot be constructed from the subsequences of source string due to the character "d" in target string.

// Example 3:
// Input: source = "xyz", target = "xzyxz"
// Output: 3
// Explanation: The target string can be constructed as follows "xz" + "y" + "xz".

// Constraints:
//     1 <= source.length, target.length <= 1000
//     source and target consist of lowercase English letters.

import "fmt"

func shortestWay(source string, target string) int {
    count := 0
    for i :=0; i <len(target); { // 以滑动窗口的方式扫 target 字符串找 source 的子序列
        prei := i
        for j:= 0; j < len(source) && i < len(target); j++ { // 窗口大小为 source 的长度，每次向右移动的距离是source子序列的长度
            if target[i] == source[j] {
                i++
            }
        }
        if prei == i {
            return -1
        }
        count++
    }
    return count
}

func main() {
    // Example 1:
    // Input: source = "abc", target = "abcbc"
    // Output: 2
    // Explanation: The target "abcbc" can be formed by "abc" and "bc", which are subsequences of source "abc".
    fmt.Println(shortestWay("abc","abcbc")) // 2
    // Example 2:
    // Input: source = "abc", target = "acdbc"
    // Output: -1
    // Explanation: The target string cannot be constructed from the subsequences of source string due to the character "d" in target string.
    fmt.Println(shortestWay("abc","acdbc")) // -1
    // Example 3:
    // Input: source = "xyz", target = "xzyxz"
    // Output: 3
    // Explanation: The target string can be constructed as follows "xz" + "y" + "xz".
    fmt.Println(shortestWay("xyz","xzyxz")) // 3
}