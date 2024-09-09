package main

// 1078. Occurrences After Bigram
// Given two strings first and second, consider occurrences in some text of the form "first second third", 
// where second comes immediately after first, and third comes immediately after second.

// Return an array of all the words third for each occurrence of "first second third".

// Example 1:
// Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
// Output: ["girl","student"]

// Example 2:
// Input: text = "we will we will rock you", first = "we", second = "will"
// Output: ["we","rock"]
 
// Constraints:
//     1 <= text.length <= 1000
//     text consists of lowercase English letters and spaces.
//     All the words in text are separated by a single space.
//     1 <= first.length, second.length <= 10
//     first and second consist of lowercase English letters.
//     text will not have any leading or trailing spaces.

import "fmt"
import "strings"

func findOcurrences(text string, first string, second string) []string {
    res := []string{}
    words := strings.Fields(text)
    for i := 2; i < len(words); i++ {
        if words[i-2] == first && words[i-1] == second { // 前两个单词都匹配，把这个词加入到结果返回
            res = append(res, words[i])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
    // Output: ["girl","student"]
    fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good")) // ["girl","student"]
    // Example 2:
    // Input: text = "we will we will rock you", first = "we", second = "will"
    // Output: ["we","rock"]
    fmt.Println(findOcurrences("we will we will rock you", "we", "will")) // ["we","rock"]
}