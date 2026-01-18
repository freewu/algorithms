package main

// 3816. Lexicographically Smallest String After Deleting Duplicate Characters
// You are given a string s that consists of lowercase English letters.

// You can perform the following operation any number of times (possibly zero times):
//     Choose any letter that appears at least twice in the current string s and delete any one occurrence.

// Return the lexicographically smallest resulting string that can be formed this way.

// Example 1:
// Input: s = "aaccb"
// Output: "aacb"
// Explanation:
// We can form the strings "acb", "aacb", "accb", and "aaccb". "aacb" is the lexicographically smallest one.
// For example, we can obtain "aacb" by choosing 'c' and deleting its first occurrence.

// Example 2:
// Input: s = "z"
// Output: "z"
// Explanation:
// We cannot perform any operations. The only string we can form is "z".

// Constraints:
//     1 <= s.length <= 10^5
//     s contains lowercase English letters only.

import "fmt"

// 单调栈
func lexSmallestAfterDeletion(s string) string {
    left := [26]int{}
    for _, v := range s {
        left[v - 'a']++
    }
    stack := []rune{}
    for _, ch := range s {
        // 如果 ch 比栈顶小，移除栈顶，可以让字典序更小
        for len(stack) > 0 && ch < stack[len(stack) - 1] && left[stack[len(stack) - 1] - 'a'] > 1 {
            left[stack[len(stack) - 1] - 'a']--
            stack = stack[: len(stack) - 1]
        }
        stack = append(stack, ch)
    }
    // 最后，移除末尾的重复字母，可以让字典序更小
    for left[stack[len(stack) - 1] - 'a'] > 1 {
        left[stack[len(stack) - 1] - 'a']--
        stack = stack[:len(stack) - 1]
    }
    return string(stack)
}

func main() {
    // Example 1:
    // Input: s = "aaccb"
    // Output: "aacb"
    // Explanation:
    // We can form the strings "acb", "aacb", "accb", and "aaccb". "aacb" is the lexicographically smallest one.
    // For example, we can obtain "aacb" by choosing 'c' and deleting its first occurrence.
    fmt.Println(lexSmallestAfterDeletion("aaccb")) // aacb
    // Example 2:
    // Input: s = "z"
    // Output: "z"
    // Explanation:
    // We cannot perform any operations. The only string we can form is "z".
    fmt.Println(lexSmallestAfterDeletion("z")) // z

    fmt.Println(lexSmallestAfterDeletion("bluefrog")) // bluefrog
    fmt.Println(lexSmallestAfterDeletion("leetcode")) // leetcod
}