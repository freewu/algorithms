package main

// 4070. Minimum Rotations to Dial a Number I
// You are given a string s of length 10 consisting of digits.

// The dial contains the digits 0 through 9 in order and is circular, so 0 and 9 are adjacent. 
// The pointer initially points to 0.

// To dial each digit of s in order, rotate the pointer until it points to that digit. 
// Each rotation moves the pointer to an adjacent digit, and you may rotate in either direction. 
// Dialing a digit that the pointer already points to requires no rotations.

// Return the minimum total number of rotations needed to dial every digit of s.

// Example 1:
// Input: s = "0192837465"
// Output: 25
// Explanation:
// Step | From | To | Rotations
// 1    | 0    | 0    | 0
// 2    | 0    | 1    | 1
// 3    | 1    | 9    | 2
// 4    | 9    | 2    | 3
// 5    | 2    | 8    | 4
// 6    | 8    | 3    | 5
// 7    | 3    | 7    | 4
// 8    | 7    | 4    | 3
// 9    | 4    | 6    | 2
// 10   | 6    | 5    | 1
// The total is 0 + 1 + 2 + 3 + 4 + 5 + 4 + 3 + 2 + 1 = 25, which is the minimum total number of rotations.

// Example 2:
// Input: s = "1200210200"
// Output: 12
// Explanation:
// Step | From | To | Rotations
// 1    | 0    | 1    | 1
// 2    | 1    | 2    | 1
// 3    | 2    | 0    | 2
// 4    | 0    | 0    | 0
// 5    | 0    | 2    | 2
// 6    | 2    | 1    | 1  
// 7    | 1    | 0    | 1
// 8    | 0    | 2    | 2
// 9    | 2    | 0    | 2
// 10   | 0    | 0    | 0
// The total is 1 + 1 + 2 + 0 + 2 + 1 + 1 + 2 + 2 + 0 = 12, which is the minimum total number of rotations.

// Constraints:
//     s.length == 10
//     s consists only of digits '0' to '9'

import "fmt"

func minRotations(s string) int {
    res, pre := 0, '0'
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func(a, b int) int { if a < b { return a; }; return b; } 
    for _, v := range s {
        d := abs(int(v) - int(pre))
        res += min(d, 10-d)
        pre = v
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "0192837465"
    // Output: 25
    // Explanation:
    // Step | From | To | Rotations
    // 1    | 0    | 0    | 0
    // 2    | 0    | 1    | 1
    // 3    | 1    | 9    | 2
    // 4    | 9    | 2    | 3
    // 5    | 2    | 8    | 4
    // 6    | 8    | 3    | 5
    // 7    | 3    | 7    | 4
    // 8    | 7    | 4    | 3
    // 9    | 4    | 6    | 2
    // 10   | 6    | 5    | 1
    // The total is 0 + 1 + 2 + 3 + 4 + 5 + 4 + 3 + 2 + 1 = 25, which is the minimum total number of rotations.
    fmt.Println(minRotations("0192837465")) // 25
    // Example 2:
    // Input: s = "1200210200"
    // Output: 12
    // Explanation:
    // Step | From | To | Rotations
    // 1    | 0    | 1    | 1
    // 2    | 1    | 2    | 1
    // 3    | 2    | 0    | 2
    // 4    | 0    | 0    | 0
    // 5    | 0    | 2    | 2
    // 6    | 2    | 1    | 1  
    // 7    | 1    | 0    | 1
    // 8    | 0    | 2    | 2
    // 9    | 2    | 0    | 2
    // 10   | 0    | 0    | 0
    // The total is 1 + 1 + 2 + 0 + 2 + 1 + 1 + 2 + 2 + 0 = 12, which is the minimum total number of rotations.
    fmt.Println(minRotations("1200210200")) // 12
    fmt.Println(minRotations("0192837465")) // 25
    fmt.Println(minRotations("1200210200")) // 12
}