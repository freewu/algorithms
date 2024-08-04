package main

// 788. Rotated Digits
// An integer x is a good if after rotating each digit individually by 180 degrees, we get a valid number that is different from x. 
// Each digit must be rotated - we cannot choose to leave it alone.

// A number is valid if each digit remains a digit after rotation. For example:
//     0, 1, and 8 rotate to themselves,
//     2 and 5 rotate to each other (in this case they are rotated in a different direction, in other words, 2 or 5 gets mirrored),
//     6 and 9 rotate to each other, and
//     the rest of the numbers do not rotate to any other number and become invalid.

// Given an integer n, return the number of good integers in the range [1, n].

// Example 1:
// Input: n = 10
// Output: 4
// Explanation: There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
// Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.

// Example 2:
// Input: n = 1
// Output: 0

// Example 3:
// Input: n = 2
// Output: 1

// Constraints:
//     1 <= n <= 10^4

import "fmt"

func rotatedDigits(n int) int {
    res, mp := 0, map[rune]rune{ '0':'0', '1':'1', '8':'8', '2':'5', '5':'2', '6':'9', '9':'6'} 
    for i := 1; i<=n; i++ {
        si, sy, flag := fmt.Sprintf("%d", i), "", true
        for _, c := range si {
            if v, ok := mp[c]; ok {
                sy += string(v)
            } else {
                flag = false
                break
            }
        }
        if flag && si != sy {
            res += 1
        }
    }
    return res
}

func rotatedDigits1(n int) int {
    res := 0
    isGood := func (i int) bool {
        b := false
        for i > 0 {
            if i % 10 == 2 || i % 10 == 5 || i % 10 == 6 || i % 10 == 9 { // 可转
                b = true
            } else if i % 10 == 3 || i % 10 == 4 || i % 10 == 7 { // 遇见不能转的
                return false
            }
            i /= 10
        }
        return b
    }
    for i := 1; i <= n; i++ {
        if isGood(i) {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 10
    // Output: 4
    // Explanation: There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
    // Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.
    fmt.Println(rotatedDigits(10)) // 4
    // Example 2:
    // Input: n = 1
    // Output: 0
    fmt.Println(rotatedDigits(1)) // 0
    // Example 3:
    // Input: n = 2
    // Output: 1
    fmt.Println(rotatedDigits(2)) // 1

    fmt.Println(rotatedDigits1(10)) // 4
    fmt.Println(rotatedDigits1(1)) // 0
    fmt.Println(rotatedDigits1(2)) // 1
}