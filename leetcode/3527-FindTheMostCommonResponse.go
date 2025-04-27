package main

// 3527. Find the Most Common Response
// You are given a 2D string array responses where each responses[i] is an array of strings representing survey responses from the ith day.

// Return the most common response across all days after removing duplicate responses within each responses[i]. 
// If there is a tie, return the lexicographically smallest response.

// Example 1:
// Input: responses = [["good","ok","good","ok"],["ok","bad","good","ok","ok"],["good"],["bad"]]
// Output: "good"
// Explanation:
// After removing duplicates within each list, responses = [["good", "ok"], ["ok", "bad", "good"], ["good"], ["bad"]].
// "good" appears 3 times, "ok" appears 2 times, and "bad" appears 2 times.
// Return "good" because it has the highest frequency.

// Example 2:
// Input: responses = [["good","ok","good"],["ok","bad"],["bad","notsure"],["great","good"]]
// Output: "bad"
// Explanation:
// After removing duplicates within each list we have responses = [["good", "ok"], ["ok", "bad"], ["bad", "notsure"], ["great", "good"]].
// "bad", "good", and "ok" each occur 2 times.
// The output is "bad" because it is the lexicographically smallest amongst the words with the highest frequency.

// Constraints:
//     1 <= responses.length <= 1000
//     1 <= responses[i].length <= 1000
//     1 <= responses[i][j].length <= 10
//     responses[i][j] consists of only lowercase English letters

import "fmt"

func findCommonResponse(responses [][]string) string {
    res, mx := "", 0
    mp := make(map[string]int)
    for _, row := range responses {
        set := make(map[string]bool)
        for _, v := range row {
            if set[v] { continue }
            set[v] = true
        }
        for k := range set {
            mp[k]++
        }
    }
    for k, v := range mp {
        if mx < v {
            res, mx = k, v
        } else if mx == v {
            if res > k {
                res = k
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: responses = [["good","ok","good","ok"],["ok","bad","good","ok","ok"],["good"],["bad"]]
    // Output: "good"
    // Explanation:
    // After removing duplicates within each list, responses = [["good", "ok"], ["ok", "bad", "good"], ["good"], ["bad"]].
    // "good" appears 3 times, "ok" appears 2 times, and "bad" appears 2 times.
    // Return "good" because it has the highest frequency.
    fmt.Println(findCommonResponse([][]string{{"good","ok","good","ok"},{"ok","bad","good","ok","ok"},{"good"},{"bad"}})) // "good"
    // Example 2:
    // Input: responses = [["good","ok","good"],["ok","bad"],["bad","notsure"],["great","good"]]
    // Output: "bad"
    // Explanation:
    // After removing duplicates within each list we have responses = [["good", "ok"], ["ok", "bad"], ["bad", "notsure"], ["great", "good"]].
    // "bad", "good", and "ok" each occur 2 times.
    // The output is "bad" because it is the lexicographically smallest amongst the words with the highest frequency.
    fmt.Println(findCommonResponse([][]string{{"good","ok","good"},{"ok","bad"},{"bad","notsure"},{"great","good"}})) // "bad"
}