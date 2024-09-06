package main

// 1528. Shuffle String
// You are given a string s and an integer array indices of the same length. 
// The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

// Return the shuffled string.

// Example 1:
// Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// Output: "leetcode"
// Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.

// Example 2:
// Input: s = "abc", indices = [0,1,2]
// Output: "abc"
// Explanation: After shuffling, each character remains in its position.

// Constraints:
//     s.length == indices.length == n
//     1 <= n <= 100
//     s consists of only lowercase English letters.
//     0 <= indices[i] < n
//     All values of indices are unique.

import "fmt"

func restoreString(s string, indices []int) string {
    res := make([]byte, len(indices))
    for i := range indices {
        res[indices[i]] = s[i]
    }
    //for i, index := range indices { res[index] = s[i] }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
    // Output: "leetcode"
    // Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.
    fmt.Println(restoreString("codeleet", []int{4,5,6,7,0,2,1,3})) // "leetcode"
    // Example 2:
    // Input: s = "abc", indices = [0,1,2]
    // Output: "abc"
    // Explanation: After shuffling, each character remains in its position.
    fmt.Println(restoreString("abc", []int{0,1,2})) // "abc"
}