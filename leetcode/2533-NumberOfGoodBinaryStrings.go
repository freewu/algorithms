package main 

// 2533. Number of Good Binary Strings
// You are given four integers minLength, maxLength, oneGroup and zeroGroup.

// A binary string is good if it satisfies the following conditions:
//     1. The length of the string is in the range [minLength, maxLength].
//     2. The size of each block of consecutive 1's is a multiple of oneGroup.
//         For example in a binary string 00110111100 sizes of each block of consecutive ones are [2,4].
//     3. The size of each block of consecutive 0's is a multiple of zeroGroup.
//         For example, in a binary string 00110111100 sizes of each block of consecutive zeros are [2,1,2].

// Return the number of good binary strings. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Note that 0 is considered a multiple of all the numbers.

// Example 1:
// Input: minLength = 2, maxLength = 3, oneGroup = 1, zeroGroup = 2
// Output: 5
// Explanation: There are 5 good binary strings in this example: "00", "11", "001", "100", and "111".
// It can be proven that there are only 5 good strings satisfying all conditions.

// Example 2:
// Input: minLength = 4, maxLength = 4, oneGroup = 4, zeroGroup = 3
// Output: 1
// Explanation: There is only 1 good binary string in this example: "1111".
// It can be proven that there is only 1 good string satisfying all conditions.

// Constraints:
//     1 <= minLength <= maxLength <= 10^5
//     1 <= oneGroup, zeroGroup <= maxLength

import "fmt"

func goodBinaryStrings(minLength int, maxLength int, oneGroup int, zeroGroup int) int {
    res, mod := 0, 1_000_000_007
    dp := make([]int, maxLength+1)
    dp[0] = 1
    for i := 1; i <= maxLength; i++ {
        if i >= oneGroup {
            dp[i] = (dp[i] + dp[i-oneGroup]) % mod
        }
        if i >= zeroGroup {
            dp[i] = (dp[i] + dp[i-zeroGroup]) % mod
        }
        if i >= minLength {
            res = (res + dp[i]) % mod
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: minLength = 2, maxLength = 3, oneGroup = 1, zeroGroup = 2
    // Output: 5
    // Explanation: There are 5 good binary strings in this example: "00", "11", "001", "100", and "111".
    // It can be proven that there are only 5 good strings satisfying all conditions.
    fmt.Println(goodBinaryStrings(2, 3, 1, 2)) // 5
    // Example 2:
    // Input: minLength = 4, maxLength = 4, oneGroup = 4, zeroGroup = 3
    // Output: 1
    // Explanation: There is only 1 good binary string in this example: "1111".
    // It can be proven that there is only 1 good string satisfying all conditions.
    fmt.Println(goodBinaryStrings(4, 4, 4, 3)) // 1

    fmt.Println(goodBinaryStrings(1, 1, 1, 1)) // 2
    fmt.Println(goodBinaryStrings(10000, 10000, 10000, 10000)) // 2
    fmt.Println(goodBinaryStrings(1, 10000, 10000, 10000)) // 2
    fmt.Println(goodBinaryStrings(10000, 10000, 1, 1)) // 905611805
    fmt.Println(goodBinaryStrings(1, 10000, 1, 1)) // 811223601
}