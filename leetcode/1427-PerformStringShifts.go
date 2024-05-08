package main

// 1427. Perform String Shifts
// You are given a string s containing lowercase English letters, 
// and a matrix shift, where shift[i] = [directioni, amounti]:
//     directioni can be 0 (for left shift) or 1 (for right shift).
//     amounti is the amount by which string s is to be shifted.
//     A left shift by 1 means remove the first character of s and append it to the end.
//     Similarly, a right shift by 1 means remove the last character of s and add it to the beginning.

// Return the final string after all operations.

// Example 1:
// Input: s = "abc", shift = [[0,1],[1,2]]
// Output: "cab"
// Explanation: 
// [0,1] means shift to left by 1. "abc" -> "bca"
// [1,2] means shift to right by 2. "bca" -> "cab"

// Example 2:
// Input: s = "abcdefg", shift = [[1,1],[1,1],[0,2],[1,3]]
// Output: "efgabcd"
// Explanation:  
// [1,1] means shift to right by 1. "abcdefg" -> "gabcdef"
// [1,1] means shift to right by 1. "gabcdef" -> "fgabcde"
// [0,2] means shift to left by 2. "fgabcde" -> "abcdefg"
// [1,3] means shift to right by 3. "abcdefg" -> "efgabcd"
 
// Constraints:
//     1 <= s.length <= 100
//     s only contains lower case English letters.
//     1 <= shift.length <= 100
//     shift[i].length == 2
//     directioni is either 0 or 1.
//     0 <= amounti <= 100

import "fmt"

func stringShift(s string, shift [][]int) string {
    flag, totalShift, n := 0, 0, len(s)
    for _, val := range shift { // 合并左右移动，左边移动减法，右边移动加法
        if val[0] == 1 {
            flag = 1
        } else {
            flag = -1
        }
        totalShift += flag * val[1]
    }
    totalShift %= n // 求余
    startIndex := (n - totalShift) % n // 计算首字母在字符串中的位置
    return s[startIndex:] + s[:startIndex]
}

func main() {
    // Example 1:
    // Input: s = "abc", shift = [[0,1],[1,2]]
    // Output: "cab"
    // Explanation: 
    // [0,1] means shift to left by 1. "abc" -> "bca"
    // [1,2] means shift to right by 2. "bca" -> "cab"
    fmt.Println(stringShift("abc",[][]int{{0,1},{1,2}})) // cab
    // Example 2:
    // Input: s = "abcdefg", shift = [[1,1],[1,1],[0,2],[1,3]]
    // Output: "efgabcd"
    // Explanation:  
    // [1,1] means shift to right by 1. "abcdefg" -> "gabcdef"
    // [1,1] means shift to right by 1. "gabcdef" -> "fgabcde"
    // [0,2] means shift to left by 2. "fgabcde" -> "abcdefg"
    // [1,3] means shift to right by 3. "abcdefg" -> "efgabcd"
    fmt.Println(stringShift("abcdefg",[][]int{{1,1},{1,1},{0,2},{1,3}})) // efgabcd
}