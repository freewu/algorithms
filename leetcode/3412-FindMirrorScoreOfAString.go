package main

// 3412. Find Mirror Score of a String
// You are given a string s.

// We define the mirror of a letter in the English alphabet as its corresponding letter when the alphabet is reversed.
// For example, the mirror of 'a' is 'z', and the mirror of 'y' is 'b'.

// Initially, all characters in the string s are unmarked.

// You start with a score of 0, and you perform the following process on the string s:
//     1. Iterate through the string from left to right.
//     2. At each index i, find the closest unmarked index j such that j < i and s[j] is the mirror of s[i]. 
//        Then, mark both indices i and j, and add the value i - j to the total score.
//     3. If no such index j exists for the index i, move on to the next index without making any changes.

// Return the total score at the end of the process.

// Example 1:
// Input: s = "aczzx"
// Output: 5
// Explanation:
// i = 0. There is no index j that satisfies the conditions, so we skip.
// i = 1. There is no index j that satisfies the conditions, so we skip.
// i = 2. The closest index j that satisfies the conditions is j = 0, so we mark both indices 0 and 2, and then add 2 - 0 = 2 to the score.
// i = 3. There is no index j that satisfies the conditions, so we skip.
// i = 4. The closest index j that satisfies the conditions is j = 1, so we mark both indices 1 and 4, and then add 4 - 1 = 3 to the score.

// Example 2:
// Input: s = "abcdef"
// Output: 0
// Explanation:
// For each index i, there is no index j that satisfies the conditions.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters.

import "fmt"

func calculateScore(s string) int64 {
    res, mp := 0, [26][]int{}
    for i := range s {
        mirror := ('z' - s[i])
        arr := mp[mirror]
        if len(arr) > 0 {
            // have mirror index -> pop the closest mirror index
            prev := mp[mirror][len(arr) - 1]
            mp[mirror] = mp[mirror][:len(arr) - 1]
            res += (i - prev)
            // marked and then proceed to the next character
            continue
        }
        // add the new character and its index
        mp[s[i] - 'a'] = append(mp[s[i] - 'a'], i)
    }
    return int64(res)
}

func calculateScore1(s string) int64 {
    res, n := 0, len(s)
    arr := make([][]int, 26)
    for i := 0; i < n; i++ {
        c := s[i] - 'a'
        if len(arr[25 - c]) > 0 {
            res += (i - arr[25 - c][len(arr[25 - c]) - 1])
            arr[25 - c] = arr[25 - c][:len(arr[25 - c]) - 1]
        } else {
            arr[c] = append(arr[c], i)
        }
    }
    return  int64(res)
}

func main() {
    // Example 1:
    // Input: s = "aczzx"
    // Output: 5
    // Explanation:
    // i = 0. There is no index j that satisfies the conditions, so we skip.
    // i = 1. There is no index j that satisfies the conditions, so we skip.
    // i = 2. The closest index j that satisfies the conditions is j = 0, so we mark both indices 0 and 2, and then add 2 - 0 = 2 to the score.
    // i = 3. There is no index j that satisfies the conditions, so we skip.
    // i = 4. The closest index j that satisfies the conditions is j = 1, so we mark both indices 1 and 4, and then add 4 - 1 = 3 to the score.
    fmt.Println(calculateScore("aczzx")) // 5
    // Example 2:
    // Input: s = "abcdef"
    // Output: 0
    // Explanation:
    // For each index i, there is no index j that satisfies the conditions.
    fmt.Println(calculateScore("abcdef")) // 0

    fmt.Println(calculateScore("bluefrog")) // 7
    fmt.Println(calculateScore("leetcode")) // 5

    fmt.Println(calculateScore1("aczzx")) // 5
    fmt.Println(calculateScore1("abcdef")) // 0
    fmt.Println(calculateScore1("bluefrog")) // 7
    fmt.Println(calculateScore1("leetcode")) // 5
}