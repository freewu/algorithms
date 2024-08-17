package main

// 833. Find And Replace in String
// You are given a 0-indexed string s that you must perform k replacement operations on. 
// The replacement operations are given as three 0-indexed parallel arrays, indices, sources, and targets, all of length k.

// To complete the ith replacement operation:
//     1. Check if the substring sources[i] occurs at index indices[i] in the original string s.
//     2. If it does not occur, do nothing.
//     3. Otherwise if it does occur, replace that substring with targets[i].

// For example, if s = "abcd", indices[i] = 0, sources[i] = "ab", and targets[i] = "eee", 
// then the result of this replacement will be "eeecd".

// All replacement operations must occur simultaneously, 
// meaning the replacement operations should not affect the indexing of each other. 
// The testcases will be generated such that the replacements will not overlap.

//     For example, a testcase with s = "abc", indices = [0, 1], and sources = ["ab","bc"] will not be generated because the "ab" and "bc" replacements overlap.

// Return the resulting string after performing all replacement operations on s.
// A substring is a contiguous sequence of characters in a string.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/12/833-ex1.png" />
// Input: s = "abcd", indices = [0, 2], sources = ["a", "cd"], targets = ["eee", "ffff"]
// Output: "eeebffff"
// Explanation:
// "a" occurs at index 0 in s, so we replace it with "eee".
// "cd" occurs at index 2 in s, so we replace it with "ffff".

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/12/833-ex2-1.png" />
// Input: s = "abcd", indices = [0, 2], sources = ["ab","ec"], targets = ["eee","ffff"]
// Output: "eeecd"
// Explanation:
// "ab" occurs at index 0 in s, so we replace it with "eee".
// "ec" does not occur at index 2 in s, so we do nothing.

// Constraints:
//     1 <= s.length <= 1000
//     k == indices.length == sources.length == targets.length
//     1 <= k <= 100
//     0 <= indexes[i] < s.length
//     1 <= sources[i].length, targets[i].length <= 50
//     s consists of only lowercase English letters.
//     sources[i] and targets[i] consist of only lowercase English letters.

import "fmt"
import "strings"
import "sort"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
    chars := strings.Split(s, "") // convert string to []string
    for i := 0; i < len(indices); i++ {
        index, source, target := indices[i], sources[i], targets[i]
        if len(source) <= len(s[index:]) && s[index:index+len(source)] == source {
            chars[index] = target
            for j := index + 1; j < index + len(source); j++ { // fill in empty string
                chars[j] = ""
            }
        }
    }
    return strings.Join(chars, "")
}

func findReplaceString1(s string, indices []int, sources []string, targets []string) string {
    type Rep struct {
        i int
        s string
        t string
    }
    reps := []*Rep{}
    for i := range indices {
        reps = append(reps, &Rep{ indices[i], sources[i], targets[i], })
    }
    sort.Slice(reps, func(i, j int) bool {
        return reps[i].i < reps[j].i
    })
    res, j := "", 0
    for i := 0; i < len(s); {
        if j == len(reps) || i != reps[j].i {
            res += string(s[i])
            i++
            continue
        }
        l := len(reps[j].s)
        if i+l <= len(s) && s[i:i+l] == reps[j].s {
            res += reps[j].t
            i += l
        }
        j++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/12/833-ex1.png" />
    // Input: s = "abcd", indices = [0, 2], sources = ["a", "cd"], targets = ["eee", "ffff"]
    // Output: "eeebffff"
    // Explanation:
    // "a" occurs at index 0 in s, so we replace it with "eee".
    // "cd" occurs at index 2 in s, so we replace it with "ffff".
    fmt.Println(findReplaceString("abcd", []int{0, 2},[]string{"a", "cd"}, []string{"eee", "ffff"})) // "eeebffff"
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/12/833-ex2-1.png" />
    // Input: s = "abcd", indices = [0, 2], sources = ["ab","ec"], targets = ["eee","ffff"]
    // Output: "eeecd"
    // Explanation:
    // "ab" occurs at index 0 in s, so we replace it with "eee".
    // "ec" does not occur at index 2 in s, so we do nothing.
    fmt.Println(findReplaceString("abcd", []int{0, 2},[]string{"ab","ec"}, []string{"eee","ffff"})) // "eeecd"

    fmt.Println(findReplaceString1("abcd", []int{0, 2},[]string{"a", "cd"}, []string{"eee", "ffff"})) // "eeebffff"
    fmt.Println(findReplaceString1("abcd", []int{0, 2},[]string{"ab","ec"}, []string{"eee","ffff"})) // "eeecd"
}