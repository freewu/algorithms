package main

// 1181. Before and After Puzzle
// Given a list of phrases, generate a list of Before and After puzzles.

// A phrase is a string that consists of lowercase English letters and spaces only. 
// No space appears in the start or the end of a phrase. There are no consecutive spaces in a phrase.

// Before and After puzzles are phrases that are formed by merging two phrases where the last word of the first phrase is the same as the first word of the second phrase.

// Return the Before and After puzzles that can be formed by every two phrases phrases[i] and phrases[j] where i != j. 
// Note that the order of matching two phrases matters, we want to consider both orders.

// You should return a list of distinct strings sorted lexicographically.

// Example 1:
// Input: phrases = ["writing code","code rocks"]
// Output: ["writing code rocks"]

// Example 2:
// Input: phrases = ["mission statement",
//                   "a quick bite to eat",
//                   "a chip off the old block",
//                   "chocolate bar",
//                   "mission impossible",
//                   "a man on a mission",
//                   "block party",
//                   "eat my words",
//                   "bar of soap"]
// Output: ["a chip off the old block party",
//          "a man on a mission impossible",
//          "a man on a mission statement",
//          "a quick bite to eat my words",
//          "chocolate bar of soap"]

// Example 3:
// Input: phrases = ["a","b","a"]
// Output: ["a"]

// Constraints:
//     1 <= phrases.length <= 100
//     1 <= phrases[i].length <= 100

import "fmt"
import "strings"
import "sort"

func beforeAndAfterPuzzles(phrases []string) []string {
    n := len(phrases)
    sp := [][]string{}
    for i := 0; i < n; i++ {
        s := strings.Split(phrases[i], " ")
        sp = append(sp, []string{s[0], s[len(s)-1]})
    }
    mp := map[string]bool{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
           if i == j {
               continue
           }
           if sp[i][0] == sp[j][1] { // 前后词一样的句子合并
                mp[phrases[j] + phrases[i][len(sp[i][0]):]] = true
           }
        }
    }
    res := make([]string, 0)
    for k := range mp{
        res = append(res, k)
    }
    sort.Strings(res)
    return res
}

func main() {
    // Example 1:
    // Input: phrases = ["writing code","code rocks"]
    // Output: ["writing code rocks"]
    fmt.Println(beforeAndAfterPuzzles([]string{"writing code","code rocks"})) // ["writing code rocks"]
    // Example 2:
    // Input: phrases = ["mission statement",
    //                   "a quick bite to eat",
    //                   "a chip off the old block",
    //                   "chocolate bar",
    //                   "mission impossible",
    //                   "a man on a mission",
    //                   "block party",
    //                   "eat my words",
    //                   "bar of soap"]
    // Output: ["a chip off the old block party",
    //          "a man on a mission impossible",
    //          "a man on a mission statement",
    //          "a quick bite to eat my words",
    //          "chocolate bar of soap"]
    arr2 := []string{
        "mission statement",
        "a quick bite to eat",
        "a chip off the old block",
        "chocolate bar",
        "mission impossible",
        "a man on a mission",
        "block party",
        "eat my words",
        "bar of soap",
    }
    fmt.Println(beforeAndAfterPuzzles(arr2)) // ["a chip off the old block party","a man on a mission impossible","a man on a mission statement","a quick bite to eat my words","chocolate bar of soap"]
    // Example 3:
    // Input: phrases = ["a","b","a"]
    // Output: ["a"]
    fmt.Println(beforeAndAfterPuzzles([]string{"a","b","a"})) // ["a"]
}