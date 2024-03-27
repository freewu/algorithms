package main

// 202. Happy Number
// Write an algorithm to determine if a number n is happy.
// A happy number is a number defined by the following process:
//     Starting with any positive integer, replace the number by the sum of the squares of its digits.
//     Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
//     Those numbers for which this process ends in 1 are happy.

// Return true if n is a happy number, and false if not.

// Example 1:
// Input: n = 19
// Output: true
// Explanation:
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1

// Example 2:
// Input: n = 2
// Output: false
 
// Constraints:
//     1 <= n <= 2^31 - 1

import "fmt"

// n = 7 有问题
// func isHappy(n int) bool {
//     if n == 1 {
//         return true
//     }
//     if n < 10 {
//         return false
//     }
//     sum, a := 0,0
//     for n > 0 {
//         a = n % 10
//         sum += a * a
//         n = n / 10
//     }
//     return isHappy(sum)
// }

// 快慢指针
func isHappy(n int) bool {
    // 快慢指针：最终会相遇，有循环，当相遇时即完成一个循环
    // 判断循环的结尾/开始是不是符合预期
    if n == 1 {
        return true
    }
    bitSquare := func (n int) int {
        if n == 1 {
            return 1
        }
        res := 0
        for n > 0 {
            x := n % 10
            res += x * x
            n = n / 10
        }
        return res
    }
    fast, slow := bitSquare(n), n 
    for fast != slow {
        fast = bitSquare(fast)
        fast = bitSquare(fast)
        slow = bitSquare(slow)
    }
    return fast == 1
}

func main() {
    // 1^2 + 9^2 = 82
    // 8^2 + 2^2 = 68
    // 6^2 + 8^2 = 100
    // 1^2 + 0^2 + 0^2 = 1
    fmt.Println(isHappy(19)) // true

    fmt.Println(isHappy(2)) // false
}