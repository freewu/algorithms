package main

// 3706. Maximum Distance Between Unequal Words in Array II
// You are given a string array words.

// Find the maximum distance between two distinct indices i and j such that:
//     1. words[i] != words[j], and
//     2. the distance is defined as j - i + 1.

// Return the maximum distance among all such pairs. If no valid pair exists, return 0.

// Example 1:
// Input: words = ["leetcode","leetcode","codeforces"]
// Output: 3
// Explanation:
// In this example, words[0] and words[2] are not equal, and they have the maximum distance 2 - 0 + 1 = 3.

// Example 2:
// Input: words = ["a","b","c","a","a"]
// Output: 4
// Explanation:
// In this example words[1] and words[4] have the largest distance of 4 - 1 + 1 = 4.

// Example 3:
// Input: words = ["z","z","z"]
// Output: 0
// Explanation:
// ​​​​​​​In this example all the words are equal, thus the answer is 0.

// Constraints:
//     1 <= words.length <= 10^5
//     1 <= words[i].length <= 10
//     words[i] consists of lowercase English letters.

import "fmt"

func maxDistance(words []string) int {
    n, a, b := len(words), words[0], words[len(words) - 1]
    for i := 0; i < n; i++ {
        if a != words[n - i - 1] || b != words[i] {
            return n - i
        }
    }
    return 0
}

func main() {
    // Example 1:
    // Input: words = ["leetcode","leetcode","codeforces"]
    // Output: 3
    // Explanation:
    // In this example, words[0] and words[2] are not equal, and they have the maximum distance 2 - 0 + 1 = 3.
    fmt.Println(maxDistance([]string{"leetcode","leetcode","codeforces"})) // 3
    // Example 2:
    // Input: words = ["a","b","c","a","a"]
    // Output: 4
    // Explanation:
    // In this example words[1] and words[4] have the largest distance of 4 - 1 + 1 = 4.
    fmt.Println(maxDistance([]string{"a","b","c","a","a"})) // 4
    // Example 3:
    // Input: words = ["z","z","z"]
    // Output: 0
    // Explanation:
    // ​​​​​​​In this example all the words are equal, thus the answer is 0.   
    fmt.Println(maxDistance([]string{"z","z","z"})) // 0

    fmt.Println(maxDistance([]string{"bluefrog","leetcode","freewu"})) // 3
}