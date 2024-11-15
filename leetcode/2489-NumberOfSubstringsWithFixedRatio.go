package main

// 2489. Number of Substrings With Fixed Ratio
// You are given a binary string s, and two integers num1 and num2. 
// num1 and num2 are coprime numbers.

// A ratio substring is a substring of s where the ratio between the number of 0's 
// and the number of 1's in the substring is exactly num1 : num2.
//     For example, if num1 = 2 and num2 = 3, then "01011" and "1110000111" are ratio substrings, while "11000" is not.

// Return the number of non-empty ratio substrings of s.

// Note that:
//     1. A substring is a contiguous sequence of characters within a string.
//     2. Two values x and y are coprime if gcd(x, y) == 1 where gcd(x, y) is the greatest common divisor of x and y.

// Example 1:
// Input: s = "0110011", num1 = 1, num2 = 2
// Output: 4
// Explanation: There exist 4 non-empty ratio substrings.
// - The substring s[0..2]: "0110011". It contains one 0 and two 1's. The ratio is 1 : 2.
// - The substring s[1..4]: "0110011". It contains one 0 and two 1's. The ratio is 1 : 2.
// - The substring s[4..6]: "0110011". It contains one 0 and two 1's. The ratio is 1 : 2.
// - The substring s[1..6]: "0110011". It contains two 0's and four 1's. The ratio is 2 : 4 == 1 : 2.
// It can be shown that there are no more ratio substrings.

// Example 2:
// Input: s = "10101", num1 = 3, num2 = 1
// Output: 0
// Explanation: There is no ratio substrings of s. We return 0.

// Constraints:
//     1 <= s.length <= 10^5
//     1 <= num1, num2 <= s.length
//     num1 and num2 are coprime integers.

import "fmt"

func fixedRatio(s string, num1 int, num2 int) int64 {
    res, one, zero := 0, 0, 0
    mp := map[int]int{ 0 : 1 }
    for _, v := range s {
        if v == '1' {
            one++
        } else {
            zero++
        }
        key := num1 * one - num2 * zero
        res += mp[key]
        mp[key]++
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: s = "0110011", num1 = 1, num2 = 2
    // Output: 4
    // Explanation: There exist 4 non-empty ratio substrings.
    // - The substring s[0..2]: "0110011". It contains one 0 and two 1's. The ratio is 1 : 2.
    // - The substring s[1..4]: "0110011". It contains one 0 and two 1's. The ratio is 1 : 2.
    // - The substring s[4..6]: "0110011". It contains one 0 and two 1's. The ratio is 1 : 2.
    // - The substring s[1..6]: "0110011". It contains two 0's and four 1's. The ratio is 2 : 4 == 1 : 2.
    // It can be shown that there are no more ratio substrings.
    fmt.Println(fixedRatio("0110011", 1, 2)) // 4
    // Example 2:
    // Input: s = "10101", num1 = 3, num2 = 1
    // Output: 0
    // Explanation: There is no ratio substrings of s. We return 0.
    fmt.Println(fixedRatio("10101", 3, 1)) // 0
}