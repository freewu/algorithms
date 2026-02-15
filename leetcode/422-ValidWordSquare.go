package main

// 422. Valid Word Square
// Given an array of strings words, return true if it forms a valid word square.
// A sequence of strings forms a valid word square 
// if the kth row and column read the same string, where 0 <= k < max(numRows, numColumns).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/validsq1-grid.jpg" />
// Input: words = ["abcd","bnrt","crmy","dtye"]
// Output: true
// Explanation:
// The 1st row and 1st column both read "abcd".
// The 2nd row and 2nd column both read "bnrt".
// The 3rd row and 3rd column both read "crmy".
// The 4th row and 4th column both read "dtye".
// Therefore, it is a valid word square.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/validsq2-grid.jpg" />
// Input: words = ["abcd","bnrt","crm","dt"]
// Output: true
// Explanation:
// The 1st row and 1st column both read "abcd".
// The 2nd row and 2nd column both read "bnrt".
// The 3rd row and 3rd column both read "crm".
// The 4th row and 4th column both read "dt".
// Therefore, it is a valid word square.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/validsq3-grid.jpg" />
// Input: words = ["ball","area","read","lady"]
// Output: false
// Explanation:
// The 3rd row reads "read" while the 3rd column reads "lead".
// Therefore, it is NOT a valid word square.
 
// Constraints:
//     1 <= words.length <= 500
//     1 <= words[i].length <= 500
//     words[i] consists of only lowercase English letters.

import "fmt"

func validWordSquare(words []string) bool {
    l := len(words)
    for i := 0; i < l; i++ {
        for j := 0; j < len(words[i]); j++ {
            // 处理长度边界问题不够直接返回 false
            if j >= l || i >= len(words[j]) || words[i][j] != words[j][i] {
                return false
            }
        }
    }
    return true
}

func validWordSquare1(words []string) bool {
    for i := range words{
        for j := range words[i] {
            if  j >= len(words) || i >= len(words[j]) || words[i][j] != words[j][i] {
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/09/validsq1-grid.jpg" />
    // Input: words = ["abcd","bnrt","crmy","dtye"]
    // Output: true
    // Explanation:
    // The 1st row and 1st column both read "abcd".
    // The 2nd row and 2nd column both read "bnrt".
    // The 3rd row and 3rd column both read "crmy".
    // The 4th row and 4th column both read "dtye".
    // Therefore, it is a valid word square.
    fmt.Println(validWordSquare([]string{"abcd","bnrt","crmy","dtye"})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/09/validsq2-grid.jpg" />
    // Input: words = ["abcd","bnrt","crm","dt"]
    // Output: true
    // Explanation:
    // The 1st row and 1st column both read "abcd".
    // The 2nd row and 2nd column both read "bnrt".
    // The 3rd row and 3rd column both read "crm".
    // The 4th row and 4th column both read "dt".
    // Therefore, it is a valid word square.
    fmt.Println(validWordSquare([]string{"abcd","bnrt","crm","dt"})) // true
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/04/09/validsq3-grid.jpg" />
    // Input: words = ["ball","area","read","lady"]
    // Output: false
    // Explanation:
    // The 3rd row reads "read" while the 3rd column reads "lead".
    // Therefore, it is NOT a valid word square.
    fmt.Println(validWordSquare([]string{"ball","area","read","lady"})) // false

    fmt.Println(validWordSquare([]string{"bluefrog","leetcode","freewu"})) // false

    fmt.Println(validWordSquare1([]string{"abcd","bnrt","crmy","dtye"})) // true
    fmt.Println(validWordSquare1([]string{"abcd","bnrt","crm","dt"})) // true
    fmt.Println(validWordSquare1([]string{"ball","area","read","lady"})) // false
    fmt.Println(validWordSquare1([]string{"bluefrog","leetcode","freewu"})) // false
}