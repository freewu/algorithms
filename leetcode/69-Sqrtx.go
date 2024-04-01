package main

// 69. Sqrt(x)
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. 
// The returned integer should be non-negative as well.

// You must not use any built-in exponent function or operator.
//     For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
 
// Example 1:
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.

// Example 2:
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
 
// Constraints:
//     0 <= x <= 2^31 - 1

import "fmt"

// 二分法
func mySqrt1(x int) int {
    l, r := 0, x
    for l < r {
        mid := (l + r + 1) / 2
        if mid * mid > x {
            r = mid - 1
        } else {
            l = mid
        }
    }
    return l
}

// 牛顿法
// http://www.cnblogs.com/AnnieKim/archive/2013/04/18/3028607.html
// xi+1=xi - (xi2 - n) / (2xi) = xi - xi / 2 + n / (2xi) = xi / 2 + n / 2xi = (xi + n/xi) / 2
func mySqrt(x int) int {
    r := x
    for r * r > x {
        r = (r + x / r) / 2
    }
    return r
}

func main() {
	fmt.Println(mySqrt(10)) // 3
	// fmt.Println(mySqrt(2))  // 1
	// fmt.Println(mySqrt(3))  // 1
	// fmt.Println(mySqrt(4))  // 2
	// fmt.Println(mySqrt(8))  // 2
	// fmt.Println(mySqrt(9))  // 3
	// fmt.Println(mySqrt(15)) // 3
	// fmt.Println(mySqrt(16)) // 4
	// fmt.Println(mySqrt(24)) // 4
	// fmt.Println(mySqrt(25)) // 5

    // Explanation: The square root of 4 is 2, so we return 2.
    fmt.Println(mySqrt(4)) // 2
    // Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
    fmt.Println(mySqrt(8)) // 2
    fmt.Println(mySqrt(9)) // 3

    // Explanation: The square root of 4 is 2, so we return 2.
    fmt.Println(mySqrt1(4)) // 2
    // Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
    fmt.Println(mySqrt1(8)) // 2
    fmt.Println(mySqrt1(9)) // 3
}
