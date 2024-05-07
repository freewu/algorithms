package main

// 1134. Armstrong Number
// Given an integer n, return true if and only if it is an Armstrong number.
// The k-digit number n is an Armstrong number if and only if the kth power of each digit sums to n.

// Example 1:
// Input: n = 153
// Output: true
// Explanation: 153 is a 3-digit number, and 153 = 13 + 53 + 33.

// Example 2:
// Input: n = 123
// Output: false
// Explanation: 123 is a 3-digit number, and 123 != 13 + 23 + 33 = 36.
 
// Constraints:
//     1 <= n <= 10^8

import "fmt"

func isArmstrong(n int) bool {
    tobits := func (n int) []int { // 拆成单数(0-9)数组
        res := []int{}
        for n > 0 {
            res = append(res, n%10)
            n /= 10
        }
        return res
    }
    sum, nbits := 0,tobits(n)
    pow := func (b int, p int) int { // 求幂
        res := 1
        for i := 0; i < p; i++ {
            res *= b
        }
        return res
    }
    for i, size := 0, len(nbits); i < size; i++ {
        sum += pow(nbits[i], size)
    }
    return sum == n
}

func main() {
    // Example 1:
    // Input: n = 153
    // Output: true
    // Explanation: 153 is a 3-digit number, and 153 = 13 + 53 + 33.
    fmt.Println(isArmstrong(153)) // true
    // Example 2:
    // Input: n = 123
    // Output: false
    // Explanation: 123 is a 3-digit number, and 123 != 13 + 23 + 33 = 36.
    fmt.Println(isArmstrong(123)) // false
}