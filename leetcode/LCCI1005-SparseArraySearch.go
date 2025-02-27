package main

// 面试题 10.05. Sparse Array Search LCCI
// Given a sorted array of strings that is interspersed with empty strings, 
// write a method to find the location of a given string.

// Example1:
// Input: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
// Output: -1
// Explanation: Return -1 if s is not in words.

// Example2:
// Input: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
// Output: 4

// Note:
//     1 <= words.length <= 1000000

import "fmt"

func findString(words []string, s string) int {
    for i, v := range words {
        if v == s { return i }
        if v == "" { continue }
        if words[i][0] > s[0] { return -1 }
    }
    return -1
}

func main() {
    // Example1:
    // Input: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
    // Output: -1
    // Explanation: Return -1 if s is not in words.
    fmt.Println(findString([]string{"at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""}, "ta")) // -1
    // Example2:
    // Input: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
    // Output: 4
    fmt.Println(findString([]string{"at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""}, "ball")) // 4
}