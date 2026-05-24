package main

// 3941. Password Strength
// You are given a string password.

// The strength of the password is calculated based on the following rules:
//     1 point for each distinct lowercase letter ('a' to 'z').
//     2 points for each distinct uppercase letter ('A' to 'Z').
//     3 points for each distinct digit ('0' to '9').
//     5 points for each distinct special character from the set "!@#$".

// Each character contributes at most once, even if it appears multiple times.

// Return an integer denoting the strength of the password.

// Example 1:
// Input: password = "aA1!"
// Output: 11
// Explanation:
// The distinct characters are 'a', 'A', '1' and '!'.
// Thus, the strength = 1 + 2 + 3 + 5 = 11.

// Example 2:
// Input: password = "bbB11#"
// Output: 11
// Explanation:
// The distinct characters are 'b', 'B', '1' and '#'.
// Thus, the strength = 1 + 2 + 3 + 5 = 11.​​​​​​​

// Constraints:
//     1 <= password.length <= 10^5
//     password consists of lowercase and uppercase English letters, digits, and special characters from "!@#$".

import "fmt"

func passwordStrength(password string) int {
    res, singlePass := 0, ""
    s := map[string]bool{}
    for _, p := range password {
        if !s[string(p)] {
            singlePass += string(p)
            s[string(p)] = true
        }
    }
    sign := "!@#$"
    mp := map[string]int{}
    for i := 'a'; i <= 'z'; i++ {
        mp[string(i)] = 1
    }
    for i := 'A'; i <= 'Z'; i++ {
        mp[string(i)] = 2
    }
    for i := '0'; i <= '9'; i++ {
        mp[string(i)] = 3
    }
    for _, v := range sign {
        mp[string(v)] = 5
    }
    for _, v := range singlePass {
        res += mp[string(v)]
    }
    return res
}

func passwordStrength1(password string) int {
    res, visited := 0, make([]bool, 256)
    for i := range password{
        b := password[i]
        if visited[b] {
            continue
        }
        visited[b] = true
        switch{
            case b >= 'a' && b <= 'z':
                res += 1
            case b >= 'A' && b <= 'Z':
                res += 2
            case b >= '0' && b <= '9':
                res += 3
            case b == '!' || b == '@' || b == '#' || b == '$':
                res += 5
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: password = "aA1!"
    // Output: 11
    // Explanation:
    // The distinct characters are 'a', 'A', '1' and '!'.
    // Thus, the strength = 1 + 2 + 3 + 5 = 11.
    fmt.Println(passwordStrength("aA1!")) // 11
    // Example 2:
    // Input: password = "bbB11#"
    // Output: 11
    // Explanation:
    // The distinct characters are 'b', 'B', '1' and '#'.
    // Thus, the strength = 1 + 2 + 3 + 5 = 11.​​​​​​​
    fmt.Println(passwordStrength("bbB11#")) // 11

    fmt.Println(passwordStrength("bluefrog")) // 8
    fmt.Println(passwordStrength("leetcode")) // 6
    fmt.Println(passwordStrength("freewu")) // 5

    fmt.Println(passwordStrength1("aA1!")) // 11
    fmt.Println(passwordStrength1("bbB11#")) // 11
    fmt.Println(passwordStrength1("bluefrog")) // 8
    fmt.Println(passwordStrength1("leetcode")) // 6
    fmt.Println(passwordStrength1("freewu")) // 5
}