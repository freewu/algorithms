package main

// 3032. Count Numbers With Unique Digits II
// Given two positive integers a and b, 
// return the count of numbers having unique digits in the range [a, b] (inclusive).

// Example 1:
// Input: a = 1, b = 20
// Output: 19
// Explanation: All the numbers in the range [1, 20] have unique digits except 11. Hence, the answer is 19.

// Example 2:
// Input: a = 9, b = 19
// Output: 10
// Explanation: All the numbers in the range [9, 19] have unique digits except 11. Hence, the answer is 10. 

// Example 3:
// Input: a = 80, b = 120
// Output: 27
// Explanation: There are 41 numbers in the range [80, 120], 27 of which have unique digits.

// Constraints:
//     1 <= a <= b <= 1000

import "fmt"

func numberCount(a int, b int) int {
    res := 0
    check := func(n int) bool {
        set := make(map[int]bool)
        for n > 0 {
            d := n % 10;
            if set[d] { return false }
            set[d] = true
            n /= 10
        }
        return true
    }
    for i := a; i <= b; i++ {
        if check(i) { res++ }
    }
    return res
}

func main() {
    // Example 1:
    // Input: a = 1, b = 20
    // Output: 19
    // Explanation: All the numbers in the range [1, 20] have unique digits except 11. Hence, the answer is 19.
    fmt.Println(numberCount(1, 20)) // 19
    // Example 2:
    // Input: a = 9, b = 19
    // Output: 10
    // Explanation: All the numbers in the range [9, 19] have unique digits except 11. Hence, the answer is 10. 
    fmt.Println(numberCount(9, 19)) // 10
    // Example 3:
    // Input: a = 80, b = 120
    // Output: 27
    // Explanation: There are 41 numbers in the range [80, 120], 27 of which have unique digits.
    fmt.Println(numberCount(80, 120)) // 27
}