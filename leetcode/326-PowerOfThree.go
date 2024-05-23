package main

// 326. Power of Three
// Given an integer n, return true if it is a power of three. Otherwise, return false.
// An integer n is a power of three, if there exists an integer x such that n == 3x.

// Example 1:
// Input: n = 27
// Output: true
// Explanation: 27 = 3^3

// Example 2:
// Input: n = 0
// Output: false
// Explanation: There is no x where 3x = 0.

// Example 3:
// Input: n = -1
// Output: false
// Explanation: There is no x where 3x = (-1).
 
// Constraints:
//     -2^31 <= n <= 2^31 - 1
 
// Follow up: Could you solve it without loops/recursion?

import "fmt"

func isPowerOfThree(n int) bool {
    for n >= 3 {
        // 每次除3 判断是否能被 3 整除
        if n % 3 == 0 {
            n = n / 3
        } else {
            return false
        }
    }
    return n == 1
}

// 每次乘 3
func isPowerOfThree1(n int) bool {
    res := 1
    for res < n {
        res *= 3
    }
    return res == n
}

// 数论
func isPowerOfThree2(n int) bool {
    return n > 0 && 1162261467 % n == 0
}

func main() {
	fmt.Println(isPowerOfThree(3)) // true 
	fmt.Println(isPowerOfThree(9)) // true 
	fmt.Println(isPowerOfThree(8)) // true 
    fmt.Println(isPowerOfThree(-1)) // false 

    fmt.Println(isPowerOfThree1(3)) // true 
	fmt.Println(isPowerOfThree1(9)) // true 
	fmt.Println(isPowerOfThree1(8)) // true 
    fmt.Println(isPowerOfThree1(-1)) // false 

    fmt.Println(isPowerOfThree2(3)) // true 
	fmt.Println(isPowerOfThree2(9)) // true 
	fmt.Println(isPowerOfThree2(8)) // true 
    fmt.Println(isPowerOfThree2(-1)) // false 
}