package main

// 3499. Maximize Active Section with Trade I
// You are given a binary string s of length n, where:
//     '1' represents an active section.
//     '0' represents an inactive section.

// You can perform at most one trade to maximize the number of active sections in s. In a trade, you:
//     1. Convert a contiguous block of '1's that is surrounded by '0's to all '0's.
//     2. Afterward, convert a contiguous block of '0's that is surrounded by '1's to all '1's.

// Return the maximum number of active sections in s after making the optimal trade.

// Note: Treat s as if it is augmented with a '1' at both ends, forming t = '1' + s + '1'. 
// The augmented '1's do not contribute to the final count.

// Example 1:
// Input: s = "01"
// Output: 1
// Explanation:
// Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.

// Example 2:
// Input: s = "0100"
// Output: 4
// Explanation:
// String "0100" → Augmented to "101001".
// Choose "0100", convert "101001" → "100001" → "111111".
// The final string without augmentation is "1111". The maximum number of active sections is 4.

// Example 3:
// Input: s = "1000100"
// Output: 7
// Explanation:
// String "1000100" → Augmented to "110001001".
// Choose "000100", convert "110001001" → "110000001" → "111111111".
// The final string without augmentation is "1111111". The maximum number of active sections is 7.

// Example 4:
// Input: s = "01010"
// Output: 4
// Explanation:
// String "01010" → Augmented to "1010101".
// Choose "010", convert "1010101" → "1000101" → "1111101".
// The final string without augmentation is "11110". The maximum number of active sections is 4.

// Constraints:
//     1 <= n == s.length <= 10^5
//     s[i] is either '0' or '1'

import "fmt"

func maxActiveSectionsAfterTrade(s string) int {
    oneCount, convertedOne, curZeroCount, lastZeroCount, zeroSegCount := 0, 0, 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range s {
        if v == '0' {
            curZeroCount++
        } else {
            if curZeroCount != 0 { 
                lastZeroCount = curZeroCount
                zeroSegCount++
            }
            curZeroCount = 0
            oneCount++
        }
        convertedOne = max(convertedOne, curZeroCount + lastZeroCount)
    }
    if curZeroCount != 0 {
        zeroSegCount++
    }
    if zeroSegCount > 1 {
        return oneCount + convertedOne
    }
    return oneCount
}

func main() {
    // Example 1:
    // Input: s = "01"
    // Output: 1
    // Explanation:
    // Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.
    fmt.Println(maxActiveSectionsAfterTrade("01")) // 1
    // Example 2:
    // Input: s = "0100"
    // Output: 4
    // Explanation:
    // String "0100" → Augmented to "101001".
    // Choose "0100", convert "101001" → "100001" → "111111".
    // The final string without augmentation is "1111". The maximum number of active sections is 4.
    fmt.Println(maxActiveSectionsAfterTrade("0100")) // 4
    // Example 3:
    // Input: s = "1000100"
    // Output: 7
    // Explanation:
    // String "1000100" → Augmented to "110001001".
    // Choose "000100", convert "110001001" → "110000001" → "111111111".
    // The final string without augmentation is "1111111". The maximum number of active sections is 7.
    fmt.Println(maxActiveSectionsAfterTrade("1000100")) // 7
    // Example 4:
    // Input: s = "01010"
    // Output: 4
    // Explanation:
    // String "01010" → Augmented to "1010101".
    // Choose "010", convert "1010101" → "1000101" → "1111101".
    // The final string without augmentation is "11110". The maximum number of active sections is 4.
    fmt.Println(maxActiveSectionsAfterTrade("01010")) // 4

    fmt.Println(maxActiveSectionsAfterTrade("0000000000")) // 0
    fmt.Println(maxActiveSectionsAfterTrade("1111111111")) // 10
    fmt.Println(maxActiveSectionsAfterTrade("0000011111")) // 5
    fmt.Println(maxActiveSectionsAfterTrade("1111100000")) // 5
    fmt.Println(maxActiveSectionsAfterTrade("0101010101")) // 7
    fmt.Println(maxActiveSectionsAfterTrade("1010101010")) // 7
}