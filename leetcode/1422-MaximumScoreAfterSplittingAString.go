package main 

// 1422. Maximum Score After Splitting a String
// Given a string s of zeros and ones, return the maximum score after splitting the string into two non-empty substrings (i.e. left substring and right substring).
// The score after splitting a string is the number of zeros in the left substring plus the number of ones in the right substring.

// Example 1:
// Input: s = "011101"
// Output: 5 
// Explanation: 
// All possible ways of splitting s into two non-empty substrings are:
// left = "0" and right = "11101", score = 1 + 4 = 5 
// left = "01" and right = "1101", score = 1 + 3 = 4 
// left = "011" and right = "101", score = 1 + 2 = 3 
// left = "0111" and right = "01", score = 1 + 1 = 2 
// left = "01110" and right = "1", score = 2 + 1 = 3

// Example 2:
// Input: s = "00111"
// Output: 5
// Explanation: When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5

// Example 3:
// Input: s = "1111"
// Output: 3
 
// Constraints:
//     2 <= s.length <= 500
//     The string s consists of characters '0' and '1' only.

import "fmt"

// func maxScore(s string) int {
//     score := 0
//     left,right := 0, 1
//     ls,rs := 0, 0
//     for right < len(s) {
//         // 判断左边的积分
//         for i := 0; i <= right; i++ {
//             if s[i] == '0' {
//                 ls++
//             }
//         }
//         // 判断右边的积分
//         for i := right; i < len(s); i++ {
//             if s[i] == '1' {
//                 rs++
//             }
//         }
//         if (rs + ls) > score {
//             score = rs + ls
//         }
//         right++
//         left++
//         ls,rs = 0, 0
//     }
//     return score
// }
func maxScore(s string) int {
    max := func (a, b int) int { if a > b { return a; }; return b; }
    var result, totalOnes, leftZeros, leftOnes int 
    for _, char := range s {
        if char == '1' {
            totalOnes++
        }
    }
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '0' {
            leftZeros++
        } else {
            leftOnes++
        }
        result = max(result, leftZeros + totalOnes - leftOnes)
    }
    return result
}

func maxScore1(s string) int {
    n := len(s)
    sums := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        v := 0
        if s[i - 1] == '1' {
            v = 1
        }
        sums[i] = sums[i - 1] + v
    }
    max := 1 - sums[1] + sums[n] - sums[1]
    for i := 2; i < n; i++ {
        // i 为分割点  左边为 0， 自己和右边为1 
        l := i - sums[i]
        r := sums[n] - sums[i]
        if l + r > max {
            max = l + r
        }
    }
    return max
}

func main() {
    // left = "0" and right = "11101", score = 1 + 4 = 5 
    // left = "01" and right = "1101", score = 1 + 3 = 4 
    // left = "011" and right = "101", score = 1 + 2 = 3 
    // left = "0111" and right = "01", score = 1 + 1 = 2 
    // left = "01110" and right = "1", score = 2 + 1 = 3
    fmt.Println(maxScore("011101")) // 5
    // When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5
    fmt.Println(maxScore("00111")) // 5
    fmt.Println(maxScore("1111")) // 3

    fmt.Println(maxScore1("011101")) // 5
    // When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5
    fmt.Println(maxScore1("00111")) // 5
    fmt.Println(maxScore1("1111")) // 3
}