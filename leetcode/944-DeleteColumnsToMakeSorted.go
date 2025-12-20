package main

// 944. Delete Columns to Make Sorted
// You are given an array of n strings strs, all of the same length.
// The strings can be arranged such that there is one on each line, making a grid.

// For example, strs = ["abc", "bce", "cae"] can be arranged as follows:
//     abc
//     bce
//     cae
// You want to delete the columns that are not sorted lexicographically. 
// In the above example (0-indexed), columns 0 ('a', 'b', 'c') and 2 ('c', 'e', 'e') are sorted, while column 1 ('b', 'c', 'a') is not, so you would delete column 1.

// Return the number of columns that you will delete.

// Example 1:
// Input: strs = ["cba","daf","ghi"]
// Output: 1
// Explanation: The grid looks as follows:
//   cba
//   daf
//   ghi
// Columns 0 and 2 are sorted, but column 1 is not, so you only need to delete 1 column.

// Example 2:
// Input: strs = ["a","b"]
// Output: 0
// Explanation: The grid looks as follows:
//   a
//   b
// Column 0 is the only column and is sorted, so you will not delete any columns.

// Example 3:
// Input: strs = ["zyx","wvu","tsr"]
// Output: 3
// Explanation: The grid looks as follows:
//   zyx
//   wvu
//   tsr
// All 3 columns are not sorted, so you will delete all 3.

// Constraints:
//     n == strs.length
//     1 <= n <= 100
//     1 <= strs[i].length <= 1000
//     strs[i] consists of lowercase English letters.

import "fmt"
import "strings"
import "sort"

func minDeletionSize(strs []string) int {
    res := 0
    SortString := func(word string) string {
        words := strings.Split(word, "")
        sort.Strings(words)
        return strings.Join(words, "")
    }
    for i := 0; i < len(strs[0]); i++ {
        word := ""
        for j := 0; j < len(strs); j++ {
            word += string(strs[j][i])
        }
        if word != SortString(word) {
            res++
        }
    }
    return res
}

func minDeletionSize1(strs []string) int {
    count := 0
    for i := 0; i < len(strs[0]); i++ {
        for j := 0; j< len(strs) - 1; j++ {
            if strs[j][i] > strs[j+1][i] {
                count++
                break
            }
        }
    }
    return count
}

func minDeletionSize2(strs []string) int {
    rows := make([][]byte, len(strs))
    for i := 0; i < len(rows); i++{
        rows[i] = []byte(strs[i])

    }
    notIncreasing := func(rows [][]byte, j int)bool {
        for i:= 1; i < len(rows); i++{
            if rows[i-1][j] > rows[i][j] { return true }
        }
        return false
    }
    count := 0
    for j:= 0; j < len(rows[0]); j++{
        if notIncreasing(rows, j) {
            count++
        }
    }
    return count
}

func main() {
    // Example 1:
    // Input: strs = ["cba","daf","ghi"]
    // Output: 1
    // Explanation: The grid looks as follows:
    //   cba
    //   daf
    //   ghi
    // Columns 0 and 2 are sorted, but column 1 is not, so you only need to delete 1 column.
    fmt.Println(minDeletionSize([]string{"cba","daf","ghi"})) // 1
    // Example 2:
    // Input: strs = ["a","b"]
    // Output: 0
    // Explanation: The grid looks as follows:
    //   a
    //   b
    // Column 0 is the only column and is sorted, so you will not delete any columns.
    fmt.Println(minDeletionSize([]string{"a","b"})) // 0
    // Example 3:
    // Input: strs = ["zyx","wvu","tsr"]
    // Output: 3
    // Explanation: The grid looks as follows:
    //   zyx
    //   wvu
    //   tsr
    // All 3 columns are not sorted, so you will delete all 3.
    fmt.Println(minDeletionSize([]string{"zyx","wvu","tsr"})) // 3

    fmt.Println(minDeletionSize([]string{"bluefrog","leetcode"})) // 6

    fmt.Println(minDeletionSize1([]string{"cba","daf","ghi"})) // 1
    fmt.Println(minDeletionSize1([]string{"a","b"})) // 0
    fmt.Println(minDeletionSize1([]string{"zyx","wvu","tsr"})) // 3
    fmt.Println(minDeletionSize1([]string{"bluefrog","leetcode"})) // 6

    fmt.Println(minDeletionSize2([]string{"cba","daf","ghi"})) // 1
    fmt.Println(minDeletionSize2([]string{"a","b"})) // 0
    fmt.Println(minDeletionSize2([]string{"zyx","wvu","tsr"})) // 3
    fmt.Println(minDeletionSize2([]string{"bluefrog","leetcode"})) // 6
}