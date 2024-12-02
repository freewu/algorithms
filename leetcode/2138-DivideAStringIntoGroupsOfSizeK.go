package main

// 2138. Divide a String Into Groups of Size k
// A string s can be partitioned into groups of size k using the following procedure:
//     1. The first group consists of the first k characters of the string, 
//        the second group consists of the next k characters of the string, and so on. 
//        Each character can be a part of exactly one group.
//     2. For the last group, if the string does not have k characters remaining, 
//        a character fill is used to complete the group.

// Note that the partition is done so that after removing the fill character from the last group (if it exists) and concatenating all the groups in order, 
// the resultant string should be s.

// Given the string s, the size of each group k and the character fill, 
// return a string array denoting the composition of every group s has been divided into, using the above procedure.

// Example 1:
// Input: s = "abcdefghi", k = 3, fill = "x"
// Output: ["abc","def","ghi"]
// Explanation:
// The first 3 characters "abc" form the first group.
// The next 3 characters "def" form the second group.
// The last 3 characters "ghi" form the third group.
// Since all groups can be completely filled by characters from the string, we do not need to use fill.
// Thus, the groups formed are "abc", "def", and "ghi".

// Example 2:
// Input: s = "abcdefghij", k = 3, fill = "x"
// Output: ["abc","def","ghi","jxx"]
// Explanation:
// Similar to the previous example, we are forming the first three groups "abc", "def", and "ghi".
// For the last group, we can only use the character 'j' from the string. To complete this group, we add 'x' twice.
// Thus, the 4 groups formed are "abc", "def", "ghi", and "jxx".

// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters only.
//     1 <= k <= 100
//     fill is a lowercase English letter.

import "fmt"
import "strings"

func divideString(s string, k int, fill byte) []string {
    res, arr := []string{}, []byte(s)
    for (len(arr) % k) != 0 {
        arr = append(arr, fill)
    }
    for i := 0; i <= len(arr) - k; i += k {
        res = append(res, string(arr[i:i + k]))
    }
    return res
}

func divideString1(s string, k int, fill byte) []string {
    res, n := make([]string, 0), len(s)
    for i := 0; i < n; i += k {
        if i + k <= n {
            res = append(res, s[i : i + k])
        } else {
            res = append(res, s[i : n] + strings.Repeat(string(fill), k - (n - i)))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcdefghi", k = 3, fill = "x"
    // Output: ["abc","def","ghi"]
    // Explanation:
    // The first 3 characters "abc" form the first group.
    // The next 3 characters "def" form the second group.
    // The last 3 characters "ghi" form the third group.
    // Since all groups can be completely filled by characters from the string, we do not need to use fill.
    // Thus, the groups formed are "abc", "def", and "ghi".
    fmt.Println(divideString("abcdefghi", 3, 'x')) // ["abc","def","ghi"]
    // Example 2:
    // Input: s = "abcdefghij", k = 3, fill = "x"
    // Output: ["abc","def","ghi","jxx"]
    // Explanation:
    // Similar to the previous example, we are forming the first three groups "abc", "def", and "ghi".
    // For the last group, we can only use the character 'j' from the string. To complete this group, we add 'x' twice.
    // Thus, the 4 groups formed are "abc", "def", "ghi", and "jxx".
    fmt.Println(divideString("abcdefghij", 3, 'x')) // ["abc","def","ghi","jxx"]

    fmt.Println(divideString1("abcdefghi", 3, 'x')) // ["abc","def","ghi"]
    fmt.Println(divideString1("abcdefghij", 3, 'x')) // ["abc","def","ghi","jxx"]
}