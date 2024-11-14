package main

// 1869. Longer Contiguous Segments of Ones than Zeros
// Given a binary string s, 
// return true if the longest contiguous segment of 1's is strictly longer than the longest contiguous segment of 0's in s, 
// or return false otherwise.
//     For example, in s = "110100010" the longest continuous segment of 1s has length 2, 
//     and the longest continuous segment of 0s has length 3.

// Note that if there are no 0's, then the longest continuous segment of 0's is considered to have a length 0. 
// The same applies if there is no 1's.

// Example 1:
// Input: s = "1101"
// Output: true
// Explanation:
// The longest contiguous segment of 1s has length 2: "[11]01"
// The longest contiguous segment of 0s has length 1: "11[0]1"
// The segment of 1s is longer, so return true.

// Example 2:
// Input: s = "111000"
// Output: false
// Explanation:
// The longest contiguous segment of 1s has length 3: "[111]000"
// The longest contiguous segment of 0s has length 3: "111[000]"
// The segment of 1s is not longer, so return false.

// Example 3:
// Input: s = "110100010"
// Output: false
// Explanation:
// The longest contiguous segment of 1s has length 2: "[11]0100010"
// The longest contiguous segment of 0s has length 3: "1101[000]10"
// The segment of 1s is not longer, so return false.

// Constraints:
//     1 <= s.length <= 100
//     s[i] is either '0' or '1'.

import "fmt"

func checkZeroOnes(s string) bool {
    mx0, mx1, count0, count1 := 0, 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range s {
        if v == '1' {
            mx0 = max(mx0, count0)
            count0 = 0
            count1++
        } else {
            mx1 = max(mx1, count1)
            count1 = 0
            count0++
        }
    }
    return max(mx1, count1) > max(mx0, count0)
}

func main() {
    // Example 1:
    // Input: s = "1101"
    // Output: true
    // Explanation:
    // The longest contiguous segment of 1s has length 2: "[11]01"
    // The longest contiguous segment of 0s has length 1: "11[0]1"
    // The segment of 1s is longer, so return true.
    fmt.Println(checkZeroOnes("1101")) // true
    // Example 2:
    // Input: s = "111000"
    // Output: false
    // Explanation:
    // The longest contiguous segment of 1s has length 3: "[111]000"
    // The longest contiguous segment of 0s has length 3: "111[000]"
    // The segment of 1s is not longer, so return false.
    fmt.Println(checkZeroOnes("111000")) // false
    // Example 3:
    // Input: s = "110100010"
    // Output: false
    // Explanation:
    // The longest contiguous segment of 1s has length 2: "[11]0100010"
    // The longest contiguous segment of 0s has length 3: "1101[000]10"
    // The segment of 1s is not longer, so return false.
    fmt.Println(checkZeroOnes("110100010")) // false
}