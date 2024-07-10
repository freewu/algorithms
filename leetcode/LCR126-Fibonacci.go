package main

// LCR 126. 斐波那契数
// 斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//     F(0) = 0，F(1) = 1
//     F(n) = F(n - 1) + F(n - 2)，其中 n > 1
// 给定 n ，请计算 F(n) 。
// 答案需要取模 1e9+7(1000000007) ，如计算初始结果为：1000000008，请返回 1。

// 示例 1：
// 输入：n = 2
// 输出：1
// 解释：F(2) = F(1) + F(0) = 1 + 0 = 1

// 示例 2：
// 输入：n = 3
// 输出：2
// 解释：F(3) = F(2) + F(1) = 1 + 1 = 2

// 示例 3：
// 输入：n = 4
// 输出：3
// 解释：F(4) = F(3) + F(2) = 2 + 1 = 3

// 提示：
// 0 <= n <= 100

import "fmt"

// 递归 超出时间限制 23 / 51 
func fib(n int) int {
    if n == 0 { return 0 }
    if n == 1 { return 1 }
    return fib(n - 1) + fib(n - 2)
}

// 动态规划法 (利用数组来存储)
func fib1(n int) int  {
    if n == 0 || n == 1 {
        return n
    }
    arr := make([]int, n + 1)
    arr[0] = 0
    arr[1] = 1
    for i := 2 ; i <= n; i++ {
        arr[i] = (arr[i - 1] + arr[i - 2]) % 1_000_000_007
    }
    return arr[n] % 1_000_000_007
}

func main() {
    // 示例 1：
    // 输入：n = 2
    // 输出：1
    // 解释：F(2) = F(1) + F(0) = 1 + 0 = 1
    fmt.Println(fib(2)) // 1
    // 示例 2：
    // 输入：n = 3
    // 输出：2
    // 解释：F(3) = F(2) + F(1) = 1 + 1 = 2
    fmt.Println(fib(3)) // 3
    // 示例 3：
    // 输入：n = 4
    // 输出：3
    // 解释：F(4) = F(3) + F(2) = 2 + 1 = 3
    fmt.Println(fib(4)) // 3

    fmt.Println(fib1(2)) // 1
    fmt.Println(fib1(3)) // 3
    fmt.Println(fib1(4)) // 3
    fmt.Println(fib1(43)) // 433494437
    fmt.Println(fib1(95)) // 407059028
}