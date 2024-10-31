package main

// 2900. Longest Unequal Adjacent Groups Subsequence I
// You are given a string array words and a binary array groups both of length n, where words[i] is associated with groups[i].

// Your task is to select the longest alternating subsequence from words. 
// A subsequence of words is alternating if for any two consecutive strings in the sequence, 
// their corresponding elements in the binary array groups differ. 
// Essentially, you are to choose strings such that adjacent elements have non-matching corresponding bits in the groups array.

// Formally, you need to find the longest subsequence of an array of indices [0, 1, ..., n - 1] denoted as [i0, i1, ..., ik-1], 
// such that groups[ij] != groups[ij+1] for each 0 <= j < k - 1 and then find the words corresponding to these indices.

// Return the selected subsequence. 
// If there are multiple answers, return any of them.

// Note: The elements in words are distinct.

// Example 1:
// Input: words = ["e","a","b"], groups = [0,0,1]
// Output: ["e","b"]
// Explanation: A subsequence that can be selected is ["e","b"] because groups[0] != groups[2]. 
// Another subsequence that can be selected is ["a","b"] because groups[1] != groups[2]. 
// It can be demonstrated that the length of the longest subsequence of indices that satisfies the condition is 2.

// Example 2:
// Input: words = ["a","b","c","d"], groups = [1,0,1,1]
// Output: ["a","b","c"]
// Explanation: A subsequence that can be selected is ["a","b","c"] because groups[0] != groups[1] and groups[1] != groups[2]. 
// Another subsequence that can be selected is ["a","b","d"] because groups[0] != groups[1] and groups[1] != groups[3]. 
// It can be shown that the length of the longest subsequence of indices that satisfies the condition is 3.

// Constraints:
//     1 <= n == words.length == groups.length <= 100
//     1 <= words[i].length <= 10
//     groups[i] is either 0 or 1.
//     words consists of distinct strings.
//     words[i] consists of lowercase English letters.

import "fmt"

// GREEDY
func getLongestSubsequence(words []string, groups []int) []string {
    res, prev := []string{words[0]}, groups[0]
    for i := 1; i < len(words); i++ {
        if prev != groups[i] {
            prev = groups[i]
            res = append(res, words[i])
        }
    }
    return res
}

// BRUTE FORCE
func getLongestSubsequence1(words []string, groups []int) []string {
    res := make([]string, 0)
    for i := 0; i < len(words); i++ {
        prev, curr := groups[i], []string{words[i]}
        for j := i; j < len(words); j++ {
            if prev != groups[j] {
                prev = groups[j]
                curr = append(curr, words[j])
            }
        }
        if len(curr) > len(res) {
            res = curr
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["e","a","b"], groups = [0,0,1]
    // Output: ["e","b"]
    // Explanation: A subsequence that can be selected is ["e","b"] because groups[0] != groups[2]. 
    // Another subsequence that can be selected is ["a","b"] because groups[1] != groups[2]. 
    // It can be demonstrated that the length of the longest subsequence of indices that satisfies the condition is 2.
    fmt.Println(getLongestSubsequence([]string{"e","a","b"}, []int{0,0,1})) // ["e","b"]
    // Example 2:
    // Input: words = ["a","b","c","d"], groups = [1,0,1,1]
    // Output: ["a","b","c"]
    // Explanation: A subsequence that can be selected is ["a","b","c"] because groups[0] != groups[1] and groups[1] != groups[2]. 
    // Another subsequence that can be selected is ["a","b","d"] because groups[0] != groups[1] and groups[1] != groups[3]. 
    // It can be shown that the length of the longest subsequence of indices that satisfies the condition is 3.
    fmt.Println(getLongestSubsequence([]string{"a","b","c","d"}, []int{1,0,1,1})) // ["a","b","c"]

    fmt.Println(getLongestSubsequence1([]string{"e","a","b"}, []int{0,0,1})) // ["e","b"]
    fmt.Println(getLongestSubsequence1([]string{"a","b","c","d"}, []int{1,0,1,1})) // ["a","b","c"]
}