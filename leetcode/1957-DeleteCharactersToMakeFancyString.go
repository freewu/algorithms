package main

// 1957. Delete Characters to Make Fancy String
// A fancy string is a string where no three consecutive characters are equal.

// Given a string s, delete the minimum possible number of characters from s to make it fancy.

// Return the final string after the deletion. It can be shown that the answer will always be unique.

// Example 1:
// Input: s = "leeetcode"
// Output: "leetcode"
// Explanation:
// Remove an 'e' from the first group of 'e's to create "leetcode".
// No three consecutive characters are equal, so return "leetcode".

// Example 2:
// Input: s = "aaabaaaa"
// Output: "aabaa"
// Explanation:
// Remove an 'a' from the first group of 'a's to create "aabaaaa".
// Remove two 'a's from the second group of 'a's to create "aabaa".
// No three consecutive characters are equal, so return "aabaa".

// Example 3:
// Input: s = "aab"
// Output: "aab"
// Explanation: No three consecutive characters are equal, so return "aab".

// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of lowercase English letters.

import "fmt"

func makeFancyString(s string) string {
    res, n := []byte{}, len(s)
    if n <= 2 { return s }
    for i := 0; i < n - 2; i++ {
        if s[i] == s[i + 1] && s[i] == s[i + 2] { // 不能 3 个相同字符
            continue
        }
        res = append(res, s[i])
    }
    // 追加后面两个字符
    res = append(res, s[n - 2])
    res = append(res, s[n - 1])
    return string(res)
}

func makeFancyString1(s string) string {
    arr := []byte(s)
    candidate := arr[0]
    count, i, j, n := 0, 0, 0, len(arr)
    for ; j < n; j++ {
        if candidate != arr[j] { // 不是连续的
            count = 0
            candidate = arr[j]
        } 
        count++
        if count > 2 { continue }
        arr[i] = arr[j]
        i++
    }
    return string(arr[:i])
}

func main() {
    // Example 1:
    // Input: s = "leeetcode"
    // Output: "leetcode"
    // Explanation:
    // Remove an 'e' from the first group of 'e's to create "leetcode".
    // No three consecutive characters are equal, so return "leetcode".
    fmt.Println(makeFancyString("leeetcode")) // leetcode
    // Example 2:
    // Input: s = "aaabaaaa"
    // Output: "aabaa"
    // Explanation:
    // Remove an 'a' from the first group of 'a's to create "aabaaaa".
    // Remove two 'a's from the second group of 'a's to create "aabaa".
    // No three consecutive characters are equal, so return "aabaa".
    fmt.Println(makeFancyString("aaabaaaa")) // aabaa
    // Example 3:
    // Input: s = "aab"
    // Output: "aab"
    // Explanation: No three consecutive characters are equal, so return "aab".
    fmt.Println(makeFancyString("aab")) // aab

    fmt.Println(makeFancyString("a")) // a

    fmt.Println(makeFancyString1("leeetcode")) // leetcode
    fmt.Println(makeFancyString1("aaabaaaa")) // aabaa
    fmt.Println(makeFancyString1("aab")) // aab
    fmt.Println(makeFancyString1("a")) // a
}