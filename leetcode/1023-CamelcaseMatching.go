package main

// 1023. Camelcase Matching
// Given an array of strings queries and a string pattern, 
// return a boolean array answer where answer[i] is true if queries[i] matches pattern, and false otherwise.

// A query word queries[i] matches pattern if you can insert lowercase English letters pattern so that it equals the query. 
// You may insert each character at any position and you may not insert any characters.

// Example 1:
// Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FB"
// Output: [true,false,true,true,false]
// Explanation: "FooBar" can be generated like this "F" + "oo" + "B" + "ar".
// "FootBall" can be generated like this "F" + "oot" + "B" + "all".
// "FrameBuffer" can be generated like this "F" + "rame" + "B" + "uffer".

// Example 2:
// Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBa"
// Output: [true,false,true,false,false]
// Explanation: "FooBar" can be generated like this "Fo" + "o" + "Ba" + "r".
// "FootBall" can be generated like this "Fo" + "ot" + "Ba" + "ll".

// Example 3:
// Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBaT"
// Output: [false,true,false,false,false]
// Explanation: "FooBarTest" can be generated like this "Fo" + "o" + "Ba" + "r" + "T" + "est".

// Constraints:
//     1 <= pattern.length, queries.length <= 100
//     1 <= queries[i].length <= 100
//     queries[i] and pattern consist of English letters.

import "fmt"
import "unicode"

func camelMatch(queries []string, pattern string) []bool {
    n, res := len(pattern), make([]bool, len(queries))
    for i, word := range queries {
        index, extraUpper := 0, true
        for j, v := range word {
            if index < n && word[j] == pattern[index] {
                index++
                continue
            }
            if unicode.IsUpper(v) {
                extraUpper = false
                break
            }
        }
        res[i] = extraUpper && index == n 
            
    }
    return res
}

func main() {
    // Example 1:
    // Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FB"
    // Output: [true,false,true,true,false]
    // Explanation: "FooBar" can be generated like this "F" + "oo" + "B" + "ar".
    // "FootBall" can be generated like this "F" + "oot" + "B" + "all".
    // "FrameBuffer" can be generated like this "F" + "rame" + "B" + "uffer".
    fmt.Println(camelMatch([]string{"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}, "FB")) // [true,false,true,true,false]
    // Example 2:
    // Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBa"
    // Output: [true,false,true,false,false]
    // Explanation: "FooBar" can be generated like this "Fo" + "o" + "Ba" + "r".
    // "FootBall" can be generated like this "Fo" + "ot" + "Ba" + "ll".
    fmt.Println(camelMatch([]string{"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}, "FoBa")) // [true,false,true,false,false]
    // Example 3:
    // Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBaT"
    // Output: [false,true,false,false,false]
    // Explanation: "FooBarTest" can be generated like this "Fo" + "o" + "Ba" + "r" + "T" + "est".
    fmt.Println(camelMatch([]string{"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}, "FoBaT")) // [false,true,false,false,false]
}