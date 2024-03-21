package main

// 50. Pow(x, n)
// Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

// Example 1:
// Input: x = 2.00000, n = 10
// Output: 1024.00000

// Example 2:
// Input: x = 2.10000, n = 3
// Output: 9.26100

// Example 3:
// Input: x = 2.00000, n = -2
// Output: 0.25000
// Explanation: 2-2 = 1/22 = 1/4 = 0.25

// Constraints:
//     -100.0 < x < 100.0
//     -2^31 <= n <= 2^31-1
//     n is an integer.
//     Either x is not zero or n > 0.
//     -10^4 <= xn <= 10^4

import "fmt"

// 时间复杂度 O(log n),空间复杂度 O(1)
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    }
    // n 为 负数  转成 1 / x
    if n < 0 {
        n = -n
        x = 1 / x
    }
    // 递归 每次取一半
    tmp := myPow(x, n / 2)
    //fmt.Println("n = ",n," x = ",x," tmp = ",tmp)
    if n % 2 == 0 {
        //fmt.Println("tmp * tmp = ",tmp * tmp)
        return tmp * tmp
    }
    //fmt.Println("tmp * tmp * x = ",tmp * tmp * x)
    return tmp * tmp * x
}

// 递归
func myPow1(x float64, n int) float64 {
    var fastmul func (x float64, n int) float64
    fastmul = func (x float64, n int) float64 {
        if n == 0 {
            return 1
        }
        // 每次取一半
        y := fastmul(x, n / 2)
        if n % 2 == 0 {
            return y * y
        }
        return y * y * x
    }
    if n >= 0 {
        return fastmul(x, n)
    }
    return 1.0 / fastmul(x, n)
}


func main() {
    fmt.Println(myPow(2.0, 10)) // 1024.00000
    fmt.Println(myPow(2.1, 3)) // 9.26100
    fmt.Println(myPow(2.0, -2)) // 0.25000

    fmt.Println(myPow1(2.0, 10)) // 1024.00000
    fmt.Println(myPow1(2.1, 3)) // 9.26100
    fmt.Println(myPow1(2.0, -2)) // 0.25000
}