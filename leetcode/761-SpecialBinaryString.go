package main

// 761. Special Binary String
// Special binary strings are binary strings with the following two properties:
//     The number of 0's is equal to the number of 1's.
//     Every prefix of the binary string has at least as many 1's as 0's.

// You are given a special binary string s.

// A move consists of choosing two consecutive, non-empty, special substrings of s, and swapping them. 
// Two strings are consecutive if the last character of the first string is exactly one index before the first character of the second string.

// Return the lexicographically largest resulting string possible after applying the mentioned operations on the string.

// Example 1:
// Input: s = "11011000"
// Output: "11100100"
// Explanation: The strings "10" [occuring at s[1]] and "1100" [at s[3]] are swapped.
// This is the lexicographically largest string possible after some number of swaps.

// Example 2:
// Input: s = "10"
// Output: "10"

// Constraints:
//     1 <= s.length <= 50
//     s[i] is either '0' or '1'.
//     s is a special binary string.

import "fmt"
import "sort"

func makeLargestSpecial(s string) string {
    memo := make(map[string]string)
    split := func(s string) []string {
        res, count, start := []string{}, 0, 0
        for end := 0; end < len(s); end++ {
            if s[end] == '1' {
                count++
            } else {
                count--
            }
            if count == 0 {
                res = append(res, s[start:end+1])
                start = end + 1
            }
        }
        return res
    }
    var largestSpecial func(s string, memo map[string]string) string 
    largestSpecial = func(s string, memo map[string]string) string {
        // Base Case: "10", or as parenthesis ()
        n := len(s)
        if n == 2 { return s }
        if _, exists := memo[s]; exists { // Check if solution exists in memory
            return memo[s]
        }
        tuples := split(s)
        // Get max for each individual tuple
        if len(tuples) != 1 {
            for i, tuple := range tuples {
                tuples[i] = largestSpecial(tuple, memo) // Recursive
            }
        } else {
            // Single tuple, break it smaller.
            // Why? Because the sub-tuples inside may NOT be optimal
            // Eg: (()(())); obviously () needs to be swapped with (())
            // resulting in ((())())
            tuple := s[1 : n-1] // Remove first and last chars
            tuples[0] = fmt.Sprintf("1%s0", largestSpecial(tuple, memo))
        }
        // Greedy approach: max value can be obtained by combining
        // tuples in descending order. Can be proven, but
        // think of it like in decimals: {"111", "9", "777"} ==>
        // Max will be always: 9 777 111
        sort.Strings(tuples)
        res := ""
        for i := len(tuples) - 1; 0 <= i; i-- {
            res += tuples[i]
        }
        memo[s] = res
        return res
    }
	return largestSpecial(s, memo)
}

func main() {
    // Example 1:
    // Input: s = "11011000"
    // Output: "11100100"
    // Explanation: The strings "10" [occuring at s[1]] and "1100" [at s[3]] are swapped.
    // This is the lexicographically largest string possible after some number of swaps.
    fmt.Println(makeLargestSpecial("11011000")) // "11100100"
    // Example 2:
    // Input: s = "10"
    // Output: "10"
    fmt.Println(makeLargestSpecial("10")) // "10"

    fmt.Println(makeLargestSpecial("1111111111")) // ""
    fmt.Println(makeLargestSpecial("0000000000")) // ""
    fmt.Println(makeLargestSpecial("1111100000")) // "1111100000"
    fmt.Println(makeLargestSpecial("0000011111")) // "1111010000"
    fmt.Println(makeLargestSpecial("0101010101")) // "0101010101"
    fmt.Println(makeLargestSpecial("1010101010")) // "1010101010"
}