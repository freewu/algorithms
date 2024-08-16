package main

// 884. Uncommon Words from Two Sentences
// A sentence is a string of single-space separated words where each word consists only of lowercase letters.
// A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

// Given two sentences s1 and s2, return a list of all the uncommon words. 
// You may return the answer in any order.

// Example 1:
// Input: s1 = "this apple is sweet", s2 = "this apple is sour"
// Output: ["sweet","sour"]

// Example 2:
// Input: s1 = "apple apple", s2 = "banana"
// Output: ["banana"]

// Constraints:
//     1 <= s1.length, s2.length <= 200
//     s1 and s2 consist of lowercase English letters and spaces.
//     s1 and s2 do not have leading or trailing spaces.
//     All the words in s1 and s2 are separated by a single space.

import "fmt"
import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
    res, mp := []string{}, make(map[string]int)
    arr1 := strings.Split(s1, " ") 
    for _, v := range arr1 { mp[v]++ }
    arr2 := strings.Split(s2, " ") 
    for _, v := range arr2 { mp[v]++ }
    for k, v := range mp {
        if v == 1 {
            res = append(res, k)
        }
    }
    return res
}

// do not use strings lib
func uncommonFromSentences1(s1 string, s2 string) []string {
    res, mp, index := []string{}, make(map[string]int), 0
    s1 += " "
    for i := 0; i < len(s1); i++ {
        if s1[i] == ' ' {
            if index < i {
                mp[s1[index:i]]++
            }
            index = i+1
        }
    }
    index = 0
    s2 += " "
    for i := 0; i < len(s2); i++ {
        if s2[i] == ' ' {
            if index < i {
                mp[s2[index:i]]++
            }
            index = i+1
        }
    }
    for k, v := range mp {
        if v == 1 {
            res = append(res, k)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s1 = "this apple is sweet", s2 = "this apple is sour"
    // Output: ["sweet","sour"]
    fmt.Println(uncommonFromSentences("this apple is sweet","this apple is sour")) // ["sweet","sour"]
    // Example 2:
    // Input: s1 = "apple apple", s2 = "banana"
    // Output: ["banana"]
    fmt.Println(uncommonFromSentences("apple apple","banana")) // ["banana"]

    fmt.Println(uncommonFromSentences1("this apple is sweet","this apple is sour")) // ["sweet","sour"]
    fmt.Println(uncommonFromSentences1("apple apple","banana")) // ["banana"]
}