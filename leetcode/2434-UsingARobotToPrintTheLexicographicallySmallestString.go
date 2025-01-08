package main

// 2434. Using a Robot to Print the Lexicographically Smallest String
// You are given a string s and a robot that currently holds an empty string t. 
// Apply one of the following operations until s and t are both empty:
//     1. Remove the first character of a string s and give it to the robot. 
//        The robot will append this character to the string t.
//     2. Remove the last character of a string t and give it to the robot. 
//        The robot will write this character on paper.

// Return the lexicographically smallest string that can be written on the paper.

// Example 1:
// Input: s = "zza"
// Output: "azz"
// Explanation: Let p denote the written string.
// Initially p="", s="zza", t="".
// Perform first operation three times p="", s="", t="zza".
// Perform second operation three times p="azz", s="", t="".

// Example 2:
// Input: s = "bac"
// Output: "abc"
// Explanation: Let p denote the written string.
// Perform first operation twice p="", s="c", t="ba". 
// Perform second operation twice p="ab", s="c", t="". 
// Perform first operation p="ab", s="", t="c". 
// Perform second operation p="abc", s="", t="".

// Example 3:
// Input: s = "bdda"
// Output: "addb"
// Explanation: Let p denote the written string.
// Initially p="", s="bdda", t="".
// Perform first operation four times p="", s="", t="bdda".
// Perform second operation four times p="addb", s="", t="".

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of only English lowercase letters.

import "fmt"
import "slices"

func robotWithString(s string) string {
    count := make([]int, 26)
    for _, v := range s {
        count[v - 'a']++
    }
    stack, res, a := []byte{}, []byte{}, byte('a')
    for i := range s {
        count[s[i] - 'a']--
        for a < 'z' && count[a - 'a'] == 0 {
            a++
        }
        stack = append(stack, s[i])
        for len(stack) > 0 && stack[len(stack) - 1] <= a {
            res = append(res, stack[len(stack) - 1])
            stack = stack[:len(stack) - 1]
        }
    }
    return string(res)
}

func robotWithString1(s string) string {
    n := len(s)
    dp := make([]byte, n)
    dp[n-1] = s[n-1]
    for i := n - 2; i >= 0; i-- {
        dp[i] = min(dp[i+1], s[i])
    }
    index, stack, res := 0, []byte{}, make([]byte, 0, n)
    for i := 0; i < n; i++ {
        for len(stack) > 0 && stack[len(stack) - 1] <= dp[index] {
            res = append(res, stack[len(stack) - 1])
            stack = stack[:len(stack) - 1]
            index = i
        }
        stack = append(stack, s[i])
    }
    slices.Reverse(stack)
    res = append(res, stack...)
    return string(res)
}

func robotWithString2(s string) string {
    low, n, count := 0, len(s), [26]uint16{}
    for _, v := range s {
        count[v - 'a']++
    }
    res, stack := make([]byte, 0, n), make([]byte, 0, n)
    for _, v := range s {
        stack = append(stack, byte(v))
        count[v - 'a']--
        for low < 25 && count[low] == 0 {
            low++
        }
        for len(stack) > 0 && int(stack[len(stack) - 1] - 'a') <= low {
            res = append(res, stack[len(stack) - 1])
            stack = stack[:len(stack) - 1]
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "zza"
    // Output: "azz"
    // Explanation: Let p denote the written string.
    // Initially p="", s="zza", t="".
    // Perform first operation three times p="", s="", t="zza".
    // Perform second operation three times p="azz", s="", t="".
    fmt.Println(robotWithString("zza")) // "azz"
    // Example 2:
    // Input: s = "bac"
    // Output: "abc"
    // Explanation: Let p denote the written string.
    // Perform first operation twice p="", s="c", t="ba". 
    // Perform second operation twice p="ab", s="c", t="". 
    // Perform first operation p="ab", s="", t="c". 
    // Perform second operation p="abc", s="", t="".
    fmt.Println(robotWithString("bac")) // "abc"
    // Example 3:
    // Input: s = "bdda"
    // Output: "addb"
    // Explanation: Let p denote the written string.
    // Initially p="", s="bdda", t="".
    // Perform first operation four times p="", s="", t="bdda".
    // Perform second operation four times p="addb", s="", t="".
    fmt.Println(robotWithString("bdda")) // "addb"

    fmt.Println(robotWithString("bluefrog")) // "befgorul"
    fmt.Println(robotWithString("leetcode")) // "cdeoteel"

    fmt.Println(robotWithString1("zza")) // "azz"
    fmt.Println(robotWithString1("bac")) // "abc"
    fmt.Println(robotWithString1("bdda")) // "addb"
    fmt.Println(robotWithString1("bluefrog")) // "befgorul"
    fmt.Println(robotWithString1("leetcode")) // "cdeoteel"

    fmt.Println(robotWithString2("zza")) // "azz"
    fmt.Println(robotWithString2("bac")) // "abc"
    fmt.Println(robotWithString2("bdda")) // "addb"
    fmt.Println(robotWithString2("bluefrog")) // "befgorul"
    fmt.Println(robotWithString2("leetcode")) // "cdeoteel"
}