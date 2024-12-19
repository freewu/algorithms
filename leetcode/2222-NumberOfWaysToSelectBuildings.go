package main

// 2222. Number of Ways to Select Buildings
// You are given a 0-indexed binary string s which represents the types of buildings along a street where:
//     s[i] = '0' denotes that the ith building is an office and
//     s[i] = '1' denotes that the ith building is a restaurant.

// As a city official, you would like to select 3 buildings for random inspection. 
// However, to ensure variety, no two consecutive buildings out of the selected buildings can be of the same type.
//     For example, given s = "001101", we cannot select the 1st, 3rd, and 5th buildings as that would form "011" which is not allowed due to having two consecutive buildings of the same type.

// Return the number of valid ways to select 3 buildings.

// Example 1:
// Input: s = "001101"
// Output: 6
// Explanation: 
// The following sets of indices selected are valid:
// - [0,2,4] from "001101" forms "010"
// - [0,3,4] from "001101" forms "010"
// - [1,2,4] from "001101" forms "010"
// - [1,3,4] from "001101" forms "010"
// - [2,4,5] from "001101" forms "101"
// - [3,4,5] from "001101" forms "101"
// No other selection is valid. Thus, there are 6 total ways.

// Example 2:
// Input: s = "11100"
// Output: 0
// Explanation: It can be shown that there are no valid selections.

// Constraints:
//     3 <= s.length <= 10^5
//     s[i] is either '0' or '1'.

import "fmt"
import "strings"

func numberOfWays(s string) int64 {
    res, c0, c1, n := 0, 0, 0, len(s)
    cnt0 := strings.Count(s, "0")
    cnt1 := n - cnt0
    for _, c := range s {
        if c == '0' {
            res += c1 * (cnt1 - c1)
            c0++
        } else {
            res += c0 * (cnt0 - c0)
            c1++
        }
    }
    return int64(res)
}

func numberOfWays1(s string) int64 {
    res, n, l, r := 0, len(s), [2]int{}, [2]int{}
    r[0] = strings.Count(s, "0")
    r[1] = n - r[0]
    for _, c := range s {
        x := int(c - '0')
        r[x]--
        res += (l[x^1] * r[x^1])
        l[x]++
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: s = "001101"
    // Output: 6
    // Explanation: 
    // The following sets of indices selected are valid:
    // - [0,2,4] from "001101" forms "010"
    // - [0,3,4] from "001101" forms "010"
    // - [1,2,4] from "001101" forms "010"
    // - [1,3,4] from "001101" forms "010"
    // - [2,4,5] from "001101" forms "101"
    // - [3,4,5] from "001101" forms "101"
    // No other selection is valid. Thus, there are 6 total ways.
    fmt.Println(numberOfWays("001101")) // 6
    // Example 2:
    // Input: s = "11100"
    // Output: 0
    // Explanation: It can be shown that there are no valid selections.
    fmt.Println(numberOfWays("11100")) // 0

    fmt.Println(numberOfWays1("001101")) // 6
    fmt.Println(numberOfWays1("11100")) // 0
}