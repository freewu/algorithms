package main

// 1987. Number of Unique Good Subsequences
// You are given a binary string binary. 
// A subsequence of binary is considered good if it is not empty and has no leading zeros (with the exception of "0").

// Find the number of unique good subsequences of binary.
//     For example, if binary = "001", then all the good subsequences are ["0", "0", "1"], so the unique good subsequences are "0" and "1". Note that subsequences "00", "01", and "001" are not good because they have leading zeros.

// Return the number of unique good subsequences of binary. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// A subsequence is a sequence that can be derived from another sequence by deleting some 
// or no elements without changing the order of the remaining elements.

// Example 1:
// Input: binary = "001"
// Output: 2
// Explanation: The good subsequences of binary are ["0", "0", "1"].
// The unique good subsequences are "0" and "1".

// Example 2:
// Input: binary = "11"
// Output: 2
// Explanation: The good subsequences of binary are ["1", "1", "11"].
// The unique good subsequences are "1" and "11".

// Example 3:
// Input: binary = "101"
// Output: 5
// Explanation: The good subsequences of binary are ["1", "0", "1", "10", "11", "101"]. 
// The unique good subsequences are "0", "1", "10", "11", and "101".

// Constraints:
//     1 <= binary.length <= 10^5
//     binary consists of only '0's and '1's.

import "fmt"

func numberOfUniqueGoodSubsequences(binary string) int {
    mod := 1_000_000_007
    last0, last1, exist0 := 0, 0, 0
    for _, v := range binary {
        if v == '1' {
            last1 = (last0 + last1 + 1) % mod
        } else{
            last0 = (last0 + last1 ) % mod
            exist0 = 1
        }
    }
    return (last0 + last1 + exist0) % mod
}

func main() {
    // Example 1:
    // Input: binary = "001"
    // Output: 2
    // Explanation: The good subsequences of binary are ["0", "0", "1"].
    // The unique good subsequences are "0" and "1".
    fmt.Println(numberOfUniqueGoodSubsequences("001")) // 2
    // Example 2:
    // Input: binary = "11"
    // Output: 2
    // Explanation: The good subsequences of binary are ["1", "1", "11"].
    // The unique good subsequences are "1" and "11".
    fmt.Println(numberOfUniqueGoodSubsequences("11")) // 2
    // Example 3:
    // Input: binary = "101"
    // Output: 5
    // Explanation: The good subsequences of binary are ["1", "0", "1", "10", "11", "101"]. 
    // The unique good subsequences are "0", "1", "10", "11", and "101".
    fmt.Println(numberOfUniqueGoodSubsequences("101")) // 5
}