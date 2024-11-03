package main

// 796. Rotate String
// Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.
// A shift on s consists of moving the leftmost character of s to the rightmost position.
// For example, if s = "abcde", then it will be "bcdea" after one shift.

// Example 1:
// Input: s = "abcde", goal = "cdeab"
// Output: true

// Example 2:
// Input: s = "abcde", goal = "abced"
// Output: false

// Constraints:
//     1 <= s.length, goal.length <= 100
//     s and goal consist of lowercase English letters.

import "fmt"
import "strings"

func rotateString(s string, goal string) bool {
    return len(s) == len(goal) && strings.Contains(s+s, goal)
}

func rotateString1(s string, goal string) bool {
    arr := []byte(s) 
    for _, v := range arr {
        arr = arr[1:] // 从头移动到尾部
        arr = append(arr, v)
        if string(arr) == goal {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: s = "abcde", goal = "cdeab"
    // Output: true
    fmt.Println(rotateString("abcde","cdeab")) // true
    // Example 2:
    // Input: s = "abcde", goal = "abced"
    // Output: false
    fmt.Println(rotateString("abcde","abced")) // false

    fmt.Println(rotateString1("abcde","cdeab")) // true
    fmt.Println(rotateString1("abcde","abced")) // false
}