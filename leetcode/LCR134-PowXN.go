package main

// LCR 134. Pow(x, n)
// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。

// 示例 1：
// 输入：x = 2.00000, n = 10
// 输出：1024.00000

// 示例 2：
// 输入：x = 2.10000, n = 3
// 输出：9.26100

// 示例 3：
// 输入：x = 2.00000, n = -2
// 输出：0.25000
// 解释：2-2 = 1/22 = 1/4 = 0.25

// 提示：
//     -100.0 < x < 100.0
//     -2^31 <= n <= 2^31-1
//     -10^4 <= xn <= 10^4

import "fmt"

// 时间复杂度 O(log n),空间复杂度 O(1)
func myPow(x float64, n int) float64 {
    if n == 0 { return 1 }
    if n == 1 { return x }
    if n < 0 { // n 为 负数  转成 1 / x
        n, x = -n, 1 / x
    }
    tmp := myPow(x, n / 2) // 递归 每次取一半
    if n % 2 == 0 {
        return tmp * tmp
    }
    return tmp * tmp * x
}

func myPow1(x float64, n int) float64 {
    var fastmul func (x float64, n int) float64
    fastmul = func (x float64, n int) float64 {
        if n == 0 { return 1 }
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