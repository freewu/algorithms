package main

// 367. Valid Perfect Square
// Given a positive integer num, return true if num is a perfect square or false otherwise.
// A perfect square is an integer that is the square of an integer. 
// In other words, it is the product of some integer with itself.

// You must not use any built-in library function, such as sqrt.

// Example 1:
// Input: num = 16
// Output: true
// Explanation: We return true because 4 * 4 = 16 and 4 is an integer.

// Example 2:
// Input: num = 14
// Output: false
// Explanation: We return false because 3.742 * 3.742 = 14 and 3.742 is not an integer.
 
// Constraints:
//     1 <= num <= 2^31 - 1

import "fmt"

func isPerfectSquare(num int) bool {
    left, right := 0, num + 1
    for left < right {
        mid := left + (right - left) / 2
        if mid * mid >= num {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left * left == num
}

func isPerfectSquare1(num int) bool {
    low, high := 1, num
    // 从 [1, n] 区间内进行二分
    for low <= high {
        mid := low + (high-low) >> 1
        // 找到一个数的平方是否可以等于待判断的数字
        if mid * mid == num { // 若能找到则返回 true
            return true
        } else if mid*mid < num {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    // 找不到就返回 false 
    return false
}

func main() {
    fmt.Println(isPerfectSquare(16)) // true
    fmt.Println(isPerfectSquare(14)) // false

    fmt.Println(isPerfectSquare1(16)) // true
    fmt.Println(isPerfectSquare1(14)) // false
}